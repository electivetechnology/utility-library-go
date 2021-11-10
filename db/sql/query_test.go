package sql

import (
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

func TestPrepareNew(t *testing.T) {
	q := NewQuery("candidates")
	expected := "SELECT * FROM candidates LIMIT 1000 OFFSET 0"
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimple(t *testing.T) {
	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates LIMIT 1000 OFFSET 0"
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimpleWithSort(t *testing.T) {
	// Prepare sorts
	sorts := make(map[string]data.Sort)

	// First Sort
	s1 := data.NewSort()
	s1.Field = "id"
	sorts["s_00"] = *s1

	// Second Sort
	s2 := data.NewSort()
	s2.Field = "email"
	s2.Direction = data.SORT_DIRECTION_DESC
	sorts["s_z"] = *s2

	// Third Sort
	s3 := data.NewSort()
	s3.Field = "first_name"
	sorts["s_a"] = *s3

	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates ORDER BY `id` ASC, `email` DESC, `first_name` ASC LIMIT 1000 OFFSET 0"

	// Add sorts
	q.Sorts = sorts

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimpleWithFilter(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "first_name",
		Operand: "eq",
		Type:    "value",
		Value:   "Kris",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "organisation",
		Operand: "eq",
		Type:    "value",
		Value:   "Ds7q0eBi2Iyy",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates WHERE (`first_name` =  CAST(:f_0_w_filter_0 AS CHAR) COLLATE utf8mb4_bin AND `organisation` =  CAST(:f_0_w_filter_1 AS CHAR) COLLATE utf8mb4_bin) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	//fmt.Printf("\nQuery Params: %v\n", q.Parameters)
	//fmt.Printf("\nSQL with Params: %s\n", q.GetSql())

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM candidates WHERE (`first_name` =  CAST(\"Kris\" AS CHAR) COLLATE utf8mb4_bin AND `organisation` =  CAST(\"Ds7q0eBi2Iyy\" AS CHAR) COLLATE utf8mb4_bin) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}
