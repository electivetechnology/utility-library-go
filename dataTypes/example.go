package dataTypes

import (
	"encoding/json"
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
	var result map[string]interface{}
	//var result types.Map

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(dataJSON), &result)

	//fmt.Print(result["VendorDateOfBirth"])
	//fmt.Print(result)

	// Run App Migrations
	ToElectiveStruct(fieldMap, result)
	//ret, entity, err := ToElectiveStruct(fieldMap, result)

	//log.Printf("Results: ", ret)
	//log.Printf("Entity: " + entity)
	//log.Printf("Error: " + err)
}