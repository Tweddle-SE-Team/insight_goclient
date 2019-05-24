package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	LABELS_PATH = "/management/labels"
)

// The Labels resource allows you to interact with Labels in your account. The following operations are supported:
// - Get details of an existing label
// - Get details of a list of all labels

// Label represents the entity used to get an existing label from the insight API
type Label struct {
	Id       string `json:"id,omitempty"`
	SN       int    `json:"sn,omitempty"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Reserved bool   `json:"reserved,omitempty"`
}

type Labels []Label

// GetLabels gets details of a list of all Labels
func (client *InsightClient) GetLabels() (*Labels, error) {
	var labels Labels
	if err := client.get(client.getLabelEndpoint(""), &labels); err != nil {
		return nil, err
	}
	return &labels, nil
}

// GetLabel gets a specific Label from an account
func (client *InsightClient) GetLabel(labelId string) (*Label, error) {
	var label Label
	if err := client.get(client.getLabelEndpoint(labelId), &label); err != nil {
		return nil, err
	}
	return &label, nil
}

// PostTag creates a new Label
func (client *InsightClient) PostLabel(body Label) (*Label, error) {
	resp, err := client.post(client.getLabelEndpoint(""), body)
	if err != nil {
		return nil, err
	}
	var label Label
	err = json.Unmarshal(resp, &label)
	if err != nil {
		return nil, err
	}
	return &label, nil
}

// PutTag updates an existing Label
func (client *InsightClient) PutLabel(body Label) (*Label, error) {
	resp, err := client.put(client.getLabelEndpoint(body.Id), body)
	if err != nil {
		return nil, err
	}
	var label Label
	err = json.Unmarshal(resp, &label)
	if err != nil {
		return nil, err
	}
	return &label, nil
}

// DeleteTag deletes a specific Label from an account.
func (client *InsightClient) DeleteLabel(labelId string) error {
	return client.delete(client.getLabelEndpoint(labelId))
}

func (client *InsightClient) getLabelEndpoint(labelId string) string {
	if labelId == "" {
		return LABELS_PATH
	} else {
		return fmt.Sprintf("%s/%s", LABELS_PATH, labelId)
	}
}
