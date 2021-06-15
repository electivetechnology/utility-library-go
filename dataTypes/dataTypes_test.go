package dataTypes

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func formatResponse (dataJSON string, fieldMapJSON string) ElectiveResponse {
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
	fieldMapJSON := `{"candidate":{"Dob":{"type":"string","field":"VendorDob","DisplayName":"Date of Birth"},"Email":{"type":"string","field":"VendorEmail","DisplayName":"Email"}}}`


	email := "dixon@awesome.co.uk"
	//firstName 				:= "chris"
	//lastName 				:= "dixon"
	//phone 					:= "+442314124124"
	//primaryLanguage 		:= "english"
	//secondaryLanguage 		:= "german"
	//tertiaryLanguage 		:= "spanish"
	//cvText 					:= "Some cv text"
	//alternativePhoneNumber 	:= "+44231412423123"
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


	dataJSON := `{"VendorDob":"` + dob + `","VendorEmail":"` + email + `"}`

	ret := formatResponse(fieldMapJSON, dataJSON)

	if ret.error != "" {
		t.Errorf(ret.error)
	}

	rep 						:= CandidateResponse{}
	rep.Email = email
	rep.Dob = dob


	assert.Equal(t, rep, ret.TransformedData)
}