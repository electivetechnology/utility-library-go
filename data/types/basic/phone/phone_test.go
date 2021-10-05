package phone

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_PhoneNumbers = []struct {
	input    string
	expected Phone
}{
	{"+447877558877", Phone{Number: "+447877558877", IsValid: true, IsMobile: true}},
	{"+442087600471", Phone{Number: "+442087600471", IsValid: true, IsMobile: false}},
	{"+4402087600471", Phone{Number: "+4402087600471", IsValid: false, IsMobile: false}},
	{"+4407877558877", Phone{Number: "+4407877558877", IsValid: false, IsMobile: false}},
}

func TestFromString(t *testing.T) {
	for i, tt := range test_PhoneNumbers {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p, _ := FromString(tt.input)

			if tt.expected.IsMobile {
				assert.True(t, p.IsMobile)
				if !p.IsMobile {
					t.Errorf("Expected number %s to be a mobile number", tt.input)
				}
			} else {
				assert.False(t, p.IsMobile)
				if p.IsMobile {
					t.Errorf("Expected number %s to be a NOT mobile number", tt.input)
				}
			}

			if tt.expected.IsValid {
				assert.True(t, p.IsValid)
				if !p.IsValid {
					t.Errorf("Expected number %s to be valid number", tt.input)
				}
			} else {
				assert.False(t, p.IsValid)
				if p.IsValid {
					t.Errorf("Expected number %s to be invalid number", tt.input)
				}
			}
		})
	}
}
