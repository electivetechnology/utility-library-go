package sql

import (
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
)

const DEFAULT_LIMIT = 1000
const DEFAULT_OFFSET = 0

type Query struct {
	Statement  string
	Filters    []*data.Filter
	Sorts      []string
	Parameters map[string]string
	Limit      int
	Offset     int
}

func (q *Query) Expand() (*Query, error) {
	sql := q.Statement

	// Build LIMIT clause
	sql += " " + GetLimitSql(q)

	// Build Filter clause
	sql += " " + GetFilterSql(q)

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
