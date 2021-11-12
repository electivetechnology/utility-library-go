package sql

import "testing"

type TestGetSafeFieldNameDataItem struct {
	Field    string
	Expected string
}

func TestGetSafeFieldName(t *testing.T) {
	fieldMap := make(map[string]string)
	fieldMap["*"] = "candidates"
	fieldMap["`*`"] = "candidates"
	fieldMap["`id`"] = "assessments"
	fieldMap["id"] = "assessments"
	fieldMap["candidate.id"] = "candidates"
	fieldMap["candidates.id"] = "candidates"
	fieldMap["organisation"] = "connect-f7e5b.staging_reporting.candidates"

	testData := []TestGetSafeFieldNameDataItem{
		{Field: "id", Expected: "`assessments`.`id`"},
		{Field: "*", Expected: "`candidates`.*"},
		{Field: "`id`", Expected: "`assessments`.`id`"},
		{Field: "candidate.id", Expected: "`candidates`.`id`"},
		{Field: "candidates.id", Expected: "`candidates`.`id`"},
		{Field: "organisation", Expected: "`connect-f7e5b.staging_reporting.candidates`.`organisation`"},
	}

	for i, item := range testData {
		ret := getSafeFieldName(item.Field, fieldMap)

		if ret != item.Expected {
			t.Errorf("getSafeFieldName() failed for data series %d, expected %v, got %v", i, item.Expected, ret)
		}
	}
}
