package controllers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

func TestOneSemesterOneProgramOneClassSuccess(t *testing.T) {
	// Mock 1 program, 1 semester, 1 class
	testUtils.ProgramRepo.MockSelectAll = func(context.Context) ([]domains.Program, error) {
		return createMockPrograms(1), nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(context.Context) ([]domains.Semester, error) {
		return createMockSemesters(1), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(context.Context, bool) ([]domains.Class, error) {
		return createMockClasses(1), nil
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
	assert.EqualValues(t, "program1_2020_spring_class1", results[0].ProgramClasses[0].Classes[0].ClassId)
}

func TestOneSemesterOneProgramOneClassFailure(t *testing.T) {
	// Mock 1 program, 1 semester, no classes created
	testUtils.ProgramRepo.MockSelectAll = func(context.Context) ([]domains.Program, error) {
		return createMockPrograms(1), nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(context.Context) ([]domains.Semester, error) {
		return createMockSemesters(1), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(context.Context, bool) ([]domains.Class, error) {
		return []domains.Class{}, appErrors.MockDbNoRowsError()
		// return []domains.Class{}, appErrors.MockMySQLUnknownError()
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classesbysemesters", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

func TestGetClassesAndProgramsBySemesterSuccess(t *testing.T) {
	// Mock 2 programs, 2 semesters, 2 classes
	testUtils.ProgramRepo.MockSelectAll = func(context.Context) ([]domains.Program, error) {
		return createMockPrograms(1, 2), nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(context.Context) ([]domains.Semester, error) {
		return createMockSemesters(1, 2), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(context.Context, bool) ([]domains.Class, error) {
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
	assert.EqualValues(t, "program1_2020_spring_class1", results[0].ProgramClasses[0].Classes[0].ClassId)
	assert.EqualValues(t, "program1_2020_spring_class2", results[0].ProgramClasses[0].Classes[1].ClassId)
	assert.EqualValues(t, "2020_summer", results[1].Semester.SemesterId)
	assert.EqualValues(t, "program1", results[1].ProgramClasses[0].ProgramObj.ProgramId)
	assert.EqualValues(t, "program1_2020_summer_final_review", results[1].ProgramClasses[0].Classes[0].ClassId)
	assert.EqualValues(t, "program2", results[1].ProgramClasses[1].ProgramObj.ProgramId)
}

func TestProgramWithNoClassSuccess(t *testing.T) {
	// Mock 1 semester, 1 program, 0 class, where program has no class
	testUtils.ProgramRepo.MockSelectAll = func(context.Context) ([]domains.Program, error) {
		return createMockPrograms(1), nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(context.Context) ([]domains.Semester, error) {
		return createMockSemesters(1), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(context.Context, bool) ([]domains.Class, error) {
		return []domains.Class{}, nil
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
}

func TestSemesterWithNoProgramsSuccess(t *testing.T) {
	// Mock one semester with no programs or classes
	testUtils.ProgramRepo.MockSelectAll = func(context.Context) ([]domains.Program, error) {
		return []domains.Program{}, nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(context.Context) ([]domains.Semester, error) {
		return createMockSemesters(1), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(context.Context, bool) ([]domains.Class, error) {
		return []domains.Class{}, nil
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
				domains.SUBJECT_MATH,
				"Description 1",
				domains.FEATURED_NONE,
			)
		case 2:
			programs[i] = testUtils.CreateMockProgram(
				"program2",
				"Program 2",
				10,
				11,
				domains.SUBJECT_MATH,
				"Description 2",
				domains.FEATURED_POPULAR,
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
				domains.SPRING,
				2020,
			)
		case 2:
			semesters[i] = testUtils.CreateMockSemester(
				domains.SUMMER,
				2020,
			)
		default:
			semesters[i] = domains.Semester{}
		}
	}
	return semesters
}
