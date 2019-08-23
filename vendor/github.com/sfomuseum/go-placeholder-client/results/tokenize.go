package results

import (
	"encoding/json"
)

type TokenizeResults [][]string

func (t *TokenizeResults) Results() [][]string {
	return *t
}

func NewTokenizeResults(body []byte) (*TokenizeResults, error) {

	var tokenize_results *TokenizeResults
	err := json.Unmarshal(body, &tokenize_results)

	if err != nil {
		return nil, err
	}

	return tokenize_results, nil
}
