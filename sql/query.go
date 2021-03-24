package sql

type Query struct {
	Statement  string
	Parameters map[string]string
}

func ExpandSimpleQuery(q *Query) (*Query, error) {

	return q, nil
}

func NewQuery(statement string) *Query {
	query := &Query{Statement: statement}

	return query
}
