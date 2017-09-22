package logentries_goclient

import (
	"net/http"
	"reflect"
	"testing"

	"fmt"

	"github.com/dikhan/logentries_goclient/testutils"
	"github.com/stretchr/testify/assert"
)

func TestTags_GetTags(t *testing.T) {

	expectedTags := []Tag{
		{
			Id:   "tag-uuid",
			Name: "Login Failure",
			Type: "Alert",
			Sources: []Source{
				{
					Id:              "source-uuid",
					Name:            "auth.log",
					RetentionPeriod: "default",
					StoredDays:      []int{},
				},
			},
			Actions: []Action{
				{
					Id:               "action-uuid",
					MinMatchesCount:  1,
					MinReportCount:   1,
					MinMatchesPeriod: "Day",
					MinReportPeriod:  "Day",
					Targets: Targets{
						{
							Id:   "",
							Type: "",
							ParamsSet: ParamsSet{
								Direct: "user@example.com",
								Teams:  "some-team",
								Users:  "user@example.com",
							},
						},
					},
					AlertContentSet: map[string]string{},
					Enabled:         true,
					Type:            "Alert",
				},
			},
			Labels: Labels{
				{
					Id:       "label-uuid",
					Name:     "Login Failure",
					Reserved: false,
					Color:    "007afb",
					SN:       1056,
				},
			},
			Patters: []string{"Power Button as"},
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
		Sources: []Source{
			{
				Id:              "source-uuid",
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []Action{
			{
				Id:               "action-uuid",
				MinMatchesCount:  1,
				MinReportCount:   1,
				MinMatchesPeriod: "Day",
				MinReportPeriod:  "Day",
				Targets: Targets{
					{
						Id:   "",
						Type: "",
						ParamsSet: ParamsSet{
							Direct: "user@example.com",
							Teams:  "some-team",
							Users:  "user@example.com",
						},
					},
				},
				AlertContentSet: map[string]string{},
				Enabled:         true,
				Type:            "Alert",
			},
		},
		Labels: Labels{
			{
				Id:       "label-uuid",
				Name:     "Login Failure",
				Reserved: false,
				Color:    "007afb",
				SN:       1056,
			},
		},
		Patters: []string{"Power Button as"},
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

	postTag := PostTag{
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
				Targets: Targets{
					{
						Type: "mailto",
						ParamsSet: ParamsSet{
							Direct: "test@test.com",
						},
					},
				},
				AlertContentSet: map[string]string{"le_context": "true"},
				Enabled:         true,
				Type:            "Alert",
			},
		},
		Labels: Labels{
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
		Name: postTag.Name,
		Type: postTag.Type,
		Sources: []Source{
			{
				Id:              postTag.Sources[0].Id,
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []Action{
			{
				Id:               "new-action-uuid",
				MinMatchesCount:  postTag.Actions[0].MinMatchesCount,
				MinReportCount:   postTag.Actions[0].MinReportCount,
				MinMatchesPeriod: postTag.Actions[0].MinMatchesPeriod,
				MinReportPeriod:  postTag.Actions[0].MinReportPeriod,
				Targets: Targets{
					{
						Id:   "new-target-uuid",
						Type: postTag.Actions[0].Targets[0].Type,
						ParamsSet: ParamsSet{
							Direct: postTag.Actions[0].Targets[0].ParamsSet.Direct,
							Teams:  postTag.Actions[0].Targets[0].ParamsSet.Teams,
							Users:  postTag.Actions[0].Targets[0].ParamsSet.Users,
						},
					},
				},
				AlertContentSet: postTag.Actions[0].AlertContentSet,
				Enabled:         postTag.Actions[0].Enabled,
				Type:            postTag.Actions[0].Type,
			},
		},
		Labels: Labels{
			{
				Id:       postTag.Labels[0].Id,
				Name:     postTag.Labels[0].Name,
				Reserved: postTag.Labels[0].Reserved,
				Color:    postTag.Labels[0].Color,
				SN:       postTag.Labels[0].SN,
			},
		},
		Patters: postTag.Patterns,
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodPost, "/management/tags", postTag, http.StatusCreated, &getTag{expectedTag})

	tags := getTagsClient(requestMatcher)

	returnedTag, err := tags.PostTag(postTag)
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
				MinMatchesCount:  1,
				MinReportCount:   1,
				MinMatchesPeriod: "Day",
				MinReportPeriod:  "Day",
				Targets: Targets{
					{
						Type: "mailto",
						ParamsSet: ParamsSet{
							Direct: "test@test.com",
						},
					},
				},
				AlertContentSet: map[string]string{"le_context": "true"},
				Enabled:         true,
				Type:            "Alert",
			},
		},
		Labels: Labels{
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
		Name: putTag.Name,
		Type: putTag.Type,
		Sources: []Source{
			{
				Id:              putTag.Sources[0].Id,
				Name:            "auth.log",
				RetentionPeriod: "default",
				StoredDays:      []int{},
			},
		},
		Actions: []Action{
			{
				Id:               "new-action-uuid",
				MinMatchesCount:  putTag.Actions[0].MinMatchesCount,
				MinReportCount:   putTag.Actions[0].MinReportCount,
				MinMatchesPeriod: putTag.Actions[0].MinMatchesPeriod,
				MinReportPeriod:  putTag.Actions[0].MinReportPeriod,
				Targets: Targets{
					{
						Id:   "new-target-uuid",
						Type: putTag.Actions[0].Targets[0].Type,
						ParamsSet: ParamsSet{
							Direct: putTag.Actions[0].Targets[0].ParamsSet.Direct,
							Teams:  putTag.Actions[0].Targets[0].ParamsSet.Teams,
							Users:  putTag.Actions[0].Targets[0].ParamsSet.Users,
						},
					},
				},
				AlertContentSet: putTag.Actions[0].AlertContentSet,
				Enabled:         putTag.Actions[0].Enabled,
				Type:            putTag.Actions[0].Type,
			},
		},
		Labels: Labels{
			{
				Id:       putTag.Labels[0].Id,
				Name:     putTag.Labels[0].Name,
				Reserved: putTag.Labels[0].Reserved,
				Color:    putTag.Labels[0].Color,
				SN:       putTag.Labels[0].SN,
			},
		},
		Patters: putTag.Patterns,
	}

	url := fmt.Sprintf("/management/tags/%s", tagId)
	requestMatcher := testutils.NewRequestMatcher(http.MethodPut, url, putTag, http.StatusOK, &getTag{expectedTag})

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

func getTagsClient(requestMatcher testutils.TestRequestMatcher) Tags {
	c := getTestClient(requestMatcher)
	return NewTags(c)
}
