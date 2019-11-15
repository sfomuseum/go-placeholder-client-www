package filters

type Filter interface {
	Key() string
	Value() string
}
