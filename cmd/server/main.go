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
	"strings"
)

func main() {

	placeholder_endpoint := flag.String("placeholder-endpoint", client.DEFAULT_ENDPOINT, "The address of the Placeholder endpoint to query.")

	var proto = flag.String("protocol", "http", "The protocol for placeholder-client server to listen on. Valid protocols are: http, lambda.")
	host := flag.String("host", "localhost", "The host to listen for requests on.")
	port := flag.Int("port", 8080, "The port to listen for requests on.")

	prefix := flag.String("prefix", "", "...")

	nextzen_apikey := flag.String("nextzen-apikey", "", "A valid Nextzen API key")
	path_templates := flag.String("templates", "", "An optional string for local templates. This is anything that can be read by the 'templates.ParseGlob' method.")

	flag.Parse()

	err := flags.SetFlagsFromEnvVars("PLACEHOLDER")

	if err != nil {
		log.Fatal(err)
	}

	cl, err := client.NewPlaceholderClient(*placeholder_endpoint)

	if err != nil {
		log.Fatal(err)
	}

	*prefix = strings.TrimRight(*prefix, "/")
	
	mux := gohttp.NewServeMux()

	t := template.New("placeholder-client").Funcs(template.FuncMap{
		"Ancestors": http.Ancestors,
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

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	nextzen_opts := nextzenjs.DefaultNextzenJSOptions()
	nextzen_opts.APIKey = *nextzen_apikey

	search_opts := &http.SearchHandlerOptions{
		PlaceholderClient: cl,
		Templates: t,
		URLPrefix: *prefix,
	}
	
	search_handler, err := http.NewSearchHandler(search_opts)

	if err != nil {
		log.Fatal(err)
	}

	search_handler = bootstrap.AppendResourcesHandlerWithPrefix(search_handler, bootstrap_opts, *prefix)
	search_handler = nextzenjs.AppendResourcesHandlerWithPrefix(search_handler, nextzen_opts, *prefix)

	err = bootstrap.AppendAssetHandlersWithPrefix(mux, *prefix)

	if err != nil {
		log.Fatal(err)
	}

	err = nextzenjs.AppendAssetHandlersWithPrefix(mux, *prefix)

	if err != nil {
		log.Fatal(err)
	}

	static_handler, err := http.StaticHandlerWithPrefix(*prefix)

	if err != nil {
		log.Fatal(err)
	}

	// TO DO: prefix hoohah...
	mux.Handle("/javascript/placeholder.client.maps.js", static_handler)
	mux.Handle("/javascript/placeholder.client.results.js", static_handler)	
	mux.Handle("/javascript/placeholder.client.init.js", static_handler)
	mux.Handle("/css/placeholder.client.css", static_handler)		
	
	// auth-y bits go here, yeah
	// "github.com/abbot/go-http-auth"

	search_path := fmt.Sprintf("%s/", *prefix)
	
	mux.Handle(search_path, search_handler)

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
