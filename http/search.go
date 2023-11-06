package http

import (
	"errors"
	"html/template"
	_ "log"
	gohttp "net/http"

	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client/results"
)

type SearchVars struct {
	StaticPrefix     string
	IsAPIGateway     bool
	Query            string
	Results          *results.SearchResults
	Error            error
	EnableReadyCheck bool
	ReadyCheckURL    string
	Account          *auth.Account
}

type SearchHandlerOptions struct {
	PlaceholderClient *client.PlaceholderClient
	IsAPIGateway      bool
	Templates         *template.Template
	StaticPrefix      string
	EnableReadyCheck  bool
	ReadyCheckURL     string
	Authenticator     auth.Authenticator
}

func NewSearchHandler(opts *SearchHandlerOptions) (gohttp.Handler, error) {

	t := opts.Templates.Lookup("search")

	if t == nil {
		return nil, errors.New("Missing search template")
	}

	t = t.Funcs(template.FuncMap{
		"Ancestors": Ancestors,
	})

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		acct, err := opts.Authenticator.GetAccountForRequest(req)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusForbidden)
			return
		}

		text, err := sanitize.GetString(req, "text")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		search_filters, err := SearchFiltersWithRequest(req)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		var search_vars SearchVars
		search_vars.StaticPrefix = opts.StaticPrefix
		search_vars.IsAPIGateway = opts.IsAPIGateway
		search_vars.EnableReadyCheck = opts.EnableReadyCheck
		search_vars.ReadyCheckURL = opts.ReadyCheckURL
		search_vars.Account = acct

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
