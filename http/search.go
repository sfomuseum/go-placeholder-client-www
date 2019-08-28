package http

import (
	"errors"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client/results"
	"log"
	"html/template"
	gohttp "net/http"
)

type SearchVars struct {
	Query   string
	Results *results.SearchResults
	Error   error
}

func NewSearchHandler(cl *client.PlaceholderClient, t *template.Template) (gohttp.Handler, error) {

	t = t.Lookup("search")

	if t == nil {
		return nil, errors.New("Missing search template")
	}

	t = t.Funcs(template.FuncMap{
		"Ancestors": Ancestors,
	})

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		log.Println("REQUEST", req.URL.Path)
		
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

		// important if we're trying to use this in a Lambda/API Gateway context

		rsp.Header().Set("Content-Type", "text/html; charset=utf-8")

		err = t.Execute(rsp, search_vars)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
		}

		return
	}

	return gohttp.HandlerFunc(fn), nil
}
