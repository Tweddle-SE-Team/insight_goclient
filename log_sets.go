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

type LogSets struct {
	client *client `json:"-"`
}

func NewLogSets(c *client) LogSets {
	return LogSets{c}
}

// Structs meant for clients
type PostLogSet struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	UserData    map[string]string `json:"user_data,omitempty"`
	LogsInfo    []PostLogSetInfo  `json:"logs_info,omitempty"`
}

type PostLogSetInfo struct {
	Id string `json:"id"`
}

type PutLogSet struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	UserData    map[string]string  `json:"user_data,omitempty"`
	LogsInfo    []LogInfo `json:"logs_info,omitempty"`
}

type LogSet struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserData    map[string]string  `json:"user_data"`
	LogsInfo    []LogInfo `json:"logs_info"`
}

type UserData struct {
	LogEntriesDistName string `json:"le_distname"`
	LogEntriesDistVer  string `json:"le_distver"`
	LogEntriesNameIntr string `json:"le_nameintr"`
}

type LogSetInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Links []Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// Structs meant for marshalling/un-marshalling purposes
type logSetCollection struct {
	LogSets []LogSet `json:"logsets"`
}

type getLogSet struct {
	LogSet LogSet `json:"logset"`
}

type postLogSet struct {
	PostLogSet PostLogSet `json:"logset"`
}

type putLogSet struct {
	PutLogSet PutLogSet `json:"logset"`
}

// GetLogSets gets details of an existing Log Set
func (l *LogSets) GetLogSets() ([]LogSet, error) {
	logSets := &logSetCollection{}
	if err := l.client.get(l.getPath(), logSets); err != nil {
		return nil, err
	}
	return logSets.LogSets, nil
}

// GetLogSet gets details of a list of all Log Sets
func (l *LogSets) GetLogSet(logSetId string) (LogSet, error) {
	if logSetId == "" {
		return LogSet{}, errors.New("logSetId input parameter is mandatory")
	}
	logSet := &getLogSet{}
	if err := l.client.get(l.getLogSetEndPoint(logSetId), logSet); err != nil {
		return LogSet{}, err
	}
	return logSet.LogSet, nil
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

func (l *LogSets) getPath() string {
	return "management/logsets"
}

func (l *LogSets) getLogSetEndPoint(logSetId string) string {
	return fmt.Sprintf("%s/%s", l.getPath(), logSetId)
}