package sql

import (
	"fmt"
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

func TestGetFilterSql(t *testing.T) {
	statement := "SELECT * FROM example"
	filter := data.NewFilter()

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "gt",
		Type:    "value",
		Value:   "1",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "ge",
		Type:    "value",
		Value:   "2",
	}

	filter.Criterions = append(filter.Criterions, c1)
	filter.Criterions = append(filter.Criterions, c2)

	q := NewQuery(statement)
	q.Filters = append(q.Filters, filter)
	t.Logf("Query Filters are %v", q.Filters)

	ret := GetFilterSql(q)
	fmt.Printf("Ret is %v", ret)
}
