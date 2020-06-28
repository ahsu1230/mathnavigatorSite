package tests_integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Achievements and GetAll(false)
func Test_CreateAchievements(t *testing.T) {
	achieve1 := createAchievement(2020, "message1", 1)
	achieve2 := createAchievement(2021, "message2", 2)
	achieve3 := createAchievement(2022, "message3", 3)
	body1 := utils.CreateJsonBody(&achieve1)
	body2 := utils.CreateJsonBody(&achieve2)
	body3 := utils.CreateJsonBody(&achieve3)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/achievements/all", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var achieves []domains.Achieve
	if err := json.Unmarshal(recorder4.Body.Bytes(), &achieves); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieves[0].Id)
	assert.EqualValues(t, 2020, achieves[0].Year)
	assert.EqualValues(t, "message1", achieves[0].Message)
	assert.EqualValues(t, 1, achieves[0].Position)
	assert.EqualValues(t, 2, achieves[1].Id)
	assert.EqualValues(t, 2021, achieves[1].Year)
	assert.EqualValues(t, "message2", achieves[1].Message)
	assert.EqualValues(t, 2, achieves[1].Position)
	assert.EqualValues(t, 3, achieves[2].Id)
	assert.EqualValues(t, 2022, achieves[2].Year)
	assert.EqualValues(t, "message3", achieves[2].Message)
	assert.EqualValues(t, 3, achieves[2].Position)
	assert.EqualValues(t, 3, len(achieves))

	utils.ResetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Test: Create 4 Achievements and GetAllGroupedByYear()
func Test_GetAllAchievementsGroupedByYear(t *testing.T) {
	achieve1 := createAchievement(2020, "message1", 1)
	achieve2 := createAchievement(2021, "message2", 2)
	achieve3 := createAchievement(2022, "message3", 3)
	achieve4 := createAchievement(2021, "message4", 4)
	body1 := utils.CreateJsonBody(&achieve1)
	body2 := utils.CreateJsonBody(&achieve2)
	body3 := utils.CreateJsonBody(&achieve3)
	body4 := utils.CreateJsonBody(&achieve4)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body3)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body4)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Call Get All!
	recorder5 := utils.SendHttpRequest(t, http.MethodGet, "/api/achievements/years", nil)
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

	utils.ResetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Test: Create 1 Achievement, Update it, GetById()
func Test_UpdateAchievement(t *testing.T) {
	// Create 1 Achievement
	achieve1 := createAchievement(2020, "message1", 1)
	body1 := utils.CreateJsonBody(&achieve1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedAchieve := createAchievement(2021, "message2", 2)
	updatedBody := utils.CreateJsonBody(&updatedAchieve)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/achievement/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/achievements/achievement/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var achieve domains.Achieve
	if err := json.Unmarshal(recorder3.Body.Bytes(), &achieve); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, achieve.Id)
	assert.EqualValues(t, 2021, achieve.Year)
	assert.EqualValues(t, "message2", achieve.Message)

	utils.ResetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Test: Create 1 Achievement, Delete it, GetById()
func Test_DeleteAchievement(t *testing.T) {
	// Create
	achieve1 := createAchievement(2020, "message1", 1)
	body1 := utils.CreateJsonBody(&achieve1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/achievements/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/achievements/achievement/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/achievements/achievement/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_ACHIEVEMENTS)
}

// Helper methods
func createAchievement(year uint, message string, position uint) domains.Achieve {
	return domains.Achieve{
		Year:     year,
		Message:  message,
		Position: position,
	}
}
