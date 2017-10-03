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
					Links: []link{
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
				Links: []link{
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

	returnedLogSet, _, err := logSets.GetLogSet(expectedLogSet.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogSet, returnedLogSet)

}

func TestLogSets_GetLogSetErrorsIfLogSetIdIsEmpty(t *testing.T) {
	logSets := LogSets{nil}
	_, _, err := logSets.GetLogSet("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logSetId input parameter is mandatory")
}

func TestLogSets_PostLogSet(t *testing.T) {

	p := PostLogSet{
		Name:        "MyLogSet2",
		Description: "some description",
		LogsInfo: []PostLogInfo{
			{
				Id: "logs-info-uuid",
			},
		},
		UserData: map[string]string{},
	}

	expectedLogSet := LogSet{
		Id:          "log-set-uuid",
		Name:        p.Name,
		Description: p.Description,
		LogsInfo: []LogInfo{
			{
				Id:   p.LogsInfo[0].Id,
				Name: "mylog",
				Links: []link{
					{
						Href: "https://rest.logentries.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: map[string]string{},
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/management/logsets", &postLogSet{p}, http.StatusCreated, &getLogSet{expectedLogSet})
	logSets := getLogSetsClient(requestMatcher)

	returnedLogSet, err := logSets.PostLogSet(p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogSet, returnedLogSet)

}

func TestLogSets_PutLogSet(t *testing.T) {

	logSetId := "log-set-uuid"

	p := PutLogSet{
		Name:        "New Name",
		Description: "updated description",
		LogsInfo: []LogInfo{
			{
				Id:   "logs-info-uuid",
				Name: "Lambda Log",
				Links: []link{
					{
						Href: "https://rest.logentries.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: map[string]string{},
	}

	expectedLogSet := LogSet{
		Id:          logSetId,
		Name:        p.Name,
		Description: p.Description,
		LogsInfo: []LogInfo{
			{
				Id:   p.LogsInfo[0].Id,
				Name: p.LogsInfo[0].Name,
				Links: []link{
					{
						Href: p.LogsInfo[0].Links[0].Href,
						Rel:  p.LogsInfo[0].Links[0].Rel,
					},
				},
			},
		},
		UserData: map[string]string{},
	}

	url := fmt.Sprintf("/management/logsets/%s", logSetId)
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, url, &putLogSet{p}, http.StatusOK, &getLogSet{expectedLogSet})
	logSets := getLogSetsClient(requestMatcher)

	returnedLogSet, err := logSets.PutLogSet(logSetId, p)
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
	return newLogSets(c)
}
