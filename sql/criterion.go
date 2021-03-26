package sql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
)

const (
	CRITERION_TYPE_FIELD = "field" // Field type comparison
	CRITERION_TYPE_VALUE = "value" // Value type comparison

	CAST_TYPE_STRING  = "STRING"
	CAST_TYPE_NUMERIC = "NUMERIC"
)

func CriterionToSqlClause(criterion data.Criterion, placeHolder string, collation bool) Clause {
	c := Clause{}

	method, _ := criterionOperandToMethod(criterion)

	clause := method(criterion, placeHolder, collation)

	c.Statement = clause.Statement

	fmt.Printf("SQL from criterion: %s", c.Statement)

	return c
}

func criterionToBoolClause(criterion data.Criterion, placeHolder string, collation bool) Clause {
	c := Clause{}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key)
	fmt.Printf("Safe filed name is: %s\n", field)

	// We can only compare true or false for boolean
	if strings.ToLower(criterion.Value) == "true" || criterion.Value == "1" {
		c.Statement = addLogic(criterion) + ` (` + field + ` IS NOT NULL AND ` + field + ` IS NOT FALSE)`
	} else if strings.ToLower(criterion.Value) == "false" || criterion.Value == "0" {
		c.Statement = addLogic(criterion) + ` (` + field + ` IS NULL OR ` + field + ` IS FALSE)`
	}

	return c
}

// criterionToDirectClause turns Criterion object to SQL clause
func criterionToDirectClause(criterion data.Criterion, placeHolder string, collation bool) Clause {
	c := Clause{}

	var op, collate, comparand string
	var caseSensitive bool
	castType := CAST_TYPE_STRING

	switch criterion.Operand {
	case "eq":
		op = "="
		collate = "utf8mb4_bin"
		caseSensitive = true

	case "ne":
		op = "!="
		collate = "utf8mb4_bin"
		caseSensitive = true

	case "eqi":
		op = "="
		collate = "utf8mb4_general_ci"
		caseSensitive = false

	case "nei":
		op = "!="
		collate = "utf8mb4_general_ci"
		caseSensitive = false
	}

	fmt.Printf("Using operand %s and collate %s\n", op, collate)

	if criterion.Type == CRITERION_TYPE_VALUE {
		// Add the static value as a parameter
		comparand = ":" + placeHolder
		c.Parameters = map[string]string{
			placeHolder: criterion.Value,
		}
	} else if criterion.Type == CRITERION_TYPE_FIELD {
		comparand = getSafeFieldName(criterion.Value)
	}

	// Check value type
	if _, err := strconv.Atoi(criterion.Value); err == nil {
		castType = CAST_TYPE_NUMERIC
	}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key)

	// Build the final clause.
	// We can't assume the value is text, so cast it to char and
	// use the relevant collation for the comparison.
	// Testing shows this doesn't affect use of keys on integer
	// fields, if they are used as part of a case sensitive filter
	if collation {
		c.Statement = addLogic(criterion) +
			" " + field + " " + op + " " +
			" CAST(" + comparand + " AS CHAR)" +
			" COLLATE " + collate
	} else {
		if caseSensitive {
			c.Statement = addLogic(criterion) +
				" CAST(" + field + " AS " + castType + ") " + op + " " +
				" CAST(" + comparand + " AS " + castType + ")"
		} else {
			if castType == CAST_TYPE_NUMERIC {
				c.Statement = addLogic(criterion) +
					" CAST(" + field + " AS " + castType + ") " + op + " " +
					" CAST(" + comparand + " AS " + castType + ")"
			} else if castType == CAST_TYPE_STRING {
				c.Statement = addLogic(criterion) +
					" LOWER(CAST(" + field + " AS " + castType + " )) " + op + " " +
					" LOWER(CAST(" + comparand + " AS " + castType + "))"
			}
		}
	}

	fmt.Printf("SQL Statement: %s\n", c.Statement)
	fmt.Printf("SQL: %s\n", c.GetSql())

	return c
}

func criterionToRelativeClause(criterion data.Criterion, placeHolder string, collation bool) Clause {
	c := Clause{}

	return c
}

func criterionToContainsClause(criterion data.Criterion, placeHolder string, collation bool) Clause {
	c := Clause{}

	return c
}

func criterionToBeginsClause(criterion data.Criterion, placeHolder string, collation bool) Clause {
	c := Clause{}

	return c
}

func criterionToRegexClause(criterion data.Criterion, placeHolder string, collation bool) Clause {
	c := Clause{}

	return c
}

func criterionToInClause(criterion data.Criterion, placeHolder string, collation bool) Clause {
	c := Clause{}

	return c
}

func criterionOperandToMethod(criterion data.Criterion) (func(criterion data.Criterion, placeHolder string, collation bool) Clause, string) {
	var method func(criterion data.Criterion, placeHolder string, collation bool) Clause
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
