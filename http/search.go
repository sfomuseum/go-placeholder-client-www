package http

import (
	"errors"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client/filters"
	"github.com/sfomuseum/go-placeholder-client/results"
	wof_sanitize "github.com/whosonfirst/go-sanitize"
	"html/template"
	_ "log"
	gohttp "net/http"
	"strings"
)

type SearchVars struct {
	URLPrefix    string
	IsAPIGateway bool
	Query        string
	Results      *results.SearchResults
	Error        error
}

type SearchHandlerOptions struct {
	PlaceholderClient *client.PlaceholderClient
	IsAPIGateway      bool
	Templates         *template.Template
	URLPrefix         string
}

func NewSearchHandler(opts *SearchHandlerOptions) (gohttp.Handler, error) {

	t := opts.Templates.Lookup("search")

	if t == nil {
		return nil, errors.New("Missing search template")
	}

	t = t.Funcs(template.FuncMap{
		"Ancestors": Ancestors,
	})

	sn_opts := wof_sanitize.DefaultOptions()

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		text, err := sanitize.GetString(req, "text")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		var search_filters []filters.Filter

		possible_filters := []string{
			"lang",
			"placetype",
			"mode",
		}

		q := req.URL.Query()

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
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			str_v := strings.Join(v, ",")

			sanitized_v, err := wof_sanitize.SanitizeString(str_v, sn_opts)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			f, err := filters.NewSearchFilter(sanitized_k, sanitized_v)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			search_filters = append(search_filters, f)
		}

		var search_vars SearchVars
		search_vars.URLPrefix = opts.URLPrefix
		search_vars.IsAPIGateway = opts.IsAPIGateway

		if text != "" {

			search_vars.Query = text
			res, err := opts.PlaceholderClient.Search(text, search_filters...)

			if err != nil {
				search_vars.Error = err
			} else {
				search_vars.Results = res
			}
		}

		// important if we're trying to use this in a Lambda/API Gateway context

		rsp.Header().Set("Content-Type", "text/html; charset=utf-8")

		err = t.Execute(rsp, search_vars)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		return
	}

	return gohttp.HandlerFunc(fn), nil
}
