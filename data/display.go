package data

type Display struct {
	Field string
	Alias string
}

func NewDisplay() *Display {
	display := &Display{}

	return display
}
