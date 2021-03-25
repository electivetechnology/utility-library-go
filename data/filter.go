package data

type Filter struct {
	Criterions []Criterion
	Logic      string
	Filters    []*Filter
}

func NewFilter() *Filter {
	filter := &Filter{Logic: "and"}

	return filter
}
