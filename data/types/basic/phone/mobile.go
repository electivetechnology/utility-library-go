package phone

import (
	"regexp"
)

func IsMobilePhone(phone Phone) bool {
	switch phone.CallingCode.CountryCode {
	case COUNTRY_CODE_GB:
		return IsGBMobilePhone(phone)
	}

	return false
}

func IsGBMobilePhone(phone Phone) bool {
	log.InfoF("Checking number %s is GB mobile number", phone.Number)
	match, _ := regexp.MatchString(`[7][0-9]{9}$`, phone.Number)

	if match {
		log.InfoF("Number is GB mobile number")
	}

	return match
}
