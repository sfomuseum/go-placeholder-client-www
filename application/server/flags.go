package server

import (
	"github.com/sfomuseum/go-flags/multi"
)

var placeholder_endpoint string

var server_uri string

var prefix string

var static_prefix string

var nextzen_apikey string

var nextzen_style_url string

var nextzen_tile_url string

var proxy_tiles bool

var proxy_tiles_url string

var proxy_tiles_cache string

var proxy_tiles_timeout int

var proxy_test_network bool

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

var cors_origins multi.MultiString
