package dataTypes

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
	AddressLine1 Field
	AddressLine2 Field
	City Field
	Postcode Field
	Country Field
	County Field
	JobTitle Field
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

type CandidateAddressResponse struct {
	AddressLine1 	string     `json:"addressLine1"`
	AddressLine2 	string     `json:"addressLine2"`
	City 			string     `json:"city"`
	Postcode 		string     `json:"postcode"`
	Country 		string     `json:"country"`
	County 			string     `json:"county"`
}

type CandidateJobResponse struct {
	JobTitle 			string     `json:"title"`
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

	address1 				:= data[candidate.AddressLine1.Field]
	address2 				:= data[candidate.AddressLine2.Field]
	city 					:= data[candidate.City.Field]
	postcode 				:= data[candidate.Postcode.Field]
	country 				:= data[candidate.Country.Field]
	county 					:= data[candidate.County.Field]

	title 					:= data[candidate.JobTitle.Field]
	location 				:= data[candidate.Location.Field]
	willRelocate 			:= data[candidate.WillRelocate.Field]
	expectedSalary 			:= data[candidate.ExpectedSalary.Field]
	salaryCurrency 			:= data[candidate.SalaryCurrency.Field]
	notice 					:= data[candidate.Notice.Field]
	noticeUnit 				:= data[candidate.NoticeUnit.Field]
	jobType 				:= data[candidate.JobType.Field]
	company 				:= data[candidate.Company.Field]
	summary 				:= data[candidate.Summary.Field]

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

	rep.CandidateJob.JobTitle 			= title
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
