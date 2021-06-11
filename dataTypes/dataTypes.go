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
	TransformedData byte
	entityType string
	error string
}

func ToElectiveStruct(fieldMap FieldMap, data map[string]interface{}) {

	//fmt.Print(fieldMap)
	//fmt.Print(data)

	CreateCandidate(fieldMap.Candidate, data)
}