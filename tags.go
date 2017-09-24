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

// NewTags creates a new Tags struct that exposes Tags CRUD operations
func NewTags(c *client) Tags {
	return Tags{c}
}

// Structs meant for clients

// PostTag represents the entity used to post new tags to logentries API
type PostTag struct {
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	Sources  PostSources `json:"sources"`
	Actions  PostActions `json:"actions"`
	Patterns []string    `json:"patterns"`
	Labels   getLabels   `json:"labels"`
}

// PostSource represents the source log that the PostTag will be associated with
type PostSource struct {
	Id string `json:"id"`
}

// PostAction represents the entity used to associate actions (e,g: alerts) with the posted PostTag
type PostAction struct {
	MinMatchesCount  int               `json:"min_matches_count"`
	MinReportCount   int               `json:"min_report_count"`
	MinMatchesPeriod string            `json:"min_matches_period"`
	MinReportPeriod  string            `json:"min_report_period"`
	Targets          PostTargets           `json:"targets"`
	Enabled          bool              `json:"enabled"`
	Type             string            `json:"type"`
}

// Tag represents the entity used to get an existing tag from the logentries API
type Tag struct {
	Id       string    `json:"id"`
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Sources  Sources   `json:"sources"`
	Actions  Actions   `json:"actions"`
	Patterns []string  `json:"patterns"`
	Labels   getLabels `json:"labels"`
}
// Source represents the source log associated with the Tag
type Source struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	RetentionPeriod string `json:"retention_period"`
	StoredDays      []int  `json:"stored_days"`
}

// Action represents the action (e,g: alerts) associated with the given Tag
type Action struct {
	Id               string            `json:"id"`
	MinMatchesCount  int               `json:"min_matches_count"`
	MinReportCount   int               `json:"min_report_count"`
	MinMatchesPeriod string            `json:"min_matches_period"`
	MinReportPeriod  string            `json:"min_report_period"`
	Targets          Targets           `json:"targets"`
	Enabled          bool              `json:"enabled"`
	Type             string            `json:"type"`
}

// Target represents the target for the configured alarm (e,g: mailto, pagerduty)
type Target struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	ParamsSet ParamsSet `json:"params_set"`
	AlertContentSet  map[string]string `json:"alert_content_set"`
}

// PostTarget represents the target for the configured alarm (e,g: mailto, pagerduty)
type PostTarget struct {
	Type      string    `json:"type"`
	ParamsSet ParamsSet `json:"params_set"`
	AlertContentSet  map[string]string `json:"alert_content_set"`
}

// ParamsSet represents a set of parameters used to configure the target
type ParamsSet struct {
	Direct string `json:"direct"`
	Teams  string `json:"teams"`
	Users  string `json:"users"`
}

type Sources []Source
type PostSources []PostSource
type Actions []Action
type PostActions []PostAction
type Targets []Target
type PostTargets []PostTarget

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

// getPath returns the rest end point for tags
func (t *Tags) getPath() string {
	return "/management/tags"
}

// getTagEndPoint returns the rest end point to retrieve an individual tag
func (t *Tags) getTagEndPoint(tagId string) string {
	return fmt.Sprintf("%s/%s", t.getPath(), tagId)
}