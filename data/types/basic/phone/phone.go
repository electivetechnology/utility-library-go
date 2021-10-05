package phone

import (
	"errors"
	"fmt"

	"github.com/electivetechnology/utility-library-go/logger"
)

var log logger.AdvancedLogging

func init() {
	// Add generic logger
	log = logger.NewLogger("data/types/basic/phone")
}

type Phone struct {
	Number      string
	CallingCode CallingCode
	IsValid     bool
	IsMobile    bool
}

func NewPhone(number string) Phone {
	return Phone{Number: number, IsValid: false, IsMobile: false}
}

func FromString(number string) (Phone, error) {
	log.InfoF("Getting PhoneNumber for %s", number)
	callingCode, err := GetCallingCodeFromString(number)

	if err != nil {
		msg := fmt.Sprintf("Could not get CallingCode for number %s. Country not suppored", number)
		log.Fatalf(msg)

		return NewPhone(number), errors.New(msg)
	}

	phone := NewPhone(number)
	phone.CallingCode = callingCode
	phone.IsValid = true
	phone.IsMobile = IsMobilePhone(phone)

	return phone, nil
}
