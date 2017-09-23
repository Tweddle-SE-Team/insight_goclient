package logentries_goclient

type Logs struct {
	client *client `json:"-"`
}

func NewLogs(c *client) Logs {
	return Logs{c}
}

// Structs meant for clients

type Log struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	LogsetsInfo []LogInfo `json:"logsets_info"`
	UserData    LogUserData  `json:"user_data"`
	Token       []string     `json:"token"`
	SourceType  string       `json:"source_type"`
	TokenSeed   string       `json:"token_seed"`
	Structures  []string     `json:"structures"`
}

type LogUserData struct {
	LogEntriesAgentFileName string `json:"le_agent_filename"`
	LogEntriesAgentFollow   string `json:"le_agent_follow"`
}

// Structs meant for marshalling/un-marshalling purposes

type logsCollection struct {
	Logs []Log
}

func (l *Logs) getPath() string {
	return "management/logs"
}

func (l *Logs) GetLogs() ([]Log, error) {
	logs := &logsCollection{}
	if err := l.client.get(l.getPath(), logs); err != nil {
		return nil, err
	}
	return logs.Logs, nil
}
