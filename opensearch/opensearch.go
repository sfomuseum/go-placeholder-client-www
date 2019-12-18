package opensearch

import (
	os "github.com/sfomuseum/go-http-opensearch"
)

type QueryDescriptionOptions struct {
	QueryParameter string
	SearchTemplate string
	SearchForm     string
	ImageURI       string
	Name           string
	Description    string
}

func QueryDescription(opts *QueryDescriptionOptions) (*os.OpenSearchDescription, error) {

	im := &os.OpenSearchImage{
		Height: os.DEFAULT_IMAGE_HEIGHT,
		Width:  os.DEFAULT_IMAGE_WIDTH,
		URI:    opts.ImageURI,
	}

	params := []*os.OpenSearchURLParameter{
		&os.OpenSearchURLParameter{
			Name:  opts.QueryParameter,
			Value: os.DEFAULT_SEARCHTERMS,
		},
	}

	u := &os.OpenSearchURL{
		Type:       os.DEFAULT_URL_TYPE,
		Method:     os.DEFAULT_URL_METHOD,
		Template:   opts.SearchTemplate,
		Parameters: params,
	}

	desc := &os.OpenSearchDescription{
		NSMoz:         os.NS_MOZ,
		NSOpenSearch:  os.NS_OPENSEARCH,
		InputEncoding: "UTF-8",
		ShortName:     opts.Name,
		Description:   opts.Description,
		Image:         im,
		URL:           u,
		SearchForm:    opts.SearchForm,
	}

	return desc, nil
}
