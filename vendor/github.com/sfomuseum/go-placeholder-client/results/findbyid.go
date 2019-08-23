package results

import (
	"encoding/json"
)

/*

TBD: the response value for findbyid queries is different than for search queries,
specifically the "lineage" property:

"101748479": {
    ...
    "lineage": [
      {
        "continent_id": 102191581,
        "country_id": 85633111,
        "county_id": 102063261,
        "locality_id": 101748479,
        "macrocounty_id": 404227567,
        "region_id": 85682571
      }
    ],
    ...
}

which results in the following:

go run cmd/findbyid/main.go 85922583 101748479
2019/08/23 22:44:48 GET http://localhost:3000/parser/findbyid?ids=85922583%2C101748479
2019/08/23 22:44:48 json: cannot unmarshal number into Go struct field PlaceholderRecord.lineage of type results.PlaceholderRecord
exit status 1

so for today we're just interface{} -ing all the thing
(20190823/thisisaaronland)

*/

// type FindByIDResults map[string]PlaceholderRecord
type FindByIDResults map[string]interface{}

// func (s *FindByIDResults) Results() []PlaceholderRecord {
func (s *FindByIDResults) Results() []interface{} {

	// results := make([]PlaceholderRecord, 0)
	results := make([]interface{}, 0)

	for _, v := range *s {
		results = append(results, v)
	}

	return results
}

func NewFindByIDResults(body []byte) (*FindByIDResults, error) {

	var id_results *FindByIDResults
	err := json.Unmarshal(body, &id_results)

	if err != nil {
		return nil, err
	}

	return id_results, nil
}
