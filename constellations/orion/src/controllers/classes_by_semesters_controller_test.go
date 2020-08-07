package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

func TestOneSemesterOneProgramOneClass_Success(t *testing.T) {
	// Mock 1 program, 1 semester, 1 class
	testUtils.ProgramRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Program, error) {
		return createMockPrograms(1), nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return createMockSemesters(1), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Class, error) {
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

func TestOneSemesterOneProgramOneClass_Failure(t *testing.T) {
	// Mock 1 program, 1 semester, no classes created
	testUtils.ProgramRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Program, error) {
		return createMockPrograms(1), nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return createMockSemesters(1), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Class, error) {
		return []domains.Class{}, errors.New("error in class")
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classesbysemesters", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

func TestGetClassesAndProgramsBySemester_Success(t *testing.T) {
	// Mock 2 programs, 2 semesters, 2 classes
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
	assert.EqualValues(t, "program1_2020_spring_class1", results[0].ProgramClasses[0].Classes[0].ClassId)
	assert.EqualValues(t, "program1_2020_spring_class2", results[0].ProgramClasses[0].Classes[1].ClassId)
	assert.EqualValues(t, "2020_summer", results[1].Semester.SemesterId)
	assert.EqualValues(t, "program1", results[1].ProgramClasses[0].ProgramObj.ProgramId)
	assert.EqualValues(t, "program1_2020_summer_final_review", results[1].ProgramClasses[0].Classes[0].ClassId)
	assert.EqualValues(t, "program2", results[1].ProgramClasses[1].ProgramObj.ProgramId)
}

func TestProgramWithNoClass_Success(t *testing.T) {
	// Mock 1 semester, 1 program, 0 class, where program has no class
	testUtils.ProgramRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Program, error) {
		return createMockPrograms(1), nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return createMockSemesters(1), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Class, error) {
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

func TestSemesterWithNoPrograms_Success(t *testing.T) {
	// Mock one semester with no programs or classes
	testUtils.ProgramRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Program, error) {
		return []domains.Program{}, nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return createMockSemesters(1), nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	testUtils.ClassRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Class, error) {
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
				1,
			)
		case 2:
			semesters[i] = testUtils.CreateMockSemester(
				"2020_summer",
				"Summer 2020",
				2,
			)
		default:
			semesters[i] = domains.Semester{}
		}
	}
	return semesters
}
