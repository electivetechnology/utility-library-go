package sanitise

import (
	"strings"
)

func Email(input string) string {
	trimmed := strings.TrimSpace(input)

	splitString := strings.FieldsFunc(trimmed, SplitByCommaSpace)
	output := buildOutput(splitString, false)

	splitString = strings.FieldsFunc(output, SplitBySlash)
	output = buildOutput(splitString, true)

	return output
}

func buildOutput(seperatedString []string, hasSlash bool) string {
	output := ""

	for _, element := range seperatedString {
		output += element
		if strings.Contains(element, "@") {
			break
		} else if hasSlash {
			output += "/"
		}
	}

	return output
}
