package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/aaronland/go-string/dsn"
	tzhttp "github.com/sfomuseum/go-http-tilezen/http"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client-www/assets/templates"
	"github.com/sfomuseum/go-placeholder-client-www/http"
	"github.com/sfomuseum/go-placeholder-client-www/server"
	"github.com/whosonfirst/go-cache"
	"github.com/whosonfirst/go-cache-blob"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"html/template"
	"log"
	gohttp "net/http"
	gourl "net/url"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	placeholder_endpoint := flag.String("placeholder-endpoint", client.DEFAULT_ENDPOINT, "The address of the Placeholder endpoint to query.")

	var proto = flag.String("protocol", "http", "The protocol for placeholder-client server to listen on. Valid protocols are: http, lambda.")
	host := flag.String("host", "localhost", "The host to listen for requests on.")
	port := flag.Int("port", 8080, "The port to listen for requests on.")

	static_prefix := flag.String("static-prefix", "", "Prepend this prefix to URLs for static assets.")

	nextzen_apikey := flag.String("nextzen-apikey", "", "A valid Nextzen API key")
	nextzen_style_url := flag.String("nextzen-style-url", "/tangram/refill-style.zip", "...")
	nextzen_tile_url := flag.String("nextzen-tile-url", tangramjs.NEXTZEN_MVT_ENDPOINT, "...")

	path_templates := flag.String("templates", "", "An optional string for local templates. This is anything that can be read by the 'templates.ParseGlob' method.")

	proxy_tiles := flag.Bool("proxy-tiles", false, "...")
	proxy_tiles_url := flag.String("proxy-tiles-url", "/tiles/", "...")
	proxy_tiles_dsn := flag.String("proxy-tiles-dsn", "cache=gocache", "...")
	proxy_tiles_timeout := flag.Int("proxy-tiles-timeout", 30, "The maximum number of seconds to allow for fetching a tile from the proxy.")
	proxy_test_network := flag.Bool("proxy-test-network", false, "Ensure outbound network connectivity for proxy tiles")

	enable_api := flag.Bool("api", false, "...")

	flag.Parse()

	err := flags.SetFlagsFromEnvVars("PLACEHOLDER")

	if err != nil {
		log.Fatal(err)
	}

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

	if *path_templates != "" {

		t, err = t.ParseGlob(*path_templates)

		if err != nil {
			log.Fatal(err)
		}

	} else {

		for _, name := range templates.AssetNames() {

			body, err := templates.Asset(name)

			if err != nil {
				log.Fatal(err)
			}

			t, err = t.Parse(string(body))

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if *static_prefix != "" {

		*static_prefix = strings.TrimRight(*static_prefix, "/")

		if !strings.HasPrefix(*static_prefix, "/") {
			log.Fatal("Invalid -static-prefix value")
		}
	}

	// handlers

	mux := gohttp.NewServeMux()

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/ping", ping_handler)

	if *proxy_tiles {

		caches := make([]cache.Cache, 0)

		dsn_map, err := dsn.StringToDSNWithKeys(*proxy_tiles_dsn, "cache")

		if err != nil {
			log.Fatal(err)
		}

		switch dsn_map["cache"] {
		case "blob":

			blob_dsn, ok := dsn_map["blob"]

			if !ok {
				log.Fatal("Missing blob DSN property")
			}

			blob_cache, err := blob.NewBlobCacheWithDSN(blob_dsn)

			if err != nil {
				log.Fatal(err)
			}

			caches = append(caches, blob_cache)

		case "gocache":

			cache_opts, err := cache.DefaultGoCacheOptions()

			if err != nil {
				log.Fatal(err)
			}

			go_cache, err := cache.NewGoCache(cache_opts)

			if err != nil {
				log.Fatal(err)
			}

			caches = append(caches, go_cache)

		case "null":

			null_cache, err := cache.NewNullCache()

			if err != nil {
				log.Fatal(err)
			}

			caches = append(caches, null_cache)

		default:
			log.Fatal("Invalid cache")
		}

		if len(caches) == 0 {
			log.Fatal("No proxy caches defined")
		}

		multi_cache, err := cache.NewMultiCache(caches)

		if err != nil {
			log.Fatal(err)
		}

		timeout := time.Duration(*proxy_tiles_timeout) * time.Second

		proxy_opts := &tzhttp.TilezenProxyHandlerOptions{
			Cache:   multi_cache,
			Timeout: timeout,
		}

		proxy_handler, err := tzhttp.TilezenProxyHandler(proxy_opts)

		if err != nil {
			log.Fatal(err)
		}

		if *proxy_test_network {

			req, err := gohttp.NewRequest("GET", "tile.nextzen.org", nil)

			if err != nil {
				log.Fatal(err)
			}

			cl := new(gohttp.Client)

			ctx, _ := context.WithTimeout(context.Background(), timeout)
			_, err = cl.Do(req.WithContext(ctx))

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

	tangramjs_opts := tangramjs.DefaultTangramJSOptions()
	tangramjs_opts.Nextzen.APIKey = *nextzen_apikey
	tangramjs_opts.Nextzen.StyleURL = *nextzen_style_url
	tangramjs_opts.Nextzen.TileURL = *nextzen_tile_url

	err = bootstrap.AppendAssetHandlersWithPrefix(mux, *static_prefix)

	if err != nil {
		log.Fatal(err)
	}

	search_opts := &http.SearchHandlerOptions{
		PlaceholderClient: cl,
		Templates:         t,
		URLPrefix:         *static_prefix,
	}

	search_handler, err := http.NewSearchHandler(search_opts)

	if err != nil {
		log.Fatal(err)
	}

	search_handler = bootstrap.AppendResourcesHandlerWithPrefix(search_handler, bootstrap_opts, *static_prefix)
	search_handler = tangramjs.AppendResourcesHandlerWithPrefix(search_handler, tangramjs_opts, *static_prefix)

	// auth-y bits go here...

	search_path := "/"

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

		api_handler, err := http.NewAPIHandler(cl)

		if err != nil {
			log.Fatal(err)
		}

		// something something something CORS
		mux.Handle("/api/", api_handler)
	}

	// end of handlers

	address := fmt.Sprintf("http://%s:%d", *host, *port)

	u, err := gourl.Parse(address)

	if err != nil {
		log.Fatal(err)
	}

	s, err := server.NewStaticServer(*proto, u)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatal(err)
	}
}
