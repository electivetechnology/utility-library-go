package sql

import (
	"fmt"
	"strings"
)

type Clause struct {
	Statement  string
	Parameters map[string]string
}

func (c Clause) GetSql() string {
	sql := c.Statement

	for key, value := range c.Parameters {
		fmt.Printf("C has parameter: %s, %v\n", key, value)
		sql = strings.ReplaceAll(sql, ":"+key, `"`+value+`"`)
	}

	return sql
}
