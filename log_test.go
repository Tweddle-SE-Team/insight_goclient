package insight_goclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLogs_GetLogs(t *testing.T) {

	expectedLogs := []Log{
		{
			Id:   "log-uuid",
			Name: "MyLogset",
			LogsetsInfo: []Info{
				{
					Id:   "log-set-uuid",
					Name: "MyLogset",
					Links: []Link{
						{
							Href: "https://eu.rest.logs.insight.rapid7.com/management/logsets/log-set-uuid",
							Rel:  "self",
						},
					},
				},
			},
			Tokens:     []string{},
			SourceType: "AGENT",
			TokenSeed:  "",
			Structures: []string{},
			UserData: LogUserData{
				LogEntriesAgentFileName: "",
				LogEntriesAgentFollow:   "",
			},
		},
	}

	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/logs", nil, http.StatusOK, expectedLogs)
	logs := getTestClient(requestMatcher)
	returnedLogs, err := logs.GetLogs()
	assert.Nil(t, err)
	assert.EqualValues(t, &expectedLogs, returnedLogs)
}

func TestLogs_GetLog(t *testing.T) {

	expectedLog := Log{
		Id:   "log-uuid",
		Name: "MyLogset",
		LogsetsInfo: []Info{
			{
				Id:   "log-set-uuid",
				Name: "MyLogset",
				Links: []Link{
					{
						Href: "https://eu.rest.logs.insight.rapid7.com/management/logsets/log-set-uuid",
						Rel:  "self",
					},
				},
			},
		},
		Tokens:     []string{},
		SourceType: "AGENT",
		TokenSeed:  "",
		Structures: []string{},
		UserData: LogUserData{
			LogEntriesAgentFileName: "",
			LogEntriesAgentFollow:   "",
		},
	}

	url := fmt.Sprintf("/management/logs/%s", expectedLog.Id)
	requestMatcher := NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, expectedLog)
	client := getTestClient(requestMatcher)
	returnedLog, err := client.GetLog(expectedLog.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, &expectedLog, returnedLog)
}

func TestLogs_GetLogErrorsIfLogsetIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/logs/", nil, http.StatusOK, Log{})
	client := getTestClient(requestMatcher)
	_, err := client.GetLog("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logId input parameter is mandatory")
}

func TestLogs_PostLog(t *testing.T) {

	p := Log{
		Name:       "My New Awesome Log",
		Structures: []string{},
		SourceType: "token",
		LogsetsInfo: []Info{
			{Id: "log-set-uuid"},
		},
		UserData: LogUserData{
			LogEntriesAgentFileName: "",
			LogEntriesAgentFollow:   "false",
		},
	}

	expectedLog := Log{
		Id:         "log-set-uuid",
		Name:       p.Name,
		Tokens:     []string{"daf42867-a82f-487e-95b7-8d10dba6c4f5"},
		Structures: []string{},
		LogsetsInfo: []Info{
			{Id: p.LogsetsInfo[0].Id},
		},
		UserData: LogUserData{
			LogEntriesAgentFileName: p.UserData.LogEntriesAgentFileName,
			LogEntriesAgentFollow:   p.UserData.LogEntriesAgentFollow,
		},
	}

	requestMatcher := NewRequestMatcher(http.MethodPost, "/management/logs", p, http.StatusCreated, expectedLog)
	client := getTestClient(requestMatcher)
	err := client.PostLog(&p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLog, p)

}

func TestLogs_PutLog(t *testing.T) {

	logId := "log-set-uuid"

	p := &Log{
		Id:         logId,
		Name:       "My New Awesome Log",
		Structures: []string{},
		SourceType: "token",
		LogsetsInfo: []Info{
			{
				Id:   "log-set-uuid",
				Name: "ibtest",
				Links: []Link{
					{
						Href: "https://eu.rest.logs.insight.rapid7.com/management/logsets/log-set-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: LogUserData{
			LogEntriesAgentFileName: "",
			LogEntriesAgentFollow:   "false",
		},
	}

	expectedLog := Log{
		Id:         logId,
		Name:       p.Name,
		Tokens:     []string{"daf42867-a82f-487e-95b7-8d10dba6c4f5"},
		Structures: []string{},
		LogsetsInfo: []Info{
			{
				Id:   p.LogsetsInfo[0].Id,
				Name: p.LogsetsInfo[0].Name,
				Links: []Link{
					{
						Href: p.LogsetsInfo[0].Links[0].Href,
						Rel:  p.LogsetsInfo[0].Links[0].Rel,
					},
				},
			},
		},
		UserData: LogUserData{
			LogEntriesAgentFileName: p.UserData.LogEntriesAgentFileName,
			LogEntriesAgentFollow:   p.UserData.LogEntriesAgentFollow,
		},
	}

	url := fmt.Sprintf("/management/logs/%s", logId)
	requestMatcher := NewRequestMatcher(http.MethodPut, url, p, http.StatusOK, expectedLog)
	client := getTestClient(requestMatcher)
	err := client.PutLog(p)
	assert.Nil(t, err)
	assert.EqualValues(t, &expectedLog, p)
}

func TestLogs_PutLogErrorsIfLogsetIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/logs/", nil, http.StatusOK, Log{})
	client := getTestClient(requestMatcher)
	err := client.PutLog(&Log{})
	assert.NotNil(t, err)
	assert.Error(t, err, "logId input parameter is mandatory")
}

func TestLogs_DeleteLog(t *testing.T) {
	logId := "log-set-uuid"
	url := fmt.Sprintf("/management/logs/%s", logId)
	requestMatcher := NewRequestMatcher(http.MethodDelete, url, nil, http.StatusNoContent, nil)
	log := getTestClient(requestMatcher)
	err := log.DeleteLog(logId)
	assert.Nil(t, err)
}

func TestLogs_DeleteLogErrorsIfLogsetIdIsEmpty(t *testing.T) {
	requestMatcher := NewRequestMatcher(http.MethodGet, "/management/logs/", nil, http.StatusOK, Log{})
	client := getTestClient(requestMatcher)
	err := client.DeleteLog("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logId input parameter is mandatory")
}
