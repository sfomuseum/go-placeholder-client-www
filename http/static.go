package http

import (
	gohttp "net/http"
	"github.com/aaronland/go-http-rewrite"
	"strings"
)

func StaticFileSystem() (gohttp.FileSystem, error) {
	fs := assetFS()
	return fs, nil
}

func StaticHandler() (gohttp.Handler, error) {

	fs := assetFS()
	return gohttp.FileServer(fs), nil
}

func StaticHandlerWithPrefix(prefix string) (gohttp.Handler, error) {

	fs_handler, err := StaticHandler()

	if err != nil {
		return nil, err
	}
	
	prefix = strings.TrimRight(prefix, "/")

	if prefix == "" {
		return fs_handler, nil
	}

	rewrite_func := func(req *gohttp.Request) (*gohttp.Request, error){
		req.URL.Path = strings.Replace(req.URL.Path, prefix, "", 1)
		return req, nil
	}

	rewrite_handler := rewrite.RewriteRequestHandler(fs_handler, rewrite_func)
	return rewrite_handler, nil	
}
