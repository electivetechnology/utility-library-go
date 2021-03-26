package sql

import (
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

type TestCriterionOperandToMethodItem struct {
	criterion data.Criterion
	expected  string
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
		{c1, "criterionToBool"},
		{c2, "criterionToDirect"},
		{c3, "criterionToDirect"},
		{c4, "criterionToDirect"},
		{c5, "criterionToDirect"},
		{c6, "criterionToRelative"},
		{c7, "criterionToRelative"},
		{c8, "criterionToRelative"},
		{c9, "criterionToRelative"},
		{c10, "criterionToContains"},
		{c11, "criterionToContains"},
		{c12, "criterionToContains"},
		{c13, "criterionToContains"},
		{c14, "criterionToBegins"},
		{c15, "criterionToBegins"},
		{c16, "criterionToBegins"},
		{c17, "criterionToBegins"},
		{c18, "criterionToRegex"},
		{c19, "criterionToIn"},
		{c20, "criterionToIn"},
		{c21, "criterionToIn"},
		{c22, "criterionToIn"},
	}

	for _, item := range testData {
		ret := criterionOperandToMethod(item.criterion)
		if ret != item.expected {
			t.Errorf("criterionOperandToMethod("+item.criterion.Operand+") failed, expected %v, got %v", item.expected, ret)
		} else {
			t.Logf("criterionOperandToMethod("+item.criterion.Operand+") success, expected %v, got %v", item.expected, ret)
		}
	}
}
