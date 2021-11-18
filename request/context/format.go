package context

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	FORMAT_BODY_RAW_JSON = "application/json"
	FORMAT_BODY_RAW_CSV  = "text/csv"
)

type Format struct {
	Format   string
	Priority float64
}

type Formats struct {
	Formats []Format
}

type FormatType interface {
	GetFormat() string
}

func GetFormat(c *gin.Context) Formats {
	formats := GetRequestedFormat(c)

	return formats
}

func NewFormat() Format {
	format := Format{Format: FORMAT_BODY_RAW_JSON, Priority: 1.0}

	return format
}

func NewFormats() Formats {
	formats := Formats{}
	formats.Formats = make([]Format, 0)

	return formats
}

func GetRequestedFormat(ctx *gin.Context) Formats {
	directives := ctx.GetHeader("Accept")
	formats := NewFormats()

	// Split all formats sent in header
	parts := strings.Split(directives, ",")

	for _, directive := range parts {
		format := NewFormat()
		// Split directive to format and priority
		elements := strings.Split(directive, ";")
		for _, element := range elements {
			if strings.Contains(element, "/") {
				format.Format = element
			}

			if strings.Contains(element, "q=") {
				element = strings.ReplaceAll(element, "q=", "")
				p, err := strconv.ParseFloat(element, 64)

				if err != nil {
					msg := fmt.Sprintf(
						"Could not parse Accept header priority into float. "+
							"Expected float type got '%v' instead", element)
					log.Fatalf(msg)
				}

				format.Priority = p
			}
		}

		// Add format to the list accepted formats
		formats.Formats = append(formats.Formats, format)
	}

	sort.SliceStable(formats.Formats, func(i, j int) bool {
		return formats.Formats[i].Priority > formats.Formats[j].Priority
	})

	return formats
}
