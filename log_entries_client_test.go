package logentries_goclient

import (
	"github.com/dikhan/logentries_goclient/testutils"
	"testing"
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
)

func getTestClient(requestMatcher testutils.TestRequestMatcher) *client {
	testClientServer := testutils.TestClientServer {
		RequestMatcher: requestMatcher,
	}
	httpClient, httpServer := testClientServer.TestClientServer(requestMatcher)
	c := &client{logEntriesUrl: httpServer.URL, api_key: "apikey", httpClient: &HttpClient{httpClient} }
	return c
}

func TestNewLogEntriesClient(t *testing.T) {
	c, err := NewLogEntriesClient(os.Getenv("API_KEY"))
	logs, err := c.Logs.GetLogs()
	fmt.Println(err)
	fmt.Println(logs)
}

func TestNewLogEntriesClientApiKeyMissing(t *testing.T) {
	_, err := NewLogEntriesClient("")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "apiKey is mandatory to initialise Logentries client")
}