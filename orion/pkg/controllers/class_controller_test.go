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
	"time"
)

//
// Test Get All
//
func TestGetAllClasses_Success(t *testing.T) {
	classService.mockGetAll = func() ([]domains.Class, error) {
		return createMockClasses(), nil
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

	assertMockClasses(t, 1, classes[0])
	assertMockClasses(t, 2, classes[1])
	assertMockClasses(t, 3, classes[2])
	assertMockClasses(t, 4, classes[3])
	assert.EqualValues(t, 4, len(classes))
}

//
// Test Get Class
//
func TestGetClass_Success(t *testing.T) {
	now := time.Now().UTC()
	later := now.Add(time.Hour * 24 * 60)
	classService.mockGetByClassId = func(classId string) (domains.Class, error) {
		class := createMockClass(
			"program1",
			"2020_spring",
			"class1",
			"program1_2020_spring_class1",
			"churchill",
			"3 pm - 5 pm",
			now,
			later,
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
	assert.EqualValues(t, "class1", class.ClassKey)
	assert.EqualValues(t, "program1_2020_spring_class1", class.ClassId)
	assert.EqualValues(t, "churchill", class.LocationId)
	assert.EqualValues(t, "3 pm - 5 pm", class.Times)
	assert.EqualValues(t, now, class.StartDate)
	assert.EqualValues(t, later, class.EndDate)
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
// Test Get Classes by other properties
//
func TestGetClassesByProgram_Success(t *testing.T) {
	classService.mockGetByProgramId = func(programId string) ([]domains.Class, error) {
		return createMockClasses(), nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/classes/program/program1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertMockClasses(t, 1, classes[0])
	assertMockClasses(t, 2, classes[1])
	assertMockClasses(t, 3, classes[2])
	assert.EqualValues(t, 3, len(classes))
}

func TestGetClassesBySemester_Success(t *testing.T) {
	classService.mockGetBySemesterId = func(semesterId string) ([]domains.Class, error) {
		return createMockClasses(), nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/classes/semester/2020_summer", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertMockClasses(t, 3, classes[0])
	assertMockClasses(t, 4, classes[1])
	assert.EqualValues(t, 2, len(classes))
}

func TestGetClassesByProgramAndSemester_Success(t *testing.T) {
	classService.mockGetByProgramAndSemesterId = func(programId, semesterId string) ([]domains.Class, error) {
		return createMockClasses(), nil
	}
	services.ClassService = &classService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/classes/program/program1/semester/2020_spring", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertMockClasses(t, 1, classes[0])
	assertMockClasses(t, 2, classes[1])
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
	later := now.Add(time.Hour * 24 * 60)
	class := createMockClass(
		"program1",
		"2020_spring",
		"class1",
		"program1_2020_spring_class1",
		"churchill",
		"3 pm - 5 pm",
		now,
		later,
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
	later := now.Add(time.Hour * 24 * 60)
	class := createMockClass( // Empty fields and end time is before start time
		"",
		"",
		"",
		"",
		"",
		"",
		later,
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
	later := now.Add(time.Hour * 24 * 30)
	class := createMockClass(
		"program2",
		"2020_summer",
		"",
		"program2_2020_summer",
		"churchill",
		"5 pm - 7 pm",
		now,
		later,
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
	later := now.Add(time.Hour * 24 * 60)
	class := createMockClass( // Empty fields and end time is before start time
		"program2",
		"",
		"",
		"",
		"",
		"",
		later,
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
	later := now.Add(time.Hour * 24 * 30)
	class := createMockClass(
		"program2",
		"2020_summer",
		"",
		"program2_2020_summer",
		"churchill",
		"5 pm - 7 pm",
		now,
		later,
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
func createMockClass(programId, semesterId, classKey, classId, locationId, times string, startDate, endDate time.Time) domains.Class {
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

func createMockClasses() []domains.Class {
	now := time.Now().UTC()
	later1 := now.Add(time.Hour * 24 * 30)
	later2 := now.Add(time.Hour * 24 * 31)
	later3 := now.Add(time.Hour * 24 * 60)
	return []domains.Class{
		{
			Id:         1,
			ProgramId:  "program1",
			SemesterId: "2020_spring",
			ClassKey:   "class1",
			ClassId:    "program1_2020_spring_class1",
			LocationId: "churchill",
			Times:      "3 pm - 5 pm",
			StartDate:  now,
			EndDate:    later1,
		},
		{
			Id:         2,
			ProgramId:  "program1",
			SemesterId: "2020_spring",
			ClassKey:   "class2",
			ClassId:    "program1_2020_spring_class2",
			LocationId: "churchill",
			Times:      "5 pm - 7 pm",
			StartDate:  now,
			EndDate:    later1,
		},
		{
			Id:         3,
			ProgramId:  "program1",
			SemesterId: "2020_summer",
			ClassKey:   "final_review",
			ClassId:    "program1_2020_summer_final_review",
			LocationId: "churchill",
			Times:      "5 pm - 8 pm",
			StartDate:  later1,
			EndDate:    later2,
		},
		{
			Id:         4,
			ProgramId:  "program2",
			SemesterId: "2020_summer",
			ClassKey:   "",
			ClassId:    "program2_2020_summer",
			LocationId: "churchill",
			Times:      "4 pm - 6 pm",
			StartDate:  later2,
			EndDate:    later3,
		},
	}
}

func assertMockClasses(t *testing.T, id int, class domains.Class) {
	now := time.Now().UTC()
	later1 := now.Add(time.Hour * 24 * 30)
	later2 := now.Add(time.Hour * 24 * 31)
	later3 := now.Add(time.Hour * 24 * 60)
	switch id {
	case 1:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_spring", class.SemesterId)
		assert.EqualValues(t, "class1", class.ClassKey)
		assert.EqualValues(t, "program1_2020_spring_class1", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "3 pm - 5 pm", class.Times)
		assert.EqualValues(t, now, class.StartDate)
		assert.EqualValues(t, later1, class.EndDate)
	case 2:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_spring", class.SemesterId)
		assert.EqualValues(t, "class2", class.ClassKey)
		assert.EqualValues(t, "program1_2020_spring_class2", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "5 pm - 7 pm", class.Times)
		assert.EqualValues(t, now, class.StartDate)
		assert.EqualValues(t, later1, class.EndDate)
	case 3:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_summer", class.SemesterId)
		assert.EqualValues(t, "final_review", class.ClassKey)
		assert.EqualValues(t, "program1_2020_summer_final_review", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "5 pm - 8 pm", class.Times)
		assert.EqualValues(t, later1, class.StartDate)
		assert.EqualValues(t, later2, class.EndDate)
	case 4:
		assert.EqualValues(t, "program2", class.ProgramId)
		assert.EqualValues(t, "2020_summer", class.SemesterId)
		assert.EqualValues(t, "", class.ClassKey)
		assert.EqualValues(t, "program2_2020_summer", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "4 pm - 6 pm", class.Times)
		assert.EqualValues(t, later2, class.StartDate)
		assert.EqualValues(t, later3, class.EndDate)
	}
}

func createBodyFromClass(class domains.Class) io.Reader {
	marshal, err := json.Marshal(class)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
