package insight_goclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getTestClient(requestMatcher testutils.TestRequestMatcher) *client {
	testClientServer := testutils.TestClientServer{
		RequestMatcher: requestMatcher,
	}
	httpClient, httpServer := testClientServer.TestClientServer()
	c := &client{insightUrl: httpServer.URL, api_key: "apikey", httpClient: &httpGoClient.HttpClient{httpClient}}
	return c
}

func TestInsightClient_NewInsightClient(t *testing.T) {
	c, err := NewInsightClient("apiKey")
	assert.Nil(t, err)
	assert.NotNil(t, c.Logs)
	assert.NotNil(t, c.Tags)
	assert.NotNil(t, c.Logsets)
}

func TestInsightClient_NewInsightClientApiKeyMissing(t *testing.T) {
	_, err := NewInsightClient("")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "apiKey is mandatory to initialise Insight client")
}

type mockObject struct {
	Data string `json:"data"`
}

func TestInsightClient_ClientGet(t *testing.T) {
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/api/testing", nil, http.StatusOK, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	err := c.get("/api/testing", expectedResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightClient_ClientGetResponseNotStatusOk(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/api/testing", nil, http.StatusUnauthorized, &mockObject{})
	c := getTestClient(requestMatcher)
	err := c.get("/api/testing", &mockObject{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightClient_ClientPost(t *testing.T) {
	mockRequestPayload := &mockObject{Data: "some req data..."}
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/api/testing", mockRequestPayload, http.StatusCreated, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	err := c.post("/api/testing", mockRequestPayload, expectedResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightClient_ClientPostResponseNotStatusCreated(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/api/testing", &mockObject{}, http.StatusUnauthorized, &mockObject{})

	c := getTestClient(requestMatcher)
	err := c.post("/api/testing", &mockObject{}, &mockObject{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightClient_ClientPut(t *testing.T) {
	mockRequestPayload := &mockObject{Data: "some req data..."}
	mockResponse := &mockObject{Data: "some data..."}
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, "/api/testing", mockRequestPayload, http.StatusOK, mockResponse)

	c := getTestClient(requestMatcher)
	expectedResponse := &mockObject{}
	err := c.put("/api/testing", mockRequestPayload, expectedResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.Data, mockResponse.Data)
}

func TestInsightClient_ClientPutResponseNotStatusCreated(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, "/api/testing", &mockObject{}, http.StatusUnauthorized, &mockObject{})

	c := getTestClient(requestMatcher)
	err := c.put("/api/testing", &mockObject{}, &mockObject{})

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}

func TestInsightClient_ClientDelete(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodDelete, "/api/testing", nil, http.StatusNoContent, nil)
	c := getTestClient(requestMatcher)
	err := c.delete("/api/testing")
	assert.Nil(t, err)
}

func TestInsightClient_ClientGetResponseNotStatusNoContent(t *testing.T) {
	requestMatcher := testutils.NewRequestMatcher(http.MethodDelete, "/api/testing", nil, http.StatusUnauthorized, nil)
	c := getTestClient(requestMatcher)
	err := c.delete("/api/testing")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("Received a non expected response status code %d", http.StatusUnauthorized))
}
