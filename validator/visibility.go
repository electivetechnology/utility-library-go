package validator

import (
	"github.com/go-playground/validator/v10"
)

const (
	PRIVATE   = "private"
	PUBLIC    = "public"
	PROTECTED = "protected"
)

func VisibilityValidator(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case PUBLIC:
		return true
	case PRIVATE:
		return true
	case PROTECTED:
		return true
	default:
		return false
	}
}
