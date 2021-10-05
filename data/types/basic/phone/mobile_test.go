package phone

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_GBMobileNumbers = []struct {
	input    string
	expected bool
}{
	{"+447877558877", true},
	{"+442087600471", false},
	{"+447579487821", true},
	{"+441087600471", false},
}

func TestIsGBMobilePhone(t *testing.T) {
	for i, tt := range test_GBMobileNumbers {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p := NewPhone(tt.input)
			assert.Equal(t, tt.expected, IsGBMobilePhone(p))
		})
	}
}
