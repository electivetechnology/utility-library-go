package sql

import "testing"

func TestExpandSimpleQuery(t *testing.T) {
	statement := "SELECT * FROM example"
	q := &Query{Statement: statement}
	ret, _ := ExpandSimpleQuery(q)

	if ret.Statement != statement {
		t.Errorf("ExpandSimpleQuery("+statement+") failed, expected %v, got %v", statement, ret.Statement)
	} else {
		t.Logf("ExpandSimpleQuery("+statement+") success, expected %v, got %v", statement, ret.Statement)
	}
}

func TestNewQuery(t *testing.T) {
	statement := "SELECT * FROM example"
	ret := NewQuery(statement)

	if ret.Statement != statement {
		t.Errorf("ExpandSimpleQuery("+statement+") failed, expected %v, got %v", statement, ret.Statement)
	} else {
		t.Logf("ExpandSimpleQuery("+statement+") success, expected %v, got %v", statement, ret.Statement)
	}
}
