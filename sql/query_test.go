package sql

import (
	"testing"
)

type TestExpandItem struct {
	statement string
	expected  string
	query     *Query
}

func TestExpand(t *testing.T) {
	statement := "SELECT * FROM example"
	q := &Query{Statement: statement}

	testData := []TestExpandItem{
		{statement, statement, q},
		{statement, statement + " LIMIT 0,1000", NewQuery(statement)},
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
