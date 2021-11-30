package sql

const (
	JOIN_INNER = "INNER"
	JOIN_LEFT  = "LEFT"
	JOIN_RIGHT = "RIGHT"
)

type Joiner struct {
	Field string
	Table string
}

type Join struct {
	Type   string
	Master Joiner
	Slave  Joiner
}

func GetJoinSql(q *Query) Clause {
	c := Clause{}

	for _, join := range q.Joins {
		clause := JoinToSqlClause(join, q.FieldMap)

		if len(clause.Statement) > 0 {
			c.Statement += " " + clause.Statement
		}
	}

	return c
}

func JoinToSqlClause(join Join, fieldMap map[string]string) Clause {
	c := Clause{}

	// Add Join statement
	c.Statement = join.Type + " JOIN " + quote(escape(join.Slave.Table))

	// Add ON statement
	c.Statement += " ON " + quote(escape(join.Master.Table)) + "." + quote(escape(join.Master.Field)) + " = " + quote(escape(join.Slave.Table)) + "." + quote(escape(join.Slave.Field))

	return c
}

func NewJoin(masterField string, masterTable string, slaveField string, slaveTable string) (Join, error) {
	return New(masterField, masterTable, slaveField, slaveTable, JOIN_INNER)
}

func NewLeftJoin(masterField string, masterTable string, slaveField string, slaveTable string) (Join, error) {
	return New(masterField, masterTable, slaveField, slaveTable, JOIN_LEFT)
}

func New(masterField string, masterTable string, slaveField string, slaveTable string, jType string) (Join, error) {
	j := Join{Type: jType}

	// Check is master field is defined in field map
	master := Joiner{Field: masterField, Table: masterTable}

	// Add Master to Join
	j.Master = master

	// Check is slave field is defined in field map
	slave := Joiner{Field: slaveField, Table: slaveTable}

	// Add Master & Slave to Join
	j.Slave = slave

	return j, nil
}
