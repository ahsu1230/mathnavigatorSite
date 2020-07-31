package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get All
//
func TestGetAllSemesters_Success(t *testing.T) {
	testUtils.SemesterRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return []domains.Semester{
			testUtils.CreateMockSemester("2020_fall", "Fall 2020", 1),
			testUtils.CreateMockSemester("2020_winter", "Winter 2020", 2),
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
	assert.EqualValues(t, 1, semesters[0].Ordering)
	assert.EqualValues(t, "2020_winter", semesters[1].SemesterId)
	assert.EqualValues(t, "Winter 2020", semesters[1].Title)
	assert.EqualValues(t, 2, semesters[1].Ordering)
	assert.EqualValues(t, 2, len(semesters))
}

//
// Test Get Published
//
func TestGetPublishedSemesters_Success(t *testing.T) {
	testUtils.SemesterRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return []domains.Semester{
			testUtils.CreateMockSemester("2020_fall", "Fall 2020", 1),
			testUtils.CreateMockSemester("2020_winter", "Winter 2020", 2),
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
	assert.EqualValues(t, 1, semesters[0].Ordering)
	assert.EqualValues(t, "2020_winter", semesters[1].SemesterId)
	assert.EqualValues(t, "Winter 2020", semesters[1].Title)
	assert.EqualValues(t, 2, semesters[1].Ordering)
	assert.EqualValues(t, 2, len(semesters))
}

//
// Test Get Semester
//
func TestGetSemester_Success(t *testing.T) {
	testUtils.SemesterRepo.MockSelectBySemesterId = func(semesterId string) (domains.Semester, error) {
		semester := testUtils.CreateMockSemester("2020_fall", "Fall 2020", 1)
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
	assert.EqualValues(t, 1, semester.Ordering)
}

func TestGetSemester_Failure(t *testing.T) {
	testUtils.SemesterRepo.MockSelectBySemesterId = func(semesterId string) (domains.Semester, error) {
		return domains.Semester{}, errors.New("not found")
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
func TestCreateSemester_Success(t *testing.T) {
	testUtils.SemesterRepo.MockInsert = func(semester domains.Semester) error {
		return nil
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester("2020_fall", "Fall 2020", 1)
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateSemester_Failure(t *testing.T) {
	// no mock needed
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester("2020_fall", "", 1) // Empty title
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateSemester_Success(t *testing.T) {
	testUtils.SemesterRepo.MockUpdate = func(semesterId string, semester domains.Semester) error {
		return nil // Successful update
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester("2020_winter", "Winter 2020", 1)
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateSemester_Invalid(t *testing.T) {
	// no mock needed
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester("2020_winter", "", 1) // Empty title
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateSemester_Failure(t *testing.T) {
	testUtils.SemesterRepo.MockUpdate = func(semesterId string, semester domains.Semester) error {
		return errors.New("not found")
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semester := testUtils.CreateMockSemester("2020_winter", "Winter 2020", 1)
	body := createBodyFromSemester(semester)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Publish
//
func TestPublishSemesters_Success(t *testing.T) {
	testUtils.SemesterRepo.MockPublish = func(semesterId []string) error {
		return nil // Return no error, successful publish!
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semesterIds := []string{"2020_fall"}
	marshal, err := json.Marshal(semesterIds)
	if err != nil {
		panic(err)
	}
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/publish", bytes.NewBuffer(marshal))

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestPublishSemesters_Failure(t *testing.T) {
	testUtils.SemesterRepo.MockPublish = func(semesterId []string) error {
		return errors.New("not found")
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	semesterIds := []string{"2020_fall"}
	marshal, err := json.Marshal(semesterIds)
	if err != nil {
		panic(err)
	}
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/semesters/publish", bytes.NewBuffer(marshal))

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteSemester_Success(t *testing.T) {
	testUtils.SemesterRepo.MockDelete = func(semesterId string) error {
		return nil // Return no error, successful delete!
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/semesters/semester/some_semester", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteSemester_Failure(t *testing.T) {
	testUtils.SemesterRepo.MockDelete = func(semesterId string) error {
		return errors.New("not found")
	}
	repos.SemesterRepo = &testUtils.SemesterRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/semesters/semester/some_semester", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
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
