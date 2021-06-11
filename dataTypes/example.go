package dataTypes

import (
	"encoding/json"
	"fmt"
	"go/types"
)

func Example() {
	email := Field{}
	email.Type = "string"
	email.Field = "VendorMergedEmail"
	email.DisplayName = "Email"

	dob := Field{}
	dob.Type = "string"
	dob.Field = "VendorDateOfBirth"
	dob.DisplayName = "Date of Birth"


	candidate := Candidate{}
	candidate.Email = email
	candidate.Dob = dob

	fieldMap := FieldMap{}
	fieldMap.Candidate = candidate

	//dataJSON := `{
	//    "VendorDateOfBirth": "11/12/13",
	//    "VendorMergedEmail": "dixon@awesome.co.uk"
	//}`

	// Declared an empty map interface
	//var result map[string]interface{}
	var result types.Map

	// Unmarshal or Decode the JSON to the interface.
	//json.Unmarshal([]byte(dataJSON), &result)

	example := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)

	json.Unmarshal(example, &result)

	fmt.Print(result)
	//fmt.Print(result)

	// Run App Migrations
	ret, entity, err := ToElectiveStruct(fieldMap, result)

	log.Printf("Results: ", ret)
	log.Printf("Entity: " + entity)
	log.Printf("Error: " + err)
}