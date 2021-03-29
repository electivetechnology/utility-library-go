package sql

import (
	"fmt"
	"strconv"

	"github.com/electivetechnology/utility-library-go/data"
)

func FiltersToSqlClause(filters map[string]*data.Filter) Clause {
	c := Clause{}
	c.Parameters = make(map[string]string)

	// Iterate over filters and turn each filter to SQL Clause
	for i, filter := range filters {
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
	c.Parameters = make(map[string]string)

	for i, criterion := range filter.Criterions {
		// Placeholder name for query binding
		placeHolder := namespace + "_" + strconv.Itoa(i)

		// Turn each Criterion into Clause
		clause := CriterionToSqlClause(criterion, placeHolder, filter.Collation)

		// Add Parametes
		for key, parameter := range clause.Parameters {
			c.Parameters[key] = parameter
		}

		// Remove Logic from first Statement
		// and Append SQL Statement
		if len(c.Statement) == 0 {
			c.Statement += clause.removeLogicFromStatement().Statement
		} else {
			c.Statement += " " + clause.Statement
		}
	}

	fmt.Printf("Filter SQL: %s", c.GetSql())

	return c
}
