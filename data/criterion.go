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

	CRITERION_LOGIC_INTERSECTION = "and" // Logic Intersection (AND A AND B AND C ...)
	CRITERION_LOGIC_UNION        = "or"  // Logic Union (OR A OR B OR C ...)

	CRITERION_OP_BOOL     = "bool"     // Boolean comparison, e.g. true or false.
	CRITERION_OP_EQ       = "eq"       // Equals comparison (case sensitive).
	CRITERION_OP_EQI      = "eqi"      // Equals comparison (case insensitive).
	CRITERION_OP_NE       = "ne"       // Not equals comparison (case sensitive).
	CRITERION_OP_NEI      = "nei"      // Not equals comparison (case insensitive).
	CRITERION_OP_LT       = "lt"       // Less than comparison.
	CRITERION_OP_GT       = "gt"       // Greater than comparison.
	CRITERION_OP_GE       = "ge"       // Greater than or equal to comparison.
	CRITERION_OP_LE       = "le"       // Less than or equal to comparison.
	CRITERION_OP_INC      = "inc"      // Includes (case sensitive).
	CRITERION_OP_INCI     = "inci"     // Includes (case insensitive).
	CRITERION_OP_NINC     = "ninc"     // Not includes (case sensitive).
	CRITERION_OP_NINCI    = "ninci"    // Not includes (case insensitive).
	CRITERION_OP_RE       = "re"       // Regular expression.
	CRITERION_OP_BEGINS   = "begins"   // Begins (case sensitive).
	CRITERION_OP_BEGINSI  = "beginsi"  // Begins (case insensitive).
	CRITERION_OP_NBEGINS  = "nbegins"  // Not Begins (case sensitive).
	CRITERION_OP_NBEGINSI = "nbeginsi" // Not Begins (case insensitive).
	CRITERION_OP_IN       = "in"       // Comma delimited list of values to match (case sensitive).
	CRITERION_OP_INI      = "ini"      // Comma delimited list of values to match (case insensitive).
	CRITERION_OP_NIN      = "nin"      // Comma delimited list of values to not match (case sensitive).
	CRITERION_OP_NINI     = "nini"     // Comma delimited list of values to not match (case insensitive).
)

var logics = map[string]string{
	CRITERION_LOGIC_INTERSECTION: CRITERION_LOGIC_INTERSECTION,
	CRITERION_LOGIC_UNION:        CRITERION_LOGIC_UNION,
}

var operands = map[string]string{
	CRITERION_OP_BOOL:     CRITERION_OP_BOOL,
	CRITERION_OP_EQ:       CRITERION_OP_EQ,
	CRITERION_OP_EQI:      CRITERION_OP_EQI,
	CRITERION_OP_NE:       CRITERION_OP_NE,
	CRITERION_OP_NEI:      CRITERION_OP_NEI,
	CRITERION_OP_LT:       CRITERION_OP_LT,
	CRITERION_OP_GT:       CRITERION_OP_GT,
	CRITERION_OP_GE:       CRITERION_OP_GE,
	CRITERION_OP_LE:       CRITERION_OP_LE,
	CRITERION_OP_INC:      CRITERION_OP_INC,
	CRITERION_OP_INCI:     CRITERION_OP_INCI,
	CRITERION_OP_NINC:     CRITERION_OP_NINC,
	CRITERION_OP_NINCI:    CRITERION_OP_NINCI,
	CRITERION_OP_RE:       CRITERION_OP_RE,
	CRITERION_OP_BEGINS:   CRITERION_OP_BEGINS,
	CRITERION_OP_BEGINSI:  CRITERION_OP_BEGINSI,
	CRITERION_OP_NBEGINS:  CRITERION_OP_NBEGINS,
	CRITERION_OP_NBEGINSI: CRITERION_OP_NBEGINSI,
	CRITERION_OP_IN:       CRITERION_OP_IN,
	CRITERION_OP_INI:      CRITERION_OP_INI,
	CRITERION_OP_NIN:      CRITERION_OP_NIN,
	CRITERION_OP_NINI:     CRITERION_OP_NINI,
}

func GetOperandOptions() []string {
	v := make([]string, 0)

	for _, value := range operands {
		v = append(v, value)
	}

	return v
}

func IsOperand(val string) bool {
	if _, ok := operands[val]; ok {
		return true
	}

	return false
}

func GetLogicOptions() []string {
	v := make([]string, 0)

	for _, value := range logics {
		v = append(v, value)
	}

	return v
}
