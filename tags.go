package logentries_goclient

import (
	"errors"
	"fmt"
)

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
	Patterns []string    `json:"patters"`
	Labels   Labels      `json:"labels"`
}

type Tag struct {
	Id      string   `json:"id"`
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Sources Sources  `json:"sources"`
	Actions Actions  `json:"actions"`
	Patters []string `json:"patters"`
	Labels  Labels   `json:"labels"`
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
	Targets          Targets           `json:"targets"`
	AlertContentSet  map[string]string `json:"alert_content_set"`
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
	AlertContentSet  map[string]string `json:"alert_content_set"`
	Enabled          bool              `json:"enabled"`
	Type             string            `json:"type"`
}

type Target struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	ParamsSet ParamsSet `json:"params_set"`
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

// Structs meant for marshalling/un-marshalling purposes
type tagsCollection struct {
	Tags []Tag `json:"tags"`
}

type getTag struct {
	Tag Tag `json:"tag"`
}

func (t *Tags) getPath() string {
	return "management/tags"
}

func (t *Tags) GetTags() ([]Tag, error) {
	tags := &tagsCollection{}
	if err := t.client.get(t.getPath(), tags); err != nil {
		return nil, err
	}
	return tags.Tags, nil
}

func (t *Tags) GetTag(tagId string) (Tag, error) {
	if tagId == "" {
		return Tag{}, errors.New("tagId input parameter is mandatory")
	}

	tagEndPoint := fmt.Sprintf("%s/%s", t.getPath(), tagId)
	tag := &getTag{}
	if err := t.client.get(tagEndPoint, tag); err != nil {
		return Tag{}, err
	}
	return tag.Tag, nil
}

func (t *Tags) PostTag(postTag PostTag) (Tag, error) {
	tag := &getTag{}
	if err := t.client.post(t.getPath(), postTag, tag); err != nil {
		return Tag{}, err
	}
	return tag.Tag, nil
}

func (t *Tags) PutTag(tagId string, postTag PostTag) (Tag, error) {
	if tagId == "" {
		return Tag{}, errors.New("tagId input parameter is mandatory")
	}

	tagEndPoint := fmt.Sprintf("%s/%s", t.getPath(), tagId)
	tag := &getTag{}
	if err := t.client.put(tagEndPoint, postTag, tag); err != nil {
		return Tag{}, err
	}
	return tag.Tag, nil
}
