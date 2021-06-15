package dataTypes

type Client struct {
	Name Field
	Overview Field
	Linkedin Field
	Website Field
	Introduction Field
	PrimaryContactFirstName Field
	PrimaryContactLastName Field
	PrimaryContactEmail Field
	PrimaryContactPhone Field
	Status Field
	BrandColour Field
}

type ClientResponse struct {
	Name  						string     `json:"name"`
	Overview    				string     `json:"overview"`
	Linkedin    				string     `json:"linkedin"`
	Website    					string     `json:"website"`
	Introduction    			string     `json:"introduction"`
	PrimaryContactFirstName    	string     `json:"primaryContactFirstName"`
	PrimaryContactLastName    	string     `json:"primaryContactLastName"`
	PrimaryContactEmail    		string     `json:"primaryContactEmail"`
	PrimaryContactPhone    		string     `json:"primaryContactPhone"`
	Status    					string     `json:"status"`
	BrandColour    				string     `json:"brandColour"`
}

func CreateClient(client Client, data map[string] string) ClientResponse {
	name 					:= data[client.Name.Field]
	overview 				:= data[client.Overview.Field]
	linkedin 				:= data[client.Linkedin.Field]
	website 				:= data[client.Website.Field]
	introduction 			:= data[client.Introduction.Field]
	primaryContactFirstName := data[client.PrimaryContactFirstName.Field]
	primaryContactLastName 	:= data[client.PrimaryContactLastName.Field]
	primaryContactEmail 	:= data[client.PrimaryContactEmail.Field]
	primaryContactPhone 	:= data[client.PrimaryContactPhone.Field]
	status 					:= data[client.Status.Field]
	brandColour 			:= data[client.BrandColour.Field]

	rep := ClientResponse{}
	rep.Name 					= name
	rep.Overview 				= overview
	rep.Linkedin 				= linkedin
	rep.Website 				= website
	rep.Introduction 			= introduction
	rep.PrimaryContactFirstName = primaryContactFirstName
	rep.PrimaryContactLastName 	= primaryContactLastName
	rep.PrimaryContactEmail 	= primaryContactEmail
	rep.PrimaryContactPhone 	= primaryContactPhone
	rep.Status 					= status
	rep.BrandColour 			= brandColour
	return rep
}
