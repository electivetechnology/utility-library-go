package sql

import (
	"strings"

	"github.com/electivetechnology/utility-library-go/logger"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("db/sql")
}

func getSafeFieldName(field string, fieldMap map[string]string) string {
	var ret string

	// Start by escaping and quoting the field name
	if field != "*" {
		ret = quote(escape(field))
	} else {
		ret = field
	}

	// If the field map is populated we can look for a table name
	if len(fieldMap) > 0 {
		var fieldName string
		table := getSafeTableName(field, fieldMap)

		// Check if field was already prepended with table name
		parts := strings.Split(ret, ".")

		if len(parts) > 1 {
			fieldName = parts[1]
		} else {
			fieldName = parts[0]
		}

		ret = table + "." + fieldName
	}

	return ret
}

func getSafeTableName(field string, fieldMap map[string]string) string {
	var table string
	var ret string

	// Find the correct table for this field according to the field map
	if fieldMap[field] != "" {
		table = fieldMap[field]
	} else {
		if fieldMap["*"] != "" {
			table = fieldMap["*"]
		} else {
			table = ""
		}
	}

	// Ignore 'having' table as it doesn't really exist.
	if table == "having" {
		table = ""
	}

	if len(table) > 0 {
		// Start by escaping and quoting the field name.
		ret = quote(escape(table))
	}

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
	num := strings.Count(name, ".")

	if num > 1 {
		// this is specific to BigQuery table names
		// we may need to find better solution for this
		return name
	} else {
		idx := strings.LastIndex(name, ".")
		if idx > 0 {
			name = name[:idx] + "`.`" + name[idx+1:]

		}
	}

	return name
}
