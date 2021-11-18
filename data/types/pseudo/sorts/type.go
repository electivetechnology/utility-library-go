package sorts

import (
	"errors"
	"fmt"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
	"github.com/electivetechnology/utility-library-go/validation"
)

func Check(sort data.Sort, requirements validation.ValidatorRequirements) (data.Sort, error) {
	log.Printf("Validating sort")

	// Check fields
	if !contains(requirements.GetFields(), sort.Field) {
		msg := fmt.Sprintf("Sort failed validation."+
			" Part 1 (field) must be one of "+strings.Join(requirements.GetFields(), ", ")+
			". '%s' given instead.", sort.Field)
		log.Fatalf(msg)
		return sort, errors.New(msg)
	}

	return sort, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
