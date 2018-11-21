// Package logentries_goclient provides a logentries client which allows the interaction with logentries rest API
// via the seamless resource interfaces exposed. Examples include:
// - LogSets
// - Logs
// - Tags
// - Labels
package logentries_goclient

import (
	"fmt"
	httpGoClient "github.com/dikhan/http_goclient"
	"io/ioutil"
	"net/http"
)

const LOG_ENTRIES_API = "https://rest.logentries.com"

type LogEntriesClient struct {
	LogSets LogSets
	Logs    Logs
	Tags    Tags
	Labels  Labels
}

// NewLogEntriesClient creates a logentries client which exposes an interface with CRUD operations for each of the
// resources provided by logentries rest API
func NewLogEntriesClient(apiKey string) (LogEntriesClient, error) {
	if apiKey == "" {
		return LogEntriesClient{}, fmt.Errorf("apiKey is mandatory to initialise Logentries client")
	}
	client := &httpGoClient.HttpClient{&http.Client{}}
	return newLogEntriesClient(apiKey, client)
}

func newLogEntriesClient(apiKey string, httpClient *httpGoClient.HttpClient) (LogEntriesClient, error) {
	c := &client{LOG_ENTRIES_API, apiKey, httpClient}
	return LogEntriesClient{
		LogSets: newLogSets(c),
		Logs:    newLogs(c),
		Tags:    newTags(c),
		Labels:  newLabels(c),
	}, nil
}

type client struct {
	logEntriesUrl string
	api_key       string
	httpClient    *httpGoClient.HttpClient
}

func (c *client) requestHeaders() map[string]string {
	headers := map[string]string{}
	headers["x-api-key"] = c.api_key
	return headers
}

func (c *client) get(path string, in interface{}) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.Get(url, c.requestHeaders(), in)
	return checkResponseStatusCode(res, err, http.StatusOK)
}

func (c *client) post(path string, in interface{}, out interface{}) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.PostJson(url, c.requestHeaders(), in, out)
	return checkResponseStatusCode(res, err, http.StatusCreated)
}

func (c *client) put(path string, in interface{}, out interface{}) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.PutJson(url, c.requestHeaders(), in, out)
	return checkResponseStatusCode(res, err, http.StatusOK)
}

func (c *client) deleteWithStatus(path string, customStatus int) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.Delete(url, c.requestHeaders())
	return checkResponseStatusCode(res, err, customStatus)
}

func (c *client) delete(path string) error {
	url := c.getLogEntriesUrl(path)

	res, err := c.httpClient.Delete(url, c.requestHeaders())
	return checkResponseStatusCode(res, err, http.StatusNoContent)
}

func (c *client) getLogEntriesUrl(path string) string {
	return fmt.Sprintf("%s%s", c.logEntriesUrl, path)
}

func checkResponseStatusCode(res *http.Response, err error, expectedResponseStatusCode int) error {
	if err != nil {
		return fmt.Errorf("\nReceived unexpected error response: '%s'", err.Error())
	}
	defer res.Body.Close()
	var bodyBytes []byte
	bodyBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("\nReceived unexpected error response: '%s'", err.Error())
	}
	bodyString := string(bodyBytes)
	if res.StatusCode != expectedResponseStatusCode {
		return fmt.Errorf("\nReceived a non expected response status code %d, expected code was %d. Response: %s", res.StatusCode, expectedResponseStatusCode, bodyString)
	}
	return nil
}
