package data

type Filter struct {
	Criterions []Criterion
	Logic      string
	Filters    []*Filter
	Collation  bool
}

func NewFilter() *Filter {
	filter := &Filter{Logic: "and", Collation: true}

	return filter
}
