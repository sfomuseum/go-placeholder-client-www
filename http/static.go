package http

import (
	"fmt"
	"github.com/sfomuseum/go-placeholder-client-www/static"
	"io/fs"
	gohttp "net/http"
	"net/url"
	"strings"
)

func StaticAssetsHandler() (gohttp.Handler, error) {
	http_fs := gohttp.FS(static.FS)
	return gohttp.FileServer(http_fs), nil
}

func StaticAssetsHandlerWithPrefix(prefix string) (gohttp.Handler, error) {

	fs_handler, err := StaticAssetsHandler()

	if err != nil {
		return nil, err
	}

	fs_handler = gohttp.StripPrefix(prefix, fs_handler)
	return fs_handler, nil
}

func AppendStaticAssetHandlers(mux *gohttp.ServeMux) error {
	return AppendStaticAssetHandlersWithPrefix(mux, "")
}

func AppendStaticAssetHandlersWithPrefix(mux *gohttp.ServeMux, prefix string) error {

	asset_handler, err := StaticAssetsHandlerWithPrefix(prefix)

	if err != nil {
		return nil
	}

	walk_func := func(path string, info fs.DirEntry, err error) error {

		if path == "." {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if prefix != "" {
			path = appendPrefix(prefix, path)
		}

		if !strings.HasPrefix(path, "/") {
			path = fmt.Sprintf("/%s", path)
		}

		// log.Println("APPEND", path)

		mux.Handle(path, asset_handler)
		return nil
	}

	return fs.WalkDir(static.FS, ".", walk_func)
}

func appendPrefix(prefix string, path string) string {

	prefix = strings.TrimRight(prefix, "/")

	if prefix != "" {
		path = strings.TrimLeft(path, "/")
		path, _ = url.JoinPath(prefix, path)
	}

	return path
}
