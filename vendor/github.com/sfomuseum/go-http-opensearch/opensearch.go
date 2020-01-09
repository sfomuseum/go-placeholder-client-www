package opensearch

import (
	"encoding/xml"
)

const DEFAULT_IMAGE_HEIGHT int = 16
const DEFAULT_IMAGE_WIDTH int = 16
const DEFAULT_URL_TYPE string = "text/html"
const DEFAULT_URL_METHOD string = "GET"
const DEFAULT_SEARCHTERMS string = "{searchTerms}"

// https://thenounproject.com/search/?q=globe&i=2115311

const DEFAULT_IMAGE_URI string = "data:image/x-icon;base64,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"


const NS_MOZ string = "http://www.mozilla.org/2006/browser/search/"
const NS_OPENSEARCH string = "http://a9.com/-/spec/opensearch/1.1/"

type OpenSearchImage struct {
	Height int    `xml:"height,attr"`
	Width  int    `xml:"width,attr"`
	URI    string `xml:",chardata"`
}

type OpenSearchURL struct {
	Type       string                    `xml:"type,attr"`
	Method     string                    `xml:"method,attr"`
	Template   string                    `xml:"template,attr"`
	Parameters []*OpenSearchURLParameter `xml:"Param"`
}

type OpenSearchURLParameter struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type OpenSearchDescription struct {
	XMLName       xml.Name         `xml:"OpenSearchDescription"`
	NSMoz         string           `xml:"xmlns:moz,attr"`
	InputEncoding string           `xml:"InputEncoding"`
	NSOpenSearch  string           `xml:"xmlns,attr"`
	ShortName     string           `xml:"ShortName"`
	Description   string           `xml:"Description"`
	Image         *OpenSearchImage `xml:"Image"`
	URL           *OpenSearchURL   `xml:"Url"`
	SearchForm    string           `xml:"moz:searchForm"`
}

func (d *OpenSearchDescription) Marshal() ([]byte, error) {

	enc, err := xml.Marshal(d)

	if err != nil {
		return nil, err
	}

	body := []byte(xml.Header)
	body = append(body, enc...)

	return body, nil
}
