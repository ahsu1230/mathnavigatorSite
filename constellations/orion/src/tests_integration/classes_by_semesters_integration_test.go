package tests_integration

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 2 semesters, 2 programs, 4 classes and sort
func TestTwoSemestersTwoProgramsFourClasses(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	createAllClasses(t)

	// Call sorting endpoint
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classesbysemesters", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var results []domains.ProgramClassesBySemester
	if err := json.Unmarshal(recorder.Body.Bytes(), &results); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_spring", results[0].Semester.SemesterId)
	assert.EqualValues(t, "program1", results[0].ProgramClasses[0].ProgramObj.ProgramId)
	assert.EqualValues(t, "program1_2020_spring_class1", results[0].ProgramClasses[0].Classes[0].ClassId)
	assert.EqualValues(t, "program1_2020_spring_class2", results[0].ProgramClasses[0].Classes[1].ClassId)
	assert.EqualValues(t, "2020_summer", results[1].Semester.SemesterId)

	resetClassTables(t)
}

// This test should work without a database call!
// Previous results from above test should've been saved in cache
// Make sure redis is running for this to work!
func TestGetSemestersProgramsClassesCacheHit(t *testing.T) {
	if os.Getenv(utils.ENV_VAR_TEST) == utils.ENV_VAR_CIRCLE_CI {
		t.Skip("Skipping test because redis not created during CI tests")
	}

	// Call sorting endpoint
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classesbysemesters", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var results []domains.ProgramClassesBySemester
	if err := json.Unmarshal(recorder.Body.Bytes(), &results); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_spring", results[0].Semester.SemesterId)
	assert.EqualValues(t, "program1", results[0].ProgramClasses[0].ProgramObj.ProgramId)
	assert.EqualValues(t, "program1_2020_spring_class1", results[0].ProgramClasses[0].Classes[0].ClassId)
	assert.EqualValues(t, "program1_2020_spring_class2", results[0].ProgramClasses[0].Classes[1].ClassId)
	assert.EqualValues(t, "2020_summer", results[1].Semester.SemesterId)
}
