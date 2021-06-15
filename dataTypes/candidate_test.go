package dataTypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
"Status":{"type":"string","field":"Status","DisplayName":"Status"},
"AddressLine1":{"type":"string","field":"AddressLine1","DisplayName":"AddressLine1"},
"AddressLine2":{"type":"string","field":"AddressLine2","DisplayName":"AddressLine2"},
"City":{"type":"string","field":"City","DisplayName":"City"},
"Postcode":{"type":"string","field":"Postcode","DisplayName":"Postcode"},
"Country":{"type":"string","field":"Country","DisplayName":"Country"},
"County":{"type":"string","field":"County","DisplayName":"County"},
"JobTitle":{"type":"string","field":"JobTitle","DisplayName":"JobTitle"},
"Location":{"type":"string","field":"Location","DisplayName":"Location"},
"WillRelocate":{"type":"string","field":"WillRelocate","DisplayName":"WillRelocate"},
"ExpectedSalary":{"type":"string","field":"ExpectedSalary","DisplayName":"ExpectedSalary"},
"SalaryCurrency":{"type":"string","field":"SalaryCurrency","DisplayName":"SalaryCurrency"},
"Notice":{"type":"string","field":"Notice","DisplayName":"Notice"},
"NoticeUnit":{"type":"string","field":"NoticeUnit","DisplayName":"NoticeUnit"},
"JobType":{"type":"string","field":"JobType","DisplayName":"JobType"},
"Company":{"type":"string","field":"Company","DisplayName":"Company"},
"Summary":{"type":"string","field":"Summary","DisplayName":"Summary"}
}}`


	email 					:= "dixon@awesome.co.uk"
	firstName 				:= "chris"
	lastName 				:= "dixon"
	phone 					:= "+442314124124"
	primaryLanguage 		:= "english"
	secondaryLanguage 		:= "german"
	tertiaryLanguage 		:= "spanish"
	cvText 					:= "Some cv text"
	title 					:= "Mr"
	alternativePhoneNumber 	:= "+44231412423123"
	dob 					:= "11/12/13"
	vendorId 				:= "some Id"
	vendorStatus 			:= "some status"
	vendorSource 			:= "some source"
	gender 					:= "male"
	status 					:= "active"

	address1 				:= "2"
	address2 				:= "Carlile Mews"
	city 					:= "Leeds"
	postcode 				:= "LS7 PBU"
	country 				:= "UK"
	county 					:= "Yorkshire"

	jobTitle 				:= "permanent"
	location 				:= "London"
	willRelocate 			:= "true"
	expectedSalary 			:= "40000"
	salaryCurrency 			:= "GBP"
	notice 					:= "4"
	noticeUnit 				:= "weeks"
	jobType 				:= "permenant"
	company 				:= "recii"
	summary 				:= "great guy"


	dataJSON := `{
"Email":"` + email + `",
"FirstName":"` + firstName + `",
"LastName":"` + lastName + `",
"Phone":"` + phone + `",
"PrimaryLanguage":"` + primaryLanguage + `",
"SecondaryLanguage":"` + secondaryLanguage + `",
"TertiaryLanguage":"` + tertiaryLanguage + `",
"CvText":"` + cvText + `",
"Title":"` + title + `",
"AlternativePhoneNumber":"` + alternativePhoneNumber + `",
"Dob":"` + dob + `",
"VendorId":"` + vendorId + `",
"VendorStatus":"` + vendorStatus + `",
"VendorSource":"` + vendorSource + `",
"Gender":"` + gender + `",
"Status":"` + status + `",
"AddressLine1":"` + address1 + `",
"AddressLine2":"` + address2 + `",
"City":"` + city + `",
"Postcode":"` + postcode + `",
"Country":"` + country + `",
"County":"` + county + `",
"JobTitle":"` + jobTitle + `",
"Location":"` + location + `",
"WillRelocate":"` + willRelocate + `",
"ExpectedSalary":"` + expectedSalary + `",
"SalaryCurrency":"` + salaryCurrency + `",
"Notice":"` + notice + `",
"NoticeUnit":"` + noticeUnit + `",
"JobType":"` + jobType + `",
"Company":"` + company + `",
"Summary":"` + summary + `"
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

	rep.CandidateAddress.AddressLine1 					= address1
	rep.CandidateAddress.AddressLine2 					= address2
	rep.CandidateAddress.City 							= city
	rep.CandidateAddress.Postcode 						= postcode
	rep.CandidateAddress.Country 						= country
	rep.CandidateAddress.County 						= county

	rep.CandidateJob.JobTitle 							= jobTitle
	rep.CandidateJob.Location 							= location
	rep.CandidateJob.WillRelocate 						= willRelocate
	rep.CandidateJob.ExpectedSalary 					= expectedSalary
	rep.CandidateJob.SalaryCurrency 					= salaryCurrency
	rep.CandidateJob.Notice 							= notice
	rep.CandidateJob.NoticeUnit 						= noticeUnit
	rep.CandidateJob.JobType 							= jobType
	rep.CandidateJob.Company 							= company
	rep.CandidateJob.Summary 							= summary

	assert.Equal(t, rep, ret.TransformedData)
}
