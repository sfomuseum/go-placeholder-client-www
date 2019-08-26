package http

import (
	"github.com/sfomuseum/go-placeholder-client/results"
)

func Ancestors(r *results.PlaceholderRecord, placetype string) []*results.PlaceholderRecord {

	ancestors := make([]*results.PlaceholderRecord, 0)
	candidates := make(map[int64]*results.PlaceholderRecord)

	if placetype != r.Placetype {
	for _, l := range r.Lineage {

		a, ok := l[placetype]

		if ok {
			candidates[a.Id] = &a
		}
	}

	for _, a := range candidates {
		ancestors = append(ancestors, a)
	}
	}
	
	return ancestors
}
