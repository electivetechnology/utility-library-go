package dataTypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToElectiveJobStruct(t *testing.T) {
	fieldMapJSON := `{"job":{
"Title":{"type":"string","field":"Title","DisplayName":"Title"},
"CandidateSummary":{"type":"string","field":"CandidateSummary","DisplayName":"CandidateSummary"},
"Brief":{"type":"string","field":"Brief","DisplayName":"Brief"},
"Type":{"type":"string","field":"Type","DisplayName":"Type"},
"Currency":{"type":"string","field":"Currency","DisplayName":"Currency"},
"Salary":{"type":"string","field":"Salary","DisplayName":"Salary"},
"SalaryUnit":{"type":"string","field":"SalaryUnit","DisplayName":"SalaryUnit"},
"StartDate":{"type":"string","field":"StartDate","DisplayName":"StartDate"},
"Location":{"type":"string","field":"Location","DisplayName":"Location"},
"Keywords":{"type":"string","field":"Keywords","DisplayName":"Keywords"},
"RemoteWorking":{"type":"string","field":"RemoteWorking","DisplayName":"RemoteWorking"},
"OtherTypes":{"type":"string","field":"OtherTypes","DisplayName":"OtherTypes"},
"SalaryFlexibility":{"type":"string","field":"SalaryFlexibility","DisplayName":"SalaryFlexibility"},
"EarlierStart":{"type":"string","field":"EarlierStart","DisplayName":"EarlierStart"},
"NoticeFlexibility":{"type":"string","field":"NoticeFlexibility","DisplayName":"NoticeFlexibility"},
"Status":{"type":"string","field":"Status","DisplayName":"Status"},
"Headline":{"type":"string","field":"Headline","DisplayName":"Headline"},
"Notes":{"type":"string","field":"Notes","DisplayName":"Notes"}
}}`


	title 					:= "Coca cola"
	candidateSummary 				:= "overview example"
	brief 				:= "linkedin"
	jobType 				:= "website"
	currency 			:= "introduction"
	salary	:= "chris"
	salaryUnit 	:= "dixon"
	startDate 	:= "dixon@awesome.co.uk"
	location 	:= "+44231412423123"
	keywords 			:= "keywords"
	remoteWorking 			:= "remoteWorking"
	otherTypes 			:= "otherTypes"
	salaryFlexibility 			:= "salaryFlexibility"
	earlierStart 			:= "earlierStart"
	noticeFlexibility 			:= "noticeFlexibility"
	status 					:= "active"
	headline 			:= "headline"
	notes 			:= "notes"

	dataJSON := `{
"Title":"` + title + `",
"CandidateSummary":"` + candidateSummary + `",
"Brief":"` + brief + `",
"Type":"` + jobType + `",
"Currency":"` + currency + `",
"Salary":"` + salary + `",
"SalaryUnit":"` + salaryUnit + `",
"StartDate":"` + startDate + `",
"Location":"` + location + `",
"Keywords":"` + keywords + `",
"RemoteWorking":"` + remoteWorking + `",
"OtherTypes":"` + otherTypes + `",
"SalaryFlexibility":"` + salaryFlexibility + `",
"EarlierStart":"` + earlierStart + `",
"NoticeFlexibility":"` + noticeFlexibility + `",
"Status":"` + status + `",
"Headline":"` + headline + `",
"Notes":"` + notes + `"
}`

	ret := formatResponse(fieldMapJSON, dataJSON)

	if ret.error != "" {
		t.Errorf(ret.error)
	}

	rep 						:= JobResponse{}
	rep.Title 					= title
	rep.CandidateSummary 		= candidateSummary
	rep.Brief 					= brief
	rep.Type 					= jobType
	rep.Currency 				= currency
	rep.Salary 					= salary
	rep.SalaryUnit 				= salaryUnit
	rep.StartDate 				= startDate
	rep.Location 				= location
	rep.Keywords 				= keywords
	rep.RemoteWorking 			= remoteWorking
	rep.OtherTypes 				= otherTypes
	rep.SalaryFlexibility 		= salaryFlexibility
	rep.EarlierStart 			= earlierStart
	rep.NoticeFlexibility 		= noticeFlexibility
	rep.Status 					= status
	rep.Headline 				= headline
	rep.Notes 					= notes

	assert.Equal(t, rep, ret.TransformedData)
}