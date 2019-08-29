package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client-www/assets/templates"
	"github.com/sfomuseum/go-placeholder-client-www/http"
	"github.com/sfomuseum/go-placeholder-client-www/server"
	"github.com/whosonfirst/go-http-nextzenjs"
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
	path_templates := flag.String("templates", "", "An optional string for local templates. This is anything that can be read by the 'templates.ParseGlob' method.")

	is_api_gateway := flag.Bool("is-api-gateway", false, "...")
	
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
	
	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	nextzen_opts := nextzenjs.DefaultNextzenJSOptions()
	nextzen_opts.APIKey = *nextzen_apikey

	err = bootstrap.AppendAssetHandlersWithPrefix(mux, *static_prefix)

	if err != nil {
		log.Fatal(err)
	}

	search_opts := &http.SearchHandlerOptions{
		PlaceholderClient: cl,
		Templates:         t,
		URLPrefix:         *static_prefix,
		IsAPIGateway:	*is_api_gateway,
	}

	search_handler, err := http.NewSearchHandler(search_opts)

	if err != nil {
		log.Fatal(err)
	}

	search_handler = bootstrap.AppendResourcesHandlerWithPrefix(search_handler, bootstrap_opts, *static_prefix)
	search_handler = nextzenjs.AppendResourcesHandlerWithPrefix(search_handler, nextzen_opts, *static_prefix)

	// auth-y bits go here...

	search_path := "/"

	mux.Handle(search_path, search_handler)

	err = nextzenjs.AppendAssetHandlersWithPrefix(mux, *static_prefix)

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
