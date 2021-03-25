package sql

import (
	"fmt"
	"strconv"

	"github.com/electivetechnology/utility-library-go/data"
)

func GetFilterSql(q *Query) string {
	ret := ""
	whereFilters := make(map[string]*data.Filter)

	for i, filter := range q.Filters {
		fmt.Printf("Filter idx %v for filter %v", i, filter)
		// @todo, iterate over criterions and compare keys with FieldMap
		// als check for HAVING filter
		whereFilters[strconv.Itoa(i)+"_w"] = filter
	}

	fmt.Printf("Filters %v", whereFilters)

	whereClause := FiltersToSqlClause(whereFilters)

	fmt.Printf("Where Clause %v", whereClause)

	return ret
}
