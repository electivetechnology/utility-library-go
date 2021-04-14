package data

const (
	SORT_DIRECTION_ASC  = "asc"
	SORT_DIRECTION_DESC = "desc"
)

type Sort struct {
	Field     string
	Direction string
}

func NewSort() *Sort {
	sort := &Sort{Direction: SORT_DIRECTION_ASC}

	return sort
}
