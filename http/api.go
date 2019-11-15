package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client/filters"
	wof_sanitize "github.com/whosonfirst/go-sanitize"
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

func NewAPIHandler(cl *client.PlaceholderClient) (gohttp.Handler, error) {

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

			term, err := getString(req, "term")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			}

			api_results, api_err = cl.Query(term)

		case "search":

			term, err := getString(req, "term")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			search_filters := make([]filters.Filter, 0)

			sn_opts := wof_sanitize.DefaultOptions()
			q := req.URL.Query()

			for k, values := range q {

				if k == "term" {
					continue
				}
				
				sanitized_k, err := wof_sanitize.SanitizeString(k, sn_opts)

				if err != nil {
					gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
					return
				}

				for _, v := range values {

					sanitized_v, err := wof_sanitize.SanitizeString(v, sn_opts)

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
			}

			api_results, api_err = cl.Search(term, search_filters...)

		case "tokenize":

			term, err := getString(req, "term")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			api_results, api_err = cl.Tokenize(term)

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
