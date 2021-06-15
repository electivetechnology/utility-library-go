package dataTypes

import (
	"encoding/json"
	"fmt"
)

func Example() {
	fieldMapJSON := `{"candidate":{"Dob":{"type":"string","field":"VendorDateOfBirth","DisplayName":"Date of Birth"},"Email":{"type":"string","field":"VendorMergedEmail","DisplayName":"Email"}}}`

	dob := `"11/12/13"`
	email := `"dixon@awesome.co.uk"`
	dataJSON := `{"VendorDateOfBirth":` + dob + `,"VendorMergedEmail":` + email + `}`

	rawData := json.RawMessage(dataJSON)
	data, err := rawData.MarshalJSON()
	if err != nil {
		panic(err)
	}

	rawField := json.RawMessage(fieldMapJSON)
	fieldMap, err := rawField.MarshalJSON()
	if err != nil {
		panic(err)
	}

	// Run App Migrations
	ret := ToElectiveStruct(fieldMap, data)
	fmt.Print(ret)
}