package request

type Filter struct {
	Criterions []Criterion
	Logic      string
}

func NewFilter() *Filter {
	filter := &Filter{Logic: "and"}

	return filter
}
