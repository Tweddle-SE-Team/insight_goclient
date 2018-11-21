package logentries_goclient

import (
	"errors"
	"fmt"
)

// The Log Set resource allows you to interact with Logs in your account. The following operations are supported:
// - Get details of an existing Log Set
// - Get details of a list of all Log Sets
// - Create a new Log Set
// - Update an existing Log Set
// - Delete a Log Set

// LogSets represents the logsets interface by which user can interact with logentries logsets API
type LogSets struct {
	client *client `json:"-"`
}

// newLogSets creates a new LogSets struct that exposes LogSets CRUD operations
func newLogSets(c *client) LogSets {
	return LogSets{c}
}

// Structs meant for clients

// PostLogSet represents the entity used to create a new logset to the logentries API
type PostLogSet struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	UserData    map[string]string `json:"user_data,omitempty"`
	LogsInfo    []PostLogInfo     `json:"logs_info,omitempty"`
}

type PostLogSetInfo struct {
	Id string `json:"id"`
}

// PostLogSet represents the entity used to update an existing logset to the logentries API
type PutLogSet struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	UserData    map[string]string `json:"user_data,omitempty"`
	LogsInfo    []LogInfo         `json:"logs_info,omitempty"`
}

// LogSet represents the entity used to get an existing log from the logentries API
type LogSet struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	UserData    map[string]string `json:"user_data"`
	LogsInfo    []LogInfo         `json:"logs_info"`
}

// userData represents user metadata
type userData struct {
	LogEntriesDistName string `json:"le_distname"`
	LogEntriesDistVer  string `json:"le_distver"`
	LogEntriesNameIntr string `json:"le_nameintr"`
}

// LogSetInfo represent information about the logset
type LogSetInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Links []link `json:"links"`
}

// link represents link metadata of the logset
type link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// Structs meant for marshalling/un-marshalling purposes

// logSetCollection represents a wrapper struct for marshalling/unmarshalling purposes
type logSetCollection struct {
	LogSets []LogSet `json:"logsets"`
}

// getLogSet represents a wrapper struct for marshalling/unmarshalling purposes
type getLogSet struct {
	LogSet LogSet `json:"logset"`
}

// postLogSet represents a wrapper struct for marshalling/unmarshalling purposes
type postLogSet struct {
	PostLogSet PostLogSet `json:"logset"`
}

// putLogSet represents a wrapper struct for marshalling/unmarshalling purposes
type putLogSet struct {
	PutLogSet PutLogSet `json:"logset"`
}

// GetLogSet gets details of a list of all Log Sets
func (l *LogSets) GetLogSets() ([]LogSet, error) {
	logSets := &logSetCollection{}
	if err := l.client.get(l.getPath(), logSets); err != nil {
		return nil, err
	}
	return logSets.LogSets, nil
}

// GetLogSets gets details of an existing Log Set
func (l *LogSets) GetLogSet(logSetId string) (LogSet, LogSetInfo, error) {
	if logSetId == "" {
		return LogSet{}, LogSetInfo{}, errors.New("logSetId input parameter is mandatory")
	}
	logSet := &getLogSet{}
	if err := l.client.get(l.getLogSetEndPoint(logSetId), logSet); err != nil {
		return LogSet{}, LogSetInfo{}, err
	}
	logSetInfo := logSet.LogSet.logSetInfo(l)
	return logSet.LogSet, logSetInfo, nil
}

// PostLogSet creates a new Log Set
func (l *LogSets) PostLogSet(p PostLogSet) (LogSet, error) {
	logSet := &getLogSet{}
	postLogSet := &postLogSet{p}
	if err := l.client.post(l.getPath(), postLogSet, logSet); err != nil {
		return LogSet{}, err
	}
	return logSet.LogSet, nil
}

// PutLogSet updates an existing Log Set
func (l *LogSets) PutLogSet(logSetId string, p PutLogSet) (LogSet, error) {
	if logSetId == "" {
		return LogSet{}, errors.New("logSetId input parameter is mandatory")
	}
	logSet := &getLogSet{}
	putLogSet := &putLogSet{p}
	if err := l.client.put(l.getLogSetEndPoint(logSetId), putLogSet, logSet); err != nil {
		return LogSet{}, err
	}
	return logSet.LogSet, nil
}

// DeleteLogSet deletes a Log Set
func (l *LogSets) DeleteLogSet(logSetId string) error {
	if logSetId == "" {
		return errors.New("logSetId input parameter is mandatory")
	}
	var err error
	if err = l.client.delete(l.getLogSetEndPoint(logSetId)); err != nil {
		return err
	}
	return nil
}

func (l *LogSet) logSetInfo(c *LogSets) LogSetInfo {
	return LogSetInfo{
		Id:   l.Id,
		Name: l.Name,
		Links: []link{
			{
				Href: fmt.Sprintf("%s%s", c.client.logEntriesUrl, c.getLogSetEndPoint(l.Id)),
				Rel:  "Self",
			},
		},
	}
}

// getPath returns the rest end point for logsets
func (l *LogSets) getPath() string {
	return "/management/logsets"
}

// getLogSetEndPoint returns the rest end point to retrieve an individual logsets
func (l *LogSets) getLogSetEndPoint(logSetId string) string {
	return fmt.Sprintf("%s/%s", l.getPath(), logSetId)
}
