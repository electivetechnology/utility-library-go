package assessments

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	assert.Equal(t, CLIENT_NAME, client.GetApiClient().GetName())
	assert.NotEmpty(t, client.GetApiClient().GetId())
	assert.Equal(t, BASE_URL, client.GetApiClient().GetBaseUrl())
	assert.False(t, client.GetApiClient().IsCacheEnabled())

	// Test env read
	basUrl := "assessments"
	isEnabled := true
	ttl := 1200
	os.Setenv(HOST_ENV, basUrl)
	os.Setenv(CLIENT_ENABLED_ENV, strconv.FormatBool(isEnabled))
	os.Setenv(CACHE_TTL_ENV, strconv.Itoa(ttl))

	// Test client
	clientTwo := NewClient()
	assert.Equal(t, basUrl, clientTwo.GetApiClient().GetBaseUrl())
	assert.Equal(t, isEnabled, clientTwo.GetApiClient().IsEnabled())
	assert.Equal(t, ttl, clientTwo.GetApiClient().GetRedisTTL())
}
