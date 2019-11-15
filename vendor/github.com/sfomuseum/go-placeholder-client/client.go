package client

import (
	"errors"
	"github.com/sfomuseum/go-placeholder-client/filters"
	"github.com/sfomuseum/go-placeholder-client/results"
	"io/ioutil"
	_ "log"
	"net/http"
	"net/url"
	"strings"
)

const DEFAULT_ENDPOINT = "http://localhost:3000"

const PATH_SEARCH string = "/parser/search"
const PATH_QUERY string = "/parser/query"
const PATH_TOKENIZE string = "/parser/tokenize"
const PATH_FINDBYID string = "/parser/findbyid"

type PlaceholderClient struct {
	endpoint    *url.URL
	http_client *http.Client
}

func NewPlaceholderClient(endpoint string) (*PlaceholderClient, error) {

	u, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	if u.Scheme != "http" {
		return nil, errors.New("Invalid scheme")
	}

	cl := PlaceholderClient{
		endpoint:    u,
		http_client: http.DefaultClient,
	}

	return &cl, nil
}

func (cl *PlaceholderClient) Search(term string, search_filters ...filters.Filter) (*results.SearchResults, error) {

	params := map[string]string{
		"text": term,
	}

	for _, f := range search_filters {
		params[f.Key()] = f.Value()
	}

	body, err := cl.ExecuteMethod(PATH_SEARCH, params)

	if err != nil {
		return nil, err
	}

	return results.NewSearchResults(body)
}

func (cl *PlaceholderClient) Query(term string) (*results.QueryResults, error) {

	params := map[string]string{
		"text": term,
	}

	body, err := cl.ExecuteMethod(PATH_QUERY, params)

	if err != nil {
		return nil, err
	}

	return results.NewQueryResults(body)
}

func (cl *PlaceholderClient) Tokenize(term string) (*results.TokenizeResults, error) {

	params := map[string]string{
		"text": term,
	}

	body, err := cl.ExecuteMethod(PATH_TOKENIZE, params)

	if err != nil {
		return nil, err
	}

	return results.NewTokenizeResults(body)
}

func (cl *PlaceholderClient) FindById(ids ...string) (*results.FindByIDResults, error) {

	params := map[string]string{
		"ids": strings.Join(ids, ","),
	}

	body, err := cl.ExecuteMethod(PATH_FINDBYID, params)

	if err != nil {
		return nil, err
	}

	return results.NewFindByIDResults(body)
}

func (cl *PlaceholderClient) ExecuteMethod(path string, params map[string]string) ([]byte, error) {

	rsp, err := cl.ExecuteMethodAsResponse(path, params)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	return ioutil.ReadAll(rsp.Body)
}

func (cl *PlaceholderClient) ExecuteMethodAsResponse(path string, params map[string]string) (*http.Response, error) {

	u, err := cl.requestURL(path, params)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	rsp, err := cl.http_client.Do(req)

	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != 200 {
		rsp.Body.Close()
		return nil, errors.New(rsp.Status)
	}

	return rsp, nil
}

func (cl *PlaceholderClient) requestURL(path string, params map[string]string) (*url.URL, error) {

	u, err := url.Parse(cl.endpoint.String())

	if err != nil {
		return nil, err
	}

	u.Path = path
	q := u.Query()

	for k, v := range params {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	return u, nil
}
