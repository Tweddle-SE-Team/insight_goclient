package logentries_goclient

import (
	"net/http"
	"reflect"
	"testing"

	"fmt"
	"github.com/dikhan/http_goclient/testutils"
	"github.com/stretchr/testify/assert"
)

func TestLogs_GetLogs(t *testing.T) {

	expectedLogs := []Log{
		{
			Id:   "log-uuid",
			Name: "MyLogSet",
			LogsetsInfo: []LogSetInfo{
				{
					Id:   "log-set-uuid",
					Name: "MyLogSet",
					Links: []link{
						{
							Href: "https://rest.logentries.com/management/logsets/log-set-uuid",
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

	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/management/logs", nil, http.StatusOK, &logsCollection{expectedLogs})
	logs := getLogsClient(requestMatcher)

	returnedLogs, err := logs.GetLogs()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedLogs, returnedLogs))
}

func TestLogs_GetLog(t *testing.T) {

	expectedLog := Log{
		Id:   "log-uuid",
		Name: "MyLogSet",
		LogsetsInfo: []LogSetInfo{
			{
				Id:   "log-set-uuid",
				Name: "MyLogSet",
				Links: []link{
					{
						Href: "https://rest.logentries.com/management/logsets/log-set-uuid",
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
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, &getLog{expectedLog})

	log := getLogsClient(requestMatcher)

	returnedLog, _, err := log.GetLog(expectedLog.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLog, returnedLog)
}

func TestLogs_GetLogErrorsIfLogSetIdIsEmpty(t *testing.T) {
	log := Logs{nil}
	_, _, err := log.GetLog("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logId input parameter is mandatory")
}

func TestLogs_PostLog(t *testing.T) {

	p := PostLog{
		Name:       "My New Awesome Log",
		Structures: []string{},
		SourceType: "token",
		LogsetsInfo: []PostLogSetInfo{
			{"log-set-uuid"},
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
		LogsetsInfo: []LogSetInfo{
			{Id: p.LogsetsInfo[0].Id},
		},
		UserData: LogUserData{
			LogEntriesAgentFileName: p.UserData.LogEntriesAgentFileName,
			LogEntriesAgentFollow:   p.UserData.LogEntriesAgentFollow,
		},
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/management/logs", &postLog{p}, http.StatusCreated, &getLog{expectedLog})
	log := getLogsClient(requestMatcher)

	returnedLog, err := log.PostLog(p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLog, returnedLog)

}

func TestLogs_PutLog(t *testing.T) {

	logId := "log-set-uuid"

	p := PutLog{
		Name:       "My New Awesome Log",
		Structures: []string{},
		SourceType: "token",
		LogsetsInfo: []LogSetInfo{
			{
				Id:   "log-set-uuid",
				Name: "ibtest",
				Links: []link{
					{
						Href: "https://rest.logentries.com/management/logsets/log-set-uuid",
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
		LogsetsInfo: []LogSetInfo{
			{
				Id:   p.LogsetsInfo[0].Id,
				Name: p.LogsetsInfo[0].Name,
				Links: []link{
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
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, url, &putLog{p}, http.StatusOK, &getLog{expectedLog})
	log := getLogsClient(requestMatcher)

	returnedLog, err := log.PutLog(logId, p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLog, returnedLog)

}

func TestLogs_PutLogErrorsIfLogSetIdIsEmpty(t *testing.T) {
	log := Logs{nil}
	_, err := log.PutLog("", PutLog{})
	assert.NotNil(t, err)
	assert.Error(t, err, "logId input parameter is mandatory")
}

func TestLogs_DeleteLog(t *testing.T) {
	logId := "log-set-uuid"

	url := fmt.Sprintf("/management/logs/%s", logId)
	requestMatcher := testutils.NewRequestMatcher(http.MethodDelete, url, nil, http.StatusNoContent, nil)
	log := getLogsClient(requestMatcher)

	err := log.DeleteLog(logId)
	assert.Nil(t, err)
}

func TestLogs_DeleteLogErrorsIfLogSetIdIsEmpty(t *testing.T) {
	log := Logs{nil}
	err := log.DeleteLog("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logId input parameter is mandatory")
}

func getLogsClient(requestMatcher testutils.TestRequestMatcher) Logs {
	c := getTestClient(requestMatcher)
	return newLogs(c)
}
