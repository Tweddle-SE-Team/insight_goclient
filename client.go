package logentries_goclient

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

const LOG_ENTRIES_API = "https://rest.logentries.com"


type LogEntriesClient struct {
	LogSets LogSets
}

func NewLogEntriesClient(apiKey string) LogEntriesClient {
	c := &client{LOG_ENTRIES_API,apiKey, &http.Client{}}
	return LogEntriesClient{
		LogSets: LogSets{c},
	}
}

type client struct {
	logEntriesUrl string
	api_key       string
	httpClient    *http.Client
}

func (c *client) prepareRequest(method, url string) (*http.Request, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", c.api_key)
	return req, nil
}

func (c *client) get(path string, data interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.logEntriesUrl, path)
	req, err := c.prepareRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK  && resp.StatusCode != http.StatusCreated  {
		return nil, fmt.Errorf("[StatusCode=%s] Response %s", resp.StatusCode, resp)
	}

	if data != nil {
		var body []byte
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		}
		if err = json.Unmarshal(body, &data); err != nil {
			return nil, err
		}
	}

	return resp, nil
}