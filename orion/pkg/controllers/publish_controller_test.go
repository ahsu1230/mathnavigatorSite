package controllers_test

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//
// Test Get Unpublished
//
func TestGetUnpublishedAchievements_Success(t *testing.T) {
	achieveService.mockGetUnpublished = func() ([]domains.Achieve, error) {
		return []domains.Achieve{
			createMockAchievement(1, 2020, "message1"),
			createMockAchievement(2, 2021, "message2"),
		}, nil
	}
	semesterService.mockGetUnpublished = func() ([]domains.Semester, error) {
		return []domains.Semester{
			createMockSemester("2020_fall", "Fall 2020"),
			createMockSemester("2020_winter", "Winter 2020"),
		}, nil
	}
	services.AchieveService = &achieveService
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, unpublishedDomains.Achieves[0].Id)
	assert.EqualValues(t, 2020, unpublishedDomains.Achieves[0].Year)
	assert.EqualValues(t, "message1", unpublishedDomains.Achieves[0].Message)
	assert.EqualValues(t, 2, unpublishedDomains.Achieves[1].Id)
	assert.EqualValues(t, 2021, unpublishedDomains.Achieves[1].Year)
	assert.EqualValues(t, "message2", unpublishedDomains.Achieves[1].Message)
	assert.EqualValues(t, 2, len(unpublishedDomains.Achieves))

	assert.EqualValues(t, "2020_fall", unpublishedDomains.Semesters[0].SemesterId)
	assert.EqualValues(t, "Fall 2020", unpublishedDomains.Semesters[0].Title)
	assert.EqualValues(t, "2020_winter", unpublishedDomains.Semesters[1].SemesterId)
	assert.EqualValues(t, "Winter 2020", unpublishedDomains.Semesters[1].Title)
	assert.EqualValues(t, 2, len(unpublishedDomains.Semesters))
}
