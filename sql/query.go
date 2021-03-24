package sql

import "strings"

const DEFAULT_LIMIT = 1000
const DEFAULT_OFFSET = 0

type Query struct {
	Statement  string
	Parameters map[string]string
	Limit      int
	Offset     int
}

func (q *Query) Expand() (*Query, error) {
	sql := q.Statement

	// Build LIMIT clause
	sql += " " + GetLimitSql(q)

	// Set Query Statement
	q.Statement = strings.TrimSpace(sql)

	return q, nil
}

func NewQuery(statement string) *Query {
	query := &Query{
		Statement: statement,
		Limit:     DEFAULT_LIMIT,
		Offset:    DEFAULT_OFFSET,
	}

	return query
}
