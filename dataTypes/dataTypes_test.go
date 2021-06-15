package dataTypes

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func formatResponse (fieldMapJSON string, dataJSON string) ElectiveResponse {
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

	return ToElectiveStruct(fieldMap, data)
}



func TestToElectiveCandidateStruct(t *testing.T) {
	fieldMapJSON := `{"candidate":{
"Email":{"type":"string","field":"Email","DisplayName":"Email"},
"FirstName":{"type":"string","field":"FirstName","DisplayName":"FirstName"},
"LastName":{"type":"string","field":"LastName","DisplayName":"LastName"},
"Phone":{"type":"string","field":"Phone","DisplayName":"Phone"},
"PrimaryLanguage":{"type":"string","field":"PrimaryLanguage","DisplayName":"PrimaryLanguage"},
"SecondaryLanguage":{"type":"string","field":"SecondaryLanguage","DisplayName":"SecondaryLanguage"},
"TertiaryLanguage":{"type":"string","field":"TertiaryLanguage","DisplayName":"TertiaryLanguage"},
"CvText":{"type":"string","field":"CvText","DisplayName":"CvText"},
"AlternativePhoneNumber":{"type":"string","field":"AlternativePhoneNumber","DisplayName":"AlternativePhoneNumber"},
"Dob":{"type":"string","field":"Dob","DisplayName":"Date of Birth"}
}}`


	email := "dixon@awesome.co.uk"
	firstName 				:= "chris"
	lastName 				:= "dixon"
	phone 					:= "+442314124124"
	primaryLanguage 		:= "english"
	secondaryLanguage 		:= "german"
	tertiaryLanguage 		:= "spanish"
	cvText 					:= "Some cv text"
	alternativePhoneNumber 	:= "+44231412423123"
	dob 					:= "11/12/13"
	//vendorId 				:= "some Id"
	//vendorStatus 			:= "soem status"
	//vendorSource 			:= "some source"
	//gender 					:= "male"
	//status 					:= "active"
	//
	//address1 				:= "2"
	//address2 				:= "Carlile Mews"
	//city 					:= "Leeds"
	//postcode 				:= "LS7 PBU"
	//country 				:= "UK"
	//county 					:= "Yorkshire"
	//
	//title 					:= "Mr"
	//location 				:= "London"
	//willRelocate 			:= "true"
	//expectedSalary 			:= "40000"
	//salaryCurrency 			:= "GBP"
	//notice 					:= "4"
	//noticeUnit 				:= "weeks"
	//jobType 				:= "permenant"
	//company 				:= "recii"
	//summary 				:= "great guy"


	dataJSON := `{
"Email":"` + email + `",
"FirstName":"` + firstName + `",
"LastName":"` + lastName + `",
"Phone":"` + phone + `",
"PrimaryLanguage":"` + primaryLanguage + `",
"SecondaryLanguage":"` + secondaryLanguage + `",
"TertiaryLanguage":"` + tertiaryLanguage + `",
"CvText":"` + cvText + `",
"AlternativePhoneNumber":"` + alternativePhoneNumber + `",
"Dob":"` + dob + `"
}`

	ret := formatResponse(fieldMapJSON, dataJSON)

	if ret.error != "" {
		t.Errorf(ret.error)
	}

	rep 						:= CandidateResponse{}
	rep.Email = email
	rep.FirstName = firstName
	rep.LastName = lastName
	rep.Phone = phone
	rep.PrimaryLanguage = primaryLanguage
	rep.SecondaryLanguage = secondaryLanguage
	rep.TertiaryLanguage = tertiaryLanguage
	rep.CvText = cvText
	rep.AlternativePhoneNumber = alternativePhoneNumber
	rep.Dob = dob


	assert.Equal(t, rep, ret.TransformedData)
}