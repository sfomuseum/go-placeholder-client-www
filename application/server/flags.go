package server

import (
	"flag"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	"github.com/sfomuseum/go-placeholder-client"
)

var placeholder_endpoint string

// A valid aaronland/go-http-server URI.
var server_uri string

// Prefix for URL mux handlers.
var url_prefix string

// Prefix for HTML paths.
var static_prefix string

var nextzen_apikey string

var nextzen_style_url string

var nextzen_tile_url string

var proxy_tiles bool

var proxy_tiles_url string

var proxy_tiles_cache string

var proxy_tiles_timeout int

var proxy_test_network bool

var enable_www bool

var enable_api bool

var enable_api_autocomplete bool

var enable_opensearch bool

var api_url string

var enable_cors bool

var opensearch_url string

var opensearch_search_template string

var opensearch_search_form string

var enable_ready bool

var ready_ttl int

var ready_url string

var cors_origins multi.MultiCSVString

var cors_allow_credentials bool

var authenticator_uri string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("placeholder-client")

	fs.StringVar(&placeholder_endpoint, "placeholder-endpoint", client.DEFAULT_ENDPOINT, "The address of the Placeholder endpoint to query.")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "...")

	fs.StringVar(&url_prefix, "url-prefix", "", "Prepend this prefix to application URLs.")
	fs.StringVar(&static_prefix, "static-prefix", "", "Prepend this prefix to URLs for static assets.")

	fs.StringVar(&nextzen_apikey, "nextzen-apikey", "", "A valid Nextzen API key")
	fs.StringVar(&nextzen_style_url, "nextzen-style-url", "/tangram/refill-style.zip", "...")
	fs.StringVar(&nextzen_tile_url, "nextzen-tile-url", tangramjs.NEXTZEN_MVT_ENDPOINT, "...")

	fs.BoolVar(&proxy_tiles, "proxy-tiles", false, "Proxy (and cache) Nextzen tiles.")
	fs.StringVar(&proxy_tiles_url, "proxy-tiles-url", "/tiles/", "The URL (a relative path) for proxied tiles.")
	fs.StringVar(&proxy_tiles_url, "proxy-tiles-dsn", "gocache://", "A valid tile proxy DSN string.")
	fs.IntVar(&proxy_tiles_timeout, "proxy-tiles-timeout", 30, "The maximum number of seconds to allow for fetching a tile from the proxy.")
	fs.BoolVar(&proxy_test_network, "proxy-test-network", false, "Ensure outbound network connectivity for proxy tiles")

	fs.BoolVar(&enable_www, "www", true, "Enable a human-facing web endpoint for Placeholder functionality.")

	fs.BoolVar(&enable_api, "api", false, "Enable an API endpoint for Placeholder functionality.")
	fs.BoolVar(&enable_api_autocomplete, "api-autocomplete", false, "Enable autocomplete for the 'search' API endpoint.")

	fs.BoolVar(&enable_opensearch, "opensearch", true, "...")

	fs.StringVar(&api_url, "api-url", "/api/", "The URL (a relative path) for the API endpoint.")
	fs.BoolVar(&enable_cors, "cors", false, "Enable CORS support for the API endpoint.")

	fs.BoolVar(&cors_allow_credentials, "cors-allow-credentials", false, "Enable Access-Control-Allow-Credentials CORS header")

	fs.StringVar(&opensearch_url, "opensearch-plugin-url", "/opensearch/", "...")
	fs.StringVar(&opensearch_search_template, "opensearch-search-template", "", "...")
	fs.StringVar(&opensearch_search_form, "opensearch-search-form", "", "...")

	fs.BoolVar(&enable_ready, "ready-check", true, "Enable the Placeholder \"ready\" check handler.")
	fs.IntVar(&ready_ttl, "ready-check-ttl", 60, "The time to live, in seconds, for the Placeholder \"check\".")
	fs.StringVar(&ready_url, "ready-check-url", "/ready/", "The URL (a relative path) for the Placeholder \"ready\" check handler.")

	fs.Var(&cors_origins, "cors-origin", "One or more hosts to restrict CORS support to on the API endpoint. If no origins are defined (and -cors is enabled) then the server will default to all hosts.")

	fs.StringVar(&authenticator_uri, "authenticator-uri", "null://", "A valid sfomuseum/go-http-auth.Authenticator URI.")

	return fs
}
