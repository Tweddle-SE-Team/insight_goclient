package logentries_goclient

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bytes"
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
	if req, err := httpClient.prepareRequest(http.MethodPost, url, headers, in); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, out)
	}
}

func (httpClient *HttpClient) prepareRequest(method, url string, headers map[string]string, in interface{}) (*http.Request, error) {

	var body []byte
	var err error
	if in != nil {
		body, err = json.Marshal(in)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBufferString(string(body)))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

func (httpClient *HttpClient) performRequest(request *http.Request, out interface{}) (*http.Response, error) {
	resp, err := httpClient.httpClient.Do(request)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK  && resp.StatusCode != http.StatusCreated  {
		return nil, fmt.Errorf("[StatusCode=%s] Response %s", resp.StatusCode, resp)
	}

	if out != nil {
		var body []byte
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &out); err != nil {
			return nil, err
		}
	}
	return resp, nil
}