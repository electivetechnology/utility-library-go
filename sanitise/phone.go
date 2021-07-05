package sanitise

import (
	"log"
	"regexp"
	"strings"
)

func Phone (input string) string {
	output := strings.TrimSpace(input)

	splitString := strings.FieldsFunc(output, SplitBySlash)
	output = splitString[0]

	reg, err := regexp.Compile("[^0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	output = reg.ReplaceAllString(output, "")

	output = strings.TrimLeftFunc(output, TrimByZero)

	return output
}
