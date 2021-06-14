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
}

type TransformedData struct {
	Candidate [2]string
}

type ElectiveResponse struct {
	TransformedData CandidateResponse
	entityType string
	error string
}

func ToElectiveStruct(fieldMap FieldMap, data map[string]string) ElectiveResponse{
	rep := ElectiveResponse{}

	ret := CreateCandidate(fieldMap.Candidate, data)

	rep.TransformedData = ret

	return rep
}