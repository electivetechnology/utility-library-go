package request

import (
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

type TestGetSortsFromQueryStringItem struct {
	query    string
	expected []*data.Sort
}

func TestGetSortsFromQueryString(t *testing.T) {
	s1 := data.Sort{
		Field:     "id",
		Direction: "asc",
	}

	s2 := data.Sort{
		Field:     "name",
		Direction: "desc",
	}

	var sorts1 []*data.Sort
	sorts1 = append(sorts1, &s1)
	sorts1 = append(sorts1, &s2)

	testData := []TestGetSortsFromQueryStringItem{
		{"sorts[]=id-asc&sorts[]=name-desc", sorts1},
	}

	for _, item := range testData {
		ret := GetSortsFromQueryString(item.query)

		if len(ret) != len(item.expected) {
			t.Errorf("GetSortsFromQueryString("+item.query+") failed, expected number of sorts %v, got %v", len(item.expected), len(ret))
		}
	}
}
