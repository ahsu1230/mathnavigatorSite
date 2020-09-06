package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get All
//
func TestGetAllSemestersSuccess(t *testing.T) {
	testUtils.SemesterRepo.MockSelectAll = func(context.Context) ([]domains.Semester, error) {
		return []domains.Semester{
			testUtils.CreateMockSemester(domains.FALL, 2020),
			testUtils.CreateMockSemester(domains.WINTER, 2020),
		}, nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/semesters/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var semesters []domains.Semester
	if err := json.Unmarshal(recorder.Body.Bytes(), &semesters); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_fall", semesters[0].SemesterId)
	assert.EqualValues(t, "Fall 2020", semesters[0].Title)
	assert.EqualValues(t, "2020_winter", semesters[1].SemesterId)
	assert.EqualValues(t, "Winter 2020", semesters[1].Title)
	assert.EqualValues(t, 2, len(semesters))
}

//
// Test Get Published
//
func TestGetPublishedSemestersSuccess(t *testing.T) {
	testUtils.SemesterRepo.MockSelectAll = func(context.Context) ([]domains.Semester, error) {
		return []domains.Semester{
			testUtils.CreateMockSemester(domains.FALL, 2020),
			testUtils.CreateMockSemester(domains.WINTER, 2020),
		}, nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/semesters/all?published=true", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var semesters []domains.Semester
	if err := json.Unmarshal(recorder.Body.Bytes(), &semesters); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_fall", semesters[0].SemesterId)
	assert.EqualValues(t, "Fall 2020", semesters[0].Title)
	assert.EqualValues(t, "2020_winter", semesters[1].SemesterId)
	assert.EqualValues(t, "Winter 2020", semesters[1].Title)
	assert.EqualValues(t, 2, len(semesters))
}

//
// Test Get Semester
//
func TestGetSemesterSuccess(t *testing.T) {
	testUtils.SemesterRepo.MockSelectBySemesterId = func(context.Context, string) (domains.Semester, error) {
		semester := testUtils.CreateMockSemester(domains.FALL, 2020)
		return semester, nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_fall", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var semester domains.Semester
	if err := json.Unmarshal(recorder.Body.Bytes(), &semester); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_fall", semester.SemesterId)
	assert.EqualValues(t, "Fall 2020", semester.Title)
}

func TestGetSemesterFailure(t *testing.T) {
	testUtils.SemesterRepo.MockSelectBySemesterId = func(context.Context, string) (domains.Semester, error) {
		return domains.Semester{}, appErrors.MockDbNoRowsError()
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/semesters/semester/2020_winter", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateSemesterSuccess(t *testing.T) {
	testUtils.SemesterRepo.MockInsert = func(context.Context, domains.Semester) error {
		return nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester(domains.FALL, 2020)
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateSemesterFailure(t *testing.T) {
	// no mock needed
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester("autumn", 2020) // Invalid season
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateSemesterSuccess(t *testing.T) {
	testUtils.SemesterRepo.MockUpdate = func(context.Context, string, domains.Semester) error {
		return nil // Successful update
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester(domains.WINTER, 2020)
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateSemesterInvalid(t *testing.T) {
	// no mock needed
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester("autumn", 2020) // Invalid season
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateSemesterFailure(t *testing.T) {
	testUtils.SemesterRepo.MockUpdate = func(context.Context, string, domains.Semester) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester(domains.WINTER, 2020)
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteSemesterSuccess(t *testing.T) {
	testUtils.SemesterRepo.MockDelete = func(context.Context, string) error {
		return nil // Return no error, successful delete!
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/semesters/semester/some_semester", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteSemesterFailure(t *testing.T) {
	testUtils.SemesterRepo.MockDelete = func(context.Context, string) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/semesters/semester/some_semester", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Helper Methods
//

func createBodyFromSemester(semester domains.Semester) io.Reader {
	marshal, err := json.Marshal(&semester)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
