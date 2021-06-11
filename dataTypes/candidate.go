package dataTypes

import (
	"go/types"
	"reflect"
)

type Candidate struct {
	Dob Field
	Email Field
}

type CandidateResponse struct {
	Dob string
	Email string
}

func CreateCandidate(candidate Candidate, data types.Map) [2]string {
	v := reflect.ValueOf(data)
	dob := v.FieldByName(candidate.Dob.Field)
	email := v.FieldByName(candidate.Email.Field)

	x := [2]string{dob.String(), email.String()}

	return x
}