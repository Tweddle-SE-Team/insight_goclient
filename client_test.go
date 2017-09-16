package logentries_goclient

func getTestClient(path string, mockResponse *MockResponse) *client {
	mockResponses := MockResponses{path: mockResponse}
	httpClient, httpServer := testClientServer(mockResponses)

	c := &client{logEntriesUrl: httpServer.URL, api_key: "apikey", httpClient: &HttpClient{httpClient} }
	return c
}