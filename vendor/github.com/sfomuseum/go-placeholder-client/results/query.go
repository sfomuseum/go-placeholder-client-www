package results

import (
	"encoding/json"
)

type QueryResults []int64

func (s *QueryResults) Results() []int64 {
	return *s
}

func NewQueryResults(body []byte) (*QueryResults, error) {

	var query_results *QueryResults
	err := json.Unmarshal(body, &query_results)

	if err != nil {
		return nil, err
	}

	return query_results, nil
}
