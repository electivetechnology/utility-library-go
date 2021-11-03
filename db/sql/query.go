package sql

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	DEFAULT_LIMIT           = 1000
	DEFAULT_OFFSET          = 0
	QUERY_FLAVOUR_MYSQL     = "mysql"
	QUERY_FLAVOUR_BIG_QUERY = "bigquery"
)

type Query struct {
	Flavour    string
	Statement  string
	Limit      int
	Offset     int
	Fields     []string
	Table      string
	Parameters map[string]string
}

func NewQuery(table string) Query {
	return Query{Flavour: QUERY_FLAVOUR_MYSQL, Limit: DEFAULT_LIMIT, Offset: DEFAULT_OFFSET, Table: table, Fields: []string{"*"}}
}

func NewSimpleQuery(query string) (Query, error) {
	var re = regexp.MustCompile(`(?mi)from (\w+)`)
	from := re.FindAllString(query, -1)

	if from == nil {
		msg := fmt.Sprintf(
			"Could not prepare simple query from: '%s'."+
				" Please check your syntax."+
				" The simple query should in format 'SELECT expressions FROM table'", query)
		log.Fatalf(msg)
		return Query{}, nil
	}

	fmt.Printf("FROM %v\n", from[len(from)])

	q := NewQuery("m")

	return q, nil
}

// GetSql returns ready to use SQL statement with parsed parameters
func (q Query) GetSql() string {
	var sql string
	for key, value := range q.Parameters {
		sql = strings.ReplaceAll(q.Statement, ":"+key, `"`+value+`"`)
	}

	return sql
}

// GetStatement returns raw SQL statement with placeholders for parameters
func (q Query) GetStatement() string {
	return q.Statement
}

func (q *Query) Prepare() {
	var s string
	// Build SELECT
	s = "SELECT " + strings.Join(q.Fields, ", ")

	// Build FROM
	s += " FROM " + q.Table

	q.Statement = s

	fmt.Printf("SELECT: %s", s)
}
