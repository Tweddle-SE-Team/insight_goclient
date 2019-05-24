package insight_goclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestLogsets_GetLogsets(t *testing.T) {

	expectedLogsets := []Logset{
		{
			Id:          "log-set-uuid",
			Name:        "MyLogset",
			Description: "some description",
			LogsInfo: []LogInfo{
				{
					Id:   "logs-info-uuid",
					Name: "MyLog",
					Links: []link{
						{
							Href: "https://eu.rest.logs.insight.rapid7.com/management/logs/logs-info-uuid",
							Rel:  "Self",
						},
					},
				},
			},
		},
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/management/logsets", nil, http.StatusOK, &logSetCollection{expectedLogsets})
	logSets := getLogsetsClient(requestMatcher)

	returnedLogsets, err := logSets.GetLogsets()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedLogsets, returnedLogsets))
}

func TestLogsets_GetLogset(t *testing.T) {

	expectedLogset := Logset{
		Id:          "log-set-uuid",
		Name:        "MyLogset",
		Description: "some description",
		LogsInfo: []LogInfo{
			{
				Id:   "logs-info-uuid",
				Name: "Lambda Log",
				Links: []link{
					{
						Href: "https://eu.rest.logs.insight.rapid7.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
	}

	url := fmt.Sprintf("/management/logsets/%s", expectedLogset.Id)
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, &getLogset{expectedLogset})

	logSets := getLogsetsClient(requestMatcher)

	returnedLogset, _, err := logSets.GetLogset(expectedLogset.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogset, returnedLogset)

}

func TestLogsets_GetLogsetErrorsIfLogsetIdIsEmpty(t *testing.T) {
	logSets := Logsets{nil}
	_, _, err := logSets.GetLogset("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logSetId input parameter is mandatory")
}

func TestLogsets_PostLogset(t *testing.T) {

	p := PostLogset{
		Name:        "MyLogset2",
		Description: "some description",
		LogsInfo: []PostLogInfo{
			{
				Id: "logs-info-uuid",
			},
		},
		UserData: map[string]string{},
	}

	expectedLogset := Logset{
		Id:          "log-set-uuid",
		Name:        p.Name,
		Description: p.Description,
		LogsInfo: []LogInfo{
			{
				Id:   p.LogsInfo[0].Id,
				Name: "mylog",
				Links: []link{
					{
						Href: "https://eu.rest.logs.insight.rapid7.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: map[string]string{},
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/management/logsets", &postLogset{p}, http.StatusCreated, &getLogset{expectedLogset})
	logSets := getLogsetsClient(requestMatcher)

	returnedLogset, err := logSets.PostLogset(p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogset, returnedLogset)

}

func TestLogsets_PutLogset(t *testing.T) {

	logSetId := "log-set-uuid"

	p := PutLogset{
		Name:        "New Name",
		Description: "updated description",
		LogsInfo: []LogInfo{
			{
				Id:   "logs-info-uuid",
				Name: "Lambda Log",
				Links: []link{
					{
						Href: "https://eu.rest.logs.insight.rapid7.com/management/logs/logs-info-uuid",
						Rel:  "Self",
					},
				},
			},
		},
		UserData: map[string]string{},
	}

	expectedLogset := Logset{
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
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, url, &putLogset{p}, http.StatusOK, &getLogset{expectedLogset})
	logSets := getLogsetsClient(requestMatcher)

	returnedLogset, err := logSets.PutLogset(logSetId, p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogset, returnedLogset)

}

func TestLogsets_PutLogsetSetErrorsIfLogsetIdIsEmpty(t *testing.T) {
	logSets := Logsets{nil}
	_, err := logSets.PutLogset("", PutLogset{})
	assert.NotNil(t, err)
	assert.Error(t, err, "logSetId input parameter is mandatory")
}

func TestLogsets_DeleteLogset(t *testing.T) {
	logSetId := "log-set-uuid"

	url := fmt.Sprintf("/management/logsets/%s", logSetId)
	requestMatcher := testutils.NewRequestMatcher(http.MethodDelete, url, nil, http.StatusNoContent, nil)
	logSets := getLogsetsClient(requestMatcher)

	err := logSets.DeleteLogset(logSetId)
	assert.Nil(t, err)
}

func TestLogsets_DeleteLogsetSetErrorsIfLogsetIdIsEmpty(t *testing.T) {
	logSets := Logsets{nil}
	err := logSets.DeleteLogset("")
	assert.NotNil(t, err)
	assert.Error(t, err, "logSetId input parameter is mandatory")
}

func getLogsetsClient(requestMatcher testutils.TestRequestMatcher) Logsets {
	c := getTestClient(requestMatcher)
	return newLogsets(c)
}
