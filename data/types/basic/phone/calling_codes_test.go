package phone

import (
	"strconv"
	"testing"
)

var callingCodes = []struct {
	input    string
	expected CallingCode
}{
	{"+447877558877", CallingCode{DialingCode: DIALING_CODE_GB, CountryCode: COUNTRY_CODE_GB}},
	{"+442087600471", CallingCode{DialingCode: DIALING_CODE_GB, CountryCode: COUNTRY_CODE_GB}},
}

func TestGetCallingCodeFromString(t *testing.T) {
	for i, tt := range callingCodes {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cc, err := GetCallingCodeFromString(tt.input)
			if err != nil {
				t.Errorf("Received error trying to parse this number %s", tt.input)
			}

			if cc.CountryCode != tt.expected.CountryCode {
				t.Errorf("Expected country code %s for number %s, got %s insted", tt.expected.CountryCode, tt.input, cc.CountryCode)
			}

			if cc.DialingCode != tt.expected.DialingCode {
				t.Errorf("Expected dialing code %s for number %s, got %s insted", tt.expected.DialingCode, tt.input, cc.DialingCode)
			}
		})
	}
}
