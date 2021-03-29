package data

type Filter struct {
	Criterions []Criterion
	Logic      string
	Filters    map[string]*Filter
	Collation  bool
	Subquery   *Subquery
}

type Subquery struct {
	IsEnabled bool
	Key       string
	Set       string
}

func NewFilter() *Filter {
	filter := &Filter{Logic: "and", Collation: true, Subquery: NewSubquery()}

	return filter
}

func NewSubquery() *Subquery {
	subquery := &Subquery{IsEnabled: false}

	return subquery
}
