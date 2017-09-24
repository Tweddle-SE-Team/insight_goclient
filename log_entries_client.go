package logentries_goclient

import (
	"fmt"
	"net/http"
)

const LOG_ENTRIES_API = "https://rest.logentries.com"

type logEntriesClient struct {
	LogSets LogSets
	Logs Logs
	Tags    Tags
}

func NewLogEntriesClient(apiKey string) (logEntriesClient, error) {
	if apiKey == "" {
		return logEntriesClient{}, fmt.Errorf("apiKey is mandatory to initialise Logentries client")
	}

	c := &client{LOG_ENTRIES_API, apiKey, &HttpClient{&http.Client{}}}
	return logEntriesClient{
		LogSets: NewLogSets(c),
		Logs: NewLogs(c),
		Tags:    NewTags(c),
	}, nil
}

type client struct {
	logEntriesUrl string
	api_key       string
	httpClient    *HttpClient
}

func (c *client) requestHeaders() map[string]string {
	headers := map[string]string{}
	headers["x-api-key"] = c.api_key
	return headers
}

func (c *client) get(path string, data interface{}) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.Get(url, c.requestHeaders(), data)
	return checkResponseStatusCode(res, err, http.StatusOK)
}

func (c *client) post(path string, in interface{}, out interface{}) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.Post(url, c.requestHeaders(), in, out)
	return checkResponseStatusCode(res, err, http.StatusCreated)
}

func (c *client) put(path string, in interface{}, out interface{}) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.Put(url, c.requestHeaders(), in, out)
	return checkResponseStatusCode(res, err, http.StatusOK)
}

func (c *client) delete(path string) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.Delete(url, c.requestHeaders())
	return checkResponseStatusCode(res, err, http.StatusNoContent)
}

func checkResponseStatusCode(res *http.Response, err error, expectedResponseStatusCode int) error {
	if err != nil {
		return fmt.Errorf("\nReceived unexpected error response: '%s'", err.Error())
	}
	if res.StatusCode != expectedResponseStatusCode {
		return fmt.Errorf("\nReceived a non expected response status code %d, expected code was %d. Response: %s", res.StatusCode, expectedResponseStatusCode, res)
	}
	return nil
}

func (c *client) getLogEntriesUrl(path string) string {
	return fmt.Sprintf("%s/%s", c.logEntriesUrl, path)
}
