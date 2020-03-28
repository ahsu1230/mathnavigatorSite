package integration_tests

import (
	"encoding/json"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// Test: Create 4 Classes and GetAll()
func Test_CreateClasses(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	createAllClasses(t)

	// Call Get All!
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertClass(t, 1, classes[0])
	assertClass(t, 2, classes[1])
	assertClass(t, 3, classes[2])
	assertClass(t, 4, classes[3])
	assert.EqualValues(t, 4, len(classes))
}

// Test: Create 2 Classes with same classId. Then GetByClassId()
func Test_UniqueClassId(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	class1 := createClass(1)
	class2 := createClass(1)
	body1 := createJsonBody(class1)
	body2 := createJsonBody(class2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusInternalServerError, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "Duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var class domains.Class
	if err := json.Unmarshal(recorder3.Body.Bytes(), &class); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertClass(t, 1, class)
}

// Test: Create 4 Classes and GetClassesByProgram()
func Test_GetClassesByProgram(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	createAllClasses(t)

	// Call GetClassesByProgram()
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/classes/program/program1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertClass(t, 1, classes[0])
	assertClass(t, 2, classes[1])
	assertClass(t, 3, classes[2])
	assert.EqualValues(t, 3, len(classes))
}

// Test: Create 4 Classes and GetClassesBySemester()
func Test_GetClassesBySemester(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	createAllClasses(t)

	// Call GetClassesBySemester()
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/classes/semester/2020_summer", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertClass(t, 3, classes[0])
	assertClass(t, 4, classes[1])
	assert.EqualValues(t, 2, len(classes))
}

// Test: Create 4 Classes and GetClassesByProgramAndSemester()
func Test_GetClassesByProgramAndSemester(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	createAllClasses(t)

	// Call GetClassesByProgramAndSemester()
	recorder := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/classes/program/program1/semester/2020_spring", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertClass(t, 1, classes[0])
	assertClass(t, 2, classes[1])
	assert.EqualValues(t, 2, len(classes))
}

// Test: Create 1 Class, Update it, GetByClassId()
func Test_UpdateClass(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	// Create 1 Class
	class1 := createClass(1)
	body1 := createJsonBody(class1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedClass := createClass(2)
	updatedBody := createJsonBody(updatedClass)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/class/program1_2020_spring_class1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program1_2020_spring_class2", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var class domains.Class
	if err := json.Unmarshal(recorder4.Body.Bytes(), &class); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 2, class)
}

// Test: Create 1 Class, Delete it, GetByClassId()
func Test_DeleteClass(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	// Create
	class1 := createClass(1)
	body1 := createJsonBody(class1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/classes/v1/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
}

// Helper methods
func createClass(id int) domains.Class {
	now := time.Now().UTC()
	later1 := now.Add(time.Hour * 24 * 30)
	later2 := now.Add(time.Hour * 24 * 31)
	later3 := now.Add(time.Hour * 24 * 60)
	switch id {
	case 1:
		return domains.Class{
			ProgramId:  "program1",
			SemesterId: "2020_spring",
			ClassKey:   "class1",
			ClassId:    "program1_2020_spring_class1",
			LocationId: "churchill",
			Times:      "3 pm - 5 pm",
			StartDate:  now,
			EndDate:    later1,
		}
	case 2:
		return domains.Class{
			ProgramId:  "program1",
			SemesterId: "2020_spring",
			ClassKey:   "class2",
			ClassId:    "program1_2020_spring_class2",
			LocationId: "churchill",
			Times:      "5 pm - 7 pm",
			StartDate:  now,
			EndDate:    later1,
		}
	case 3:
		return domains.Class{
			ProgramId:  "program1",
			SemesterId: "2020_summer",
			ClassKey:   "final_review",
			ClassId:    "program1_2020_summer_final_review",
			LocationId: "churchill",
			Times:      "5 pm - 8 pm",
			StartDate:  later1,
			EndDate:    later2,
		}
	case 4:
		return domains.Class{
			ProgramId:  "program2",
			SemesterId: "2020_summer",
			ClassKey:   "",
			ClassId:    "program2_2020_summer",
			LocationId: "churchill",
			Times:      "4 pm - 6 pm",
			StartDate:  later2,
			EndDate:    later3,
		}
	default:
		return domains.Class{}
	}
}

func createAllClasses(t *testing.T) {
	for i := 1; i < 5; i++ {
		class := createClass(i)
		body := createJsonBody(class)
		recorder := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
}

func assertClass(t *testing.T, id int, class domains.Class) {
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
