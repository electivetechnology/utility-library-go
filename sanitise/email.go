package sanitise

import (
	"strings"
)

func SanitiseEmail (input string) string {
	trimmed := strings.TrimSpace(input)

	splitString := strings.FieldsFunc(trimmed, SplitByCommaSpace)
	output := BuildOutput(splitString, false)

	splitString = strings.FieldsFunc(output, SplitBySlash)
	output = BuildOutput(splitString, true)

	return output
}

func SplitByCommaSpace(r rune) bool {
	return r == ',' || r == ' '
}

func SplitBySlash(r rune) bool {
	return r == '/'
}

func BuildOutput(seperatedString []string, hasSlash bool) string {
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