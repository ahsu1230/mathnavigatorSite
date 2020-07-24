package controllers_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

// Test getting programs and classes by semester endpoint
func TestGetClassesAndProgramsBySemester_Success(t *testing.T) {
	// Mock selecting all programs, semesters, classes
	testUtils.ProgramRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Program, error) {
		return []domains.Program{
			{
				Id:          1,
				ProgramId:   "ap_calculus",
				Name:        "AP Calculus",
				Grade1:      2,
				Grade2:      3,
				Description: "Description1",
				Featured:    0,
			},
			{
				Id:          2,
				ProgramId:   "ap_java",
				Name:        "AP Java",
				Grade1:      2,
				Grade2:      3,
				Description: "Description2",
				Featured:    1,
			},
		}, nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	testUtils.SemesterRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return []domains.Semester{
			testUtils.CreateMockSemester("2020_fall", "Fall 2020"),
			testUtils.CreateMockSemester("2020_winter", "Winter 2020"),
		}, nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	var now = time.Now().UTC()
	var later1 = now.Add(time.Hour * 24 * 30)
	testUtils.ClassRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Class, error) {
		return []domains.Class{
			testUtils.CreateMockClass(
				"prog1",
				"2020_fall",
				"class1",
				"churchill",
				"3 pm - 5 pm",
				now,
				later1,
			),
			testUtils.CreateMockClass(
				"prog2",
				"2020_winter",
				"class2",
				"churchill",
				"5 pm - 7 pm",
				now,
				later1,
			),
		}, nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classesbysemesters/group", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}
