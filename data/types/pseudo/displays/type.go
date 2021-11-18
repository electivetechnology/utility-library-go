package displays

import (
	"errors"
	"fmt"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
	"github.com/electivetechnology/utility-library-go/validation"
)

func Check(display data.Display, requirements validation.ValidatorRequirements) (data.Display, error) {
	log.Printf("Validating display")

	// Check fields
	if !contains(requirements.GetFields(), display.Field) {
		msg := fmt.Sprintf("Display failed validation."+
			" Part 1 (field) must be one of "+strings.Join(requirements.GetFields(), ", ")+
			". '%s' given instead.", display.Field)
		log.Fatalf(msg)
		return display, errors.New(msg)
	}

	return display, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
