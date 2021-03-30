package sql

import (
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

type TestExpandItem struct {
	statement string
	expected  string
	query     *Query
}

func TestExpand(t *testing.T) {
	statement := "SELECT * FROM example"
	q := &Query{Statement: statement}
	q2 := NewQuery(statement)
	q2.Filters = append(q2.Filters, data.NewFilter())

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "organisation",
		Operand: "eq",
		Type:    "value",
		Value:   "fRnFPWTQyLMl",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "job_id",
		Operand: "eq",
		Type:    "value",
		Value:   "rb6RpwrSLbiV",
	}
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Filters = make(map[string]*data.Filter)
	f1.Collation = false

	q3 := NewQuery(statement)
	q3.Filters = append(q3.Filters, f1)

	testData := []TestExpandItem{
		{statement, statement, q},
		{statement, statement + " LIMIT 1000 OFFSET 0", NewQuery(statement)},
		{statement, statement + " LIMIT 1000 OFFSET 0", q2},
		{statement, "SELECT * FROM example WHERE (CAST(`organisation` AS STRING) =  CAST(:0_w_filter_0 AS STRING) AND CAST(`job_id` AS STRING) =  CAST(:0_w_filter_1 AS STRING)) LIMIT 1000 OFFSET 0", q3},
	}

	for _, item := range testData {
		ret, _ := item.query.Expand()
		if ret.Statement != item.expected {
			t.Errorf("Expand("+item.statement+") failed, expected %v, got %v", item.expected, ret.Statement)
		} else {
			t.Logf("Expand("+item.statement+") success, expected %v, got %v", item.expected, ret.Statement)
		}
	}
}

func TestNewQuery(t *testing.T) {
	statement := "SELECT * FROM example"
	ret := NewQuery(statement)

	if ret.Statement != statement {
		t.Errorf("NewQuery("+statement+") failed, expected %v, got %v", statement, ret.Statement)
	} else {
		t.Logf("NewQuery("+statement+") success, expected %v, got %v", statement, ret.Statement)
	}

	if ret.Limit != DEFAULT_LIMIT {
		t.Errorf("NewQuery("+statement+") failed, expected default limit %d, got %v", DEFAULT_LIMIT, ret.Limit)
	} else {
		t.Logf("NewQuery("+statement+") success, expected %d, got %v", DEFAULT_LIMIT, ret.Limit)
	}

	if ret.Offset != DEFAULT_OFFSET {
		t.Errorf("NewQuery("+statement+") failed, expected default offset %d, got %v", DEFAULT_OFFSET, ret.Limit)
	} else {
		t.Logf("NewQuery("+statement+") success, expected %d, got %v", DEFAULT_OFFSET, ret.Offset)
	}
}
