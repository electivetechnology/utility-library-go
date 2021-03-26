package sql

import (
	"fmt"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
)

func CriterionToSqlClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	method, _ := criterionOperandToMethod(criterion)

	clause := method(criterion, placeHolder)

	c.Statement = clause.Statement

	return c
}

func criterionToBoolClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key)
	fmt.Printf("Safe filed name is: %s\n", field)

	// We can only compare true or false for boolean
	if strings.ToLower(criterion.Value) == "true" || criterion.Value == "1" {
		c.Statement = addLogic(criterion) + ` (` + field + ` IS NOT NULL AND ` + field + ` !="")`
	} else if strings.ToLower(criterion.Value) == "false" || criterion.Value == "0" {
		c.Statement = addLogic(criterion) + ` (` + field + ` IS NULL OR ` + field + ` = "")`
	}

	return c
}

func criterionToDirectClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	return c
}

func criterionToRelativeClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	return c
}

func criterionToContainsClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	return c
}

func criterionToBeginsClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	return c
}

func criterionToRegexClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	return c
}

func criterionToInClause(criterion data.Criterion, placeHolder string) Clause {
	c := Clause{}

	return c
}

func criterionOperandToMethod(criterion data.Criterion) (func(criterion data.Criterion, placeHolder string) Clause, string) {
	var method func(criterion data.Criterion, placeHolder string) Clause
	var methodName string

	switch criterion.Operand {
	// Boolean check to see if the value is set or not.
	case "bool":
		method = criterionToBoolClause
		methodName = "criterionToBoolClause"

	// Direct comparison checks.
	case "eq", // equals
		"ne",  // does not equal
		"eqi", // equals (case insensitive)
		"nei": // does not equal (case insensitive)

		method = criterionToDirectClause
		methodName = "criterionToDirectClause"

	// Relative comparison checks.
	case "gt", // greater than
		"ge", // greater than or equal to
		"lt", // less than
		"le": // less than or equal to

		method = criterionToRelativeClause
		methodName = "criterionToRelativeClause"

	// Wildcard comparison checks (contains/includes)
	case "inc", // includes
		"ninc",  // does not include
		"inci",  // includes (case insensitive)
		"ninci": // does not include (case insensitive)

		method = criterionToContainsClause
		methodName = "criterionToContainsClause"

	// Wildcard comparison checks (begins with)
	case "begins", // begins with
		"nbegins",  // does not begin with
		"beginsi",  // begins with (case insensitive)
		"nbeginsi": // does not begin with (case insensitive)
		method = criterionToBeginsClause
		methodName = "criterionToBeginsClause"

	// Regex match
	case "re": // matches regex string
		method = criterionToRegexClause
		methodName = "criterionToRegexClause"

	// Check for a list of values (match or no match).
	case "in", // is in the list
		"nin",  // is not in the list
		"ini",  // is in the list (case insensitive)
		"nini": // is not in the list (case insensitive)
		method = criterionToInClause
		methodName = "criterionToInClause"
	}

	return method, methodName
}

func getSafeFieldName(field string) string {
	// Start by escaping and quoting the field name.
	ret := quote(escape(field))

	return ret
}

// Escape a SQL string for use with MySQL
func escape(name string) string {
	name = strings.ReplaceAll(name, "\\", "\\\\")
	name = strings.ReplaceAll(name, `\0`, "\\0")
	name = strings.ReplaceAll(name, "\n", "\\n")
	name = strings.ReplaceAll(name, "\r", "\\r")
	name = strings.ReplaceAll(name, "'", "\\'")
	name = strings.ReplaceAll(name, `"`, `\\"`)
	name = strings.ReplaceAll(name, "x1a", "\\Z")

	return name
}

// Quote SQL string for use with MySQL
func quote(name string) string {

	name = "`" + strings.ReplaceAll(name, "`", "") + "`"
	name = strings.ReplaceAll(name, ".", "`.`")

	return name
}

func addLogic(criterion data.Criterion) string {
	return strings.ToUpper(criterion.Logic)
}
