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
}

func NewQuery(table string) Query {
	return Query{Flavour: QUERY_FLAVOUR_MYSQL, Limit: DEFAULT_LIMIT, Offset: DEFAULT_OFFSET, Table: table, Fields: []string{"*"}}
}

func NewSimpleQuery(query string) (Query, error) {
	var re = regexp.MustCompile(`(?P<select>SELECT )((\W)) (?P<from>FROM )(?P<table>\w*)`)
	match := re.FindStringSubmatch(query)

	if match == nil {
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
			result[name] = match[i]
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
func (q Query) GetSql() string {
	sql := q.Statement
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
	// Build SELECT
	s = "SELECT " + strings.Join(q.Fields, ", ")

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

	//fmt.Printf("SELECT: %s", s)
}
