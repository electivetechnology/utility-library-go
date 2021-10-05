package phone

import (
	"errors"
	"regexp"
)

const (
	DIALING_CODE_GB = "+44"
	COUNTRY_CODE_GB = "GB"
)

type CallingCode struct {
	DialingCode string
	CountryCode string
}

var countryCallingCodePatterns = []struct {
	pattern     string
	callingCode CallingCode
}{
	{`(?m)^\+44[1-9][0-9]{6,10}$`, CallingCode{DialingCode: DIALING_CODE_GB, CountryCode: COUNTRY_CODE_GB}},
	{`(?m)^44[1-9][0-9]{6,10}$`, CallingCode{DialingCode: DIALING_CODE_GB, CountryCode: COUNTRY_CODE_GB}},
	{`(?m)^0044[1-9][0-9]{6,10}$`, CallingCode{DialingCode: DIALING_CODE_GB, CountryCode: COUNTRY_CODE_GB}},
}

func GetCallingCodeFromString(number string) (CallingCode, error) {
	for _, tt := range countryCallingCodePatterns {
		match, _ := regexp.MatchString(tt.pattern, number)
		if match {
			return tt.callingCode, nil
		}
	}

	return CallingCode{}, errors.New("could not determine Calling Code from given number")
}
