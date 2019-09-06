package results

import (
	"encoding/json"
)

type FindByIDResults map[string]*PlaceholderRecord

type FindByIDResultsRaw map[string]*FindByIdRecord

// FindByIdRecord differs from PlaceholderRecord in that it has a subtly different Lineage
// property - we use the handy AsPlaceholderRecord() method to make these results look the
// "same" as search results (20190905/thisisaaronland)

type FindByIdRecord struct {
	Id                int64               `json:"id"`
	Name              string              `json:"name"`
	Placetype         string              `json:"placetype"`
	Abbreviation      string              `json:"abbr,omitempty"`
	Rank              Rank                `json:"rank,omitempty"`
	Population        int                 `json:"population,omitempty"`
	LanguageDefaulted bool                `json:"languageDefaulted"`
	Lineage           []map[string]int64  `json:"lineage,omitempty"`
	Names             map[string][]string `json:"names,omitempty"`
	Geometry          Geometry            `json:"geom,omitempty"`
}

func (r FindByIdRecord) AsPlaceholderRecord() *PlaceholderRecord {

	ph_lineage := make([]map[string]*PlaceholderRecord, 0)

	for _, l := range r.Lineage {

		ph_l := make(map[string]*PlaceholderRecord)

		for k, id := range l {
			ph_l[k] = &PlaceholderRecord{
				Id: id,
			}
		}

		ph_lineage = append(ph_lineage, ph_l)
	}

	ph_record := &PlaceholderRecord{
		Id:                r.Id,
		Name:              r.Name,
		Placetype:         r.Placetype,
		Abbreviation:      r.Abbreviation,
		Rank:              r.Rank,
		Population:        r.Population,
		LanguageDefaulted: r.LanguageDefaulted,
		Names:             r.Names,
		Geometry:          r.Geometry,
		Lineage:           ph_lineage,
	}

	return ph_record
}

func (s *FindByIDResults) Results() []*PlaceholderRecord {

	results := make([]*PlaceholderRecord, 0)

	for _, r := range *s {
		results = append(results, r)
	}

	return results
}

func NewFindByIDResults(body []byte) (*FindByIDResults, error) {

	var id_results_raw *FindByIDResultsRaw
	err := json.Unmarshal(body, &id_results_raw)

	if err != nil {
		return nil, err
	}

	id_results_map := make(map[string]*PlaceholderRecord)

	for k, r := range *id_results_raw {
		id_results_map[k] = r.AsPlaceholderRecord()
	}

	// the following hoop-jumping is to account for these errors - if there's a way around them I don't know what it is...
	// results/findbyid.go:81:14: invalid operation: id_results[k] (type *FindByIDResults does not support indexing)

	enc, err := json.Marshal(id_results_map)

	if err != nil {
		return nil, err
	}

	var id_results *FindByIDResults
	err = json.Unmarshal(enc, &id_results)

	if err != nil {
		return nil, err
	}

	return id_results, nil
}
