package sql

import (
	"fmt"
	"strconv"

	"github.com/electivetechnology/utility-library-go/data"
)

func GetFilterSql(q *Query) Clause {
	c := Clause{}
	var collation bool
	c.Parameters = make(map[string]string)

	whereFilters := make(map[string]*data.Filter)

	for i, filter := range q.Filters {
		// @todo, iterate over criterions and compare keys with FieldMap

		// Set filter collation based on query flavour
		fmt.Printf("QUERY flavour is %s \n", q.Flavour)
		if q.Flavour == QUERY_FLAVOUR_MYSQL {
			collation = true
		} else if q.Flavour == QUERY_FLAVOUR_BIG_QUERY {
			collation = false
		}

		filter = OverrideCollation(filter, collation)

		fmt.Printf("Filter for query is: %v\n", filter)

		// also check for HAVING filter
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

func OverrideCollation(filter *data.Filter, collation bool) *data.Filter {
	if len(filter.Filters) > 0 {
		for _, f := range filter.Filters {
			f = OverrideCollation(f, collation)
		}
	}

	filter.Collation = collation

	return filter
}
