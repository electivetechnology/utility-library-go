package dataTypes

import (
	"fmt"
)

type Candidate struct {
	Dob Field
	Email Field
}

type CandidateResponse struct {
	Dob interface{}
	//Dob string
	Email interface{}
}

func CreateCandidate(candidate Candidate, data map[string]interface{}) CandidateResponse {
	dob := data[candidate.Dob.Field]
	email := data[candidate.Email.Field]

	fmt.Print(dob)
	fmt.Print(email)

	rep := CandidateResponse{}
	rep.Email = email
	rep.Dob = dob
}