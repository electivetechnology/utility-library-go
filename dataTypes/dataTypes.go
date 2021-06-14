package dataTypes

import (
	"encoding/json"
	"fmt"
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

func ToElectiveStruct(fieldMap []byte, data []byte) ElectiveResponse{
	ret := ElectiveResponse{}
	transformedData := TransformedData{}


	var fieldFinal map[string]interface{}
	err := json.Unmarshal(fieldMap, &fieldFinal)
	if err != nil {
		panic(err)
	}

	var dataFinal map[string]interface{}
	err = json.Unmarshal(data, &dataFinal)
	if err != nil {
		panic(err)
	}

	fmt.Print(dataFinal)
	fmt.Print(fieldFinal)


	//candidate := CreateCandidate(fieldMap.Candidate, data)
	//transformedData.Candidate = candidate
	//
	//job := CreateJob(fieldMap.Job, data)
	//transformedData.Job = job

	ret.TransformedData = transformedData

	return ret
}