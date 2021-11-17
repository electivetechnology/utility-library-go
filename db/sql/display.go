package sql

import (
	"sort"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
)

func DisplaysToSqlClause(displays map[string]data.Display, fieldMap map[string]string) Clause {
	c := Clause{}
	ds := make([]string, 0)

	for _, display := range displays {

		field := getSafeFieldName(display.Field, fieldMap)
		if display.Alias != "" {
			field += " AS " + display.Alias
		}
		ds = append(ds, field)
	}

	// Now that we have all displays
	// Let's turn them into string and add alaias if necessary
	sort.Strings(ds)
	c.Statement = strings.Join(ds, ", ")

	return c
}
