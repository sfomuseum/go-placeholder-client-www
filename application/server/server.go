package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-ping/v2"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/rs/cors"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-http-auth"
	os "github.com/sfomuseum/go-http-opensearch"
	oshttp "github.com/sfomuseum/go-http-opensearch/http"
	tzhttp "github.com/sfomuseum/go-http-tilezen/http"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client-www/http"
	"github.com/sfomuseum/go-placeholder-client-www/templates/html"
	"github.com/whosonfirst/go-cache"
	_ "github.com/whosonfirst/go-cache-blob"
	"html/template"
	"io/fs"
	"log"
	gohttp "net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type AppendHandlersFunc func(context.Context, *gohttp.ServeMux) error

type WrapHandlerFunc func(gohttp.Handler) gohttp.Handler

type RunOptions struct {
	Logger                *log.Logger
	FlagSet               *flag.FlagSet
	Templates             []fs.FS
	AppendHandlersFunc    AppendHandlersFunc
	WrapSearchHandlerFunc WrapHandlerFunc
}

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, flgst *flag.FlagSet, logger *log.Logger) error {

	opts := &RunOptions{
		Logger:    logger,
		FlagSet:   flgst,
		Templates: []fs.FS{html.FS},
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	flagset.Parse(opts.FlagSet)

	err := flagset.SetFlagsFromEnvVars(opts.FlagSet, "PLACEHOLDER")

	if err != nil {
		return fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	search_url := "/"
	ping_url := "/ping"

	if url_prefix != "" {

		search_url, _ = url.JoinPath(url_prefix, search_url)
		ping_url, _ = url.JoinPath(url_prefix, ping_url)

		api_url, _ = url.JoinPath(url_prefix, api_url)
		opensearch_url, _ = url.JoinPath(url_prefix, opensearch_url)
		ready_url, _ = url.JoinPath(url_prefix, ready_url)

		proxy_tiles_url, _ = url.JoinPath(url_prefix, proxy_tiles_url)

		// We shouldn't have to do this but since it's not possible to serve local Nextzen
		// styles in an AWS Lambda/API Gateway configuration we're just going to support
		// remote URLs...

		if !strings.HasPrefix(nextzen_tile_url, "http") {
			nextzen_tile_url, _ = url.JoinPath(url_prefix, nextzen_tile_url)
		}
	}

	cl, err := client.NewPlaceholderClient(placeholder_endpoint)

	if err != nil {
		return fmt.Errorf("Failed to create new Placeholder client, %w", err)
	}

	t := template.New("placeholder-client").Funcs(template.FuncMap{
		"Add": func(i int, offset int) int {
			return i + offset
		},
		"Ancestors": http.Ancestors,
		"Join": func(root string, path string) string {

			root = strings.TrimRight(root, "/")

			if root != "" {
				path = filepath.Join(root, path)
			}

			return path
		},
		// For example: {{ if (IsAvailable "Account" .) }}
		"IsAvailable": func(name string, data interface{}) bool {
			v := reflect.ValueOf(data)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() != reflect.Struct {
				return false
			}
			return v.FieldByName(name).IsValid()
		},
	})

	for _, fs := range opts.Templates {

		t, err = t.ParseFS(fs, "*.html")

		if err != nil {
			return fmt.Errorf("Failed to parse templates, %w", err)
		}
	}

	if static_prefix != "" {

		static_prefix = strings.TrimRight(static_prefix, "/")

		if !strings.HasPrefix(static_prefix, "/") {
			return fmt.Errorf("Invalid -static-prefix value")
		}
	}

	/*

		create a sfomuseum/go-http-auth instance for handling authentication
		the default authenticator is null:// which allows all requests - if
		you need something more resrictive you will need to create a custom
		implementation of the auth.Authenticator interface and load it like
		this (for example):

		package main

		import (
			"context"
			_ "github.com/{YOU}/{YOUR_AUTHENTICATOR_IMPLEMENTATION}"
			"github.com/sfomuseum/go-placeholder-client-www/application/server"
		)

		func main() {
			ctx := context.Background()
			server.Run(ctx)
		}

	*/

	authenticator, err := auth.NewAuthenticator(ctx, authenticator_uri)

	if err != nil {
		return fmt.Errorf("Failed to create authenticator, %w", err)
	}

	// handlers

	mux := gohttp.NewServeMux()

	ping_handler, err := ping.PingPongHandler()

	if err != nil {
		return fmt.Errorf("Failed to create ping handler, %w", err)
	}

	mux.Handle(ping_url, ping_handler)
	// logger.Printf("Installed ping handler on %s", ping_url)

	if enable_www {

		if proxy_tiles {

			ctx := context.Background()
			tile_cache, err := cache.NewCache(ctx, proxy_tiles_cache)

			if err != nil {
				return fmt.Errorf("Failed to create proxy tile cache, %w", err)
			}

			timeout := time.Duration(proxy_tiles_timeout) * time.Second

			proxy_opts := &tzhttp.TilezenProxyHandlerOptions{
				Cache:   tile_cache,
				Timeout: timeout,
			}

			proxy_handler, err := tzhttp.TilezenProxyHandler(proxy_opts)

			if err != nil {
				return fmt.Errorf("Failed to create proxy tile handler, %w", err)
			}

			if proxy_test_network {

				test_ctx, cancel := context.WithTimeout(ctx, timeout)
				defer cancel()

				req, err := gohttp.NewRequest("GET", "tile.nextzen.org", nil)

				if err != nil {
					return fmt.Errorf("Failed to create request for tile.nextzen.org, %w", err)
				}

				cl := new(gohttp.Client)

				_, err = cl.Do(req.WithContext(test_ctx))

				if err != nil {
					return fmt.Errorf("Failed to contact tile.nextzen.org, %w", err)
				}

			}

			// the order here is important - we don't have a general-purpose "add to
			// mux with prefix" function here, like we do in the tangram handler so
			// we set the nextzen tile url with proxy_tiles_url and then update it
			// (proxy_tiles_url) with a prefix if necessary (20190911/thisisaaronland)

			nextzen_tile_url = fmt.Sprintf("%s{z}/{x}/{y}.mvt", proxy_tiles_url)

			if static_prefix != "" {

				proxy_tiles_url = filepath.Join(static_prefix, proxy_tiles_url)

				if !strings.HasSuffix(proxy_tiles_url, "/") {
					proxy_tiles_url = fmt.Sprintf("%s/", proxy_tiles_url)
				}
			}

			mux.Handle(proxy_tiles_url, proxy_handler)
			// logger.Printf("Installed proxy tiles URL on %s", proxy_tiles_url)
		}

		bootstrap_opts := bootstrap.DefaultBootstrapOptions()

		bootstrap_opts.JS = []string{
			"/javascript/bootstrap.min.js",
		}

		tangramjs_opts := tangramjs.DefaultTangramJSOptions()
		tangramjs_opts.NextzenOptions.APIKey = nextzen_apikey
		tangramjs_opts.NextzenOptions.StyleURL = nextzen_style_url
		tangramjs_opts.NextzenOptions.TileURL = nextzen_tile_url

		err = bootstrap.AppendAssetHandlersWithPrefix(mux, url_prefix)

		if err != nil {
			return fmt.Errorf("Failed to append Bootstrap assets, %w", err)
		}

		if enable_ready {

			ready_t := time.Now().Add(time.Duration(ready_ttl) * time.Second)
			ready_handler, err := http.PlaceholderReadyHandler(placeholder_endpoint, ready_t)

			if err != nil {
				return fmt.Errorf("Failed to create Placeholder ready handler, %v", err)
			}

			mux.Handle(ready_url, ready_handler)
			// logger.Printf("Install ready handler on %s", ready_url)
		}

		search_opts := &http.SearchHandlerOptions{
			PlaceholderClient: cl,
			Templates:         t,
			StaticPrefix:      static_prefix,
			Authenticator:     authenticator,
		}

		if enable_ready {

			static_ready_url := ready_url

			if static_prefix != "" {
				static_ready_url, _ = url.JoinPath(static_prefix, static_ready_url)
			}

			search_opts.EnableReadyCheck = enable_ready
			search_opts.ReadyCheckURL = static_ready_url
		}

		search_handler, err := http.NewSearchHandler(search_opts)

		if err != nil {
			return fmt.Errorf("Failed to create search handler, %w", err)
		}

		search_handler = bootstrap.AppendResourcesHandlerWithPrefix(search_handler, bootstrap_opts, static_prefix)
		search_handler = tangramjs.AppendResourcesHandlerWithPrefix(search_handler, tangramjs_opts, static_prefix)

		if enable_opensearch {

			if opensearch_search_template == "" {
				opensearch_search_template = filepath.Join(server_uri, search_url)
			}

			if opensearch_search_form == "" {
				opensearch_search_form = filepath.Join(server_uri, search_url)
			}

			os_desc_opts := &os.BasicDescriptionOptions{
				QueryParameter: "text",
				SearchTemplate: opensearch_search_template,
				SearchForm:     opensearch_search_form,
				ImageURI:       os.DEFAULT_IMAGE_URI,
				Name:           "Placeholder",
				Description:    "Search Placeholder",
			}

			os_desc, err := os.BasicDescription(os_desc_opts)

			if err != nil {
				return fmt.Errorf("Failed to create opensearc description, %w", err)
			}

			os_handler_opts := &oshttp.OpenSearchHandlerOptions{
				Description: os_desc,
			}

			os_handler, err := oshttp.OpenSearchHandler(os_handler_opts)

			if err != nil {
				return fmt.Errorf("Failed to create opensearch handler, %w", err)
			}

			mux.Handle(opensearch_url, os_handler)
			// logger.Printf("Install opensearch handler on %s", opensearch_url)

			os_plugins := map[string]*os.OpenSearchDescription{
				opensearch_url: os_desc,
			}

			os_plugins_opts := &oshttp.AppendPluginsOptions{
				Plugins: os_plugins,
			}

			search_handler = oshttp.AppendPluginsHandler(search_handler, os_plugins_opts)
		}

		if opts.WrapSearchHandlerFunc != nil {
			search_handler = opts.WrapSearchHandlerFunc(search_handler)
		}

		search_handler = authenticator.WrapHandler(search_handler)

		mux.Handle(search_url, search_handler)
		// logger.Printf("Installed search handler on %s", search_url)

		err = tangramjs.AppendAssetHandlersWithPrefix(mux, url_prefix)

		if err != nil {
			return fmt.Errorf("Failed to append Tangram assets, %w", err)
		}

		err = http.AppendStaticAssetHandlersWithPrefix(mux, static_prefix)

		if err != nil {
			return fmt.Errorf("Failed to append application assets, %w", err)
		}
	}

	if enable_api {

		api_opts := http.DefaultAPIHandlerOptions()
		api_opts.EnableSearchAutoComplete = enable_api_autocomplete
		api_opts.Authenticator = authenticator

		api_handler, err := http.NewAPIHandler(cl, api_opts)

		if err != nil {
			return fmt.Errorf("Failed to create API handler, %w", err)
		}

		if enable_cors {
			cors_wrapper := cors.New(cors.Options{
				AllowedOrigins:   cors_origins,
				AllowCredentials: cors_allow_credentials,
			})

			api_handler = cors_wrapper.Handler(api_handler)
		}

		api_handler = authenticator.WrapHandler(api_handler)
		mux.Handle(api_url, api_handler)

		// logger.Printf("Installed API handler on %s", api_url)
	}

	// end of handlers

	if opts.AppendHandlersFunc != nil {

		err := opts.AppendHandlersFunc(ctx, mux)

		if err != nil {
			return fmt.Errorf("Failed to append handlers, %w", err)
		}
	}

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new server, %w", err)
	}

	opts.Logger.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to serve requests, %w", err)
	}

	return nil
}
