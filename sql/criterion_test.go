package sql

import (
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

type TestCriterionOperandToMethodItem struct {
	criterion data.Criterion
	expected  string
}

type TestCriterionToDirectClauseItem struct {
	criterion   data.Criterion
	placeHolder string
	collation   bool
	expected    string
}

func TestCriterionOperandToMethod(t *testing.T) {
	c1 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "bool",
		Type:    "value",
		Value:   "true",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "eq",
		Type:    "value",
		Value:   "2",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "ne",
		Type:    "value",
		Value:   "2",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "eqi",
		Type:    "value",
		Value:   "2",
	}

	c5 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "nei",
		Type:    "value",
		Value:   "2",
	}

	c6 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "gt",
		Type:    "value",
		Value:   "2",
	}

	c7 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "ge",
		Type:    "value",
		Value:   "2",
	}

	c8 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "lt",
		Type:    "value",
		Value:   "2",
	}

	c9 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "le",
		Type:    "value",
		Value:   "2",
	}

	c10 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "inc",
		Type:    "value",
		Value:   "true",
	}

	c11 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "ninc",
		Type:    "value",
		Value:   "true",
	}

	c12 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "inci",
		Type:    "value",
		Value:   "true",
	}

	c13 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "ninci",
		Type:    "value",
		Value:   "true",
	}

	c14 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "begins",
		Type:    "value",
		Value:   "true",
	}

	c15 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "nbegins",
		Type:    "value",
		Value:   "true",
	}

	c16 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "beginsi",
		Type:    "value",
		Value:   "true",
	}

	c17 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "nbeginsi",
		Type:    "value",
		Value:   "true",
	}

	c18 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "re",
		Type:    "value",
		Value:   "true",
	}

	c19 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "in",
		Type:    "value",
		Value:   "true",
	}

	c20 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "nin",
		Type:    "value",
		Value:   "true",
	}

	c21 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "ini",
		Type:    "value",
		Value:   "true",
	}

	c22 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "nini",
		Type:    "value",
		Value:   "true",
	}

	testData := []TestCriterionOperandToMethodItem{
		{c1, "criterionToBoolClause"},
		{c2, "criterionToDirectClause"},
		{c3, "criterionToDirectClause"},
		{c4, "criterionToDirectClause"},
		{c5, "criterionToDirectClause"},
		{c6, "criterionToRelativeClause"},
		{c7, "criterionToRelativeClause"},
		{c8, "criterionToRelativeClause"},
		{c9, "criterionToRelativeClause"},
		{c10, "criterionToContainsClause"},
		{c11, "criterionToContainsClause"},
		{c12, "criterionToContainsClause"},
		{c13, "criterionToContainsClause"},
		{c14, "criterionToBeginsClause"},
		{c15, "criterionToBeginsClause"},
		{c16, "criterionToBeginsClause"},
		{c17, "criterionToBeginsClause"},
		{c18, "criterionToRegexClause"},
		{c19, "criterionToInClause"},
		{c20, "criterionToInClause"},
		{c21, "criterionToInClause"},
		{c22, "criterionToInClause"},
	}

	for _, item := range testData {
		_, method := criterionOperandToMethod(item.criterion)

		if method != item.expected {
			t.Errorf("criterionOperandToMethod("+item.criterion.Operand+") failed, expected %v, got %v", item.expected, method)
		} else {
			t.Logf("criterionOperandToMethod("+item.criterion.Operand+") success, expected %v, got %v", item.expected, method)
		}
	}
}

func TestCriterionToDirectClause(t *testing.T) {
	c1 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "eq",
		Type:    "value",
		Value:   "hello",
	}
	c1_0Expected := "AND `id` =  CAST(:filter_0_0 AS CHAR) COLLATE utf8mb4_bin"
	c1_1Expected := "AND CAST(`id` AS STRING) =  CAST(:filter_0_0 AS STRING)"

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "ne",
		Type:    "value",
		Value:   "hello",
	}
	c2_0Expected := "AND `id` !=  CAST(:filter_0_0 AS CHAR) COLLATE utf8mb4_bin"
	c2_1Expected := "AND CAST(`id` AS STRING) !=  CAST(:filter_0_0 AS STRING)"

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "eqi",
		Type:    "value",
		Value:   "hello",
	}
	c3_0Expected := "AND `id` =  CAST(:filter_0_0 AS CHAR) COLLATE utf8mb4_general_ci"
	c3_1Expected := "AND LOWER(CAST(`id` AS STRING )) =  LOWER(CAST(:filter_0_0 AS STRING))"

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "nei",
		Type:    "value",
		Value:   "2",
	}
	c4_0Expected := "AND `id` !=  CAST(:filter_0_0 AS CHAR) COLLATE utf8mb4_general_ci"
	c4_1Expected := "AND CAST(`id` AS NUMERIC) !=  CAST(:filter_0_0 AS NUMERIC)"

	c5 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "eq",
		Type:    "field",
		Value:   "engagement_id",
	}
	c5_0Expected := "AND `id` =  CAST(`engagement_id` AS CHAR) COLLATE utf8mb4_bin"
	c5_1Expected := "AND CAST(`id` AS STRING) =  CAST(`engagement_id` AS STRING)"

	c6 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "eqi",
		Type:    "field",
		Value:   "engagement_id",
	}
	c6_0Expected := "AND `id` =  CAST(`engagement_id` AS CHAR) COLLATE utf8mb4_general_ci"
	c6_1Expected := "AND LOWER(CAST(`id` AS STRING )) =  LOWER(CAST(`engagement_id` AS STRING))"

	testData := []TestCriterionToDirectClauseItem{
		{c1, "filter_0_0", true, c1_0Expected},
		{c1, "filter_0_0", false, c1_1Expected},
		{c2, "filter_0_0", true, c2_0Expected},
		{c2, "filter_0_0", false, c2_1Expected},
		{c3, "filter_0_0", true, c3_0Expected},
		{c3, "filter_0_0", false, c3_1Expected},
		{c4, "filter_0_0", true, c4_0Expected},
		{c4, "filter_0_0", false, c4_1Expected},
		{c5, "filter_0_0", true, c5_0Expected},
		{c5, "filter_0_0", false, c5_1Expected},
		{c6, "filter_0_0", true, c6_0Expected},
		{c6, "filter_0_0", false, c6_1Expected},
	}

	for _, item := range testData {
		ret := criterionToDirectClause(item.criterion, item.placeHolder, item.collation)
		if ret.Statement != item.expected {
			t.Errorf("criterionToDirectClause() failed, expected %v, got %v", item.expected, ret.Statement)
		} else {
			t.Logf("criterionToDirectClause() success, expected %v, got %v", item.expected, ret.Statement)
		}
	}
}
