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
	if err := client.get(LABELS_PATH, &labels); err != nil {
		return nil, err
	}
	return &labels, nil
}

// GetLabel gets a specific Label from an account
func (client *InsightClient) GetLabel(labelId string) (*Label, error) {
	var label Label
	endpoint, err := client.getLabelEndpoint(labelId)
	if err != nil {
		return nil, err
	}
	if err := client.get(endpoint, &label); err != nil {
		return nil, err
	}
	return &label, nil
}

// GetLabel gets a specific Label from an account by name
func (client *InsightClient) GetLabelsByName(name, color string) (*Labels, error) {
	var result Labels
	labels, err := client.GetLabels()
	if err != nil {
		return nil, err
	}
	for _, label := range *labels {
		if label.Name == name {
			if color != "" && label.Color != color {
				continue
			}
			result = append(result, label)
		}
	}
	return &result, nil
}

// PostTag creates a new Label
func (client *InsightClient) PostLabel(label *Label) error {
	resp, err := client.post(LABELS_PATH, label)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &label)
	if err != nil {
		return err
	}
	return nil
}

// PutTag updates an existing Label
func (client *InsightClient) PutLabel(label *Label) error {
	endpoint, err := client.getLabelEndpoint(label.Id)
	if err != nil {
		return err
	}
	resp, err := client.put(endpoint, label)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &label)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a specific Label from an account.
func (client *InsightClient) DeleteLabel(labelId string) error {
	endpoint, err := client.getLabelEndpoint(labelId)
	if err != nil {
		return err
	}
	return client.delete(endpoint)
}

func (client *InsightClient) getLabelEndpoint(labelId string) (string, error) {
	if labelId == "" {
		return "", fmt.Errorf("labelId input parameter is mandatory")
	} else {
		return fmt.Sprintf("%s/%s", LABELS_PATH, labelId), nil
	}
}
