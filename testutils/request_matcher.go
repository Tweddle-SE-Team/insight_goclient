package testutils

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExpectedRequest struct {
	HttpMethod string
	Url string
	Payload interface{}
}

type Response struct {
	HttpStatusCode int
	Payload interface{}
}

type TestRequestMatcher struct {
	ExpectedRequest ExpectedRequest
	Response Response
}

func NewRequestMatcher(expectedHttpMethod, expectedPath string, expectedPayload interface{}, responseStatusCode int, response interface{}) TestRequestMatcher {
	return 	TestRequestMatcher{
		ExpectedRequest: ExpectedRequest {
			HttpMethod: expectedHttpMethod,
			Url:        expectedPath,
			Payload:    expectedPayload,
		},
		Response: Response{responseStatusCode, response,},
	}
}

func (rm *TestRequestMatcher) match(r *http.Request) error {
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