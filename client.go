package logentries_goclient

import (
	"net/http"
	"fmt"
)

const LOG_ENTRIES_API = "https://rest.logentries.com"


type LogEntriesClient struct {
	LogSets LogSets
}

func NewLogEntriesClient(apiKey string) LogEntriesClient {
	c := &client{LOG_ENTRIES_API,apiKey, &HttpClient{&http.Client{}}}
	return LogEntriesClient{
		LogSets: LogSets{c},
	}
}

type client struct {
	logEntriesUrl string
	api_key       string
	httpClient    *HttpClient
}

func (c *client) requestHeaders() map[string]string{
	headers := map[string]string{}
	headers["x-api-key"] = c.api_key
	return headers
}

func (c *client) get(path string, data interface{}) (*http.Response, error) {
	url := c.getLogEntriesUrl(path)
	reqHeaders := c.requestHeaders()

	resp, err := c.httpClient.Get(url, reqHeaders, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *client) post(path string, in interface{}, out interface{}) (*http.Response, error){
	url := c.getLogEntriesUrl(path)
	reqHeaders := c.requestHeaders()

	resp, err := c.httpClient.Post(url, reqHeaders, in, out)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *client) getLogEntriesUrl(path string) string {
	return fmt.Sprintf("%s/%s", c.logEntriesUrl, path)
}