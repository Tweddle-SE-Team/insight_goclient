package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	LOGSETS_PATH = "/management/logsets"
)

// The Log Set resource allows you to interact with Logs in your account. The following operations are supported:
// - Get details of an existing Log Set
// - Get details of a list of all Log Sets
// - Create a new Log Set
// - Update an existing Log Set
// - Delete a Log Set
// Structs meant for clients

// PostLogset represents the entity used to create a new logset to the insight API
type Logset struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	LogsInfo    []Info `json:"logs_info,omitempty"`
}

// LogsetInfo represent information about the logset
type Info struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Logsets []Logset

// GetLogset gets details of a list of all Log Sets
func (client *InsightClient) GetLogsets() (*Logsets, error) {
	var logsets Logsets
	if err := client.get(client.getLogsetEndpoint(""), &logsets); err != nil {
		return nil, err
	}
	return &logsets, nil
}

// GetLogsets gets details of an existing Log Set
func (client *InsightClient) GetLogset(logsetId string) (*Logset, error) {
	var logset Logset
	if err := client.get(client.getLogsetEndpoint(logsetId), &logset); err != nil {
		return nil, err
	}
	return &logset, nil
}

// PostLogset creates a new LogSet
func (client *InsightClient) PostLogset(body Logset) (*Logset, error) {
	resp, err := client.post(client.getLogEndpoint(""), body)
	if err != nil {
		return nil, err
	}
	var logset Logset
	err = json.Unmarshal(resp, &logset)
	if err != nil {
		return nil, err
	}
	return &logset, nil
}

// PutTag updates an existing Logset
func (client *InsightClient) PutLogset(body Logset) (*Logset, error) {
	resp, err := client.put(client.getLogsetEndpoint(body.Id), body)
	if err != nil {
		return nil, err
	}
	var logset Logset
	err = json.Unmarshal(resp, &logset)
	if err != nil {
		return nil, err
	}
	return &logset, nil
}

// DeleteTag deletes a specific Logset from an account.
func (client *InsightClient) DeleteLogset(logsetId string) error {
	return client.delete(client.getLogsetEndpoint(logsetId))
}

// getLogEndpoint returns the rest end point to retrieve an individual log
func (client *InsightClient) getLogsetEndpoint(logsetId string) string {
	if logsetId == "" {
		return LOGSETS_PATH
	} else {
		return fmt.Sprintf("%s/%s", LOGSETS_PATH, logsetId)
	}
}
