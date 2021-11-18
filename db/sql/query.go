package sql

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
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
	Sorts      map[string]data.Sort
	Filters    map[string]data.Filter
	Displays   map[string]data.Display
	FieldMap   map[string]string
}

func NewQuery(table string) Query {
	return Query{Flavour: QUERY_FLAVOUR_MYSQL, Limit: DEFAULT_LIMIT, Offset: DEFAULT_OFFSET, Table: table, Fields: []string{"*"}}
}

func NewSimpleQuery(query string) (Query, error) {
	var re = regexp.MustCompile(`(?mi)From \s*(?P<table>.*?)\s*( |$)`)
	matches := re.FindStringSubmatch(query)

	if matches == nil {
		msg := fmt.Sprintf(
			"Could not prepare simple query from: '%s'."+
				" Please check your syntax."+
				" The simple query should in format 'SELECT * FROM table'", query)
		log.Fatalf(msg)
		return Query{}, nil
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	if _, ok := result["table"]; !ok {
		msg := fmt.Sprintf(
			"Could not prepare simple query from: '%s'."+
				" Please check your syntax."+
				" The simple query should in format 'SELECT * FROM table'", query)
		log.Fatalf(msg)
		return Query{}, nil
	}

	q := NewQuery(result["table"])

	return q, nil
}

// GetSql returns ready to use SQL statement with parsed parameters
// This will internally run Prepare() first
func (q Query) GetSql() string {
	sql := q.GetStatement()
	for key, value := range q.Parameters {
		sql = strings.ReplaceAll(sql, ":"+key, `"`+value+`"`)
	}

	return sql
}

// GetStatement returns raw SQL statement with placeholders for parameters
func (q Query) GetStatement() string {
	return q.Statement
}

func (q *Query) Prepare() {
	var s string
	// Build Select clause from Displays
	selectClause := GetSelectSql(q)

	// Build SELECT
	s = selectClause.Statement

	// Build FROM
	s += " FROM " + q.Table

	// Build Filter clause
	filterClause := GetFilterSql(q)
	if len(filterClause.Statement) > 0 {
		s += " " + filterClause.Statement
	}

	// Add Filter Parameters
	q.Parameters = make(map[string]string)
	q.Parameters = filterClause.Parameters

	// Build Sort clause
	sortClause := GetSortSql(q)
	if len(sortClause) > 0 {
		s += " " + sortClause
	}

	// Build LIMIT and Offset clause
	s += " " + GetLimitSql(q)

	q.Statement = s
}

func GetSelectSql(q *Query) Clause {
	c := Clause{}

	// If there are no displays use q.Fields
	if len(q.Displays) == 0 {
		// Build SELECT
		c.Statement = strings.Join(q.Fields, ", ")
	} else {
		displayClause := DisplaysToSqlClause(q.Displays, q.FieldMap)
		c.Statement = displayClause.Statement
	}

	// Prepend with SELECT
	if len(c.Statement) > 0 {
		c.Statement = "SELECT " + c.Statement
	}

	return c
}
