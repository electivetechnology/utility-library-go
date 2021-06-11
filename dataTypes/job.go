package dataTypes

type Job struct {
	Title Field
	Headline Field
}

type JobResponse struct {
	Title string
	Headline string
}

func CreateJob(job Job, data map[string] string) JobResponse {
	dob := data[job.Title.Field]
	email := data[job.Headline.Field]

	rep := JobResponse{}
	rep.Title = email
	rep.Headline = dob

	return rep
}