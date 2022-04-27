package connect

type ApiQuery struct {
	Limit  int
	Offset int
}

func NewApiQuery() ApiQuery {
	return ApiQuery{
		Offset: 0,
		Limit:  0,
	}
}

func (q ApiQuery) GetLimit() int {
	return q.Limit
}

func (q *ApiQuery) SetLimit(limit int) *ApiQuery {
	q.Limit = limit

	return q
}

func (q ApiQuery) GetOffset() int {
	return q.Offset
}

func (q *ApiQuery) SetOffset(offset int) *ApiQuery {
	q.Offset = offset

	return q
}
