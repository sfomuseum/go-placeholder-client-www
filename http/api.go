package http

import (
	"github.com/sfomuseum/go-placeholder-client"
	gohttp "net/http"
)

func NewAPIHandler(cl *client.PlaceholderClient) (gohttp.HandlerFunc, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		return
	}

	return gohttp.HandlerFunc(fn), nil
}
