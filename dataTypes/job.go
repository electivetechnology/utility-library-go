package dataTypes

type Job struct {
	Title Field
	Headline Field
}

type JobResponse struct {
	Title  			string     `json:"title"`
	Headline    	string     `json:"headline"`
}

func CreateJob(job Job, data map[string] string) JobResponse {
	title := data[job.Title.Field]
	headline := data[job.Headline.Field]

	rep := JobResponse{}
	rep.Title = title
	rep.Headline = headline

	return rep
}