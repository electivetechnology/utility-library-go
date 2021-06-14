package dataTypes

import (
	"encoding/json"
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

type ElectiveResponse struct {
	TransformedData interface{}
	entityType string
	error string
}

func ToElectiveStruct(fieldMap []byte, data []byte) ElectiveResponse{
	ret := ElectiveResponse{}

	var fieldInfo FieldMap
	err := json.Unmarshal(fieldMap, &fieldInfo)
	if err != nil {
		panic(err)
	}

	var dataInfo map[string] string
	err = json.Unmarshal(data, &dataInfo)
	if err != nil {
		panic(err)
	}

	candidate := CreateCandidate(fieldInfo.Candidate, dataInfo)
	ret.TransformedData = candidate

	job := CreateJob(fieldInfo.Job, dataInfo)
	ret.TransformedData = job

	return ret
}