package results

import (
	"fmt"
)

type Geometry struct {
	BoundingBox string  `json:"bbox"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
}

type Rank struct {
	Min uint `json:"min"`
	Max uint `json:"max"`
}

type PlaceholderRecord struct {
	Id                int64                           `json:"id"`
	Name              string                          `json:"name"`
	Abbreviation      string                          `json:"abbr,omitempty"`
	Placetype         string                          `json:"placetype"`
	Rank              Rank                            `json:"rank,omitempty"`
	Population        int                             `json:"population,omitempty"`
	LanguageDefaulted bool                            `json:"languageDefaulted"`
	Lineage           []map[string]*PlaceholderRecord `json:"lineage,omitempty"`
	Names             map[string][]string             `json:"names,omitempty"`
	Geometry          Geometry                        `json:"geom,omitempty"`
}

func (r *PlaceholderRecord) String() string {
	return fmt.Sprintf("%s, %s (%d)", r.Name, r.Placetype, r.Id)
}
