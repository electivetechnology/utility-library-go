package sql

import (
	"strings"
)

func GetSortSql(q *Query) string {

	var sql string

	for _, sort := range q.Sorts {
		field := getSafeFieldName(sort.Field, q.FieldMap)
		sql += field + " " + strings.ToUpper(sort.Direction) + ", "
	}

	if len(sql) > 0 {
		sql = "ORDER BY " + strings.TrimSuffix(sql, ", ")
	}

	return sql
}
