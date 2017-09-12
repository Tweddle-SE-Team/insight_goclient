package logentries_goclient

type LogSets struct {
	Client *client `json:"-"`
}

type LogSetCollection struct {
	LogSets []LogSet `json:"logsets"`
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

func (l *LogSets) getPath() string {
	return "management/logsets"
}

func (l *LogSets) GetLogSets() (*LogSetCollection, error) {
	logSets := &LogSetCollection{}
	if _, err := l.Client.get(l.getPath(), logSets); err != nil {
		return nil, err
	}
	return logSets, nil
}
