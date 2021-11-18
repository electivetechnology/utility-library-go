package filters

import (
	"errors"
	"fmt"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
	"github.com/electivetechnology/utility-library-go/validation"
)

func Check(filter data.Filter, requirements validation.ValidatorRequirements) (data.Filter, error) {
	log.Printf("Validating filter")

	for i, c := range filter.Criterions {
		// Check fields
		if !contains(requirements.GetFields(), c.Key) {
			msg := fmt.Sprintf("Criterion with index %d failed validation."+
				" Part 2 (field) must be one of "+strings.Join(requirements.GetFields(), ", ")+
				". '%s' given instead.", i, c.Key)
			log.Fatalf(msg)
			return filter, errors.New(msg)
		}
	}

	return filter, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
