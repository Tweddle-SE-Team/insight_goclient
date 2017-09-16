package logentries_goclient

import (
	"fmt"
)

type LogSets struct {
	Client *client `json:"-"`
}

// Structs meant for clients
type PostLogSet struct {
	Name string `json:"name"`
	Description string `json:"description"`
	UserData UserData `json:"user_data"`
	LogsInfo []PostLogInfo `json:"logs_info"`
}

type PostLogInfo struct {
	Id string `json:"id"`
}

type PutLogSet struct {
	Name string `json:"name"`
	Description string `json:"description"`
	UserData UserData `json:"user_data"`
	LogsInfo []LogInfo `json:"logs_info"`
}

type LogSet struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	UserData UserData `json:"user_data"`
	LogsInfo []LogInfo `json:"logs_info"`
}

type UserData struct {
	LogEntriesDistName string `json:"le_distname"`
	LogEntriesDistVer string `json:"le_distver"`
	LogEntriesNameIntr string `json:"le_nameintr"`
}

type LogInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Links []Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel string `json:"rel"`
}

// Structs meant for marshalling/un-marshalling purposes
type logSetCollection struct {
	LogSets []LogSet `json:"logsets"`
}

type getLogSet struct {
	LogSet LogSet `json:"logset"`
}

func (l *LogSets) getPath() string {
	return "management/logsets"
}

func (l *LogSets) GetLogSets() ([]LogSet, error) {
	logSets := &logSetCollection{}
	if err := l.Client.get(l.getPath(), logSets); err != nil {
		return nil, err
	}
	return logSets.LogSets, nil
}

func (l *LogSets) GetLogSet(id string) (LogSet, error) {
	logSetEndPoint := fmt.Sprintf("%s/%s", l.getPath(), id)

	logSet := &getLogSet{}
	if err := l.Client.get(logSetEndPoint, logSet); err != nil {
		return LogSet{}, err
	}
	return logSet.LogSet, nil
}

func (l *LogSets) PostLogSet(postLogSet PostLogSet) (LogSet, error) {
	logSet := &getLogSet{}
	if err := l.Client.post(l.getPath(), postLogSet, logSet); err != nil {
		return LogSet{}, err
	}
	return logSet.LogSet, nil
}

func (l *LogSets) PutLogSet(logSetId string, putLogSet PutLogSet) (LogSet, error) {
	logSetEndPoint := fmt.Sprintf("%s/%s", l.getPath(), logSetId)
	logSet := &getLogSet{}
	if err := l.Client.put(logSetEndPoint, putLogSet, logSet); err != nil {
		return LogSet{}, err
	}
	return logSet.LogSet, nil
}

func (l *LogSets) DeleteLogSet(logSetId string) (error) {
	logSetEndPoint := fmt.Sprintf("%s/%s", l.getPath(), logSetId)
	var err error
	if err = l.Client.delete(logSetEndPoint); err != nil {
		return err
	}
	return nil
}