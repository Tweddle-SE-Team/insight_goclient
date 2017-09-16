package logentries_goclient

import "github.com/dikhan/logentries_goclient/testutils"

func getTestClient(requestMatcher testutils.RequestMatcher) *client {
	httpClient, httpServer := testutils.TestClientServer(requestMatcher)
	c := &client{logEntriesUrl: httpServer.URL, api_key: "apikey", httpClient: &HttpClient{httpClient} }
	return c
}