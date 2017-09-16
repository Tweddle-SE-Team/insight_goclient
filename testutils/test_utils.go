package testutils

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type ExpectedRequest struct {
	HttpMethod string
	Url string
	Payload interface{}
}

type RequestMatcher struct {
	ExpectedRequest ExpectedRequest
	Response interface{}
}

func TestNewRequestMatcher(expectedReqHttpMethod, expectedPath string, response interface{}) RequestMatcher{
	return 	RequestMatcher{
		ExpectedRequest: ExpectedRequest {
			HttpMethod: expectedReqHttpMethod,
			Url: expectedPath,
		},
		Response: response,
	}
}

func (rm *RequestMatcher) match(r *http.Request) error {
	var body []byte
	var err error
	if rm.ExpectedRequest.HttpMethod == r.Method && rm.ExpectedRequest.Url == r.URL.Path {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if len(body) == 0 {
			return nil
		} else {
			expectedRequest, err := json.Marshal(rm.ExpectedRequest.Payload)
			if err != nil {
				return err
			}
			if string(expectedRequest) == string(body) {
				return nil
			}
		}
	}
	return fmt.Errorf("NO MATCHING EXPECTED REQUEST FOUND\n- [ExpectedRequest=%s]\n- [ActualRequest=%s %s %s]\n", rm.ExpectedRequest, r.Method, r.URL.Path, string(body))
}

func matchRequest(rm RequestMatcher, w http.ResponseWriter, r *http.Request) error {
	if err := rm.match(r); err != nil {
		return err
	} else {
		resp, err := json.Marshal(rm.Response)
		if err != nil {
			return fmt.Errorf("ERROR THROWN WHILE MARHSALING THE MOCK RESPONSE [%s] - Error: %s", resp, err)
		}
		if resp != nil {
			w.Write([]byte(resp))
		}
		return nil
	}
}

func TestClientServer(requestMatcher RequestMatcher) (*http.Client, *httptest.Server){

	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if err := matchRequest(requestMatcher, w, r); err != nil {
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