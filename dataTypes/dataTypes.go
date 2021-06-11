package dataTypes

import (
	"github.com/electivetechnology/utility-library-go/logger"
	"go/types"
)

var log logger.Logging

type Field struct {
	Type string
	Field string
	DisplayName string
}

type FieldMap struct {
	candidate Candidate
}

type TransformedData struct {
	Candidate [2]string
}

type ElectiveResponse struct {
	TransformedData byte
	entityType string
	error string
}

func ToElectiveStruct(fieldMap FieldMap, data types.Map) (transformedData TransformedData, entityType string, error string) {
	candidate := CreateCandidate(fieldMap.candidate, data)

	ret := TransformedData{}
	ret.Candidate = candidate

	return ret,  "candidate", ""

}