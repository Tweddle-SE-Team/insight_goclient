package testutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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
			return fmt.Errorf("error thrown while marshalling the mock repsonse [%s] - Error: %s", resp, err)
		}
		if resp != nil {
			w.WriteHeader(t.RequestMatcher.Response.HttpStatusCode)
			w.Write([]byte(resp))
		}
		return nil
	}
}

func (t *TestClientServer) TestClientServer(requestMatcher TestRequestMatcher) (*http.Client, *httptest.Server) {

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if err := t.matchRequest(w, r); err != nil {
			fmt.Println("unexpected error thrown - " + err.Error())
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
