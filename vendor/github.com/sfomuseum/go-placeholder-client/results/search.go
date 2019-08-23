package results

import (
	"encoding/json"
)

type SearchResults []*PlaceholderRecord

func (s *SearchResults) Results() []*PlaceholderRecord {
	return *s
}

func NewSearchResults(body []byte) (*SearchResults, error) {

	var search_results *SearchResults
	err := json.Unmarshal(body, &search_results)

	if err != nil {
		return nil, err
	}

	return search_results, nil
}
