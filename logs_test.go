package logentries_goclient

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/dikhan/logentries_goclient/testutils"
	"github.com/stretchr/testify/assert"
)

func TestLogs_GetLogs(t *testing.T) {

	expectedLogs := []Log{
		{
			Id:   "log-set-uuid",
			Name: "MyLogSet",
			LogsetsInfo: []LogInfo{
				{
					Id:   "log-set-uuid",
					Name: "MyLogSet",
					Links: []Link{
						{
							Href: "https://rest.logentries.com/management/logsets/log-set-uuid",
							Rel:  "self",
						},
					},
				},
			},
			Token:      []string{},
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

func getLogsClient(requestMatcher testutils.TestRequestMatcher) Logs {
	c := getTestClient(requestMatcher)
	return NewLogs(c)
}
