package sanitise

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var phone = "7810362111"
var country = "44"

func TestToSanitise(t *testing.T) {
	input := "0000( " + phone + "+) adasdasd/  8273192837"

	ret := Phone(input, "")

	assert.Equal(t, phone, ret)
}

func TestToSanitiseWithCountry(t *testing.T) {
	input := "0000( " + phone + "+) adasdasd/  8273192837"

	ret := Phone(input, country)

	assert.Equal(t, country+phone, ret)
}
