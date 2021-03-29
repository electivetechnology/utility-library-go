package sql

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

type TestFilterToSqlClauseItem struct {
	filter   *data.Filter
	expected string
}

func TestFilterToSqlClause(t *testing.T) {
	c1 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "eq",
		Type:    "value",
		Value:   "Hello",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "eqi",
		Type:    "value",
		Value:   "World",
	}

	f1 := data.NewFilter()

	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)

	f2 := data.NewFilter()
	f2.Collation = false

	f2.Criterions = append(f2.Criterions, c1)
	f2.Criterions = append(f2.Criterions, c2)

	testData := []TestFilterToSqlClauseItem{
		{f1, "`id` =  CAST(:f0_0 AS CHAR) COLLATE utf8mb4_bin AND `id` =  CAST(:f0_1 AS CHAR) COLLATE utf8mb4_general_ci"},
		{f2, "CAST(`id` AS STRING) =  CAST(:f1_0 AS STRING) AND LOWER(CAST(`id` AS STRING )) =  LOWER(CAST(:f1_1 AS STRING))"},
	}

	for i, item := range testData {

		clause := FilterToSqlClause(item.filter, "f"+strconv.Itoa(i))

		fmt.Printf("SQL:>>%v<<\n", clause.GetSql())
		if clause.Statement != item.expected {
			t.Errorf("FilterToSqlClause() failed, expected %v, got %v", item.expected, clause.Statement)
		} else {
			t.Logf("FilterToSqlClause() success, expected %v, got %v", item.expected, clause.Statement)
		}
	}
}
