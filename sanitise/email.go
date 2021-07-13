package sanitise

import (
	"strings"
)

func Email(input string) string {
	log.Printf("Input: %v", input)

	output := strings.TrimSpace(input)
	log.Printf("Output after TrimSpace: %v", output)

	splitString := strings.FieldsFunc(output, SplitByCommaSpace)
	output = buildOutput(splitString, false)
	log.Printf("Output after SplitByCommaSpace and buildOutput: %v", output)

	splitString = strings.FieldsFunc(output, SplitBySlash)
	output = buildOutput(splitString, true)
	log.Printf("Output after SplitBySlash and buildOutput: %v", output)

	return output
}

func buildOutput(seperatedString []string, hasSlash bool) string {
	output := ""

	// build output string, ignore everything after element with @
	for index, element := range seperatedString {
		output += element
		if strings.Contains(element, "@") {
			break
		} else if hasSlash && index != 0 {
			output += "/"
		}
	}

	return output
}
