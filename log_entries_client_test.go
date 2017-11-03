package logentries_goclient

import (
	"fmt"
	httpGoClient "github.com/dikhan/http_goclient"
	"github.com/dikhan/http_goclient/testutils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getTestClient(requestMatcher testutils.TestRequestMatcher) *client {
	testClientServer := testutils.TestClientServer{
		RequestMatcher: requestMatcher,
	}
	httpClient, httpServer := testClientServer.TestClientServer()
	c := &client{logEntriesUrl: httpServer.URL, api_key: "apikey", httpClient: &httpGoClient.HttpClient{httpClient}}
	return c
}

func TestLogentriesClient_NewLogEntriesClient(t *testing.T) {
	c, err := NewLogEntriesClient("apiKey")
	assert.Nil(t, err)
	assert.NotNil(t, c.Logs)
	assert.NotNil(t, c.Tags)
	assert.NotNil(t, c.LogSets)
}

func TestLogentriesClient_NewLogEntriesClientApiKeyMissing(t *testing.T) {
	_, err := NewLogEntriesClient("")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "apiKey is mandatory to initialise Logentries client")
}

type mockObject struct {
	Data string `json:"data"`
}

func TestLogentriesClient_ClientGet(t *testing.T) {
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/api/testing", nil, http.StatusOK, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	err := c.get("/api/testing", expectedResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestLogentriesClient_ClientGetResponseNotStatusOk(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/api/testing", nil, http.StatusUnauthorized, &mockObject{})
	c := getTestClient(requestMatcher)
	err := c.get("/api/testing", &mockObject{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestLogentriesClient_ClientPost(t *testing.T) {
	mockRequestPayload := &mockObject{Data: "some req data..."}
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/api/testing", mockRequestPayload, http.StatusCreated, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	err := c.post("/api/testing", mockRequestPayload, expectedResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestLogentriesClient_ClientPostResponseNotStatusCreated(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/api/testing", &mockObject{}, http.StatusUnauthorized, &mockObject{})

	c := getTestClient(requestMatcher)
	err := c.post("/api/testing", &mockObject{}, &mockObject{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestLogentriesClient_ClientPut(t *testing.T) {
	mockRequestPayload := &mockObject{Data: "some req data..."}
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, "/api/testing", mockRequestPayload, http.StatusOK, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	err := c.put("/api/testing", mockRequestPayload, expectedResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestLogentriesClient_ClientPutResponseNotStatusCreated(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, "/api/testing", &mockObject{}, http.StatusUnauthorized, &mockObject{})

	c := getTestClient(requestMatcher)
	err := c.put("/api/testing", &mockObject{}, &mockObject{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestLogentriesClient_ClientDelete(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodDelete, "/api/testing", nil, http.StatusNoContent, nil)
	c := getTestClient(requestMatcher)
	err := c.delete("/api/testing")
	assert.Nil(t, err)
}

func TestLogentriesClient_ClientGetResponseNotStatusNoContent(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodDelete, "/api/testing", nil, http.StatusUnauthorized, nil)
	c := getTestClient(requestMatcher)
	err := c.delete("/api/testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}
