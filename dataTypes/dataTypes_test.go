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

func TestToElectiveClientStruct(t *testing.T) {
	fieldMapJSON := `{"client":{
"Name":{"type":"string","field":"Name","DisplayName":"Name"},
"Overview":{"type":"string","field":"Overview","DisplayName":"Overview"},
"Linkedin":{"type":"string","field":"Linkedin","DisplayName":"Linkedin"},
"Website":{"type":"string","field":"Website","DisplayName":"Website"},
"Introduction":{"type":"string","field":"Introduction","DisplayName":"Introduction"},
"PrimaryContactFirstName":{"type":"string","field":"PrimaryContactFirstName","DisplayName":"PrimaryContactFirstName"},
"PrimaryContactLastName":{"type":"string","field":"PrimaryContactLastName","DisplayName":"PrimaryContactLastName"},
"PrimaryContactEmail":{"type":"string","field":"PrimaryContactEmail","DisplayName":"PrimaryContactEmail"},
"PrimaryContactPhone":{"type":"string","field":"PrimaryContactPhone","DisplayName":"PrimaryContactPhone"},
"Status":{"type":"string","field":"Status","DisplayName":"Status"},
"BrandColour":{"type":"string","field":"BrandColour","DisplayName":"BrandColour"}
}}`


	name 					:= "Coca cola"
	overview 				:= "overview example"
	linkedin 				:= "linkedin"
	website 				:= "website"
	introduction 			:= "introduction"
	primaryContactFirstName	:= "chris"
	primaryContactLastName 	:= "dixon"
	primaryContactEmail 	:= "dixon@awesome.co.uk"
	primaryContactPhone 	:= "+44231412423123"
	status 					:= "active"
	brandColour 			:= "red"

	dataJSON := `{
"Name":"` + name + `",
"Overview":"` + overview + `",
"Linkedin":"` + linkedin + `",
"Website":"` + website + `",
"Introduction":"` + introduction + `",
"PrimaryContactFirstName":"` + primaryContactFirstName + `",
"PrimaryContactLastName":"` + primaryContactLastName + `",
"PrimaryContactEmail":"` + primaryContactEmail + `",
"PrimaryContactPhone":"` + primaryContactPhone + `",
"Status":"` + status + `",
"BrandColour":"` + brandColour + `"
}`

	ret := formatResponse(fieldMapJSON, dataJSON)

	if ret.error != "" {
		t.Errorf(ret.error)
	}

	rep 						:= ClientResponse{}
	rep.Name 					= name
	rep.Overview 				= overview
	rep.Linkedin 				= linkedin
	rep.Website 				= website
	rep.Introduction 			= introduction
	rep.PrimaryContactFirstName = primaryContactFirstName
	rep.PrimaryContactLastName 	= primaryContactLastName
	rep.PrimaryContactEmail 	= primaryContactEmail
	rep.PrimaryContactPhone 	= primaryContactPhone
	rep.Status 					= status
	rep.BrandColour 			= brandColour

	assert.Equal(t, rep, ret.TransformedData)
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
"Dob":{"type":"string","field":"Dob","DisplayName":"Date of Birth"},
"VendorId":{"type":"string","field":"VendorId","DisplayName":"VendorId"},
"VendorStatus":{"type":"string","field":"VendorStatus","DisplayName":"VendorStatus"},
"VendorSource":{"type":"string","field":"VendorSource","DisplayName":"VendorSource"},
"Gender":{"type":"string","field":"Gender","DisplayName":"Gender"},
"Status":{"type":"string","field":"Status","DisplayName":"Status"}
}}`


	email 					:= "dixon@awesome.co.uk"
	firstName 				:= "chris"
	lastName 				:= "dixon"
	phone 					:= "+442314124124"
	primaryLanguage 		:= "english"
	secondaryLanguage 		:= "german"
	tertiaryLanguage 		:= "spanish"
	cvText 					:= "Some cv text"
	alternativePhoneNumber 	:= "+44231412423123"
	dob 					:= "11/12/13"
	vendorId 				:= "some Id"
	vendorStatus 			:= "soem status"
	vendorSource 			:= "some source"
	gender 					:= "male"
	status 					:= "active"

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
"Dob":"` + dob + `",
"VendorId":"` + vendorId + `",
"VendorStatus":"` + vendorStatus + `",
"VendorSource":"` + vendorSource + `",
"Gender":"` + gender + `",
"Status":"` + status + `"
}`

	ret := formatResponse(fieldMapJSON, dataJSON)

	if ret.error != "" {
		t.Errorf(ret.error)
	}

	rep 						:= CandidateResponse{}
	rep.Email 					= email
	rep.FirstName 				= firstName
	rep.LastName 				= lastName
	rep.Phone 					= phone
	rep.PrimaryLanguage 		= primaryLanguage
	rep.SecondaryLanguage 		= secondaryLanguage
	rep.TertiaryLanguage 		= tertiaryLanguage
	rep.CvText 					= cvText
	rep.AlternativePhoneNumber 	= alternativePhoneNumber
	rep.Dob 					= dob
	rep.VendorId 				= vendorId
	rep.VendorStatus 			= vendorStatus
	rep.VendorSource 			= vendorSource
	rep.Gender 					= gender
	rep.Status 					= status


	assert.Equal(t, rep, ret.TransformedData)
}