package context

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
	DisplaysType "github.com/electivetechnology/utility-library-go/data/types/pseudo/displays"
	"github.com/electivetechnology/utility-library-go/validation"
	"github.com/gin-gonic/gin"
)

type Display struct {
	ID          string
	Directive   string
	DataDisplay *data.Display
}

type Displays struct {
	Displays map[string]Display
}

type DisplayType interface {
	GetDisplays() map[string]Display
	GetDisplay(key string) (Display, error)
	GetDataDisplays() map[string]data.Display
}

func NewDisplays() Displays {
	displays := Displays{}
	displays.Displays = make(map[string]Display)

	return displays
}

func GetDisplays(c *gin.Context, requirements validation.ValidatorRequirements) (Displays, error) {
	displays, err := GetDisplaysFromContext(c)
	if err != nil {
		log.Fatalf(err.Error())
		return displays, err
	}

	for idx, display := range displays.GetDataDisplays() {
		_, err := ValidateDisplay(display, idx, requirements)
		if err != nil {
			log.Fatalf(err.Error())
			return displays, err
		}
	}

	return displays, nil
}

func ValidateDisplay(display data.Display, namespace string, requirements validation.ValidatorRequirements) (data.Display, error) {
	// Check if display is valid
	_, err := DisplaysType.Check(display, requirements)
	if err != nil {
		msg := fmt.Sprintf("Display %s failed validation with message: %s", namespace, err.Error())
		log.Fatalf(msg)
		return display, errors.New(msg)
	}

	return display, nil
}

func GetDisplaysFromContext(ctx *gin.Context) (Displays, error) {
	displays := NewDisplays()

	// Get anonymous
	anonymous, err := GetAnonymousDisplaysFromContext(ctx)
	if err != nil {
		log.Fatalf(err.Error())
		return displays, err
	}

	for key, display := range anonymous.Displays {
		displays.Displays[key] = display
	}

	return displays, nil
}

func GetAnonymousDisplaysFromContext(ctx *gin.Context) (Displays, error) {
	d, _ := ctx.GetQueryArray("displays[]")
	displays := NewDisplays()

	for idx, directive := range d {
		display := Display{}
		display.ID = getSafeDisplayName("0" + strconv.Itoa(idx))
		display.Directive = directive
		dataDisplay, err := DirectiveToDataDisplay(directive, idx, display.ID)
		// Check for errors
		if err != nil {
			log.Fatalf(err.Error())
			return displays, err
		}

		// Assign values
		display.DataDisplay = dataDisplay
		displays.Displays[display.ID] = display
	}

	return displays, nil
}

func DirectiveToDataDisplay(directive string, index int, name string) (*data.Display, error) {
	display := data.NewDisplay()
	var msg string

	// Check all parts of directive exist
	parts := strings.Split(directive, "-")

	// Check all elements of directive are present
	if len(parts) > 2 || len(parts) == 0 {
		msg = fmt.Sprintf(
			"value for displays %s index %d must be string in format of "+
				" {field} or {field}-{alias}"+
				" Example: id, firstName-name."+
				" Following value given instead: '%s'", name, index, directive)
		log.Fatalf(msg)
		return display, errors.New(msg)
	}

	// Set Display details
	display.Field = parts[0]
	if len(parts) == 2 {
		display.Alias = parts[1]
	}

	return display, nil
}

func getSafeDisplayName(key string) string {
	key = strings.ReplaceAll(key, "[", "")
	key = strings.ReplaceAll(key, "]", "")

	return "d_" + key
}

func (d Displays) GetDisplays() map[string]Display {
	return d.Displays
}

func (d Displays) GetDisplay(key string) (Display, error) {
	display, found := d.Displays[key]

	if found {
		return display, nil
	}

	msg := fmt.Sprintf("display with name %s does not exists", key)
	log.Fatalf(msg)

	return Display{}, errors.New(msg)
}

func (d Displays) GetDataDisplays() map[string]data.Display {
	displays := make(map[string]data.Display)
	for _, display := range d.GetDisplays() {
		displays[display.ID] = *display.DataDisplay
	}

	return displays
}
