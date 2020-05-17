package integration_tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Semesters and GetAll(false)
func Test_CreateSemesters(t *testing.T) {
	semester1 := createSemester("2020_spring", "Spring 2020")
	semester2 := createSemester("2020_fall", "Fall 2020")
	semester3 := createSemester("2020_winter", "Winter 2020")
	body1 := createJsonBody(&semester1)
	body2 := createJsonBody(&semester2)
	body3 := createJsonBody(&semester3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/semesters/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/semesters/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/semesters/all", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var semesters []domains.Semester
	if err := json.Unmarshal(recorder4.Body.Bytes(), &semesters); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_spring", semesters[0].SemesterId)
	assert.EqualValues(t, "Spring 2020", semesters[0].Title)
	assert.EqualValues(t, "2020_fall", semesters[1].SemesterId)
	assert.EqualValues(t, "Fall 2020", semesters[1].Title)
	assert.EqualValues(t, "2020_winter", semesters[2].SemesterId)
	assert.EqualValues(t, "Winter 2020", semesters[2].Title)
	assert.EqualValues(t, 3, len(semesters))

	resetTable(t, domains.TABLE_SEMESTERS)
}

// Test: Create 2 Semesters with same semesterId. Then GetBySemesterId()
func Test_UniqueSemesterId(t *testing.T) {
	semester1 := createSemester("2020_spring", "Spring 2020")
	semester2 := createSemester("2020_spring", "Fall 2020") // Same semesterId
	body1 := createJsonBody(&semester1)
	body2 := createJsonBody(&semester2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/semesters/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusInternalServerError, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "Duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var semester domains.Semester
	if err := json.Unmarshal(recorder3.Body.Bytes(), &semester); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_spring", semester.SemesterId)
	assert.EqualValues(t, "Spring 2020", semester.Title)

	resetTable(t, domains.TABLE_SEMESTERS)
}

// Test: Create 1 Semester, Update it, GetBySemesterId()
func Test_UpdateSemester(t *testing.T) {
	// Create 1 Semester
	semester1 := createSemester("2020_spring", "Spring 2020")
	body1 := createJsonBody(&semester1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedSemester := createSemester("2020_fall", "Fall 2020")
	updatedBody := createJsonBody(&updatedSemester)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/semesters/semester/2020_spring", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_fall", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var semester domains.Semester
	if err := json.Unmarshal(recorder4.Body.Bytes(), &semester); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_fall", semester.SemesterId)
	assert.EqualValues(t, "Fall 2020", semester.Title)

	resetTable(t, domains.TABLE_SEMESTERS)
}

// Test: Create 1 Semester, Delete it, GetBySemesterId()
func Test_DeleteSemester(t *testing.T) {
	// Create
	semester1 := createSemester("2020_spring", "Spring 2020")
	body1 := createJsonBody(&semester1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/semesters/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_spring", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	resetTable(t, domains.TABLE_SEMESTERS)
}

// Helper methods
func createSemester(semesterId string, title string) domains.Semester {
	return domains.Semester{
		SemesterId: semesterId,
		Title:      title,
	}
}
