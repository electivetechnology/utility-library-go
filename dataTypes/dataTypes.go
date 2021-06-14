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

type ElectiveResponse struct {
	TransformedData interface{}
	entityType string
	error string
}

func ToElectiveStruct(fieldMap []byte, data []byte) ElectiveResponse{
	ret := ElectiveResponse{}

	var fieldFinal FieldMap
	err := json.Unmarshal(fieldMap, &fieldFinal)
	if err != nil {
		panic(err)
	}

	var dataFinal map[string] string
	err = json.Unmarshal(data, &dataFinal)
	if err != nil {
		panic(err)
	}

	candidate := CreateCandidate(fieldFinal.Candidate, dataFinal)
	ret.TransformedData = candidate

	fmt.Print(fieldFinal.Job)
	//
	//job := CreateJob(fieldMap.Job, data)
	//transformedData.Job = job

	return ret
}