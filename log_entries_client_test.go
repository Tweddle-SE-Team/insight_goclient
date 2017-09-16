package logentries_goclient

import "github.com/dikhan/logentries_goclient/testutils"

func getTestClient(requestMatcher testutils.TestRequestMatcher) *client {
	testClientServer := testutils.TestClientServer {
		RequestMatcher: requestMatcher,
	}
	httpClient, httpServer := testClientServer.TestClientServer(requestMatcher)
	c := &client{logEntriesUrl: httpServer.URL, api_key: "apikey", httpClient: &HttpClient{httpClient} }
	return c
}