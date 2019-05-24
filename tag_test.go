package insight_goclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestTags_GetTags(t *testing.T) {

	expectedTags := []Tag{
		{
			Id:   "tag-uuid",
			Name: "Login Failure",
			Type: "Alert",
			Sources: []source{
				{
					Id:              "source-uuid",
					Name:            "auth.log",
					RetentionPeriod: "default",
					StoredDays:      []int{},
				},
			},
			Actions: []action{
				{
					Id:               "action-uuid",
					MinMatchesCount:  1,
					MinReportCount:   1,
					MinMatchesPeriod: "Day",
					MinReportPeriod:  "Day",
					Targets: []target{
						{
							Id:   "",
							Type: "",
							ParamsSet: map[string]string{
								"direct": "user@example.com",
								"teams":  "some-team",
								"users":  "user@example.com",
							},
							AlertContentSet: map[string]string{},
						},
					},
					Enabled: true,
					Type:    "Alert",
				},
			},
			Labels: GetLabels{
				{
					Id:       "label-uuid",
					Name:     "Login Failure",
					Reserved: false,
					Color:    "007afb",
					SN:       1056,
				},
			},
			Patterns: []string{"Power Button as"},
		},
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/management/tags", nil, http.StatusOK, &tagsCollection{expectedTags})
	tags := getTagsClient(requestMatcher)

	returnedTags, err := tags.GetTags()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedTags, returnedTags))
}

func TestTags_GetTag(t *testing.T) {

	expectedTag := Tag{
		Id:   "tag-uuid",
		Name: "Login Failure",
		Type: "Alert",
		Sources: []source{
			{
				Id:              "source-uuid",
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []action{
			{
				Id:               "action-uuid",
				MinMatchesCount:  1,
				MinReportCount:   1,
				MinMatchesPeriod: "Day",
				MinReportPeriod:  "Day",
				Targets: []target{
					{
						Id:   "",
						Type: "",
						ParamsSet: map[string]string{
							"direct": "user@example.com",
							"teams":  "some-team",
							"users":  "user@example.com",
						},
						AlertContentSet: map[string]string{},
					},
				},
				Enabled: true,
				Type:    "Alert",
			},
		},
		Labels: GetLabels{
			{
				Id:       "label-uuid",
				Name:     "Login Failure",
				Reserved: false,
				Color:    "007afb",
				SN:       1056,
			},
		},
		Patterns: []string{"Power Button as"},
	}

	url := fmt.Sprintf("/management/tags/%s", expectedTag.Id)
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, &getTag{expectedTag})

	tags := getTagsClient(requestMatcher)

	returnedTag, err := tags.GetTag(expectedTag.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedTag, returnedTag)

}

func TestTags_GetTagErrorsIfTagIdIsEmpty(t *testing.T) {
	tags := Tags{nil}
	_, err := tags.GetTag("")
	assert.NotNil(t, err)
	assert.Error(t, err, "tagId input parameter is mandatory")
}

func TestTags_PostTag(t *testing.T) {

	p := PostTag{
		Name: "Foo Bar Tag",
		Type: "Alert",
		Sources: []PostSource{
			{
				Id: "source-uuid",
			},
		},
		Actions: []PostAction{
			{
				MinMatchesCount:  1,
				MinReportCount:   1,
				MinMatchesPeriod: "Day",
				MinReportPeriod:  "Day",
				Targets: []PostTarget{
					{
						Type: "mailto",
						ParamsSet: map[string]string{
							"direct": "test@test.com",
						},
						AlertContentSet: map[string]string{"le_context": "true"},
					},
				},
				Enabled: true,
				Type:    "Alert",
			},
		},
		Labels: GetLabels{
			{
				Id:       "label-uuid",
				Name:     "Login Failure",
				Reserved: false,
				Color:    "007afb",
				SN:       1056,
			},
		},
		Patterns: []string{"/Foo Bar/"},
	}

	expectedTag := Tag{
		Id:   "new-tag-uuid",
		Name: p.Name,
		Type: p.Type,
		Sources: []source{
			{
				Id:              p.Sources[0].Id,
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []action{
			{
				Id:               "new-action-uuid",
				MinMatchesCount:  p.Actions[0].MinMatchesCount,
				MinReportCount:   p.Actions[0].MinReportCount,
				MinMatchesPeriod: p.Actions[0].MinMatchesPeriod,
				MinReportPeriod:  p.Actions[0].MinReportPeriod,
				Targets: []target{
					{
						Id:   "new-target-uuid",
						Type: p.Actions[0].Targets[0].Type,
						ParamsSet: map[string]string{
							"direct": p.Actions[0].Targets[0].ParamsSet["direct"],
							"teams":  p.Actions[0].Targets[0].ParamsSet["teams"],
							"users":  p.Actions[0].Targets[0].ParamsSet["users"],
						},
						AlertContentSet: p.Actions[0].Targets[0].AlertContentSet,
					},
				},
				Enabled: p.Actions[0].Enabled,
				Type:    p.Actions[0].Type,
			},
		},
		Labels: GetLabels{
			{
				Id:       p.Labels[0].Id,
				Name:     p.Labels[0].Name,
				Reserved: p.Labels[0].Reserved,
				Color:    p.Labels[0].Color,
				SN:       p.Labels[0].SN,
			},
		},
		Patterns: p.Patterns,
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/management/tags", &postTag{p}, http.StatusCreated, &getTag{expectedTag})

	tags := getTagsClient(requestMatcher)

	returnedTag, err := tags.PostTag(p)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedTag, returnedTag)
}

func TestTags_PutTag(t *testing.T) {

	tagId := "tagId"

	putTag := PostTag{
		Name: "Foo Bar Tag",
		Type: "Alert",
		Sources: []PostSource{
			{
				Id: "source-uuid",
			},
		},
		Actions: []PostAction{
			{
				MinMatchesCount:  0,
				MinReportCount:   1,
				MinMatchesPeriod: "Hour",
				MinReportPeriod:  "Hour",
				Targets: []PostTarget{
					{
						Type: "mailto",
						ParamsSet: map[string]string{
							"Direct": "test@test.com",
						},
						AlertContentSet: map[string]string{"le_context": "true"},
					},
				},
				Enabled: true,
				Type:    "Alert",
			},
		},
		Labels: GetLabels{
			{
				Id:       "label-uuid",
				Name:     "Test Label",
				Reserved: false,
				Color:    "3498db",
				SN:       1025,
			},
		},
		Patterns: []string{"/Foo Bar/"},
	}

	expectedTag := Tag{
		Id:   "new-tag-uuid",
		Name: putTag.Name,
		Type: putTag.Type,
		Sources: []source{
			{
				Id:              putTag.Sources[0].Id,
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []action{
			{
				Id:               "new-action-uuid",
				MinMatchesCount:  putTag.Actions[0].MinMatchesCount,
				MinReportCount:   putTag.Actions[0].MinReportCount,
				MinMatchesPeriod: putTag.Actions[0].MinMatchesPeriod,
				MinReportPeriod:  putTag.Actions[0].MinReportPeriod,
				Targets: []target{
					{
						Id:   "new-target-uuid",
						Type: putTag.Actions[0].Targets[0].Type,
						ParamsSet: map[string]string{
							"direct": putTag.Actions[0].Targets[0].ParamsSet["direct"],
							"teams":  putTag.Actions[0].Targets[0].ParamsSet["teams"],
							"users":  putTag.Actions[0].Targets[0].ParamsSet["users"],
						},
						AlertContentSet: putTag.Actions[0].Targets[0].AlertContentSet,
					},
				},
				Enabled: putTag.Actions[0].Enabled,
				Type:    putTag.Actions[0].Type,
			},
		},
		Labels: GetLabels{
			{
				Id:       putTag.Labels[0].Id,
				Name:     putTag.Labels[0].Name,
				Reserved: putTag.Labels[0].Reserved,
				Color:    putTag.Labels[0].Color,
				SN:       putTag.Labels[0].SN,
			},
		},
		Patterns: putTag.Patterns,
	}

	url := fmt.Sprintf("/management/tags/%s", tagId)
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, url, &postTag{putTag}, http.StatusOK, &getTag{expectedTag})

	tags := getTagsClient(requestMatcher)

	returnedTag, err := tags.PutTag(tagId, putTag)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedTag, returnedTag)
}

func TestTags_PutTagErrorsIfTagIdIsEmpty(t *testing.T) {
	tags := Tags{nil}
	_, err := tags.PutTag("", PostTag{})
	assert.NotNil(t, err)
	assert.Error(t, err, "tagId input parameter is mandatory")
}

func TestTags_DeleteTag(t *testing.T) {
	tagId := "tag-uuid"

	url := fmt.Sprintf("/management/tags/%s", tagId)
	requestMatcher := testutils.NewRequestMatcher(http.MethodDelete, url, nil, http.StatusNoContent, nil)
	log := getTagsClient(requestMatcher)

	err := log.DeleteTag(tagId)
	assert.Nil(t, err)
}

func TestTags_DeleteTagErrorsIfTagIdIsEmpty(t *testing.T) {
	tags := Tags{nil}
	err := tags.DeleteTag("")
	assert.NotNil(t, err)
	assert.Error(t, err, "tagId input parameter is mandatory")
}

func getTagsClient(requestMatcher testutils.TestRequestMatcher) Tags {
	c := getTestClient(requestMatcher)
	return newTags(c)
}
