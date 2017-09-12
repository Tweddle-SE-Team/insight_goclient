package logentries_goclient

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"net/http"
)

type MockResponses map[string]*MockResponse

type MockResponse struct {
	Data string
}

func mockResponse(responses MockResponses, w http.ResponseWriter, r *http.Request) {
	response := responses[r.URL.Path]
	if response != nil {
		w.Write([]byte(response.Data))
	} else {
		panic(fmt.Sprintf("No matching mock response for the given path %s", r.URL.Path))
	}
}

func testClientServer(responses MockResponses) (*http.Client, *httptest.Server){

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		mockResponse(responses, w, r)
	}))


	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: func(request *http.Request) (*url.URL, error) {
				return url.Parse(httpServer.URL)
			},
		},
	}

	return httpClient, httpServer
}
