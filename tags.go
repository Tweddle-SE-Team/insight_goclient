package logentries_goclient

import (
	"errors"
	"fmt"
)

// The Tags resource allows you to interact with Tags in your account. The following operations are supported:
// - Get details of an existing Tag and Alert
// - Get details of a list of all Tags and Alerts
// - Create a new Tag and Alert
// - Update an existing Tag and Alert

// Tags represents the tags interface by which user can interact with logentries tags API
type Tags struct {
	client *client `json:"-"`
}

// newTags creates a new Tags struct that exposes Tags CRUD operations
func newTags(c *client) Tags {
	return Tags{c}
}

// Structs meant for clients

// PostTag represents the entity used to post new tags to logentries API
type PostTag struct {
	Type     string       `json:"type"`
	Name     string       `json:"name"`
	Sources  []PostSource `json:"sources"`
	Actions  []PostAction `json:"actions"`
	Patterns []string     `json:"patterns"`
	Labels   GetLabels    `json:"labels"`
}

// PostSource represents the source log that the PostTag will be associated with
type PostSource struct {
	Id string `json:"id"`
}

// PostAction represents the entity used to associate actions (e,g: alerts) with the posted PostTag
type PostAction struct {
	MinMatchesCount  int          `json:"min_matches_count" mapstructure:"min_matches_count"`
	MinReportCount   int          `json:"min_report_count" mapstructure:"min_report_count"`
	MinMatchesPeriod string       `json:"min_matches_period" mapstructure:"min_matches_period"`
	MinReportPeriod  string       `json:"min_report_period" mapstructure:"min_report_period"`
	Targets          []PostTarget `json:"targets"`
	Enabled          bool         `json:"enabled"`
	Type             string       `json:"type"`
}

// Tag represents the entity used to get an existing tag from the logentries API
type Tag struct {
	Id       string    `json:"id"`
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Sources  []source  `json:"sources"`
	Actions  []action  `json:"actions"`
	Patterns []string  `json:"patterns"`
	Labels   GetLabels `json:"labels"`
}

// source represents the source log associated with the Tag
type source struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	RetentionPeriod string `json:"retention_period"`
	StoredDays      []int  `json:"stored_days"`
}

// action represents the action (e,g: alerts) associated with the given Tag
type action struct {
	Id               string   `json:"id"`
	MinMatchesCount  int      `json:"min_matches_count"`
	MinReportCount   int      `json:"min_report_count"`
	MinMatchesPeriod string   `json:"min_matches_period"`
	MinReportPeriod  string   `json:"min_report_period"`
	Targets          []target `json:"targets"`
	Enabled          bool     `json:"enabled"`
	Type             string   `json:"type"`
}

// target represents the target for the configured alarm (e,g: mailto, pagerduty)
type target struct {
	Id              string            `json:"id"`
	Type            string            `json:"type"`
	ParamsSet       map[string]string `json:"params_set"`
	AlertContentSet map[string]string `json:"alert_content_set"`
}

// PostTarget represents the target for the configured alarm (e,g: mailto, pagerduty)
type PostTarget struct {
	Type            string            `json:"type"`
	ParamsSet       map[string]string `json:"params_set" mapstructure:"params_set"`
	AlertContentSet map[string]string `json:"alert_content_set" mapstructure:"alert_content_set"`
}

// Structs meant for marshalling/un-marshalling purposes

// tagsCollection represents a wrapper struct for marshalling/unmarshalling purposes
type tagsCollection struct {
	Tags []Tag `json:"tags"`
}

// getTag represents a wrapper struct for marshalling/unmarshalling purposes
type getTag struct {
	Tag Tag `json:"tag"`
}

// postTag represents a wrapper struct for marshalling/unmarshalling purposes
type postTag struct {
	PostTag PostTag `json:"tag"`
}

// GetTags gets details of an existing Tag and Alert
func (t *Tags) GetTags() ([]Tag, error) {
	tags := &tagsCollection{}
	if err := t.client.get(t.getPath(), tags); err != nil {
		return nil, err
	}
	return tags.Tags, nil
}

// GetTag gets details of a list of all Tags and Alerts
func (t *Tags) GetTag(tagId string) (Tag, error) {
	if tagId == "" {
		return Tag{}, errors.New("tagId input parameter is mandatory")
	}
	tag := &getTag{}
	if err := t.client.get(t.getTagEndPoint(tagId), tag); err != nil {
		return Tag{}, err
	}
	return tag.Tag, nil
}

// PostTag creates a new Tag and Alert
func (t *Tags) PostTag(p PostTag) (Tag, error) {
	tag := &getTag{}
	postTag := postTag{p}
	if err := t.client.post(t.getPath(), postTag, tag); err != nil {
		return Tag{}, err
	}
	return tag.Tag, nil
}

// PutTag updates an existing Tag and Alert
func (t *Tags) PutTag(tagId string, p PostTag) (Tag, error) {
	if tagId == "" {
		return Tag{}, errors.New("tagId input parameter is mandatory")
	}

	tag := &getTag{}
	postTag := postTag{p}
	if err := t.client.put(t.getTagEndPoint(tagId), postTag, tag); err != nil {
		return Tag{}, err
	}
	return tag.Tag, nil
}

// DeleteTag deletes a specific Tag from an account.
func (t *Tags) DeleteTag(tagId string) error {
	if tagId == "" {
		return errors.New("tagId input parameter is mandatory")
	}
	var err error
	if err = t.client.delete(t.getTagEndPoint(tagId)); err != nil {
		return err
	}
	return nil
}

// getPath returns the rest end point for tags
func (t *Tags) getPath() string {
	return "/management/tags"
}

// getTagEndPoint returns the rest end point to retrieve an individual tag
func (t *Tags) getTagEndPoint(tagId string) string {
	return fmt.Sprintf("%s/%s", t.getPath(), tagId)
}
