package sql

import (
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
)

func FiltersToSqlClause(filters map[string]*data.Filter) Clause {
	c := Clause{}
	c.Parameters = make(map[string]string)

	// Iterate over filters and turn each filter to SQL Clause
	for i, filter := range filters {
		// Turn each filter into SQL Clause
		clause := FilterToSqlClause(filter, i+"_filter")

		// Add filter Logic
		if len(c.Statement) != 0 {
			c.Statement += " " + strings.ToUpper(filter.Logic) + " "
		}

		// Process filter subquery
		if filter.Subquery.IsEnabled && len(filter.Subquery.Key) > 0 && len(filter.Subquery.Set) > 0 {
			// Append SQL Statement
			c.Statement += getSafeFieldName(filter.Subquery.Key) +
				" IN (SELECT " +
				getSafeFieldName(filter.Subquery.Key) +
				" FROM " +
				getSafeTableName(filter.Subquery.Set) +
				" WHERE " +
				clause.Statement + ")"
		} else {
			// Append SQL Statement
			if len(clause.Statement) > 0 {
				c.Statement += "(" + clause.Statement + ")"
			}
		}

		// Add Parametes
		for key, parameter := range clause.Parameters {
			c.Parameters[key] = parameter
		}
	}

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

	if len(filter.Filters) > 0 {
		clause := FiltersToSqlClause(filter.Filters)
		c.Statement += " AND " + clause.Statement

		// Add Parametes
		for key, parameter := range clause.Parameters {
			c.Parameters[key] = parameter
		}
	}

	return c
}
