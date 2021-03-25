package sql

type Clause struct {
	Statement  string
	Parameters map[string]string
}
