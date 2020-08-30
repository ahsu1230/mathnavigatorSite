package tests_integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Semesters and GetAll(false)
func TestCreateSemesters(t *testing.T) {
	semester1 := createSemester(domains.FALL, 2019)
	semester2 := createSemester(domains.WINTER, 2020)
	semester3 := createSemester(domains.SPRING, 2020)
	body1 := utils.CreateJsonBody(&semester1)
	body2 := utils.CreateJsonBody(&semester2)
	body3 := utils.CreateJsonBody(&semester3)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/all", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var semesters []domains.Semester
	if err := json.Unmarshal(recorder4.Body.Bytes(), &semesters); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, "2019_fall", semesters[0].SemesterId)
	assert.EqualValues(t, "Fall 2019", semesters[0].Title)

	assert.EqualValues(t, "2020_winter", semesters[1].SemesterId)
	assert.EqualValues(t, "Winter 2020", semesters[1].Title)

	assert.EqualValues(t, "2020_spring", semesters[2].SemesterId)
	assert.EqualValues(t, "Spring 2020", semesters[2].Title)

	assert.EqualValues(t, 3, len(semesters))

	utils.ResetTable(t, domains.TABLE_SEMESTERS)
}

// Test: Create 2 Semesters with same semesterId. Then GetBySemesterId()
func TestUniqueSemesterId(t *testing.T) {
	semester1 := createSemester(domains.SPRING, 2020)
	semester2 := createSemester(domains.SPRING, 2020) // Same semesterId
	body1 := utils.CreateJsonBody(&semester1)
	body2 := utils.CreateJsonBody(&semester2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusBadRequest, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var semester domains.Semester
	if err := json.Unmarshal(recorder3.Body.Bytes(), &semester); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_spring", semester.SemesterId)
	assert.EqualValues(t, "Spring 2020", semester.Title)

	utils.ResetTable(t, domains.TABLE_SEMESTERS)
}

// Test: Create 1 Semester, Update it, GetBySemesterId()
func TestUpdateSemester(t *testing.T) {
	// Create 1 Semester
	semester1 := createSemester(domains.SPRING, 2020)
	body1 := utils.CreateJsonBody(&semester1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedSemester := createSemester(domains.FALL, 2020)
	updatedBody := utils.CreateJsonBody(&updatedSemester)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/semester/2020_spring", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_fall", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var semester domains.Semester
	if err := json.Unmarshal(recorder4.Body.Bytes(), &semester); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_fall", semester.SemesterId)
	assert.EqualValues(t, "Fall 2020", semester.Title)

	utils.ResetTable(t, domains.TABLE_SEMESTERS)
}

// Test: Create 1 Semester, Delete it, GetBySemesterId()
func TestDeleteSemester(t *testing.T) {
	// Create
	semester1 := createSemester(domains.SPRING, 2020)
	body1 := utils.CreateJsonBody(&semester1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_SEMESTERS)
}

// Helper methods
func createSemester(season string, year uint) domains.Semester {
	semesterId := fmt.Sprintf("%d_%s", year, season)
	title := strings.Title(fmt.Sprintf("%s %d", season, year))
	return domains.Semester{
		SemesterId: semesterId,
		Season:     season,
		Year:       year,
		Title:      title,
	}
}
