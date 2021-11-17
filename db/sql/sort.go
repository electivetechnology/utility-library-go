package sql

import (
	"sort"
	"strings"
)

func GetSortSql(q *Query) string {

	var sql string

	// Sort Keys
	keys := make([]string, 0)
	for k, _ := range q.Sorts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {

		field := getSafeFieldName(q.Sorts[k].Field, q.FieldMap)
		sql += field + " " + strings.ToUpper(q.Sorts[k].Direction) + ", "
	}

	if len(sql) > 0 {
		sql = "ORDER BY " + strings.TrimSuffix(sql, ", ")
	}

	return sql
}
