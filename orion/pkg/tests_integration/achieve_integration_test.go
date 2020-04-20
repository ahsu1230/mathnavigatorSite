package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 3 Achievements and GetAll(false)
func Test_CreateAchievements(t *testing.T) {
	achieve1 := createAchievement(2020, "message1")
	achieve2 := createAchievement(2021, "message2")
	achieve3 := createAchievement(2022, "message3")
	body1 := createJsonBody(&achieve1)
	body2 := createJsonBody(&achieve2)
	body3 := createJsonBody(&achieve3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/all", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var achieves []domains.Achieve
	if err := json.Unmarshal(recorder4.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves[0].Id)
	assert.EqualValues(t, 2020, achieves[0].Year)
	assert.EqualValues(t, "message1", achieves[0].Message)
	assert.EqualValues(t, 2, achieves[1].Id)
	assert.EqualValues(t, 2021, achieves[1].Year)
	assert.EqualValues(t, "message2", achieves[1].Message)
	assert.EqualValues(t, 3, achieves[2].Id)
	assert.EqualValues(t, 2022, achieves[2].Year)
	assert.EqualValues(t, "message3", achieves[2].Message)
	assert.EqualValues(t, 3, len(achieves))

	resetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Test: Create 4 Achievements and GetAllGroupedByYear()
func Test_GetAllAchievementsGroupedByYear(t *testing.T) {
	achieve1 := createAchievement(2020, "message1")
	achieve2 := createAchievement(2021, "message2")
	achieve3 := createAchievement(2022, "message3")
	achieve4 := createAchievement(2021, "message4")
	body1 := createJsonBody(&achieve1)
	body2 := createJsonBody(&achieve2)
	body3 := createJsonBody(&achieve3)
	body4 := createJsonBody(&achieve4)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body3)
	recorder4 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body4)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Call Get All!
	recorder5 := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/years", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var achieves []domains.AchieveYearGroup
	if err := json.Unmarshal(recorder5.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 3, achieves[0].Achievements[0].Id)
	assert.EqualValues(t, 2022, achieves[0].Achievements[0].Year)
	assert.EqualValues(t, "message3", achieves[0].Achievements[0].Message)
	assert.EqualValues(t, 2, achieves[1].Achievements[0].Id)
	assert.EqualValues(t, 2021, achieves[1].Achievements[0].Year)
	assert.EqualValues(t, "message2", achieves[1].Achievements[0].Message)
	assert.EqualValues(t, 4, achieves[1].Achievements[1].Id)
	assert.EqualValues(t, 2021, achieves[1].Achievements[1].Year)
	assert.EqualValues(t, "message4", achieves[1].Achievements[1].Message)
	assert.EqualValues(t, 1, achieves[2].Achievements[0].Id)
	assert.EqualValues(t, 2020, achieves[2].Achievements[0].Year)
	assert.EqualValues(t, "message1", achieves[2].Achievements[0].Message)
	assert.EqualValues(t, 3, len(achieves))

	resetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Test: Create 1 Achievement, Update it, GetById()
func Test_UpdateAchievement(t *testing.T) {
	// Create 1 Achievement
	achieve1 := createAchievement(2020, "message1")
	body1 := createJsonBody(&achieve1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedAchieve := createAchievement(2021, "message2")
	updatedBody := createJsonBody(&updatedAchieve)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/achievement/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/achievement/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var achieve domains.Achieve
	if err := json.Unmarshal(recorder3.Body.Bytes(), &achieve); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieve.Id)
	assert.EqualValues(t, 2021, achieve.Year)
	assert.EqualValues(t, "message2", achieve.Message)

	resetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Test: Create 1 Achievement, Delete it, GetById()
func Test_DeleteAchievement(t *testing.T) {
	// Create
	achieve1 := createAchievement(2020, "message1")
	body1 := createJsonBody(&achieve1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/achievements/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/achievements/v1/achievement/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/achievements/v1/achievement/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	resetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Helper methods
func createAchievement(year uint, message string) domains.Achieve {
	return domains.Achieve{
		Year:    year,
		Message: message,
	}
}
