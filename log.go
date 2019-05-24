package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	LOGS_PATH = "/management/logs"
)

// The Logs resource allows you to interact with Logs in your account. The following operations are supported:
// - Get details of an existing Log
// - Get details of a list of all Logs
// - Create a new Log
// - Update an existing Log
// - Delete a Log

// Log represents the entity used to get an existing log from the insight API
type Log struct {
	Id              string      `json:"id,omitempty"`
	Name            string      `json:"name"`
	LogsetsInfo     []Info      `json:"logsets_info"`
	UserData        LogUserData `json:"user_data"`
	Tokens          []string    `json:"tokens"`
	SourceType      string      `json:"source_type"`
	TokenSeed       string      `json:"token_seed"`
	Structures      []string    `json:"structures"`
	RetentionPeriod string      `json:"retention_period"`
	Links           []Link      `json:"links,omitempty"`
}

// LogUserData represents user metadata
type LogUserData struct {
	LogEntriesAgentFileName string `json:"le_agent_filename"`
	LogEntriesAgentFollow   string `json:"le_agent_follow"`
}

type Logs []Log

// GetLogs lists all Logs for an account
func (client *InsightClient) GetLogs() (*Logs, error) {
	var logs Logs
	if err := client.get(client.getLogEndpoint(""), &logs); err != nil {
		return nil, err
	}
	return &logs, nil
}

// GetLog gets a specific Log from an account
func (client *InsightClient) GetLog(logId string) (*Log, error) {
	var log Log
	if err := client.get(client.getLogEndpoint(logId), &log); err != nil {
		return nil, err
	}
	return &log, nil
}

// PostTag creates a new Log
func (client *InsightClient) PostLog(body Log) (*Log, error) {
	resp, err := client.post(client.getLogEndpoint(""), body)
	if err != nil {
		return nil, err
	}
	var log Log
	err = json.Unmarshal(resp, &log)
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// PutTag updates an existing Log
func (client *InsightClient) PutLog(body Log) (*Log, error) {
	resp, err := client.put(client.getLogEndpoint(body.Id), body)
	if err != nil {
		return nil, err
	}
	var log Log
	err = json.Unmarshal(resp, &log)
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// DeleteTag deletes a specific Log from an account.
func (client *InsightClient) DeleteLog(logId string) error {
	return client.delete(client.getLogEndpoint(logId))
}

func (client *InsightClient) getLogEndpoint(logId string) string {
	if logId == "" {
		return LOGS_PATH
	} else {
		return fmt.Sprintf("%s/%s", LOGS_PATH, logId)
	}
}
