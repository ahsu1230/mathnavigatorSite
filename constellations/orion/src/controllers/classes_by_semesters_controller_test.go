package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

// Test getting programs and classes by semester endpoint
func TestGetClassesAndProgramsBySemester_Success(t *testing.T) {
	// Mock selecting all programs, semesters, classes
	testUtils.ProgramRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Program, error) {
		return createMockPrograms(1, 2), nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return createMockSemesters(1, 2), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Class, error) {
		return createMockClasses(1, 2, 3, 4), nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classesbysemesters", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var results []domains.ProgramClassesBySemester
	if err := json.Unmarshal(recorder.Body.Bytes(), &results); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_spring", results[0].Semester.SemesterId)
	assert.EqualValues(t, "program1", results[0].ProgramClasses[0].ProgramObj.ProgramId)
	assert.EqualValues(t, "class1", results[0].ProgramClasses[0].Classes[0].ClassId)
	assert.EqualValues(t, "2020_summer", results[1].Semester.SemesterId)
	assert.EqualValues(t, "program1", results[1].ProgramClasses[0].ProgramObj.ProgramId)
	assert.EqualValues(t, "final_review", results[1].ProgramClasses[0].Classes[0].ClassId)
}

// Helper functions
func createMockPrograms(ids ...int) []domains.Program {
	programs := make([]domains.Program, len(ids))

	for i, id := range ids {
		switch id {
		case 1:
			programs[i] = testUtils.CreateMockProgram(
				"program1",
				"Program 1",
				9,
				12,
				"Description 1",
				0,
			)
		case 2:
			programs[i] = testUtils.CreateMockProgram(
				"program2",
				"Program 2",
				10,
				11,
				"Description 2",
				1,
			)
		default:
			programs[i] = domains.Program{}
		}
	}
	return programs
}

func createMockSemesters(ids ...int) []domains.Semester {
	semesters := make([]domains.Semester, len(ids))

	for i, id := range ids {
		switch id {
		case 1:
			semesters[i] = testUtils.CreateMockSemester(
				"2020_spring",
				"Spring 2020",
			)
		case 2:
			semesters[i] = testUtils.CreateMockSemester(
				"2020_summer",
				"Summer 2020",
			)
		default:
			semesters[i] = domains.Semester{}
		}
	}
	return semesters
}
