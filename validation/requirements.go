package validation

type ValidatorRequirements interface {
	GetFields() []string
	SetFields(fields []string)
}

type Requirements struct {
	Fields []string
}

func (v Requirements) GetFields() []string {
	return v.Fields
}

func (v Requirements) SetFields(fields []string) {
	v.Fields = fields
}
