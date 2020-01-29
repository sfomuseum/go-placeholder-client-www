package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-placeholder-client"
	_ "log"
	gohttp "net/http"
	"path/filepath"
)

func getString(req *gohttp.Request, param string) (string, error) {

	value, err := sanitize.GetString(req, param)

	if err != nil {
		return "", err
	}

	if value == "" {
		msg := fmt.Sprintf("Missing '%s' parameter", param)
		return "", errors.New(msg)
	}

	return value, nil
}

type APIHandlerOptions struct {
	EnableSearchAutoComplete bool
}

func DefaultAPIHandlerOptions() *APIHandlerOptions {

	opts := APIHandlerOptions{
		EnableSearchAutoComplete: false,
	}

	return &opts
}

func NewAPIHandler(cl *client.PlaceholderClient, opts *APIHandlerOptions) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		var api_results interface{}
		var api_err error

		path := req.URL.Path
		method := filepath.Base(path)

		switch method {

		case "findbyid":

			id, err := getString(req, "id")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			api_results, api_err = cl.FindById(id)

		case "query":

			text, err := getString(req, "text")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			}

			api_results, api_err = cl.Query(text)

		case "search":

			text, err := getString(req, "text")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			search_filters, err := SearchFiltersWithRequest(req)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			q := req.URL.Query()

			if q.Get("mode") == "live" && !opts.EnableSearchAutoComplete {
				gohttp.Error(rsp, "Autocomplete is disabled.", gohttp.StatusServiceUnavailable)
				return
			}
			
			api_results, api_err = cl.Search(text, search_filters...)

		case "tokenize":

			text, err := getString(req, "text")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			api_results, api_err = cl.Tokenize(text)

		default:
			api_err = errors.New("Invalid API endpoint")
		}

		if api_err != nil {
			gohttp.Error(rsp, api_err.Error(), gohttp.StatusBadRequest)
			return
		}

		body, err := json.Marshal(api_results)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-type", "application/json")

		rsp.Write(body)
		return
	}

	return gohttp.HandlerFunc(fn), nil
}
