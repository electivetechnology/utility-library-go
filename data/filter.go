package data

type Filter struct {
	Criterions []Criterion
	Logic      string
	Filters    map[string]*Filter
	Collation  bool
	Subquery   *Subquery
}

type Subquery struct {
	IsEnabled  bool
	Key        string
	Set        string
	IsDistinct bool
}

const (
	FILTER_LOGIC_INTERSECTION = "AND" // Logic Intersection (AND A AND B AND C ...)
	FILTER_LOGIC_UNION        = "OR"  // Logic Union (OR A OR B OR C ...)
)

func NewFilter() *Filter {
	filter := &Filter{Logic: FILTER_LOGIC_INTERSECTION, Collation: true, Subquery: NewSubquery()}

	return filter
}

func NewSubquery() *Subquery {
	subquery := &Subquery{IsEnabled: false, IsDistinct: false}

	return subquery
}
