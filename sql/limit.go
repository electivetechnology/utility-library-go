package sql

import "strconv"

func GetLimitSql(q *Query) string {
	ret := ""

	if q.Limit > 0 && q.Offset >= 0 {
		ret = "LIMIT " + strconv.Itoa(q.Offset) + "," + strconv.Itoa(q.Limit)
	}

	return ret
}
