package sql

import (
	"strconv"
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

type TestFilterToSqlClauseItem struct {
	filter   *data.Filter
	expected string
}

type TestFiltersToSqlClauseItem struct {
	filters  map[string]*data.Filter
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

		if clause.Statement != item.expected {
			t.Errorf("FilterToSqlClause() failed, expected %v, got %v", item.expected, clause.Statement)
		} else {
			t.Logf("FilterToSqlClause() success, expected %v, got %v", item.expected, clause.Statement)
		}
	}
}

func TestFiltersToSqlClause(t *testing.T) {
	c1 := data.Criterion{
		Logic:   "or",
		Key:     "id",
		Operand: "eq",
		Type:    "value",
		Value:   "1",
	}

	c2 := data.Criterion{
		Logic:   "or",
		Key:     "id",
		Operand: "eq",
		Type:    "value",
		Value:   "2",
	}

	c3 := data.Criterion{
		Logic:   "or",
		Key:     "id",
		Operand: "eq",
		Type:    "value",
		Value:   "3",
	}

	c4 := data.Criterion{
		Logic:   "or",
		Key:     "id",
		Operand: "eq",
		Type:    "value",
		Value:   "4",
	}

	f1 := data.NewFilter()

	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)

	f2 := data.NewFilter()

	f2.Criterions = append(f2.Criterions, c3)
	f2.Criterions = append(f2.Criterions, c4)

	var filters1 = make(map[string]*data.Filter)
	filters1["first"] = f1
	filters1["second"] = f2

	testData := []TestFiltersToSqlClauseItem{
		{filters1, "(`id` =  CAST(:first_filter_0 AS CHAR) COLLATE utf8mb4_bin OR `id` =  CAST(:first_filter_1 AS CHAR) COLLATE utf8mb4_bin) AND (`id` =  CAST(:second_filter_0 AS CHAR) COLLATE utf8mb4_bin OR `id` =  CAST(:second_filter_1 AS CHAR) COLLATE utf8mb4_bin)"},
	}

	for _, item := range testData {
		clause := FiltersToSqlClause(item.filters)

		if clause.Statement != item.expected {
			t.Errorf("FiltersToSqlClause() failed, expected %v, got %v", item.expected, clause.Statement)
		} else {
			t.Logf("FiltersToSqlClause() success, expected %v, got %v", item.expected, clause.Statement)
		}
	}
}

func TestFiltersToSqlClauseNestedFilters(t *testing.T) {
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

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "question_id",
		Operand: "eq",
		Type:    "value",
		Value:   "N1Romm5N1wgq",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "intent",
		Operand: "eq",
		Type:    "value",
		Value:   "no_intent",
	}

	c5 := data.Criterion{
		Logic:   "and",
		Key:     "entity",
		Operand: "eq",
		Type:    "value",
		Value:   "Duration",
	}

	c6 := data.Criterion{
		Logic:   "or",
		Key:     "keyword",
		Operand: "eq",
		Type:    "value",
		Value:   "a day",
	}

	c7 := data.Criterion{
		Logic:   "or",
		Key:     "keyword",
		Operand: "eq",
		Type:    "value",
		Value:   "1 year",
	}

	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Filters = make(map[string]*data.Filter)

	f2 := data.NewFilter()
	f2.Filters = make(map[string]*data.Filter)
	f2.Criterions = append(f2.Criterions, c3)
	f2.Criterions = append(f2.Criterions, c4)

	f3 := data.NewFilter()
	f3.Criterions = append(f3.Criterions, c5)
	f3.Filters = make(map[string]*data.Filter)

	f4 := data.NewFilter()
	f4.Criterions = append(f4.Criterions, c6)
	f4.Criterions = append(f4.Criterions, c7)

	f3.Filters["4"] = f4
	f2.Filters["3"] = f3
	f1.Filters["2"] = f2

	var filters = make(map[string]*data.Filter)

	filters["1"] = f1

	clause := FiltersToSqlClause(filters)
	expected := "(`organisation` =  CAST(:1_filter_0 AS CHAR) COLLATE utf8mb4_bin AND `job_id` =  CAST(:1_filter_1 AS CHAR) COLLATE utf8mb4_bin AND (`question_id` =  CAST(:2_filter_0 AS CHAR) COLLATE utf8mb4_bin AND `intent` =  CAST(:2_filter_1 AS CHAR) COLLATE utf8mb4_bin AND (`entity` =  CAST(:3_filter_0 AS CHAR) COLLATE utf8mb4_bin AND (`keyword` =  CAST(:4_filter_0 AS CHAR) COLLATE utf8mb4_bin OR `keyword` =  CAST(:4_filter_1 AS CHAR) COLLATE utf8mb4_bin))))"

	if clause.Statement != expected {
		t.Errorf("FiltersToSqlClause() failed, expected %v, got %v", expected, clause.Statement)
	} else {
		t.Logf("FiltersToSqlClause() success, expected %v, got %v", expected, clause.Statement)
	}
}

func TestFiltersToSqlClauseNestedFiltersWithSubquery(t *testing.T) {
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

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "question_id",
		Operand: "eq",
		Type:    "value",
		Value:   "N1Romm5N1wgq",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "intent",
		Operand: "eq",
		Type:    "value",
		Value:   "no_intent",
	}

	c5 := data.Criterion{
		Logic:   "and",
		Key:     "entity",
		Operand: "eq",
		Type:    "value",
		Value:   "Duration",
	}

	c6 := data.Criterion{
		Logic:   "or",
		Key:     "keyword",
		Operand: "eq",
		Type:    "value",
		Value:   "a day",
	}

	c7 := data.Criterion{
		Logic:   "or",
		Key:     "keyword",
		Operand: "eq",
		Type:    "value",
		Value:   "1 year",
	}

	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Filters = make(map[string]*data.Filter)
	f1.Collation = false

	f2 := data.NewFilter()
	f2.Collation = false
	f2.Subquery.IsEnabled = true
	f2.Subquery.Key = "engagement_id"
	f2.Subquery.Set = "connect-f7e5b.staging_reporting.transcripts"
	f2.Filters = make(map[string]*data.Filter)
	f2.Criterions = append(f2.Criterions, c3)
	f2.Criterions = append(f2.Criterions, c4)

	f3 := data.NewFilter()
	f3.Collation = false
	f3.Criterions = append(f3.Criterions, c5)
	f3.Filters = make(map[string]*data.Filter)

	f4 := data.NewFilter()
	f4.Collation = false
	f4.Criterions = append(f4.Criterions, c6)
	f4.Criterions = append(f4.Criterions, c7)

	f3.Filters["4"] = f4
	f2.Filters["3"] = f3
	f1.Filters["2"] = f2

	var filters = make(map[string]*data.Filter)

	filters["1"] = f1

	clause := FiltersToSqlClause(filters)
	expected := "(CAST(`organisation` AS STRING) =  CAST(:1_filter_0 AS STRING) AND CAST(`job_id` AS STRING) =  CAST(:1_filter_1 AS STRING) AND `engagement_id` IN (SELECT `engagement_id` FROM `connect-f7e5b`.`staging_reporting`.`transcripts` WHERE CAST(`question_id` AS STRING) =  CAST(:2_filter_0 AS STRING) AND CAST(`intent` AS STRING) =  CAST(:2_filter_1 AS STRING) AND (CAST(`entity` AS STRING) =  CAST(:3_filter_0 AS STRING) AND (CAST(`keyword` AS STRING) =  CAST(:4_filter_0 AS STRING) OR CAST(`keyword` AS STRING) =  CAST(:4_filter_1 AS STRING)))))"

	if clause.Statement != expected {
		t.Errorf("FiltersToSqlClause() failed, expected %v, got %v", expected, clause.Statement)
	} else {
		t.Logf("FiltersToSqlClause() success, expected %v, got %v", expected, clause.Statement)
	}
}
