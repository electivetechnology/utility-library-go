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
	Client Client
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
		ret.error = err.Error()
		return ret
	}

	var dataInfo map[string] string
	err = json.Unmarshal(data, &dataInfo)
	if err != nil {
		ret.error = err.Error()
		return ret
	}


	if(fieldInfo.Candidate != Candidate{}){
		candidate := CreateCandidate(fieldInfo.Candidate, dataInfo)
		ret.TransformedData = candidate
		ret.entityType = "candidate"
		return ret
	}

	if(fieldInfo.Job != Job{}){
		job := CreateJob(fieldInfo.Job, dataInfo)
		ret.TransformedData = job
		ret.entityType = "job"
		return ret
	}

	if(fieldInfo.Client != Client{}){
		client := CreateClient(fieldInfo.Client, dataInfo)
		ret.TransformedData = client
		ret.entityType = "client"
		return ret
	}

	return ret
}