package data

type Criterion struct {
	Logic   string
	Key     string
	Operand string
	Type    string
	Value   string
}

const (
	CRITERION_TYPE_FIELD = "field"
	CRITERION_TYPE_VALUE = "value"

	CRITERION_LOGIC_INTERSCTION = "and" // Logic Intersection (AND A AND B AND C ...)
	CRITERION_LOGIC_UNION       = "or"  // Logic Union (OR A OR B OR C ...)
)
