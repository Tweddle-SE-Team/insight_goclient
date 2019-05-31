package insight_goclient

import (
	"encoding/json"
	"fmt"
)

const (
	TAGS_PATH = "/management/tags"
)

// The Tags resource allows you to interact with Tags in your account. The following operations are supported:
// - Get details of an existing Tag and Alert
// - Get details of a list of all Tags and Alerts
// - Create a new Tag and Alert
// - Update an existing Tag and Alert

// Tag represents the entity used to get an existing tag from the insight API
type Tag struct {
	Id       string   `json:"id,omitempty"`
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Sources  []Source `json:"sources"`
	Actions  []Action `json:"actions"`
	Patterns []string `json:"patterns"`
	Labels   Labels   `json:"labels"`
}

// source represents the source log associated with the Tag
type Source struct {
	Id              string `json:"id"`
	Name            string `json:"name,omitempty"`
	RetentionPeriod string `json:"retention_period,omitempty"`
	StoredDays      []int  `json:"stored_days"`
}

type Tags []Tag

// GetTags gets details of an existing Tag and Alert
func (client *InsightClient) GetTags() (*Tags, error) {
	var tags Tags
	if err := client.get(TAGS_PATH, &tags); err != nil {
		return nil, err
	}
	return &tags, nil
}

// GetTag gets details of a list of all Tags and Alerts
func (client *InsightClient) GetTag(tagId string) (*Tag, error) {
	var tag Tag
	endpoint, err := client.getTagEndpoint(tagId)
	if err != nil {
		return nil, err
	}
	if err := client.get(endpoint, &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

// PostTag creates a new Tag and Alert
func (client *InsightClient) PostTag(body Tag) (*Tag, error) {
	resp, err := client.post(TAGS_PATH, body)
	if err != nil {
		return nil, err
	}
	var tag Tag
	err = json.Unmarshal(resp, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// PutTag updates an existing Tag and Alert
func (client *InsightClient) PutTag(tag *Tag) error {
	endpoint, err := client.getTagEndpoint(tag.Id)
	if err != nil {
		return err
	}
	resp, err := client.put(endpoint, tag)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &tag)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag deletes a specific Tag from an account.
func (client *InsightClient) DeleteTag(tagId string) error {
	endpoint, err := client.getTagEndpoint(tagId)
	if err != nil {
		return err
	}
	return client.delete(endpoint)
}

// getTagEndPoint returns the rest end point to retrieve an individual tag
func (client *InsightClient) getTagEndpoint(tagId string) (string, error) {
	if tagId == "" {
		return "", fmt.Errorf("tagId input parameter is mandatory")
	} else {
		return fmt.Sprintf("%s/%s", TAGS_PATH, tagId), nil
	}
}
