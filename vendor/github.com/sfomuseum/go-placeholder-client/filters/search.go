package filters

import (
	"errors"
	"fmt"
	"strings"
)

type SearchFilter struct {
	Filter
	key   string
	value string
}

func (f *SearchFilter) Key() string {
	return f.key
}

func (f *SearchFilter) Value() string {
	return f.value
}

func NewSearchFilter(key string, value string) (Filter, error) {

	switch key {
	case "lang", "placetype", "mode":
		// pass
	default:
		msg := fmt.Sprintf("Invalid search filter '%s'", key)
		return nil, errors.New(msg)
	}

	sf := SearchFilter{
		key:   key,
		value: value,
	}

	return &sf, nil
}

type SearchFilters []Filter

func (f *SearchFilters) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *SearchFilters) Set(value string) error {

	value = strings.Trim(value, " ")
	kv := strings.Split(value, "=")

	if len(kv) != 2 {
		return errors.New("Invalid search filter")
	}

	sf, err := NewSearchFilter(kv[0], kv[1])

	if err != nil {
		return err
	}

	*f = append(*f, sf)
	return nil
}
