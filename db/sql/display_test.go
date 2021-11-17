package sql

import (
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

type TestDisplaysToSqlClauseDataItem struct {
	Displays map[string]data.Display
	Expected string
}

func TestDisplaysToSqlClause(t *testing.T) {
	fieldMap := make(map[string]string)
	fieldMap["*"] = "candidates"
	fieldMap["`*`"] = "candidates"
	fieldMap["`id`"] = "assessments"
	fieldMap["id"] = "assessments"
	fieldMap["candidate.id"] = "candidates"
	fieldMap["candidates.id"] = "candidates"
	fieldMap["organisation"] = "connect-f7e5b.staging_reporting.candidates"
	fieldMap["candidate.firstName"] = "candidates"

	d1 := data.NewDisplay()
	d1.Field = "id"

	d2 := data.NewDisplay()
	d2.Field = "candidate.id"

	d3 := data.NewDisplay()
	d3.Field = "candidate.firstName"
	d3.Alias = "fName"

	ds1 := make(map[string]data.Display)
	ds1["d1"] = *d1

	ds2 := make(map[string]data.Display)
	ds2["d1"] = *d1
	ds2["d2"] = *d2

	ds3 := make(map[string]data.Display)
	ds3["d1"] = *d1
	ds3["d2"] = *d2
	ds3["d3"] = *d3

	testData := []TestDisplaysToSqlClauseDataItem{
		{Displays: ds1, Expected: "`assessments`.`id`"},
		{Displays: ds2, Expected: "`assessments`.`id`, `candidates`.`id`"},
		{Displays: ds3, Expected: "`assessments`.`id`, `candidates`.`firstName` AS fName, `candidates`.`id`"},
	}

	for i, item := range testData {
		ret := DisplaysToSqlClause(item.Displays, fieldMap)

		if ret.Statement != item.Expected {
			t.Errorf("DisplaysToSqlClause() failed for data series %d, expected %s, got %s", i, item.Expected, ret.Statement)
		}
	}
}
