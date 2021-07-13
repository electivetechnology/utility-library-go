package sanitise

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func TestToCSV(t *testing.T) {

	f, err := os.Open("./correct_numbers.csv")

	if err != nil {
		log.Printf("err: %v", err)
	}

	csvr := csv.NewReader(f)
	csvr.FieldsPerRecord = -1

	for {
		record, err := csvr.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("err: %v", err)
		}

		for value := range record {
			code := FindCountryCode(record[value])

			if code != 0 {
				codeString := strconv.FormatInt(int64(code), 10)

				res := strings.Replace(record[value], codeString, "", 1)
				ret := Phone(res, codeString)

				assert.Equal(t, record[value], ret)

			} else {
				log.Fatalf("fatal error", code)
			}
		}
	}
}

func FindCountryCode(input string) int {

	codeList := CodeList()

	for code, regex := range codeList {
		reg, err := regexp.Compile(regex)

		if err != nil {
			log.Fatalf("fatal error", err)
		}

		find := reg.MatchString(input)

		if find {
			return code
		}
	}

	return 0
}
