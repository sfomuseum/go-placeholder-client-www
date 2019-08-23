package filters

import (
	"errors"
	"fmt"
	"strings"
)

type SearchFilter struct {
	Key   string
	Value string
}

type SearchFilters []*SearchFilter

func (f *SearchFilters) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *SearchFilters) Set(value string) error {

	value = strings.Trim(value, " ")
	kv := strings.Split(value, "=")

	if len(kv) != 2 {
		return errors.New("Invalid search filter")
	}

	// validate key here...

	sf := SearchFilter{
		Key:   kv[0],
		Value: kv[1],
	}

	*f = append(*f, &sf)
	return nil
}
