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

	// Build Filter clause
	filterClause := GetFilterSql(q)
	if len(filterClause.Statement) > 0 {
		sql += " " + filterClause.Statement
	}

	// Add Filter Parameters
	q.Parameters = make(map[string]string)
	q.Parameters = filterClause.Parameters

	// Build LIMIT clause
	sql += " " + GetLimitSql(q)

	// Set Query Statement
	q.Statement = strings.TrimSpace(sql)

	return q, nil
}

func (q Query) GetSql() string {
	sql := q.Statement

	for key, value := range q.Parameters {
		sql = strings.ReplaceAll(sql, ":"+key, `"`+value+`"`)
	}

	return sql
}

func NewQuery(statement string) *Query {
	query := &Query{
		Statement: statement,
		Limit:     DEFAULT_LIMIT,
		Offset:    DEFAULT_OFFSET,
	}

	return query
}
