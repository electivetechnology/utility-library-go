package sanitise

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
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

	r := csv.NewReader(f)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("err: %v", err)
		}

		for value := range record {
			ret := Phone(record[value], "")
			assert.Equal(t, record[value], ret)
		}
	}
}
