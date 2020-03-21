package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

//
// Test Get All
//
func TestGetAllClasses_Success(t *testing.T) {
	now := time.Now().UTC()
	classService.mockGetAll = func() ([]domains.Class, error) {
		return []domains.Class{
			{
				Id:         1,
				ProgramId:  "program1",
				SemesterId: "2020_spring",
				ClassKey:   sql.NullString{String: "class1", Valid: true},
				ClassId:    "program1_2020_spring_class1",
				LocationId: "churchill",
				Times:      "3 pm - 5 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 60),
			},
			{
				Id:         2,
				ProgramId:  "program2",
				SemesterId: "2020_summer",
				ClassKey:   sql.NullString{},
				ClassId:    "program2_2020_summer",
				LocationId: "churchill",
				Times:      "5 pm - 7 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 30),
			},
		}, nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "program1", classes[0].ProgramId)
	assert.EqualValues(t, "2020_spring", classes[0].SemesterId)
	assert.EqualValues(t, sql.NullString{String: "class1", Valid: true}, classes[0].ClassKey)
	assert.EqualValues(t, "program1_2020_spring_class1", classes[0].ClassId)
	assert.EqualValues(t, "churchill", classes[0].LocationId)
	assert.EqualValues(t, "3 pm - 5 pm", classes[0].Times)
	assert.EqualValues(t, now, classes[0].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*60), classes[0].EndDate)

	assert.EqualValues(t, "program2", classes[1].ProgramId)
	assert.EqualValues(t, "2020_summer", classes[1].SemesterId)
	assert.EqualValues(t, sql.NullString{}, classes[1].ClassKey)
	assert.EqualValues(t, "program2_2020_summer", classes[1].ClassId)
	assert.EqualValues(t, "churchill", classes[1].LocationId)
	assert.EqualValues(t, "5 pm - 7 pm", classes[1].Times)
	assert.EqualValues(t, now, classes[1].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*30), classes[1].EndDate)

	assert.EqualValues(t, 2, len(classes))
}

//
// Test Get Class
//
func TestGetClass_Success(t *testing.T) {
	now := time.Now().UTC()
	classService.mockGetByClassId = func(classId string) (domains.Class, error) {
		class := createMockClass(
			"program1",
			"2020_spring",
			sql.NullString{String: "class1", Valid: true},
			"program1_2020_spring_class1",
			"churchill",
			"3 pm - 5 pm",
			now,
			now.Add(time.Hour*24*60),
		)
		return class, nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program1_2020_spring_class1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var class domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &class); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "program1", class.ProgramId)
	assert.EqualValues(t, "2020_spring", class.SemesterId)
	assert.EqualValues(t, sql.NullString{String: "class1", Valid: true}, class.ClassKey)
	assert.EqualValues(t, "program1_2020_spring_class1", class.ClassId)
	assert.EqualValues(t, "churchill", class.LocationId)
	assert.EqualValues(t, "3 pm - 5 pm", class.Times)
	assert.EqualValues(t, now, class.StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*60), class.EndDate)
}

func TestGetClass_Failure(t *testing.T) {
	classService.mockGetByClassId = func(classId string) (domains.Class, error) {
		return domains.Class{}, errors.New("not found")
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program2_2020_summer", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Get Classes By Program and Semester
//
func TestGetClassesByProgram_Success(t *testing.T) {
	now := time.Now().UTC()
	classService.mockGetByProgramId = func() ([]domains.Class, error) {
		return []domains.Class{
			{
				Id:         1,
				ProgramId:  "program1",
				SemesterId: "2020_spring",
				ClassKey:   sql.NullString{String: "class1", Valid: true},
				ClassId:    "program1_2020_spring_class1",
				LocationId: "churchill",
				Times:      "3 pm - 5 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 60),
			},
			{
				Id:         2,
				ProgramId:  "program2",
				SemesterId: "2020_summer",
				ClassKey:   sql.NullString{},
				ClassId:    "program2_2020_summer",
				LocationId: "churchill",
				Times:      "5 pm - 7 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 30),
			},
			{
				Id:         3,
				ProgramId:  "program1",
				SemesterId: "2020_fall",
				ClassKey:   sql.NullString{String: "class2", Valid: true},
				ClassId:    "program1_2020_fall_class2",
				LocationId: "churchill",
				Times:      "5 pm - 7 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 60),
			},
		}, nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program/program1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "program1", classes[0].ProgramId)
	assert.EqualValues(t, "2020_spring", classes[0].SemesterId)
	assert.EqualValues(t, sql.NullString{String: "class1", Valid: true}, classes[0].ClassKey)
	assert.EqualValues(t, "program1_2020_spring_class1", classes[0].ClassId)
	assert.EqualValues(t, "churchill", classes[0].LocationId)
	assert.EqualValues(t, "3 pm - 5 pm", classes[0].Times)
	assert.EqualValues(t, now, classes[0].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*60), classes[0].EndDate)

	assert.EqualValues(t, "program1", classes[1].ProgramId)
	assert.EqualValues(t, "2020_fall", classes[1].SemesterId)
	assert.EqualValues(t, sql.NullString{String: "class2", Valid: true}, classes[1].ClassKey)
	assert.EqualValues(t, "program1_2020_fall_class2", classes[1].ClassId)
	assert.EqualValues(t, "churchill", classes[1].LocationId)
	assert.EqualValues(t, "5 pm - 7 pm", classes[1].Times)
	assert.EqualValues(t, now, classes[1].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*60), classes[1].EndDate)

	assert.EqualValues(t, 2, len(classes))
}

func TestGetClassesBySemester_Success(t *testing.T) {
	now := time.Now().UTC()
	classService.mockGetByProgramId = func() ([]domains.Class, error) {
		return []domains.Class{
			{
				Id:         1,
				ProgramId:  "program1",
				SemesterId: "2020_spring",
				ClassKey:   sql.NullString{String: "class1", Valid: true},
				ClassId:    "program1_2020_spring_class1",
				LocationId: "churchill",
				Times:      "3 pm - 5 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 60),
			},
			{
				Id:         2,
				ProgramId:  "program2",
				SemesterId: "2020_summer",
				ClassKey:   sql.NullString{},
				ClassId:    "program2_2020_summer",
				LocationId: "churchill",
				Times:      "5 pm - 7 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 30),
			},
			{
				Id:         3,
				ProgramId:  "program1",
				SemesterId: "2020_summer",
				ClassKey:   sql.NullString{String: "final_review", Valid: true},
				ClassId:    "program1_2020_summer_final_review",
				LocationId: "churchill",
				Times:      "3 pm - 5 pm",
				StartDate:  now.Add(time.Hour * 24 * 60),
				EndDate:    now.Add(time.Hour * 24 * 61),
			},
		}, nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/semester/2020_summer", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "program2", classes[0].ProgramId)
	assert.EqualValues(t, "2020_summer", classes[0].SemesterId)
	assert.EqualValues(t, sql.NullString{}, classes[0].ClassKey)
	assert.EqualValues(t, "program2_2020_summer", classes[0].ClassId)
	assert.EqualValues(t, "churchill", classes[0].LocationId)
	assert.EqualValues(t, "5 pm - 7 pm", classes[0].Times)
	assert.EqualValues(t, now, classes[0].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*30), classes[0].EndDate)

	assert.EqualValues(t, "program1", classes[1].ProgramId)
	assert.EqualValues(t, "2020_summer", classes[1].SemesterId)
	assert.EqualValues(t, sql.NullString{String: "final_review", Valid: true}, classes[1].ClassKey)
	assert.EqualValues(t, "program1_2020_summer_final_review", classes[1].ClassId)
	assert.EqualValues(t, "churchill", classes[1].LocationId)
	assert.EqualValues(t, "3 pm - 5 pm", classes[1].Times)
	assert.EqualValues(t, now.Add(time.Hour*24*60), classes[1].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*61), classes[1].EndDate)

	assert.EqualValues(t, 2, len(classes))
}

func TestGetClassesByProgramAndSemester_Success(t *testing.T) {
	now := time.Now().UTC()
	classService.mockGetByProgramId = func() ([]domains.Class, error) {
		return []domains.Class{
			{
				Id:         1,
				ProgramId:  "program1",
				SemesterId: "2020_spring",
				ClassKey:   sql.NullString{String: "class1", Valid: true},
				ClassId:    "program1_2020_spring_class1",
				LocationId: "churchill",
				Times:      "3 pm - 5 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 60),
			},
			{
				Id:         2,
				ProgramId:  "program1",
				SemesterId: "2020_summer",
				ClassKey:   sql.NullString{},
				ClassId:    "program1_2020_summer",
				LocationId: "churchill",
				Times:      "5 pm - 7 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 30),
			},
			{
				Id:         3,
				ProgramId:  "program1",
				SemesterId: "2020_summer",
				ClassKey:   sql.NullString{String: "final_review", Valid: true},
				ClassId:    "program1_2020_summer_final_review",
				LocationId: "churchill",
				Times:      "5 pm - 8 pm",
				StartDate:  now.Add(time.Hour * 24 * 30),
				EndDate:    now.Add(time.Hour * 24 * 31),
			},
			{
				Id:         3,
				ProgramId:  "program2",
				SemesterId: "2020_summer",
				ClassKey:   sql.NullString{},
				ClassId:    "program2_2020_summer",
				LocationId: "churchill",
				Times:      "4 pm - 6 pm",
				StartDate:  now,
				EndDate:    now.Add(time.Hour * 24 * 90),
			},
		}, nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program1/2020_summer", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "program1", classes[0].ProgramId)
	assert.EqualValues(t, "2020_summer", classes[0].SemesterId)
	assert.EqualValues(t, sql.NullString{}, classes[0].ClassKey)
	assert.EqualValues(t, "program1_2020_summer", classes[0].ClassId)
	assert.EqualValues(t, "churchill", classes[0].LocationId)
	assert.EqualValues(t, "5 pm - 7 pm", classes[0].Times)
	assert.EqualValues(t, now, classes[0].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*30), classes[0].EndDate)

	assert.EqualValues(t, "program1", classes[1].ProgramId)
	assert.EqualValues(t, "2020_summer", classes[1].SemesterId)
	assert.EqualValues(t, sql.NullString{String: "final_review", Valid: true}, classes[1].ClassKey)
	assert.EqualValues(t, "program1_2020_summer_final_review", classes[1].ClassId)
	assert.EqualValues(t, "churchill", classes[1].LocationId)
	assert.EqualValues(t, "5 pm - 8 pm", classes[1].Times)
	assert.EqualValues(t, now.Add(time.Hour*24*30), classes[1].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*31), classes[1].EndDate)

	assert.EqualValues(t, 2, len(classes))
}

//
// Test Create
//
func TestCreateClass_Success(t *testing.T) {
	classService.mockCreate = func(class domains.Class) error {
		return nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	class := createMockClass(
		"program1",
		"2020_spring",
		sql.NullString{String: "class1", Valid: true},
		"program1_2020_spring_class1",
		"churchill",
		"3 pm - 5 pm",
		now,
		now.Add(time.Hour*24*60),
	)
	marshal, _ := json.Marshal(class)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateClass_Failure(t *testing.T) {
	// no mock needed
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	class := createMockClass( // Empty fields and end time is before start time
		"",
		"",
		sql.NullString{},
		"",
		"",
		"",
		now.Add(time.Hour*24*60),
		now,
	)
	marshal, _ := json.Marshal(class)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateClass_Success(t *testing.T) {
	classService.mockUpdate = func(classId string, class domains.Class) error {
		return nil // Successful update
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	class := createMockClass(
		"program2",
		"2020_summer",
		sql.NullString{},
		"program2_2020_summer",
		"churchill",
		"5 pm - 7 pm",
		now,
		now.Add(time.Hour*24*30),
	)
	body := createBodyFromClass(class)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/class/program1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateClass_Invalid(t *testing.T) {
	// no mock needed
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	class := createMockClass( // Empty fields and end time is before start time
		"program2",
		"",
		sql.NullString{},
		"",
		"",
		"",
		now.Add(time.Hour*24*60),
		now,
	)
	body := createBodyFromClass(class)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/class/program1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateClass_Failure(t *testing.T) {
	classService.mockUpdate = func(classId string, class domains.Class) error {
		return errors.New("not found")
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	class := createMockClass(
		"program2",
		"2020_summer",
		sql.NullString{},
		"program2_2020_summer",
		"churchill",
		"5 pm - 7 pm",
		now,
		now.Add(time.Hour*24*30),
	)
	body := createBodyFromClass(class)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/class/program1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteClass_Success(t *testing.T) {
	classService.mockDelete = func(classId string) error {
		return nil // Return no error, successful delete!
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/classes/v1/class/some_class", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteClass_Failure(t *testing.T) {
	classService.mockDelete = func(classId string) error {
		return errors.New("not found")
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/classes/v1/class/some_class", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockClass(programId, semesterId string, classKey sql.NullString, classId, locationId, times string, startDate, endDate time.Time) domains.Class {
	return domains.Class{
		ProgramId:  programId,
		SemesterId: semesterId,
		ClassKey:   classKey,
		ClassId:    classId,
		LocationId: locationId,
		Times:      times,
		StartDate:  startDate,
		EndDate:    endDate,
	}
}

func createBodyFromClass(class domains.Class) io.Reader {
	marshal, err := json.Marshal(class)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
