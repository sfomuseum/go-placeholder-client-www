package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-string/dsn"	
	tzhttp "github.com/sfomuseum/go-http-tilezen/http"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client-www/assets/templates"
	"github.com/sfomuseum/go-placeholder-client-www/http"
	"github.com/sfomuseum/go-placeholder-client-www/server"
	"github.com/whosonfirst/go-cache"
	"github.com/whosonfirst/go-cache-blob"	
	"github.com/aaronland/go-http-tangramjs"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"html/template"
	"log"
	gohttp "net/http"
	gourl "net/url"
	"path/filepath"
	"strings"
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

	is_api_gateway := flag.Bool("is-api-gateway", false, "...")

	proxy_tiles := flag.Bool("proxy-tiles", false, "...")

	var proxy_caches flags.MultiString
	flag.Var(&proxy_caches, "proxy-cache-dsn", "...")
		
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

		for _, dsn_str := range proxy_caches {

			dsn_map, err := dsn.StringToDSNWithKeys(dsn_str, "cache")

			if err != nil {
				log.Fatal(err)
			}

			switch dsn_map["cache"] {
			case "blob":

				blob_cache, err := blob.NewBlobCacheWithDSN(dsn_str)

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
		}

		if len(caches) == 0 {
			log.Fatal("No proxy caches defined")
		}
		
		multi_cache, err := cache.NewMultiCache(caches)

		if err != nil {
			log.Fatal(err)
		}
		
		proxy_opts := &tzhttp.TilezenProxyHandlerOptions{
			Cache: multi_cache,
		}

		proxy_handler, err := tzhttp.TilezenProxyHandler(proxy_opts)

		if err != nil {
			log.Fatal(err)
		}

		// prefix...		
		proxy_tiles_url := "/tiles/"			
		mux.Handle(proxy_tiles_url, proxy_handler)

		*nextzen_tile_url = proxy_tiles_url		
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
		IsAPIGateway:      *is_api_gateway,
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
