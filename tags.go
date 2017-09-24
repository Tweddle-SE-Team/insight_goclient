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

type Tags struct {
	client *client `json:"-"`
}

func NewTags(c *client) Tags {
	return Tags{c}
}

// Structs meant for clients
type PostTag struct {
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	Sources  PostSources `json:"sources"`
	Actions  PostActions `json:"actions"`
	Patterns []string    `json:"patterns"`
	Labels   Labels      `json:"labels"`
}

type Tag struct {
	Id       string   `json:"id"`
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Sources  Sources  `json:"sources"`
	Actions  Actions  `json:"actions"`
	Patterns []string `json:"patterns"`
	Labels   Labels   `json:"labels"`
}

type PostSource struct {
	Id string `json:"id"`
}

type Source struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	RetentionPeriod string `json:"retention_period"`
	StoredDays      []int  `json:"stored_days"`
}

type PostAction struct {
	MinMatchesCount  int               `json:"min_matches_count"`
	MinReportCount   int               `json:"min_report_count"`
	MinMatchesPeriod string            `json:"min_matches_period"`
	MinReportPeriod  string            `json:"min_report_period"`
	Targets          PostTargets           `json:"targets"`
	Enabled          bool              `json:"enabled"`
	Type             string            `json:"type"`
}

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

type Target struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	ParamsSet ParamsSet `json:"params_set"`
	AlertContentSet  map[string]string `json:"alert_content_set"`
}

type PostTarget struct {
	Type      string    `json:"type"`
	ParamsSet ParamsSet `json:"params_set"`
	AlertContentSet  map[string]string `json:"alert_content_set"`
}

type ParamsSet struct {
	Direct string `json:"direct"`
	Teams  string `json:"teams"`
	Users  string `json:"users"`
}

type Label struct {
	Id       string `json:"id"`
	SN       int    `json:"sn"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Reserved bool   `json:"reserved"`
}

type Sources []Source
type PostSources []PostSource
type Actions []Action
type PostActions []PostAction
type Labels []Label
type Targets []Target
type PostTargets []PostTarget

// Structs meant for marshalling/un-marshalling purposes
type tagsCollection struct {
	Tags []Tag `json:"tags"`
}

type getTag struct {
	Tag Tag `json:"tag"`
}

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

func (t *Tags) getPath() string {
	return "management/tags"
}

func (t *Tags) getTagEndPoint(tagId string) string {
	return fmt.Sprintf("%s/%s", t.getPath(), tagId)
}