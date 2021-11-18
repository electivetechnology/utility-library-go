package sql

import (
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

func CriterionToSqlClause(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause {
	c := Clause{}

	method, _ := criterionOperandToMethod(criterion)

	clause := method(criterion, placeHolder, fieldMap, collation)

	c.Statement = clause.Statement

	c.Parameters = clause.Parameters

	return c
}

func criterionToBoolClause(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause {
	c := Clause{}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key, fieldMap)

	// We can only compare true or false for boolean
	if strings.ToLower(criterion.Value) == "true" || criterion.Value == "1" {
		c.Statement = addLogic(criterion) + ` (` + field + ` IS NOT NULL AND ` + field + ` IS NOT FALSE)`
	} else if strings.ToLower(criterion.Value) == "false" || criterion.Value == "0" {
		c.Statement = addLogic(criterion) + ` (` + field + ` IS NULL OR ` + field + ` IS FALSE)`
	}

	return c
}

// criterionToDirectClause turns Criterion object to SQL clause
func criterionToDirectClause(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause {
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

	if criterion.Type == CRITERION_TYPE_VALUE {
		// Add the static value as a parameter
		comparand = ":" + placeHolder
		c.Parameters = map[string]string{
			placeHolder: criterion.Value,
		}
	} else if criterion.Type == CRITERION_TYPE_FIELD {
		comparand = getSafeFieldName(criterion.Value, fieldMap)
	}

	// Check value type
	if _, err := strconv.Atoi(criterion.Value); err == nil {
		castType = CAST_TYPE_NUMERIC
	}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key, fieldMap)

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

	return c
}

func criterionToRelativeClause(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause {
	c := Clause{}
	var op, comparand string
	castType := CAST_TYPE_STRING

	switch criterion.Operand {
	case "gt":
		op = ">"

	case "ge":
		op = ">="

	case "lt":
		op = "<"

	case "le":
		op = "<="
	}

	if criterion.Type == CRITERION_TYPE_VALUE {
		// Add the static value as a parameter
		comparand = ":" + placeHolder
		c.Parameters = map[string]string{
			placeHolder: criterion.Value,
		}
	} else if criterion.Type == CRITERION_TYPE_FIELD {
		comparand = getSafeFieldName(criterion.Value, fieldMap)
	}

	// Check value type
	if _, err := strconv.Atoi(criterion.Value); err == nil {
		castType = CAST_TYPE_NUMERIC
	}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key, fieldMap)

	// Build the final clause.
	if castType == CAST_TYPE_NUMERIC {
		c.Statement = addLogic(criterion) + " CAST(" + field + " AS " + castType + ") " + op + " CAST(" + comparand + " AS " + castType + ")"
	} else {
		c.Statement = addLogic(criterion) + " " + field + " " + op + " " + comparand
	}

	return c
}

func criterionToContainsClause(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause {
	c := Clause{}
	var op, collate, comparand string
	var caseSensitive bool
	castType := CAST_TYPE_STRING

	switch criterion.Operand {
	case "inc": // includes
		op = "LIKE"
		collate = "utf8mb4_bin"
		caseSensitive = true

	case "ninc": // does not include
		op = "NOT LIKE"
		collate = "utf8mb4_bin"
		caseSensitive = true

	case "inci": // includes (case insensitive)
		op = "LIKE"
		collate = "utf8mb4_general_ci"
		caseSensitive = false

	case "ninci": // does not include (case insensitive)
		op = "NOT LIKE"
		collate = "utf8mb4_general_ci"
		caseSensitive = false
	}

	if criterion.Type == CRITERION_TYPE_VALUE {
		// Add the static value as a parameter
		comparand = ":" + placeHolder
		c.Parameters = map[string]string{
			placeHolder: criterion.Value,
		}
	} else if criterion.Type == CRITERION_TYPE_FIELD {
		comparand = getSafeFieldName(criterion.Value, fieldMap)
	}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key, fieldMap)

	// Build the final clause.
	if collation {
		c.Statement = addLogic(criterion) +
			" " + field + " " + op + " " +
			" CONCAT(\"%\", CAST(" + comparand + " AS CHAR), \"%\")" +
			" COLLATE " + collate
	} else {
		if caseSensitive {
			c.Statement = addLogic(criterion) +
				" CAST(" + field + " AS " + castType + ") " + op + " " +
				" CONCAT(\"%\", CAST(" + comparand + " AS " + castType + "), \"%\")"
		} else {
			c.Statement = addLogic(criterion) +
				" LOWER(CAST(" + field + " AS " + castType + ")) " + op + " " +
				" CONCAT(\"%\", LOWER(CAST(" + comparand + " AS " + castType + ")), \"%\")"
		}
	}

	return c
}

func criterionToBeginsClause(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause {
	c := Clause{}
	var op, collate, comparand string
	var caseSensitive bool
	castType := CAST_TYPE_STRING

	switch criterion.Operand {
	case "begins": // begins with
		op = "LIKE"
		collate = "utf8mb4_bin"
		caseSensitive = true

	case "nbegins": // does not begin with
		op = "NOT LIKE"
		collate = "utf8mb4_bin"
		caseSensitive = true

	case "beginsi": // begins with (case insensitive)
		op = "LIKE"
		collate = "utf8mb4_general_ci"
		caseSensitive = false

	case "nbeginsi": // does not begin with (case insensitive)
		op = "NOT LIKE"
		collate = "utf8mb4_general_ci"
		caseSensitive = false
	}

	if criterion.Type == CRITERION_TYPE_VALUE {
		// Add the static value as a parameter
		comparand = ":" + placeHolder
		c.Parameters = map[string]string{
			placeHolder: criterion.Value,
		}
	} else if criterion.Type == CRITERION_TYPE_FIELD {
		comparand = getSafeFieldName(criterion.Value, fieldMap)
	}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key, fieldMap)

	// Build the final clause.
	if collation {
		c.Statement = addLogic(criterion) +
			" " + field + " " + op + " " +
			" CONCAT(CAST(" + comparand + " AS CHAR), \"%\")" +
			" COLLATE " + collate
	} else {
		if caseSensitive {
			c.Statement = addLogic(criterion) +
				" CAST(" + field + " AS " + castType + ") " + op + " " +
				" CONCAT(CAST(" + comparand + " AS " + castType + "), \"%\")"
		} else {
			c.Statement = addLogic(criterion) +
				" LOWER(CAST(" + field + " AS " + castType + ")) " + op + " " +
				" CONCAT(LOWER(CAST(" + comparand + " AS " + castType + ")), \"%\")"
		}
	}

	return c
}

func criterionToRegexClause(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause {
	c := Clause{}
	var comparand string
	castType := CAST_TYPE_STRING

	if criterion.Type == CRITERION_TYPE_VALUE {
		// Add the static value as a parameter
		comparand = ":" + placeHolder
		c.Parameters = map[string]string{
			placeHolder: criterion.Value,
		}
	} else if criterion.Type == CRITERION_TYPE_FIELD {
		comparand = getSafeFieldName(criterion.Value, fieldMap)
	}

	// Escape, quote and qualify the field name for security.
	field := getSafeFieldName(criterion.Key, fieldMap)

	if collation {
		c.Statement = addLogic(criterion) +
			" " + field + " REGEXP " +
			" " + comparand
	} else {
		c.Statement = addLogic(criterion) +
			" REGEXP_CONTAINS(CAST(" + field + " AS " + castType + "), " +
			" CAST(" + comparand + " AS " + castType + "))"
	}

	return c
}

func criterionToInClause(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause {
	c := Clause{}

	return c
}

func criterionOperandToMethod(criterion data.Criterion) (func(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause, string) {
	var method func(criterion data.Criterion, placeHolder string, fieldMap map[string]string, collation bool) Clause
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

func addLogic(criterion data.Criterion) string {
	return strings.ToUpper(criterion.Logic)
}

func (c Clause) removeLogicFromStatement() Clause {
	isAnd := strings.HasPrefix(strings.ToUpper(c.Statement), "AND")

	if isAnd {
		c.Statement = strings.TrimPrefix(c.Statement, "AND ")
	} else {
		c.Statement = strings.TrimPrefix(c.Statement, "OR ")
	}

	return c
}
