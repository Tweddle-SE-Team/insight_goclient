package logentries_goclient

import (
	"errors"
	"fmt"
)

// The Logs resource allows you to interact with Logs in your account. The following operations are supported:
// - Get details of an existing Log
// - Get details of a list of all Logs
// - Create a new Log
// - Update an existing Log
// - Delete a Log

// Logs represents the logs interface by which user can interact with logentries logs API
type Logs struct {
	client *client `json:"-"`
}

// newLogs creates a new Logs struct that exposes Logs CRUD operations
func newLogs(c *client) Logs {
	return Logs{c}
}

// Structs meant for clients

// Log represents the entity used to get an existing log from the logentries API
type Log struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	LogsetsInfo []LogSetInfo `json:"logsets_info"`
	UserData    LogUserData  `json:"user_data"`
	Tokens      []string     `json:"tokens"`
	SourceType  string       `json:"source_type"`
	TokenSeed   string       `json:"token_seed"`
	Structures  []string     `json:"structures"`
}

// LogUserData represents user metadata
type LogUserData struct {
	LogEntriesAgentFileName string `json:"le_agent_filename"`
	LogEntriesAgentFollow   string `json:"le_agent_follow"`
}

// PostLog represents the entity used to create a new log to the logentries API
type PostLog struct {
	Name        string           `json:"name"`
	LogsetsInfo []PostLogSetInfo `json:"logsets_info"`
	UserData    LogUserData      `json:"user_data"`
	Tokens      []string         `json:"tokens"`
	SourceType  string           `json:"source_type"`
	TokenSeed   string           `json:"token_seed"`
	Structures  []string         `json:"structures"`
}

type PostLogInfo struct {
	Id string `json:"id"`
}

// PutLog represents the entity used to update a log to the logentries API
type PutLog struct {
	Name        string       `json:"name"`
	LogsetsInfo []LogSetInfo `json:"logsets_info"`
	UserData    LogUserData  `json:"user_data"`
	Tokens      []string     `json:"tokens"`
	SourceType  string       `json:"source_type"`
	TokenSeed   string       `json:"token_seed"`
	Structures  []string     `json:"structures"`
}

// LogInfo represents information about the log
type LogInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Links []link `json:"links"`
}

// Structs meant for marshalling/un-marshalling purposes

// logsCollection represents a wrapper struct for marshalling/unmarshalling purposes
type logsCollection struct {
	Logs []Log
}

// getLog represents a wrapper struct for marshalling/unmarshalling purposes
type getLog struct {
	Log Log `json:"log"`
}

// postLog represents a wrapper struct for marshalling/unmarshalling purposes
type postLog struct {
	PostLog PostLog `json:"log"`
}

// putLog represents a wrapper struct for marshalling/unmarshalling purposes
type putLog struct {
	PutLog PutLog `json:"log"`
}

// GetLogs lists all Logs for an account
func (l *Logs) GetLogs() ([]Log, error) {
	logs := &logsCollection{}
	if err := l.client.get(l.getPath(), logs); err != nil {
		return nil, err
	}
	return logs.Logs, nil
}

// GetLog gets a specific Log from an account
func (l *Logs) GetLog(logId string) (Log, LogInfo, error) {
	if logId == "" {
		return Log{}, LogInfo{}, errors.New("logId input parameter is mandatory")
	}
	log := &getLog{}
	if err := l.client.get(l.getLogEndPoint(logId), log); err != nil {
		return Log{}, LogInfo{}, err
	}
	logInfo := log.Log.logInfo(l)
	return log.Log, logInfo, nil
}

// PostLog adds a log to a given account.
func (l *Logs) PostLog(p PostLog) (Log, error) {
	logSet := &getLog{}
	postLog := &postLog{p}
	if err := l.client.post(l.getPath(), postLog, logSet); err != nil {
		return Log{}, err
	}
	return logSet.Log, nil
}

// PutLog updates a specific Log for an account
func (l *Logs) PutLog(logId string, p PutLog) (Log, error) {
	if logId == "" {
		return Log{}, errors.New("logId input parameter is mandatory")
	}
	log := &getLog{}
	putLogSet := &putLog{p}
	if err := l.client.put(l.getLogEndPoint(logId), putLogSet, log); err != nil {
		return Log{}, err
	}
	return log.Log, nil
}

// DeleteLog deletes a specific Log from an account.
func (l *Logs) DeleteLog(logId string) error {
	if logId == "" {
		return errors.New("logId input parameter is mandatory")
	}
	var err error
	if err = l.client.delete(l.getLogEndPoint(logId)); err != nil {
		return err
	}
	return nil
}

func (l *Log) logInfo(c *Logs) LogInfo {
	return LogInfo{
		Id:   l.Id,
		Name: l.Name,
		Links: []link{
			{
				Href: fmt.Sprintf("%s%s", c.client.logEntriesUrl, c.getLogEndPoint(l.Id)),
				Rel:  "Self",
			},
		},
	}
}

// getPath returns the rest end point for logs
func (l *Logs) getPath() string {
	return "/management/logs"
}

// getLogEndPoint returns the rest end point to retrieve an individual log
func (l *Logs) getLogEndPoint(logId string) string {
	return fmt.Sprintf("%s/%s", l.getPath(), logId)
}
