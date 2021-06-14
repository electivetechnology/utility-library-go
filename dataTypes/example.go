package dataTypes

import (
	"encoding/json"
)
type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
func Example() {
	fieldMapJSON := `{"candidate":{"Dob":{"type":"string","field":"VendorDateOfBirth","DisplayName":"Date of Birth"},"Email":{"type":"string","field":"VendorMergedEmail","DisplayName":"Email"}}}`

	dataJSON := `{"VendorDateOfBirth":"11/12/13","VendorMergedEmail":"dixon@awesome.co.uk"}`

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
	
	ToElectiveStruct(fieldMap, data)

	//// Run App Migrations
	//ret := ToElectiveStruct(fieldMap, data)
	//fmt.Print(ret)
}