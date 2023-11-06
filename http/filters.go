package http

import (
	gohttp "net/http"
	"strings"

	"github.com/sfomuseum/go-placeholder-client/filters"
	wof_sanitize "github.com/whosonfirst/go-sanitize"
)

func SearchFiltersWithRequest(req *gohttp.Request) ([]filters.Filter, error) {

	var search_filters []filters.Filter

	// this should really come from go-placeholder-client/filters
	// (20200129/thisisaaronland)

	possible_filters := []string{
		"lang",
		"placetype",
		"mode",
	}

	q := req.URL.Query()

	sn_opts := wof_sanitize.DefaultOptions()

	for _, k := range possible_filters {

		v, ok := q[k]

		if !ok {
			continue
		}

		if len(v) == 0 {
			continue
		}

		sanitized_k, err := wof_sanitize.SanitizeString(k, sn_opts)

		if err != nil {
			return nil, err
		}

		str_v := strings.Join(v, ",")

		sanitized_v, err := wof_sanitize.SanitizeString(str_v, sn_opts)

		if err != nil {
			return nil, err
		}

		f, err := filters.NewSearchFilter(sanitized_k, sanitized_v)

		if err != nil {
			return nil, err
		}

		search_filters = append(search_filters, f)
	}

	return search_filters, nil
}
