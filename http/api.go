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

func NewAPIHandler(cl *client.PlaceholderClient) (gohttp.HandlerFunc, error) {

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
			}

			api_results, api_err = cl.Search(term)

		case "tokenize":

			term, err := getString(req, "term")

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
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
