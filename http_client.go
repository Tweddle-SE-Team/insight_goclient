package logentries_goclient

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	"fmt"
)

type HttpClient struct {
	httpClient    *http.Client
}

func (httpClient *HttpClient) Get(url string, headers map[string]string, out interface{}) (*http.Response, error) {
	if req, err := httpClient.prepareRequest(http.MethodGet, url, headers, nil); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, out)
	}
}

func (httpClient *HttpClient) Post(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error) {
	headers["Content-Type"] = "application/json"
	if req, err := httpClient.prepareRequest(http.MethodPost, url, headers, in); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, out)
	}
}

func (httpClient *HttpClient) Put(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error) {
	headers["Content-Type"] = "application/json"
	if req, err := httpClient.prepareRequest(http.MethodPut, url, headers, in); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, out)
	}
}

func (httpClient *HttpClient) Delete(url string, headers map[string]string) (*http.Response, error) {
	if req, err := httpClient.prepareRequest(http.MethodDelete, url, headers, nil); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, nil)
	}
}

func (httpClient *HttpClient) prepareRequest(method, url string, headers map[string]string, in interface{}) (*http.Request, error) {

	var body []byte
	var err error
	if in != nil {
		body, err = json.Marshal(in)
		fmt.Println(string(body))
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

func (httpClient *HttpClient) performRequest(req *http.Request, out interface{}) (*http.Response, error) {
	resp, err := httpClient.httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("request %s %s %s failed. Response Error [%s]: '%s'", req.Method, req.URL, req.Proto, resp.Status, err.Error())
	}

	if out != nil {
		var body []byte
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &out); err != nil {
			return nil, fmt.Errorf("unable to unmarshal response body ['%s'] for request = '%s %s %s'. Response = '%s'",  err.Error(), req.Method, req.URL, req.Proto, resp.Status)
		}
	}
	return resp, nil
}