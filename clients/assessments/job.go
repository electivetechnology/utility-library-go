package assessments

import "github.com/electivetechnology/utility-library-go/clients/connect"

type JobResponse struct {
	ApiResponse *connect.ApiResponse
	Job         *Job
}

type Job struct {
	Id               string   `json:"id"`
	Title            string   `json:"title"`
	CandidateSummary string   `json:"candidateSummary"`
	Brief            string   `json:"brief"`
	Type             string   `json:"type"`
	Currency         string   `json:"currency"`
	Salary           string   `json:"salary"`
	SalaryUnit       string   `json:"salaryUnit"`
	Location         string   `json:"location"`
	Keywords         []string `json:"keywords"`
	ClientId         string   `json:"clientId"`
	ClientName       string   `json:"clientName"`
	Status           string   `json:"status"`
	Headline         string   `json:"headline"`
	Notes            string   `json:"notes"`
}

func (client Client) GetJobByVendor(vendor string, vendorId string, token string) (JobResponse, error) {
	return JobResponse{}, nil
}
