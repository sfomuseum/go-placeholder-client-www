package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client-www/assets/templates"
	"github.com/sfomuseum/go-placeholder-client-www/http"
	"github.com/whosonfirst/go-http-nextzenjs"
	"html/template"
	"log"
	gohttp "net/http"
)

func main() {

	placeholder_endpoint := flag.String("placeholder-endpoint", client.DEFAULT_ENDPOINT, "...")

	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	nextzen_apikey := flag.String("nextzen-apikey", "", "...")
	path_templates := flag.String("templates", "", "...")

	flag.Parse()

	cl, err := client.NewPlaceholderClient(*placeholder_endpoint)

	if err != nil {
		log.Fatal(err)
	}

	mux := gohttp.NewServeMux()

	var t *template.Template

	if *path_templates != "" {

		t, err = template.ParseGlob(*path_templates)

		if err != nil {
			log.Fatal(err)
		}

	} else {

		t = template.New("placeholder")

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

	search_handler, err := http.NewSearchHandler(cl, t)

	if err != nil {
		log.Fatal(err)
	}

	search_handler = bootstrap.AppendResourcesHandler(search_handler, bootstrap_opts)

	search_handler, err = nextzenjs.NextzenJSHandler(search_handler, nextzen_opts)

	if err != nil {
		log.Fatal(err)
	}

	// auth-y bits go here, yeah
	// "github.com/abbot/go-http-auth"

	mux.Handle("/", search_handler)

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatal(err)
	}

	err = nextzenjs.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatal(err)
	}

	www_endpoint := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening for requests on %s\n", www_endpoint)

	err = gohttp.ListenAndServe(www_endpoint, mux)

	if err != nil {
		log.Fatal(err)
	}
}
