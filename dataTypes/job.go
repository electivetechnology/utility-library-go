package dataTypes

import (
	"strconv"
	"time"
)

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
	Client Client
}

type JobResponse struct {
	Title  					string     `json:"title"`
	CandidateSummary    	string     `json:"candidateSummary"`
	Brief    				string     `json:"brief"`
	Type    				string     `json:"type"`
	Currency    			string     `json:"currency"`
	Salary    				float32     `json:"salary"`
	SalaryUnit    			string     `json:"salaryUnit"`
	StartDate    			time.Time     `json:"startDate"`
	Location    			string     `json:"location"`
	Keywords    			string     `json:"keywords"`
	RemoteWorking    		bool     `json:"remoteWorking"`
	OtherTypes    			bool     `json:"otherTypes"`
	SalaryFlexibility    	int     `json:"salaryFlexibility"`
	EarlierStart    		bool     `json:"earlierStart"`
	NoticeFlexibility    	int     `json:"noticeFlexibility"`
	Status    				string     `json:"status"`
	Headline    			string     `json:"headline"`
	Notes    				string     `json:"notes"`
	Client    				Client     `json:"client"`
}

func CreateJob(job Job, data map[string] string) JobResponse {
	title 					:= data[job.Title.Field]
	candidateSummary 		:= data[job.CandidateSummary.Field]
	brief 					:= data[job.Brief.Field]
	jobType 				:= data[job.Type.Field]
	currency 				:= data[job.Currency.Field]
	salary, _				:= strconv.ParseFloat(data[job.Salary.Field], 32)
	salaryUnit 				:= data[job.SalaryUnit.Field]
	startDate, _ 			:= time.Parse(time.UnixDate, data[job.StartDate.Field])
	location 				:= data[job.Location.Field]
	keywords 				:= data[job.Keywords.Field]
	remoteWorking, _ 		:= strconv.ParseBool(data[job.RemoteWorking.Field])
	otherTypes, _ 			:= strconv.ParseBool(data[job.OtherTypes.Field])
	salaryFlexibility, _	:= strconv.Atoi(data[job.SalaryFlexibility.Field])
	earlierStart, _ 		:= strconv.ParseBool(data[job.EarlierStart.Field])
	noticeFlexibility, _ 	:= strconv.Atoi(data[job.NoticeFlexibility.Field])
	status 					:= data[job.Status.Field]
	headline 				:= data[job.Headline.Field]
	notes 					:= data[job.Notes.Field]

	rep := JobResponse{}
	rep.Title 				= title
	rep.CandidateSummary 	= candidateSummary
	rep.Brief 				= brief
	rep.Type 				= jobType
	rep.Currency 			= currency
	rep.Salary				= float32(salary)
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
