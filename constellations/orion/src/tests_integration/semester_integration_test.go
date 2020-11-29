package tests_integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Semesters and GetAll(false)
func TestE2ECreateSemesters(t *testing.T) {
	utils.SendCreateSemester(t, true, domains.FALL, 2019)
	utils.SendCreateSemester(t, true, domains.WINTER, 2020)
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)

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
func TestE2EUniqueSemesterId(t *testing.T) {
	_, recorder1 := utils.SendCreateSemester(t, false, domains.SPRING, 2020)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	_, recorder2 := utils.SendCreateSemester(t, false, domains.SPRING, 2020) // Same semesterId
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
func TestE2EUpdateSemester(t *testing.T) {
	// Create 1 Semester
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)

	// Update
	updatedSemester := domains.Semester{
		Season: domains.FALL,
		Year:   2020,
	}
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
func TestE2EDeleteSemester(t *testing.T) {
	// Create
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_SEMESTERS)
}

// Test: Create 1 Semester, Archive it, GetBySemesterId()
func TestE2EArchiveSemester(t *testing.T) {
	// Create
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)

	// Archive
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/semesters/archive/2020_spring", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_SEMESTERS)
}
