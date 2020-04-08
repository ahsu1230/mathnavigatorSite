package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

//
// Test Get All
//
func TestGetAllSemesters_Success(t *testing.T) {
	semesterService.mockGetAll = func() ([]domains.Semester, error) {
		return []domains.Semester{
			createMockSemester("2020_fall", "Fall 2020"),
			createMockSemester("2020_winter", "Winter 2020"),
		}, nil
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/all", nil)

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
func TestGetPublishedSemesters_Success(t *testing.T) {
	semesterService.mockGetAll = func(publishedOnly bool) ([]domains.Semester, error) {
		return []domains.Semester{
			createMockSemester("2020_fall", "Fall 2020"),
			createMockSemester("2020_winter", "Winter 2020"),
		}, nil
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/all?published=true", nil)

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
func TestGetSemester_Success(t *testing.T) {
	semesterService.mockGetBySemesterId = func(semesterId string) (domains.Semester, error) {
		semester := createMockSemester("2020_fall", "Fall 2020")
		return semester, nil
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/semester/2020_fall", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var semester domains.Semester
	if err := json.Unmarshal(recorder.Body.Bytes(), &semester); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "2020_fall", semester.SemesterId)
	assert.EqualValues(t, "Fall 2020", semester.Title)
}

func TestGetSemester_Failure(t *testing.T) {
	semesterService.mockGetBySemesterId = func(semesterId string) (domains.Semester, error) {
		return domains.Semester{}, errors.New("not found")
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/semesters/v1/semester/2020_winter", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateSemester_Success(t *testing.T) {
	semesterService.mockCreate = func(semester domains.Semester) error {
		return nil
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	semester := createMockSemester("2020_fall", "Fall 2020")
	marshal, _ := json.Marshal(semester)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateSemester_Failure(t *testing.T) {
	// no mock needed
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	semester := createMockSemester("2020_fall", "") // Empty title
	marshal, _ := json.Marshal(semester)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateSemester_Success(t *testing.T) {
	semesterService.mockUpdate = func(semesterId string, semester domains.Semester) error {
		return nil // Successful update
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	semester := createMockSemester("2020_winter", "Winter 2020")
	body := createBodyFromSemester(semester)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateSemester_Invalid(t *testing.T) {
	// no mock needed
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	semester := createMockSemester("2020_winter", "") // Empty title
	body := createBodyFromSemester(semester)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateSemester_Failure(t *testing.T) {
	semesterService.mockUpdate = func(semesterId string, semester domains.Semester) error {
		return errors.New("not found")
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	semester := createMockSemester("2020_winter", "Winter 2020")
	body := createBodyFromSemester(semester)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/semester/2020_fall", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteSemester_Success(t *testing.T) {
	semesterService.mockDelete = func(semesterId string) error {
		return nil // Return no error, successful delete!
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/semesters/v1/semester/some_semester", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteSemester_Failure(t *testing.T) {
	semesterService.mockDelete = func(semesterId string) error {
		return errors.New("not found")
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/semesters/v1/semester/some_semester", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Publish
//
func TestPublishSemesters_Success(t *testing.T) {
	semesterService.mockPublish = func(ids []uint) error {
		return nil // Return no error, successful publish!
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	semesterIds := []string{"2020_fall"}
	marshal, err := json.Marshal(semesterIds)
	if err != nil {
		panic(err)
	}
	recorder := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/publish", bytes.NewBuffer(marshal))

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestPublishSemesters_Failure(t *testing.T) {
	semesterService.mockPublish = func(ids []uint) error {
		return errors.New("not found")
	}
	services.SemesterService = &semesterService

	// Create new HTTP request to endpoint
	semesterIds := []string{"2020_fall"}
	marshal, err := json.Marshal(semesterIds)
	if err != nil {
		panic(err)
	}
	recorder := sendHttpRequest(t, http.MethodPost, "/api/semesters/v1/publish", bytes.NewBuffer(marshal))

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockSemester(semesterId string, title string) domains.Semester {
	return domains.Semester{
		SemesterId: semesterId,
		Title:      title,
	}
}

func createBodyFromSemester(semester domains.Semester) io.Reader {
	marshal, err := json.Marshal(semester)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
