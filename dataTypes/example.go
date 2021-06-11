package dataTypes

import (
	"encoding/json"
	"fmt"
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

	dataJSON := `{
	   "VendorDateOfBirth": "11/12/13",
	   "VendorMergedEmail": "dixon@awesome.co.uk"
	}`

	// Declared an empty map interface
	var result map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(dataJSON), &result)

	// Run App Migrations
	ret := ToElectiveStruct(fieldMap, result)
	fmt.Print(ret)
}