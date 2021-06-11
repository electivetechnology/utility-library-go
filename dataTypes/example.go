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

	title := Field{}
	title.Type = "string"
	title.Field = "VendorMergedTitle"
	title.DisplayName = "Title"

	headline := Field{}
	headline.Type = "string"
	headline.Field = "VendorHeadline"
	headline.DisplayName = "Headline"


	job := Job{}
	job.Title = title
	job.Headline = headline

	fieldMap := FieldMap{}
	fieldMap.Candidate = candidate
	fieldMap.Job = job

	dataJSON := `{
	   	"VendorDateOfBirth": "11/12/13",
	   	"VendorMergedEmail": "dixon@awesome.co.uk",
		"VendorMergedTitle": "Sales Rep",
		"VendorHeadline": "Read all about it"
	}`

	// Declared an empty map interface
	var result map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(dataJSON), &result)

	// Run App Migrations
	ret := ToElectiveStruct(fieldMap, result)
	fmt.Print(ret)
}