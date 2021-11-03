package sql

import "testing"

func TestPrepareNew(t *testing.T) {
	q := NewQuery("candidates")
	expected := "SELECT * FROM candidates"
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimple(t *testing.T) {
	q, _ := NewSimpleQuery("SELECT id, name FROM candidates")
	expected := "SELECT id, name FROM candidates"
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}
