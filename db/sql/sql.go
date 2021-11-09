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

func getSafeFieldName(field string) string {
	// Start by escaping and quoting the field name.
	ret := quote(escape(field))

	return ret
}

func getSafeTableName(name string) string {
	// Start by escaping and quoting the field name.
	ret := quote(escape(name))

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
