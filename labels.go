package logentries_goclient

import (
	"errors"
	"fmt"
)

// The Labels resource allows you to interact with Labels in your account. The following operations are supported:
// - Get details of an existing label
// - Get details of a list of all labels

// Labels represents the label interface by which user can interact with logentries labels API
type Labels struct {
	client *client `json:"-"`
}

// newLabels creates a new Labels struct that exposes Labels CRUD operations
func newLabels(c *client) Labels {
	return Labels{c}
}

// Structs meant for clients

// Label represents the entity used to get an existing label from the logentries API
type Label struct {
	Id       string `json:"id"`
	SN       int    `json:"sn"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Reserved bool   `json:"reserved"`
}

// Label represents the entity used to get an existing label from the logentries API
type PostLabel struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type GetLabels []Label

// Structs meant for marshalling/un-marshalling purposes

// labelsCollection represents a wrapper struct for marshalling/unmarshalling purposes
type labelsCollection struct {
	Labels []Label `json:"labels"`
}

// getLabel represents a wrapper struct for marshalling/unmarshalling purposes
type getLabel struct {
	Label Label `json:"label"`
}

type postLabel struct {
	Label PostLabel `json:"label"`
}

// GetLabels gets details of a list of all Labels
func (l *Labels) GetLabels() ([]Label, error) {
	labels := &labelsCollection{}
	if err := l.client.get(l.getPath(), labels); err != nil {
		return nil, err
	}
	return labels.Labels, nil
}

// GetLabel gets details of an existing Label
func (l *Labels) GetLabel(labelId string) (Label, error) {
	if labelId == "" {
		return Label{}, errors.New("labelId input parameter is mandatory")
	}
	label := &getLabel{}
	if err := l.client.get(l.getLabelEndPoint(labelId), label); err != nil {
		return Label{}, err
	}
	return label.Label, nil
}

// DeleteLabel delete existing label
func (l *Labels) DeleteLabel(labelId string) error {
	if labelId == "" {
		return errors.New("labelId input parameter is mandatory")
	}
	var err error
	if err = l.client.delete(l.getLabelEndPoint(labelId)); err != nil {
		return err
	}
	return nil
}

// PostLabel create a new label
func (l *Labels) PostLabel(p PostLabel) (Label, error) {
	getlabel := &getLabel{}
	postlabel := &postLabel{p}

	if err := l.client.post(l.getPath(), postlabel, getlabel); err != nil {
		return Label{}, err
	}
	return getlabel.Label, nil
}

// getPath returns the rest end point for labels
func (l *Labels) getPath() string {
	return "/management/labels"
}

// getLabelEndPoint returns the rest end point to retrieve an individual labels
func (l *Labels) getLabelEndPoint(labelId string) string {
	return fmt.Sprintf("%s/%s", l.getPath(), labelId)
}
