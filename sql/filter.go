package sql

import (
	"fmt"
	"strconv"

	"github.com/electivetechnology/utility-library-go/data"
)

func FiltersToSqlClause(filters map[string]*data.Filter) Clause {
	c := Clause{}

	// Iterate over filters and turn each filter to SQL Clause
	for i, filter := range filters {
		fmt.Printf("Filter name: %s\n", i)
		fmt.Printf("Filter value: %v\n", filter)

		// Turn each filter into SQL Clause
		clause := FilterToSqlClause(filter, i+"_filter")

		// Append SQL Statement
		c.Statement += clause.Statement

		// Add Parametes
		for key, parameter := range clause.Parameters {
			c.Parameters[key] = parameter
		}
	}

	fmt.Printf("SQL for filters: %s\n", c.Statement)

	return c
}

func FilterToSqlClause(filter *data.Filter, namespace string) Clause {
	c := Clause{}

	for i, criterion := range filter.Criterions {
		// Placeholder name for query binding
		placeHolder := namespace + "_" + strconv.Itoa(i)

		// Turn each Criterion into Clause
		clause := CriterionToSqlClause(criterion, placeHolder)

		// Append SQL Statement
		c.Statement += clause.Statement

		// Add Parametes
		for key, parameter := range clause.Parameters {
			c.Parameters[key] = parameter
		}
	}

	return c
}
