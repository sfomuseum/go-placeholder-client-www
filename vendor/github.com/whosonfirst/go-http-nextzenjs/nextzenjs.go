package nextzenjs

import (
	"github.com/aaronland/go-http-rewrite"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	_ "log"
	"net/http"
	"path/filepath"
	"strings"
)

type NextzenJSOptions struct {
	AppendAPIKey bool
	AppendJS     bool
	AppendCSS    bool
	APIKey       string
	JS           []string
	CSS          []string
}

func DefaultNextzenJSOptions() *NextzenJSOptions {

	opts := NextzenJSOptions{
		AppendAPIKey: true,
		AppendJS:     true,
		AppendCSS:    true,
		APIKey:       "nextzen-xxxxxx",
		JS:           []string{"/javascript/nextzen.min.js"},
		CSS:          []string{"/css/nextzen.js.css"},
	}

	return &opts
}

func NextzenJSHandler(next http.Handler, opts *NextzenJSOptions) (http.Handler, error) {
	return AppendResourcesHandlerWithPrefix(next, opts, ""), nil
}

func AppendResourcesHandler(next http.Handler, opts *NextzenJSOptions) http.Handler {
	return AppendResourcesHandlerWithPrefix(next, opts, "")
}

func AppendResourcesHandlerWithPrefix(next http.Handler, opts *NextzenJSOptions, prefix string) http.Handler {

	js := opts.JS
	css := opts.CSS

	if prefix != "" {

		for i, path := range js {
			js[i] = appendPrefix(prefix, path)
		}

		for i, path := range css {
			css[i] = appendPrefix(prefix, path)
		}
	}

	var cb rewrite.RewriteHTMLFunc

	cb = func(n *html.Node, w io.Writer) {

		if n.Type == html.ElementNode && n.Data == "head" {

			if opts.AppendJS {

				for _, js := range js {

					script_type := html.Attribute{"", "type", "text/javascript"}
					script_src := html.Attribute{"", "src", js}

					script := html.Node{
						Type:      html.ElementNode,
						DataAtom:  atom.Script,
						Data:      "script",
						Namespace: "",
						Attr:      []html.Attribute{script_type, script_src},
					}

					n.AppendChild(&script)
				}

			}

			if opts.AppendCSS {

				for _, css := range css {
					link_type := html.Attribute{"", "type", "text/css"}
					link_rel := html.Attribute{"", "rel", "stylesheet"}
					link_href := html.Attribute{"", "href", css}

					link := html.Node{
						Type:      html.ElementNode,
						DataAtom:  atom.Link,
						Data:      "link",
						Namespace: "",
						Attr:      []html.Attribute{link_type, link_rel, link_href},
					}

					n.AppendChild(&link)
				}
			}
		}

		if n.Type == html.ElementNode && n.Data == "body" {

			if opts.AppendAPIKey {
				api_key_ns := ""
				api_key_key := "data-nextzen-api-key"
				api_key_value := opts.APIKey

				api_key_attr := html.Attribute{api_key_ns, api_key_key, api_key_value}
				n.Attr = append(n.Attr, api_key_attr)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			cb(c, w)
		}
	}

	return rewrite.RewriteHTMLHandler(next, cb)
}

func NextzenJSAssetsHandler() (http.Handler, error) {

	fs := assetFS()
	return http.FileServer(fs), nil
}

func NextzenJSAssetsHandlerWithPrefix(prefix string) (http.Handler, error) {

	fs_handler, err := NextzenJSAssetsHandler()

	if err != nil {
		return nil, err
	}

	prefix = strings.TrimRight(prefix, "/")
	
	if prefix == "" {
		return fs_handler, nil
	}

	rewrite_func := func(req *http.Request) (*http.Request, error){
		req.URL.Path = strings.Replace(req.URL.Path, prefix, "", 1)
		return req, nil
	}

	rewrite_handler := rewrite.RewriteRequestHandler(fs_handler, rewrite_func)
	return rewrite_handler, nil
}

func AppendAssetHandlers(mux *http.ServeMux) error {
	return AppendAssetHandlersWithPrefix(mux, "")
}

func AppendAssetHandlersWithPrefix(mux *http.ServeMux, prefix string) error {

	asset_handler, err := NextzenJSAssetsHandlerWithPrefix(prefix)

	if err != nil {
		return nil
	}

	for _, path := range AssetNames() {

		path := strings.Replace(path, "static", "", 1)

		if prefix != "" {
			path = appendPrefix(prefix, path)
		}

		mux.Handle(path, asset_handler)
	}

	return nil
}

func appendPrefix(prefix string, path string) string {

	prefix = strings.TrimRight(prefix, "/")

	if prefix != "" {
		path = strings.TrimLeft(path, "/")
		path = filepath.Join(prefix, path)
	}

	return path
}
