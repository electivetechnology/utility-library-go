package assessments

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJobByVendor(t *testing.T) {
	// Start a local HTTP server
	id := "def123"
	title := "Test Job"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/v1/jobs/vendor/bullhorn/abc123")
		// Send response to be tested
		rw.Write([]byte(`{"id":"` + id + `", "title":"` + title + `"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	fmt.Printf("Server %v", server)

	// Test env read
	isEnabled := true
	ttl := 0
	os.Setenv(CLIENT_ENABLED_ENV, strconv.FormatBool(isEnabled))
	os.Setenv(CACHE_TTL_ENV, strconv.Itoa(ttl))
	os.Setenv(HOST_ENV, server.URL)

	assessmentClient := NewClient()
	assessmentClient.GetApiClient().SetHttpClient(server.Client())

	ret, err := assessmentClient.GetJobByVendor("bullhorn", "abc123", "abc")

	assert.Equal(t, id, ret.Job.Id)
	assert.Equal(t, title, ret.Job.Title)
	assert.Nil(t, err)
}
