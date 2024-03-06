package assessments

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/electivetechnology/utility-library-go/clients/connect"
)

const (
	INVITATION_TAG_PREFIX = "invitations_"
)

type InvitationResponse struct {
	ApiResponse *connect.ApiResponse
	Invitation  *Invitation
}

type Invitation struct {
	Id                    string            `json:"id"`
	AssessmentId          string            `json:"assessmentId"`
	AssessmentName        string            `json:"assessmentName"`
	Name                  string            `json:"name"`
	Email                 string            `json:"email"`
	Phone                 string            `json:"phone"`
	EngagementId          string            `json:"engagementId"`
	CandidateId           string            `json:"candidateId"`
	Status                string            `json:"status"`
	Questions             map[string]string `json:"questions"`
	ScoreTotal            string            `json:"scoreTotal"`
	ScoreTotalPercent     string            `json:"scoreTotalPercent"`
	ScoreOptimised        string            `json:"scoreOptimised"`
	ScoreOptimisedPercent string            `json:"scoreOptimisedPercent"`
	CreatedAt             string            `json:"createdAt"`
	UpdatedAt             string            `json:"updatedAt"`
	DeletedAt             string            `json:"deletedAt"`
	StartedAt             string            `json:"startedAt"`
	ArchivedAt            string            `json:"archivedAt"`
	IsSms                 bool              `json:"isSms"`
	IsWebchat             bool              `json:"isWebchat"`
	IsDialer              bool              `json:"isDialer"`
	IsWhatsApp            bool              `json:"isWhatsApp"`
	Transcripts           map[string]string `json:"transcripts"`
	Organisation          string            `json:"organisation"`
}

func (client Client) GetInvitationById(id string, token string) (InvitationResponse, error) {
	log.Printf("Will request invitation details for invitation id %s", id)

	// Generate new path replacer
	r := strings.NewReplacer(":invitation", id)
	path := r.Replace(GET_INVITATION_URL)
	log.Printf("New path generated for request %s", path)

	request, _ := http.NewRequest(http.MethodGet, client.ApiClient.GetBaseUrl()+path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(INVITATION_TAG_PREFIX + path + token)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Invitation details: %v", err)
		return InvitationResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	log.Printf("Got response from server: %v", res)
	response := InvitationResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return InvitationResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		invitation := Invitation{}
		json.Unmarshal(data, &invitation)

		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, INVITATION_TAG_PREFIX+invitation.Id)
			tags = append(tags, key)
			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.Invitation = &invitation

		return response, nil

	case http.StatusNotFound:
		msg := fmt.Sprintf("Could not find Invitation for id: %s", id)
		return response, errors.New(msg)
	default:
		msg := fmt.Sprintf("Could not get Invitation details")
		return response, errors.New(msg)
	}
}
