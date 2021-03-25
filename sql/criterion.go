package sql

import (
	"fmt"

	"github.com/electivetechnology/utility-library-go/data"
)

func CriterionToSqlClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	method := criterionOperandToMethod(criterion)

	fmt.Printf("Criterion method is: %s\n", method)

	return c
}

func criterionOperandToMethod(criterion data.Criterion) string {
	fmt.Printf("Evaluating Criterion method for operand: %s\n", criterion.Operand)
	var method string
	switch criterion.Operand {
	// Boolean check to see if the value is set or not.
	case "bool":
		method = "criterionToBool"

	// Direct comparison checks.
	case "eq", // equals
		"ne",  // does not equal
		"eqi", // equals (case insensitive)
		"nei": // does not equal (case insensitive)

		method = "criterionToDirect"

	// Relative comparison checks.
	case "gt", // greater than
		"ge", // greater than or equal to
		"lt", // less than
		"le": // less than or equal to

		method = "criterionToRelative"

	// Wildcard comparison checks (contains/includes)
	case "inc", // includes
		"ninc",  // does not include
		"inci",  // includes (case insensitive)
		"ninci": // does not include (case insensitive)

		method = "criterionToContains"

	// Wildcard comparison checks (begins with)
	case "begins", // begins with
		"nbegins",  // does not begin with
		"beginsi",  // begins with (case insensitive)
		"nbeginsi": // does not begin with (case insensitive)
		method = "criterionToBegins"

	// Regex match
	case "re": // matches regex string
		method = "criterionToRegex"

	// Check for a list of values (match or no match).
	case "in", // is in the list
		"nin",  // is not in the list
		"ini",  // is in the list (case insensitive)
		"nini": // is not in the list (case insensitive)
		method = "criterionToIn"

	}

	return method
}
