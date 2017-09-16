package logentries_goclient

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"reflect"
	"net/http"
	"github.com/dikhan/logentries_goclient/testutils"
)

func TestLogSets_GetLogSets(t *testing.T) {

	expectedLogSets := []LogSet{
		{
			Id: "log-set-uuid",
			Name: "MyLogSet",
			Description: "some description",
			LogsInfo: []LogInfo{
				{
					Id: "logs-info-uuid",
					Name: "MyLog",
					Links: []Link{
						{
							Href:  "https://rest.logentries.com/management/logs/logs-info-uuid",
							Rel: "Self",
						},
					},
				},
			},
		},
	}

	requestMatcher := testutils.RequestMatcher{
		ExpectedRequest: testutils.ExpectedRequest {
			HttpMethod: http.MethodGet,
			Url: "/management/logsets",
		},
		Response: &logSetCollection{
			expectedLogSets,
		},
	}

	logSets := getLogSetsClient(requestMatcher)

	returnedLogSets, err := logSets.GetLogSets()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedLogSets, returnedLogSets))
}

func TestLogSets_GetLogSet(t *testing.T) {

	expectedLogSet := LogSet{
			Id: "log-set-uuid",
			Name: "MyLogSet",
			Description: "some description",
			LogsInfo: []LogInfo{
				{
					Id: "logs-info-uuid",
					Name: "Lambda Log",
					Links: []Link{
						{
							Href:  "https://rest.logentries.com/management/logs/logs-info-uuid",
							Rel: "Self",
						},
					},
				},
			},
	}

	requestMatcher := testutils.RequestMatcher{
		ExpectedRequest: testutils.ExpectedRequest {
			HttpMethod: http.MethodGet,
			Url: fmt.Sprintf("/management/logset/%s", expectedLogSet.Id),
		},
		Response: &getLogSet{
			expectedLogSet,
		},
	}

	logSets := getLogSetsClient(requestMatcher)

	returnedLogSet, err := logSets.GetLogSet(expectedLogSet.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogSet, returnedLogSet)

}

func TestLogSets_PostLogSet(t *testing.T) {

	postLogSet := PostLogSet{
		Name: "MyLogSet2",
		Description: "some description",
		LogsInfo: []PostLogInfo{
			{
				Id: "logs-info-uuid",
			},
		},
		UserData: UserData{},
	}

	expectedLogSet := LogSet{
		Id: "log-set-uuid",
		Name: postLogSet.Name,
		Description: postLogSet.Description,
		LogsInfo: []LogInfo{
			{
				Id: postLogSet.LogsInfo[0].Id,
				Name: "mylog",
				Links: []Link{
					{
						Href:  "https://rest.logentries.com/management/logs/logs-info-uuid",
						Rel: "Self",
					},
				},
			},
		},
		UserData: UserData{},
	}

	requestMatcher := testutils.RequestMatcher{
		ExpectedRequest: testutils.ExpectedRequest {
			HttpMethod: http.MethodPost,
			Url: "/management/logsets",
			Payload: postLogSet,
		},
		Response: &getLogSet{
			expectedLogSet,
		},
	}

	logSets := getLogSetsClient(requestMatcher)

	returnedLogSet, err := logSets.PostLogSet(postLogSet)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogSet, returnedLogSet)

}

func getLogSetsClient(requestMatcher testutils.RequestMatcher) LogSets {
	c := getTestClient(requestMatcher)
	return LogSets{c}
}