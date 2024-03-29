package sql

import (
	"sort"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
)

func GetFilterSql(q *Query) Clause {
	c := Clause{}
	var collation bool
	c.Parameters = make(map[string]string)

	whereFilters := make(map[string]data.Filter)
	havingFilters := make(map[string]data.Filter)

	for i, filter := range q.Filters {
		// Only add filter if there are criterions
		if len(filter.Criterions) > 0 {
			// Set filter collation based on query flavour
			if q.Flavour == QUERY_FLAVOUR_MYSQL {
				collation = true
			} else if q.Flavour == QUERY_FLAVOUR_BIG_QUERY {
				collation = false
			}

			modifiedFilter := OverrideCollation(filter, collation, q.FieldMap)

			whereFilters[i+"_w"] = GetWhereFilters(modifiedFilter, q.FieldMap)
			// also check for HAVING filter
			havingFilters[i+"_w"] = GetHavingFilters(modifiedFilter, q.FieldMap)
		}
	}

	whereClause := FiltersToSqlClause(whereFilters, q.FieldMap)
	havingClause := FiltersToSqlClause(havingFilters, q.FieldMap)

	// Copy parameters
	c.Parameters = whereClause.Parameters
	for k, v := range havingClause.Parameters {
		c.Parameters[k] = v
	}

	if len(whereClause.Statement) > 0 {
		// Remove whitespace
		whereClause.Statement = strings.TrimLeft(whereClause.Statement, " ")

		// Trim first AND|OR and prepend with WHERE
		c.Statement = "WHERE " + whereClause.removeLogicFromStatement().Statement
	}

	if len(havingClause.Statement) > 0 {
		// Remove whitespace
		havingClause.Statement = strings.TrimLeft(havingClause.Statement, " ")

		// Trim first AND|OR and prepend with WHERE

		c.Statement += " HAVING " + havingClause.removeLogicFromStatement().Statement
	}

	return c
}

func FiltersToSqlClause(filters map[string]data.Filter, fieldMap map[string]string) Clause {
	c := Clause{}
	c.Parameters = make(map[string]string)

	// Sort Keys to preserve order of the filters (important)
	keys := make([]string, 0)
	for k, _ := range filters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Iterate over filters and turn each filter to SQL Clause
	for _, k := range keys {
		filter := filters[k]
		// Turn each filter into SQL Clause
		clause := FilterToSqlClause(filter, fieldMap, k+"_filter")

		// Process filter subquery
		if filter.Subquery.IsEnabled && len(filter.Subquery.Key) > 0 && len(filter.Subquery.Set) > 0 {
			// Append SQL Statement
			c.Statement += getSafeFieldName(filter.Subquery.Key, fieldMap) +
				" IN (SELECT " +
				getSafeFieldName(filter.Subquery.Key, fieldMap) +
				" FROM " +
				getSafeTableName(filter.Subquery.Set, fieldMap) +
				" WHERE " +
				clause.Statement + ")"
		} else {
			// Append SQL Statement
			c.Statement += clause.Statement
		}

		// Add Parametes
		for key, parameter := range clause.Parameters {
			c.Parameters[key] = parameter
		}
	}

	return c
}

func FilterToSqlClause(filter data.Filter, fieldMap map[string]string, namespace string) Clause {
	c := Clause{}
	c.Parameters = make(map[string]string)

	// Add logic
	c.Statement += " " + strings.ToUpper(filter.Logic) + " ("

	for i, criterion := range filter.Criterions {
		// Placeholder name for query binding
		placeHolder := namespace + "_" + strconv.Itoa(i)

		// Turn each Criterion into Clause
		clause := CriterionToSqlClause(criterion, placeHolder, fieldMap, filter.Collation)

		// Add Parametes
		for key, parameter := range clause.Parameters {
			c.Parameters[key] = parameter
		}

		// Remove Logic from first Statement
		// and Append SQL Statement
		if i == 0 {
			c.Statement += clause.removeLogicFromStatement().Statement
		} else {
			c.Statement += " " + clause.Statement
		}
	}

	if len(filter.Filters) > 0 {
		filters := make(map[string]data.Filter)
		for i, f := range filter.Filters {
			filters[i] = *f
		}

		clause := FiltersToSqlClause(filters, fieldMap)
		c.Statement += clause.Statement

		// Add Parametes
		for key, parameter := range clause.Parameters {
			c.Parameters[key] = parameter
		}
	}

	// Close logic
	c.Statement += ")"

	// Remove empty statement
	if c.Statement == " AND ()" {
		c.Statement = ""
	}

	return c
}

func OverrideCollation(filter data.Filter, collation bool, fieldMap map[string]string) data.Filter {
	if len(filter.Filters) > 0 {
		for _, f := range filter.Filters {
			modified := OverrideCollation(*f, collation, fieldMap)
			f = &modified
		}
	}

	filter.Collation = collation

	return filter
}

func GetWhereFilters(filter data.Filter, fieldMap map[string]string) data.Filter {
	if len(filter.Filters) > 0 {
		for _, f := range filter.Filters {
			modified := GetWhereFilters(*f, fieldMap)
			f = &modified
		}
	}

	criterions := make([]data.Criterion, 0)
	for _, criterion := range filter.Criterions {
		if strings.ToLower(fieldMap[criterion.Key]) != "having" {
			criterions = append(criterions, criterion)
		}
	}

	filter.Criterions = criterions

	return filter
}

func GetHavingFilters(filter data.Filter, fieldMap map[string]string) data.Filter {
	if len(filter.Filters) > 0 {
		for _, f := range filter.Filters {
			modified := GetWhereFilters(*f, fieldMap)
			f = &modified
		}
	}

	criterions := make([]data.Criterion, 0)
	for _, criterion := range filter.Criterions {
		if strings.ToLower(fieldMap[criterion.Key]) == "having" {
			criterions = append(criterions, criterion)
		}
	}

	filter.Criterions = criterions

	return filter
}
