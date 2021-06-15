package dataTypes

import (
	"encoding/json"
)

func formatResponse (fieldMapJSON string, dataJSON string) ElectiveResponse {
	rawData := json.RawMessage(dataJSON)
	data, err := rawData.MarshalJSON()
	if err != nil {
		panic(err)
	}

	rawField := json.RawMessage(fieldMapJSON)
	fieldMap, err := rawField.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return ToElectiveStruct(fieldMap, data)
}