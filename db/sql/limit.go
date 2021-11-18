package sql

import "strconv"

func GetLimitSql(q *Query) string {
	ret := ""

	// Apply default limit
	if q.Limit == 0 {
		q.Limit = DEFAULT_LIMIT
	}

	if q.Limit > 0 && q.Offset >= 0 {
		ret = "LIMIT " + strconv.Itoa(q.Limit) + " OFFSET " + strconv.Itoa(q.Offset)
	}

	return ret
}
