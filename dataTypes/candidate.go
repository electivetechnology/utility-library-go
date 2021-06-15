package dataTypes

type CandidateAddress struct {
	AddressLine1 Field
	AddressLine2 Field
	City Field
	Postcode Field
	Country Field
	County Field
}

type CandidateJob struct {
	Title Field
	Location Field
	WillRelocate Field
	ExpectedSalary Field
	SalaryCurrency Field
	Notice Field
	NoticeUnit Field
	JobType Field
	Company Field
	Summary Field
}

type Candidate struct {
	Email Field
	FirstName Field
	LastName Field
	Phone Field
	PrimaryLanguage Field
	SecondaryLanguage Field
	TertiaryLanguage Field
	CvText Field
	Title Field
	AlternativePhoneNumber Field
	Dob Field
	VendorId Field
	VendorStatus Field
	VendorSource Field
	Gender Field
	Status Field
	CandidateAddress CandidateAddress
	CandidateJob CandidateJob
}


type CandidateAddressResponse struct {
	AddressLine1 	string     `json:"addressLine1"`
	AddressLine2 	string     `json:"addressLine2"`
	City 			string     `json:"city"`
	Postcode 		string     `json:"postcode"`
	Country 		string     `json:"country"`
	County 			string     `json:"county"`
}

type CandidateJobResponse struct {
	Title 			string     `json:"title"`
	Location 		string     `json:"location"`
	WillRelocate 	string     `json:"willRelocate"`
	ExpectedSalary 	string     `json:"expectedSalary"`
	SalaryCurrency 	string     `json:"salaryCurrency"`
	Notice 			string     `json:"notice"`
	NoticeUnit 		string     `json:"noticeUnit"`
	JobType 		string     `json:"jobType"`
	Company 		string     `json:"company"`
	Summary 		string     `json:"summary"`
}

type CandidateResponse struct {
	Email  					string   					`json:"email"`
	FirstName    			string     					`json:"firstName"`
	LastName    			string     					`json:"lastName"`
	Phone    				string     					`json:"phone"`
	PrimaryLanguage    		string     					`json:"primaryLanguage"`
	SecondaryLanguage    	string     					`json:"secondaryLanguage"`
	TertiaryLanguage    	string     					`json:"tertiaryLanguage"`
	CvText    				string     					`json:"cvText"`
	AlternativePhoneNumber  string     					`json:"alternativePhoneNumber"`
	Dob    					string     					`json:"dob"`
	VendorId    			string     					`json:"vendorId"`
	VendorStatus    		string     					`json:"vendorStatus"`
	VendorSource    		string     					`json:"vendorSource"`
	Gender    				string     					`json:"gender"`
	Status    				string     					`json:"status"`
	CandidateAddress    	CandidateAddressResponse    `json:"address"`
	CandidateJob    		CandidateJobResponse     	`json:"job"`
}

func CreateCandidate(candidate Candidate, data map[string] string) CandidateResponse {
	email 					:= data[candidate.Email.Field]
	firstName 				:= data[candidate.FirstName.Field]
	lastName 				:= data[candidate.LastName.Field]
	phone 					:= data[candidate.Phone.Field]
	primaryLanguage 		:= data[candidate.PrimaryLanguage.Field]
	secondaryLanguage 		:= data[candidate.SecondaryLanguage.Field]
	tertiaryLanguage 		:= data[candidate.TertiaryLanguage.Field]
	cvText 					:= data[candidate.CvText.Field]
	alternativePhoneNumber 	:= data[candidate.AlternativePhoneNumber.Field]
	dob 					:= data[candidate.Dob.Field]
	vendorId 				:= data[candidate.VendorId.Field]
	vendorStatus 			:= data[candidate.VendorStatus.Field]
	vendorSource 			:= data[candidate.VendorSource.Field]
	gender 					:= data[candidate.Gender.Field]
	status 					:= data[candidate.Status.Field]

	address1 				:= data[candidate.CandidateAddress.AddressLine1.Field]
	address2 				:= data[candidate.CandidateAddress.AddressLine2.Field]
	city 					:= data[candidate.CandidateAddress.City.Field]
	postcode 				:= data[candidate.CandidateAddress.Postcode.Field]
	country 				:= data[candidate.CandidateAddress.Country.Field]
	county 					:= data[candidate.CandidateAddress.County.Field]

	title 					:= data[candidate.CandidateJob.Title.Field]
	location 				:= data[candidate.CandidateJob.Location.Field]
	willRelocate 			:= data[candidate.CandidateJob.WillRelocate.Field]
	expectedSalary 			:= data[candidate.CandidateJob.ExpectedSalary.Field]
	salaryCurrency 			:= data[candidate.CandidateJob.SalaryCurrency.Field]
	notice 					:= data[candidate.CandidateJob.Notice.Field]
	noticeUnit 				:= data[candidate.CandidateJob.NoticeUnit.Field]
	jobType 				:= data[candidate.CandidateJob.JobType.Field]
	company 				:= data[candidate.CandidateJob.Company.Field]
	summary 				:= data[candidate.CandidateJob.Summary.Field]

	rep 						:= CandidateResponse{}
	rep.Email 					= email
	rep.FirstName 				= firstName
	rep.LastName 				= lastName
	rep.Phone 					= phone
	rep.PrimaryLanguage 		= primaryLanguage
	rep.SecondaryLanguage 		= secondaryLanguage
	rep.TertiaryLanguage 		= tertiaryLanguage
	rep.CvText 					= cvText
	rep.AlternativePhoneNumber 	= alternativePhoneNumber
	rep.Dob 					= dob
	rep.VendorId 				= vendorId
	rep.VendorStatus 			= vendorStatus
	rep.VendorSource 			= vendorSource
	rep.Gender 					= gender
	rep.Status 					= status

	rep.CandidateAddress.AddressLine1 			= address1
	rep.CandidateAddress.AddressLine2 			= address2
	rep.CandidateAddress.City 					= city
	rep.CandidateAddress.Postcode 				= postcode
	rep.CandidateAddress.Country 				= country
	rep.CandidateAddress.County 				= county

	rep.CandidateJob.Title 				= title
	rep.CandidateJob.Location 			= location
	rep.CandidateJob.WillRelocate 		= willRelocate
	rep.CandidateJob.ExpectedSalary 	= expectedSalary
	rep.CandidateJob.SalaryCurrency		= salaryCurrency
	rep.CandidateJob.Notice 			= notice
	rep.CandidateJob.NoticeUnit 		= noticeUnit
	rep.CandidateJob.JobType 			= jobType
	rep.CandidateJob.Company 			= company
	rep.CandidateJob.Summary 			= summary
	return rep
}