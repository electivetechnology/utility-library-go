package sanitise

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var email = "m/y@email.com"

func TestToCommaSanitise(t *testing.T) {
	input	:= email + ",2@email.com /2@email.com 2@email.com"

	ret := SanitiseEmail(input)

	assert.Equal(t, email, ret)
}

func TestToSpaceSanitise(t *testing.T) {
	input	:= email + " 2@email.com /2@email.com, 2@email.com"

	ret := SanitiseEmail(input)

	assert.Equal(t, email, ret)
}


func TestToSlashSanitise(t *testing.T) {
	input	:= email + "/2@email.com /2@email.com, 2@email.com"

	ret := SanitiseEmail(input)

	assert.Equal(t, email, ret)
}

func TestToAllSanitise(t *testing.T) {
	input	:= email + " /,  2@email.com /2@email.com, 2@email.com"

	ret := SanitiseEmail(input)

	assert.Equal(t, email, ret)
}