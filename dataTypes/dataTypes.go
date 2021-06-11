package dataTypes

import (
	"github.com/electivetechnology/utility-library-go/logger"
)

var log logger.Logging

type Field struct {
	Type string
	Field string
	DisplayName string
}

type FieldMap struct {
	Candidate Candidate
	Job Job
}

type TransformedData struct {
	Candidate CandidateResponse
	Job JobResponse
}

type ElectiveResponse struct {
	TransformedData TransformedData
	entityType string
	error string
}

func ToElectiveStruct(fieldMap FieldMap, data map[string]string) ElectiveResponse{
	ret := ElectiveResponse{}
	transformedData := TransformedData{}

	candidate := CreateCandidate(fieldMap.Candidate, data)
	transformedData.Candidate = candidate

	job := CreateJob(fieldMap.Job, data)
	transformedData.Job = job

	ret.TransformedData = transformedData

	return ret
}