package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get All
//
func TestGetAllAchievementsSuccess(t *testing.T) {
	testUtils.AchieveRepo.MockSelectAll = func(context.Context) ([]domains.Achieve, error) {
		return []domains.Achieve{
			testUtils.CreateMockAchievement(1, 2020, "message1"),
			testUtils.CreateMockAchievement(2, 2021, "message2"),
		}, nil
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/achievements/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieves []domains.Achieve
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves[0].Id)
	assert.EqualValues(t, 2020, achieves[0].Year)
	assert.EqualValues(t, "message1", achieves[0].Message)
	assert.EqualValues(t, 2, achieves[1].Id)
	assert.EqualValues(t, 2021, achieves[1].Year)
	assert.EqualValues(t, "message2", achieves[1].Message)
	assert.EqualValues(t, 2, len(achieves))
}

//
// Test Get All Grouped By Year
//
func TestGetAllAchievementsGroupedByYearSuccess(t *testing.T) {
	testUtils.AchieveRepo.MockSelectAllGroupedByYear = func(context.Context) ([]domains.AchieveYearGroup, error) {
		return []domains.AchieveYearGroup{
			{
				Year: 2021,
				Achievements: []domains.Achieve{
					testUtils.CreateMockAchievement(1, 2021, "message1"),
				},
			},
			{
				Year: 2020,
				Achievements: []domains.Achieve{
					testUtils.CreateMockAchievement(2, 2020, "message2"),
				},
			},
		}, nil
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/achievements/years", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieves []domains.AchieveYearGroup
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves[0].Achievements[0].Id)
	assert.EqualValues(t, 2021, achieves[0].Achievements[0].Year)
	assert.EqualValues(t, "message1", achieves[0].Achievements[0].Message)
	assert.EqualValues(t, 2, achieves[1].Achievements[0].Id)
	assert.EqualValues(t, 2020, achieves[1].Achievements[0].Year)
	assert.EqualValues(t, "message2", achieves[1].Achievements[0].Message)
	assert.EqualValues(t, 2, len(achieves))
}

//
// Test Get Achievement
//
func TestGetAchievementSuccess(t *testing.T) {
	testUtils.AchieveRepo.MockSelectById = func(context.Context, uint) (domains.Achieve, error) {
		achieve := testUtils.CreateMockAchievement(1, 2020, "message1")
		return achieve, nil
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/achievements/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var achieve domains.Achieve
	if err := json.Unmarshal(recorder.Body.Bytes(), &achieve); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieve.Id)
	assert.EqualValues(t, 2020, achieve.Year)
	assert.EqualValues(t, "message1", achieve.Message)
}

func TestGetAchievementFailure(t *testing.T) {
	testUtils.AchieveRepo.MockSelectById = func(context.Context, uint) (domains.Achieve, error) {
		return domains.Achieve{}, appErrors.MockDbNoRowsError()
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/achievements/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateAchievementSuccess(t *testing.T) {
	testUtils.AchieveRepo.MockInsert = func(context.Context, domains.Achieve) error {
		return nil
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	achieve := testUtils.CreateMockAchievement(1, 2020, "message1")
	body := createBodyFromAchieve(achieve)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAchievementFailure(t *testing.T) {
	// no mock needed
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	achieve := testUtils.CreateMockAchievement(1, 0, "")
	body := createBodyFromAchieve(achieve)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateAchievementSuccess(t *testing.T) {
	testUtils.AchieveRepo.MockUpdate = func(context.Context, uint, domains.Achieve) error {
		return nil // Successful update
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	achieve := testUtils.CreateMockAchievement(1, 2020, "message1")
	body := createBodyFromAchieve(achieve)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/achievements/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAchievementInvalid(t *testing.T) {
	// no mock needed
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	achieve := testUtils.CreateMockAchievement(1, 0, "")
	body := createBodyFromAchieve(achieve)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/achievements/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAchievementFailure(t *testing.T) {
	testUtils.AchieveRepo.MockUpdate = func(context.Context, uint, domains.Achieve) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	achieve := testUtils.CreateMockAchievement(1, 2020, "message1")
	body := createBodyFromAchieve(achieve)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/achievements/achievement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteAchievementSuccess(t *testing.T) {
	testUtils.AchieveRepo.MockDelete = func(context.Context, uint) error {
		return nil // Return no error, successful delete!
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/achievements/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteAchievementFailure(t *testing.T) {
	testUtils.AchieveRepo.MockDelete = func(context.Context, uint) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.AchieveRepo = &testUtils.AchieveRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/achievements/achievement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Helper Methods
//

func createBodyFromAchieve(achieve domains.Achieve) io.Reader {
	marshal, err := json.Marshal(&achieve)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
