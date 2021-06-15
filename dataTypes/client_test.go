package dataTypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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