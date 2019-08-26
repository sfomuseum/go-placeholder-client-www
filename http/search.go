package http

import (
	"errors"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client/results"
	"html/template"
	gohttp "net/http"
)

type SearchVars struct {
	Query   string
	Results *results.SearchResults
	Error   error
}

func NewSearchHandler(cl *client.PlaceholderClient, t *template.Template) (gohttp.HandlerFunc, error) {

	t = t.Lookup("search")

	if t == nil {
		return nil, errors.New("Missing search template")
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		text, err := sanitize.GetString(req, "text")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
		}

		var search_vars SearchVars

		if text != "" {

			search_vars.Query = text
			res, err := cl.Search(text)

			if err != nil {
				search_vars.Error = err
			} else {
				search_vars.Results = res
			}
		}

		err = t.Execute(rsp, search_vars)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
		}

		return
	}

	return gohttp.HandlerFunc(fn), nil
}
