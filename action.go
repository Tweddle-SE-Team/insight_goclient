package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	ACTIONS_PATH = "/management/actions"
)

// The Actions resource allows you to interact with Actions in your account. The following operations are supported:
// - Get details of an existing action
// - Get details of a list of all actions

// Action represents the entity used to get an existing action from the insight API
type Action struct {
	Id               string   `json:"id,omitempty"`
	MinMatchesCount  int      `json:"min_matches_count"`
	MinReportCount   int      `json:"min_report_count"`
	MinMatchesPeriod string   `json:"min_matches_period"`
	MinReportPeriod  string   `json:"min_report_period"`
	Targets          []Target `json:"targets"`
	Enabled          bool     `json:"enabled"`
	Type             string   `json:"type"`
}

type Actions []Action

// GetActions gets details of a list of all Actions
func (client *InsightClient) GetActions() (*Actions, error) {
	var actions Actions
	if err := client.get(ACTIONS_PATH, &actions); err != nil {
		return nil, err
	}
	return &actions, nil
}

// GetAction gets a specific Action from an account
func (client *InsightClient) GetAction(actionId string) (*Action, error) {
	var action Action
	endpoint, err := client.getActionEndpoint(actionId)
	if err != nil {
		return nil, err
	}
	if err := client.get(endpoint, &action); err != nil {
		return nil, err
	}
	return &action, nil
}

// PostTag creates a new Action
func (client *InsightClient) PostAction(action *Action) error {
	resp, err := client.post(ACTIONS_PATH, action)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &action)
	if err != nil {
		return err
	}
	return nil
}

// PutTag updates an existing Action
func (client *InsightClient) PutAction(action *Action) error {
	endpoint, err := client.getActionEndpoint(action.Id)
	if err != nil {
		return err
	}
	resp, err := client.put(endpoint, action)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &action)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a specific Action from an account.
func (client *InsightClient) DeleteAction(actionId string) error {
	endpoint, err := client.getActionEndpoint(actionId)
	if err != nil {
		return err
	}
	return client.delete(endpoint)
}

func (client *InsightClient) getActionEndpoint(actionId string) (string, error) {
	if actionId == "" {
		return "", fmt.Errorf("actionId input parameter is mandatory")
	} else {
		return fmt.Sprintf("%s/%s", ACTIONS_PATH, actionId), nil
	}
}
