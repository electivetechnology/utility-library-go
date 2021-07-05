package sanitise

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//type Regex struct {
//	Pattern regexp.Regexp
//}
type Regex struct {
	Pattern string
}
type Code struct {
	Country Regex
}

func Phone(input string) string {
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

func PhoneCode(input string) string {
	output := input
	var codes map[int]string

	codes[998] = "/^\\971[1-7,9][0-9]{7,8}$/"
	codes[998] = "/^\\971[1-7,9][0-9]{7,8}$/"

	for code, regex := range codes {
		reg, err := regexp.Compile(regex)

		if err != nil {
			log.Fatal(err)
		}

		find := reg.FindString(input)
		fmt.Println(code)
		fmt.Println(find)
		//if reg.FindString(input){}
	}

	return output
}
