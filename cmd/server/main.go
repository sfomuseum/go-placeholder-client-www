package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client-www/http"	
	"github.com/whosonfirst/go-http-nextzenjs"
	"log"
	gohttp "net/http"
)

func main() {

	placeholder_endpoint := flag.String("placeholder-endpoint", client.DEFAULT_ENDPOINT, "...")

	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")
	
	flag.Parse()
	
	cl, err := client.NewPlaceholderClient(*placeholder_endpoint)

	if err != nil {
		log.Fatal(err)
	}
	
	mux := gohttp.NewServeMux()

	search_handler, err := http.NewSearchHandler(cl)

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
