package dataTypes

type Job struct {
	Title Field
	CandidateSummary Field
	Brief Field
	Type Field
	Currency Field
	Salary Field
	SalaryUnit Field
	StartDate Field
	Location Field
	Keywords Field
	RemoteWorking Field
	OtherTypes Field
	SalaryFlexibility Field
	EarlierStart Field
	NoticeFlexibility Field
	Status Field
	Headline Field
	Notes Field
}

type JobResponse struct {
	Title  					string     `json:"title"`
	CandidateSummary    	string     `json:"candidateSummary"`
	Brief    				string     `json:"brief"`
	Type    				string     `json:"type"`
	Currency    			string     `json:"currency"`
	Salary    				string     `json:"salary"`
	SalaryUnit    			string     `json:"salaryUnit"`
	StartDate    			string     `json:"startDate"`
	Location    			string     `json:"location"`
	Keywords    			string     `json:"keywords"`
	RemoteWorking    		string     `json:"remoteWorking"`
	OtherTypes    			string     `json:"otherTypes"`
	SalaryFlexibility    	string     `json:"salaryFlexibility"`
	EarlierStart    		string     `json:"earlierStart"`
	NoticeFlexibility    	string     `json:"noticeFlexibility"`
	Status    				string     `json:"status"`
	Headline    			string     `json:"headline"`
	Notes    				string     `json:"notes"`
}

func CreateJob(job Job, data map[string] string) JobResponse {
	title 				:= data[job.Title.Field]
	candidateSummary 	:= data[job.CandidateSummary.Field]
	brief 				:= data[job.Brief.Field]
	jobType 			:= data[job.Type.Field]
	currency 			:= data[job.Currency.Field]
	salary 				:= data[job.Salary.Field]
	salaryUnit 			:= data[job.SalaryUnit.Field]
	startDate 			:= data[job.StartDate.Field]
	location 			:= data[job.Location.Field]
	keywords 			:= data[job.Keywords.Field]
	remoteWorking 		:= data[job.RemoteWorking.Field]
	otherTypes 			:= data[job.OtherTypes.Field]
	salaryFlexibility 	:= data[job.SalaryFlexibility.Field]
	earlierStart 		:= data[job.EarlierStart.Field]
	noticeFlexibility 	:= data[job.NoticeFlexibility.Field]
	status 				:= data[job.Status.Field]
	headline 			:= data[job.Headline.Field]
	notes 				:= data[job.Notes.Field]

	rep := JobResponse{}
	rep.Title 				= title
	rep.CandidateSummary 	= candidateSummary
	rep.Brief 				= brief
	rep.Type 				= jobType
	rep.Currency 			= currency
	rep.Salary 				= salary
	rep.SalaryUnit 			= salaryUnit
	rep.StartDate 			= startDate
	rep.Location 			= location
	rep.Keywords 			= keywords
	rep.RemoteWorking 		= remoteWorking
	rep.OtherTypes 			= otherTypes
	rep.SalaryFlexibility 	= salaryFlexibility
	rep.EarlierStart 		= earlierStart
	rep.NoticeFlexibility 	= noticeFlexibility
	rep.Status 				= status
	rep.Headline 			= headline
	rep.Notes 				= notes

	return rep
}
