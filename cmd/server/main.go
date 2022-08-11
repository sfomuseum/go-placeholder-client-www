package main

import (
	"context"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-ping/v2"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/rs/cors"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	os "github.com/sfomuseum/go-http-opensearch"
	oshttp "github.com/sfomuseum/go-http-opensearch/http"
	tzhttp "github.com/sfomuseum/go-http-tilezen/http"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client-www/http"
	"github.com/sfomuseum/go-placeholder-client-www/templates/html"
	"github.com/whosonfirst/go-cache"
	_ "github.com/whosonfirst/go-cache-blob"
	"html/template"
	"log"
	gohttp "net/http"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	fs := flagset.NewFlagSet("placeholder-client")

	placeholder_endpoint := fs.String("placeholder-endpoint", client.DEFAULT_ENDPOINT, "The address of the Placeholder endpoint to query.")

	server_uri := fs.String("server-uri", "http://localhost:8080", "...")

	static_prefix := fs.String("static-prefix", "", "Prepend this prefix to URLs for static assets.")

	nextzen_apikey := fs.String("nextzen-apikey", "", "A valid Nextzen API key")
	nextzen_style_url := fs.String("nextzen-style-url", "/tangram/refill-style.zip", "...")
	nextzen_tile_url := fs.String("nextzen-tile-url", tangramjs.NEXTZEN_MVT_ENDPOINT, "...")

	proxy_tiles := fs.Bool("proxy-tiles", false, "Proxy (and cache) Nextzen tiles.")
	proxy_tiles_url := fs.String("proxy-tiles-url", "/tiles/", "The URL (a relative path) for proxied tiles.")
	proxy_tiles_cache := fs.String("proxy-tiles-dsn", "gocache://", "A valid tile proxy DSN string.")
	proxy_tiles_timeout := fs.Int("proxy-tiles-timeout", 30, "The maximum number of seconds to allow for fetching a tile from the proxy.")
	proxy_test_network := fs.Bool("proxy-test-network", false, "Ensure outbound network connectivity for proxy tiles")

	enable_api := fs.Bool("api", false, "Enable an API endpoint for Placeholder functionality.")
	enable_api_autocomplete := fs.Bool("api-autocomplete", false, "Enable autocomplete for the 'search' API endpoint.")

	enable_opensearch := fs.Bool("opensearch", true, "...")

	api_url := fs.String("api-url", "/api/", "The URL (a relative path) for the API endpoint.")
	enable_cors := fs.Bool("cors", false, "Enable CORS support for the API endpoint.")

	opensearch_url := fs.String("opensearch-plugin-url", "/opensearch/", "...")
	opensearch_search_template := fs.String("opensearch-search-template", "", "...")
	opensearch_search_form := fs.String("opensearch-search-form", "", "...")

	enable_ready := fs.Bool("ready-check", true, "Enable the Placeholder \"ready\" check handler.")
	ready_ttl := fs.Int("ready-check-ttl", 60, "The time to live, in seconds, for the Placeholder \"check\".")
	ready_url := fs.String("ready-check-url", "/ready/", "The URL (a relative path) for the Placeholder \"ready\" check handler.")

	var cors_origins multi.MultiString

	fs.Var(&cors_origins, "cors-origin", "One or more hosts to restrict CORS support to on the API endpoint. If no origins are defined (and -cors is enabled) then the server will default to all hosts.")

	flagset.Parse(fs)

	ctx := context.Background()

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, "PLACEHOLDER", true)

	if err != nil {
		log.Fatal(err)
	}

	search_path := "/"

	cl, err := client.NewPlaceholderClient(*placeholder_endpoint)

	if err != nil {
		log.Fatal(err)
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
	})

	t, err = t.ParseFS(html.FS, "*.html")

	if err != nil {
		log.Fatal(err)
	}

	if *static_prefix != "" {

		*static_prefix = strings.TrimRight(*static_prefix, "/")

		if !strings.HasPrefix(*static_prefix, "/") {
			log.Fatal("Invalid -static-prefix value")
		}
	}

	// handlers

	mux := gohttp.NewServeMux()

	ping_handler, err := ping.PingPongHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/ping", ping_handler)

	if *proxy_tiles {

		ctx := context.Background()
		tile_cache, err := cache.NewCache(ctx, *proxy_tiles_cache)

		if err != nil {
			log.Fatal(err)
		}

		timeout := time.Duration(*proxy_tiles_timeout) * time.Second

		proxy_opts := &tzhttp.TilezenProxyHandlerOptions{
			Cache:   tile_cache,
			Timeout: timeout,
		}

		proxy_handler, err := tzhttp.TilezenProxyHandler(proxy_opts)

		if err != nil {
			log.Fatal(err)
		}

		if *proxy_test_network {

			test_ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			req, err := gohttp.NewRequest("GET", "tile.nextzen.org", nil)

			if err != nil {
				log.Fatal(err)
			}

			cl := new(gohttp.Client)

			_, err = cl.Do(req.WithContext(test_ctx))

			if err != nil {
				log.Fatal(err)
			}

		}

		// the order here is important - we don't have a general-purpose "add to
		// mux with prefix" function here, like we do in the tangram handler so
		// we set the nextzen tile url with *proxy_tiles_url and then update it
		// (*proxy_tiles_url) with a prefix if necessary (20190911/thisisaaronland)

		*nextzen_tile_url = fmt.Sprintf("%s{z}/{x}/{y}.mvt", *proxy_tiles_url)

		if *static_prefix != "" {

			*proxy_tiles_url = filepath.Join(*static_prefix, *proxy_tiles_url)

			if !strings.HasSuffix(*proxy_tiles_url, "/") {
				*proxy_tiles_url = fmt.Sprintf("%s/", *proxy_tiles_url)
			}
		}

		mux.Handle(*proxy_tiles_url, proxy_handler)
	}

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	bootstrap_opts.JS = []string{
		"/javascript/bootstrap.min.js",
	}

	tangramjs_opts := tangramjs.DefaultTangramJSOptions()
	tangramjs_opts.NextzenOptions.APIKey = *nextzen_apikey
	tangramjs_opts.NextzenOptions.StyleURL = *nextzen_style_url
	tangramjs_opts.NextzenOptions.TileURL = *nextzen_tile_url

	err = bootstrap.AppendAssetHandlersWithPrefix(mux, *static_prefix)

	if err != nil {
		log.Fatal(err)
	}

	if *enable_ready {

		ready_t := time.Now().Add(time.Duration(*ready_ttl) * time.Second)
		ready_handler, err := http.PlaceholderReadyHandler(*placeholder_endpoint, ready_t)

		if err != nil {
			log.Fatalf("Failed to create Placeholder ready handler, %v", err)
		}

		ready_path := *ready_url
		mux.Handle(ready_path, ready_handler)
	}

	search_opts := &http.SearchHandlerOptions{
		PlaceholderClient: cl,
		Templates:         t,
		URLPrefix:         *static_prefix,
	}

	if *enable_ready {
		search_opts.EnableReadyCheck = true
		search_opts.ReadyCheckURL = *ready_url
	}

	search_handler, err := http.NewSearchHandler(search_opts)

	if err != nil {
		log.Fatal(err)
	}

	search_handler = bootstrap.AppendResourcesHandlerWithPrefix(search_handler, bootstrap_opts, *static_prefix)
	search_handler = tangramjs.AppendResourcesHandlerWithPrefix(search_handler, tangramjs_opts, *static_prefix)

	if *enable_opensearch {

		if *opensearch_search_template == "" {
			*opensearch_search_template = filepath.Join(*server_uri, search_path)
		}

		if *opensearch_search_form == "" {
			*opensearch_search_form = filepath.Join(*server_uri, search_path)
		}

		os_desc_opts := &os.BasicDescriptionOptions{
			QueryParameter: "text",
			SearchTemplate: *opensearch_search_template,
			SearchForm:     *opensearch_search_form,
			ImageURI:       os.DEFAULT_IMAGE_URI,
			Name:           "Placeholder",
			Description:    "Search Placeholder",
		}

		os_desc, err := os.BasicDescription(os_desc_opts)

		if err != nil {
			log.Fatal(err)
		}

		os_handler_opts := &oshttp.OpenSearchHandlerOptions{
			Description: os_desc,
		}

		os_handler, err := oshttp.OpenSearchHandler(os_handler_opts)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle(*opensearch_url, os_handler)

		os_plugins := map[string]*os.OpenSearchDescription{
			*opensearch_url: os_desc,
		}

		os_plugins_opts := &oshttp.AppendPluginsOptions{
			Plugins: os_plugins,
		}

		search_handler = oshttp.AppendPluginsHandler(search_handler, os_plugins_opts)
	}

	// auth-y bits go here...

	mux.Handle(search_path, search_handler)

	err = tangramjs.AppendAssetHandlersWithPrefix(mux, *static_prefix)

	if err != nil {
		log.Fatal(err)
	}

	err = http.AppendStaticAssetHandlersWithPrefix(mux, *static_prefix)

	if err != nil {
		log.Fatal(err)
	}

	if *enable_api {

		api_opts := http.DefaultAPIHandlerOptions()
		api_opts.EnableSearchAutoComplete = *enable_api_autocomplete

		api_handler, err := http.NewAPIHandler(cl, api_opts)

		if err != nil {
			log.Fatal(err)
		}

		if *enable_cors {
			cors_wrapper := cors.New(cors.Options{
				AllowedOrigins: cors_origins,
			})

			api_handler = cors_wrapper.Handler(api_handler)
		}

		mux.Handle(*api_url, api_handler)
	}

	// end of handlers

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatal(err)
	}
}
