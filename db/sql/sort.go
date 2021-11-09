package sql

import (
	"fmt"
	"strings"
)

func GetSortSql(q *Query) string {

	var sql string

	for _, sort := range q.Sorts {
		fmt.Printf("Sort: %v", sort)
		field := getSafeFieldName(sort.Field)
		sql += field + " " + strings.ToUpper(sort.Direction) + ", "
	}

	if len(sql) > 0 {
		sql = "ORDER BY " + strings.TrimSuffix(sql, ", ")
	}

	return sql
}
