package sql

import (
	"strconv"

	"github.com/electivetechnology/utility-library-go/data"
)

func GetFilterSql(q *Query) Clause {
	c := Clause{}
	c.Parameters = make(map[string]string)

	whereFilters := make(map[string]*data.Filter)

	for i, filter := range q.Filters {
		// @todo, iterate over criterions and compare keys with FieldMap
		// als check for HAVING filter
		whereFilters[strconv.Itoa(i)+"_w"] = filter
	}

	clause := FiltersToSqlClause(whereFilters)

	// Copy parameters
	c.Parameters = clause.Parameters

	if len(clause.Statement) > 0 {
		c.Statement = "WHERE " + clause.Statement
	}

	return c
}
