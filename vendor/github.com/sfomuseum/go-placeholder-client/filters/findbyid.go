package filters

import (
	"errors"
	"fmt"
	"strings"
)

type FindByIdFilter struct {
	Filter
	key   string
	value string
}

func (f *FindByIdFilter) Key() string {
	return f.key
}

func (f *FindByIdFilter) Value() string {
	return f.value
}

func NewFindByIdFilter(key string, value string) (Filter, error) {

	switch key {
	case "lang":
		// pass
	default:
		msg := fmt.Sprintf("Invalid findbyid filter '%s'", key)
		return nil, errors.New(msg)
	}

	ff := FindByIdFilter{
		key:   key,
		value: value,
	}

	return &ff, nil
}

type FindByIdFilters []Filter

func (f *FindByIdFilters) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *FindByIdFilters) Set(value string) error {

	value = strings.Trim(value, " ")
	kv := strings.Split(value, "=")

	if len(kv) != 2 {
		return errors.New("Invalid search filter")
	}

	ff, err := NewFindByIdFilter(kv[0], kv[1])

	if err != nil {
		return err
	}

	*f = append(*f, ff)
	return nil
}
