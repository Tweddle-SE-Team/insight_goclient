package logentries_goclient

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLogSets_GetLogSets(t *testing.T) {

	expectedLogSet := &LogSetCollection{LogSets:[]LogSet{
		{
			Id: "log-set-uuid",
			Name: "log-set-name",
			Description: "some description",
			LogsInfo: []LogInfo{
				{
					Id: "logs-info-uuid",
					Name: "test.log",
					Links: []Link{
						{
							Href:  "https://rest.logentries.com/management/logs/logs-info-uuid",
							Rel: "Self",
						},
					},
				},
			},
		},
	}}

	mockLogSetResponse := `{
    "logsets": [
        {
			"name": "log-set-name",
            "description": "some description",
            "id": "log-set-uuid",
            "logs_info": [
                {
                    "id": "logs-info-uuid",
                    "links": [
                        {
                            "href": "https://rest.logentries.com/management/logs/logs-info-uuid",
                            "rel": "Self"
                        }
                    ],
                    "name": "test.log"
                }
			],
            "user_data": {}
		}
	]}`

	mockResponses := MockResponses{"/management/logsets": &MockResponse{mockLogSetResponse}}
	httpClient, httpServer := testClientServer(mockResponses)

	c := &client{logEntriesUrl: httpServer.URL, api_key: "apikey", httpClient: httpClient}
	logSets := LogSets{c}

	logSetsCollection, err := logSets.GetLogSets()
	assert.Nil(t, err)
	assert.Equal(t, expectedLogSet, logSetsCollection)
}
