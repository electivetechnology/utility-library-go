package sql

import "strings"

type Clause struct {
	Statement  string
	Parameters map[string]string
}

func (c Clause) GetSql() string {
	sql := c.Statement

	for key, value := range c.Parameters {
		sql = strings.ReplaceAll(sql, ":"+key, value)
	}

	return sql
}
