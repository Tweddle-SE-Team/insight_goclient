package testutils

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"net/http"
	"encoding/json"
)

type TestClientServer struct {
	RequestMatcher TestRequestMatcher
}

func (t *TestClientServer) matchRequest(w http.ResponseWriter, r *http.Request) error {
	if err := t.RequestMatcher.match(r); err != nil {
		return err
	} else {
		resp, err := json.Marshal(t.RequestMatcher.Response.Payload)
		if err != nil {
			return fmt.Errorf("ERROR THROWN WHILE MARHSALING THE MOCK RESPONSE [%s] - Error: %s", resp, err)
		}
		if resp != nil {
			w.WriteHeader(t.RequestMatcher.Response.HttpStatusCode)
			w.Write([]byte(resp))
		}
		return nil
	}
}

func (t *TestClientServer) TestClientServer(requestMatcher TestRequestMatcher) (*http.Client, *httptest.Server){

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if err := t.matchRequest(w, r); err != nil {
			fmt.Println("UNEXPECTED ERROR THROWN - " + err.Error())
		}
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