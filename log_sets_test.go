package logentries_goclient

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dikhan/logentries_goclient/testutils"
	"github.com/stretchr/testify/assert"
)

func TestLogSets_GetLogSets(t *testing.T) {

	expectedLogSets := []LogSet{
		{
			Id:          "log-set-uuid",
			Name:        "MyLogSet",
			Description: "some description",
			LogsInfo: []LogInfo{
				{
					Id:   "logs-info-uuid",
					Name: "MyLog",
					Links: []Link{
						{
							Href: "https://rest.logentries.com/management/logs/logs-info-uuid",
							Rel:  "Self",
						},
					},
				},
			},
		},
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/management/logsets", nil, http.StatusOK, &logSetCollection{expectedLogSets})
	logSets := getLogSetsClient(requestMatcher)

	returnedLogSets, err := logSets.GetLogSets()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedLogSets, returnedLogSets))
}

func TestLogSets_GetLogSet(t *testing.T) {

	expectedLogSet := LogSet{
		Id:          "log-set-uuid",
		Name:        "MyLogSet",
		Description: "some description",
		LogsInfo: []LogInfo{
			{
				Id:   "logs-info-uuid",
				Name: "Lambda Log",
				Links: []Link{
					{
						Href: "https://rest.logentries.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
	}

	url := fmt.Sprintf("/management/logsets/%s", expectedLogSet.Id)
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, &getLogSet{expectedLogSet})

	logSets := getLogSetsClient(requestMatcher)

	returnedLogSet, err := logSets.GetLogSet(expectedLogSet.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogSet, returnedLogSet)

}

func TestLogSets_GetLogSetErrorsIfLogSetIdIsEmpty(t *testing.T) {
	logSets := LogSets{nil}
	_, err := logSets.GetLogSet("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logSetId input parameter is mandatory")
}

func TestLogSets_PostLogSet(t *testing.T) {

	postLogSet := PostLogSet{
		Name:        "MyLogSet2",
		Description: "some description",
		LogsInfo: []PostLogInfo{
			{
				Id: "logs-info-uuid",
			},
		},
		UserData: UserData{},
	}

	expectedLogSet := LogSet{
		Id:          "log-set-uuid",
		Name:        postLogSet.Name,
		Description: postLogSet.Description,
		LogsInfo: []LogInfo{
			{
				Id:   postLogSet.LogsInfo[0].Id,
				Name: "mylog",
				Links: []Link{
					{
						Href: "https://rest.logentries.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: UserData{},
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/management/logsets", postLogSet, http.StatusCreated, &getLogSet{expectedLogSet})
	logSets := getLogSetsClient(requestMatcher)

	returnedLogSet, err := logSets.PostLogSet(postLogSet)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogSet, returnedLogSet)

}

func TestLogSets_PutLogSet(t *testing.T) {

	logSetId := "log-set-uuid"

	putLogSet := PutLogSet{
		Name:        "New Name",
		Description: "updated description",
		LogsInfo: []LogInfo{
			{
				Id:   "logs-info-uuid",
				Name: "Lambda Log",
				Links: []Link{
					{
						Href: "https://rest.logentries.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: UserData{},
	}

	expectedLogSet := LogSet{
		Id:          logSetId,
		Name:        putLogSet.Name,
		Description: putLogSet.Description,
		LogsInfo: []LogInfo{
			{
				Id:   putLogSet.LogsInfo[0].Id,
				Name: putLogSet.LogsInfo[0].Name,
				Links: []Link{
					{
						Href: putLogSet.LogsInfo[0].Links[0].Href,
						Rel:  putLogSet.LogsInfo[0].Links[0].Rel,
					},
				},
			},
		},
		UserData: UserData{},
	}

	url := fmt.Sprintf("/management/logsets/%s", logSetId)
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, url, putLogSet, http.StatusOK, &getLogSet{expectedLogSet})
	logSets := getLogSetsClient(requestMatcher)

	returnedLogSet, err := logSets.PutLogSet(logSetId, putLogSet)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogSet, returnedLogSet)

}

func TestLogSets_PutLogSetSetErrorsIfLogSetIdIsEmpty(t *testing.T) {
	logSets := LogSets{nil}
	_, err := logSets.PutLogSet("", PutLogSet{})
	assert.NotNil(t, err)
	assert.Error(t, err, "logSetId input parameter is mandatory")
}

func TestLogSets_DeleteLogSet(t *testing.T) {
	logSetId := "log-set-uuid"

	url := fmt.Sprintf("/management/logsets/%s", logSetId)
	requestMatcher := testutils.NewRequestMatcher(http.MethodDelete, url, nil, http.StatusNoContent, nil)
	logSets := getLogSetsClient(requestMatcher)

	err := logSets.DeleteLogSet(logSetId)
	assert.Nil(t, err)
}

func TestLogSets_DeleteLogSetSetErrorsIfLogSetIdIsEmpty(t *testing.T) {
	logSets := LogSets{nil}
	err := logSets.DeleteLogSet("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logSetId input parameter is mandatory")
}

func getLogSetsClient(requestMatcher testutils.TestRequestMatcher) LogSets {
	c := getTestClient(requestMatcher)
	return NewLogSets(c)
}
