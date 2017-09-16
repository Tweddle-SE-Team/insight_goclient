package logentries_goclient

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"reflect"
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

	mockLogSetResponse := fmt.Sprintf(`{
    "logsets": [
        {
			"id": "%s",
			"name": "%s",
            "description": "%s",
            "logs_info": [
                {
                    "id": "%s",
					"name": "%s",
                    "links": [
                        {
                            "href": "%s",
                            "rel": "%s"
                        }
                    ]
                }
			],
            "user_data": {}
		}
	]}`, expectedLogSets[0].Id, expectedLogSets[0].Name, expectedLogSets[0].Description,
		expectedLogSets[0].LogsInfo[0].Id, expectedLogSets[0].LogsInfo[0].Name,
			expectedLogSets[0].LogsInfo[0].Links[0].Href, expectedLogSets[0].LogsInfo[0].Links[0].Rel)

	logSets := getLogSetsClient("/management/logsets", mockLogSetResponse)

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

	mockLogSetResponse := fmt.Sprintf(
		`{
			"logset": {
				"id": "%s",
				"name": "%s",
				"description": "%s",
				"user_data": {},
				"logs_info": [
					{
						"id": "%s",
						"name": "%s",
						"links": [
							{
								"href": "%s",
								"rel": "%s"
							}
						]
					}
				]
			}
		}`, expectedLogSet.Id, expectedLogSet.Name, expectedLogSet.Description, expectedLogSet.LogsInfo[0].Id,
			expectedLogSet.LogsInfo[0].Name, expectedLogSet.LogsInfo[0].Links[0].Href, expectedLogSet.LogsInfo[0].Links[0].Rel)


	logSets := getLogSetsClient(fmt.Sprintf("/management/logsets/%s", expectedLogSet.Id), mockLogSetResponse)

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

	mockLogSetResponse := fmt.Sprintf(
		`{
			"logset": {
				"id": "%s",
				"name": "%s",
				"description": "%s",
				"user_data": {},
				"logs_info": [
					{
						"id": "%s",
						"name": "%s",
						"links": [
							{
								"href": "%s",
								"rel": "%s"
							}
						]
					}
				]
			}
		}`, expectedLogSet.Id, expectedLogSet.Name, expectedLogSet.Description, expectedLogSet.LogsInfo[0].Id,
		expectedLogSet.LogsInfo[0].Name, expectedLogSet.LogsInfo[0].Links[0].Href, expectedLogSet.LogsInfo[0].Links[0].Rel)


	logSets := getLogSetsClient("/management/logsets", mockLogSetResponse)

	returnedLogSet, err := logSets.PostLogSet(postLogSet)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLogSet, returnedLogSet)

}

func getLogSetsClient(path, mockLogSetResponse string) LogSets {
	c := getTestClient(path, &MockResponse{mockLogSetResponse})
	return LogSets{c}
}