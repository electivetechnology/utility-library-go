package sanitise

import (
	"regexp"
	"strings"
)

func Phone(input string, defaultCountry string) string {
	log.Printf("Input: %v , DefaultCountry: %v", input, defaultCountry)

	output := strings.TrimSpace(input)
	log.Printf("Output after TrimSpace: %v", output)

	splitString := strings.FieldsFunc(output, SplitBySlash)
	output = splitString[0]
	log.Printf("Output after SplitBySlash: %v", output)

	reg, err := regexp.Compile("[^0-9]+")

	if err != nil {
		log.Fatalf("Output after error", err)
	}

	output = reg.ReplaceAllString(output, "")
	log.Printf("Output after ReplaceAllString: %v", output)

	output = strings.TrimLeftFunc(output, TrimByZero)
	log.Printf("Output after TrimLeftFunc: %v", output)

	withDefault := defaultCountry + output
	if !HasCountryCode(output) && defaultCountry != "" && HasCountryCode(withDefault) {
		output = withDefault
	}

	log.Printf("Output after HasCountryCode: %v", output)

	return output
}

func HasCountryCode(input string) bool {

	codeList := CodeList()

	for _, regex := range codeList {
		reg, err := regexp.Compile(regex)

		if err != nil {
			log.Fatalf("fatal error", err)
		}

		find := reg.MatchString(input)

		if find {
			return true
		}
	}

	return false
}

func CodeList() map[int]string {
	codes := make(map[int]string)

	codes[971] = "^971[1-7,9][0-9]{7,8}$"
	codes[998] = "^998[1-7,9][0-9]{7,8}$"
	codes[44] = "^44[1-9][0-9]{6,10}$"

	return codes
}
