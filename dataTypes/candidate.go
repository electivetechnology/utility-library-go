package dataTypes

type Candidate struct {
	Dob Field
	Email Field
}

type CandidateResponse struct {
	Dob  		string     `json:"dob"`
	Email    	string     `json:"email"`
}

func CreateCandidate(candidate Candidate, data map[string] string) CandidateResponse {
	dob := data[candidate.Dob.Field]
	email := data[candidate.Email.Field]

	rep := CandidateResponse{}
	rep.Email = email
	rep.Dob = dob

	return rep
}