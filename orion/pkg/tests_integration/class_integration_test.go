package integration_tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// Test: Create 3 Classes and GetAll()
func Test_CreateClasses(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	now := time.Now().UTC()
	class1 := createClass(
		"program1",
		"2020_spring",
		sql.NullString{String: "class1", Valid: true},
		"program1_2020_spring_class1",
		"churchill",
		"3 pm - 5 pm",
		now,
		now.Add(time.Hour*24*60),
	)
	class2 := createClass(
		"program2",
		"2020_summer",
		sql.NullString{},
		"program2_2020_summer",
		"churchill",
		"5 pm - 7 pm",
		now,
		now.Add(time.Hour*24*30),
	)
	class3 := createClass(
		"program1",
		"2020_summer",
		sql.NullString{String: "class2", Valid: true},
		"program1_2020_summer_class2",
		"churchill",
		"5 pm - 7 pm",
		now.Add(time.Hour*24*60),
		now.Add(time.Hour*24*120),
	)

	body1 := createJsonBody(class1)
	body2 := createJsonBody(class2)
	body3 := createJsonBody(class3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder4.Body.Bytes(), &classes); err != nil {
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

	assert.EqualValues(t, "program1", classes[2].ProgramId)
	assert.EqualValues(t, "2020_summer", classes[2].SemesterId)
	assert.EqualValues(t, sql.NullString{String: "class2", Valid: true}, classes[2].ClassKey)
	assert.EqualValues(t, "program1_2020_summer_class2", classes[2].ClassId)
	assert.EqualValues(t, "churchill", classes[2].LocationId)
	assert.EqualValues(t, "5 pm - 7 pm", classes[2].Times)
	assert.EqualValues(t, now.Add(time.Hour*24*60), classes[2].StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*120), classes[2].EndDate)

	assert.EqualValues(t, 3, len(classes))
}

// Test: Create 2 Classes with same classId. Then GetByClassId()
func Test_UniqueClassId(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	now := time.Now().UTC()
	class1 := createClass(
		"program1",
		"2020_spring",
		sql.NullString{String: "class1", Valid: true},
		"program1_2020_spring_class1",
		"churchill",
		"3 pm - 5 pm",
		now,
		now.Add(time.Hour*24*60),
	)
	class2 := createClass( // Same classId
		"program1",
		"2020_spring",
		sql.NullString{String: "class1", Valid: true},
		"program1_2020_spring_class1",
		"churchill",
		"3 pm - 5 pm",
		now,
		now.Add(time.Hour*24*60),
	)
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
	assert.EqualValues(t, "program1", class.ProgramId)
	assert.EqualValues(t, "2020_spring", class.SemesterId)
	assert.EqualValues(t, sql.NullString{String: "class1", Valid: true}, class.ClassKey)
	assert.EqualValues(t, "program1_2020_spring_class1", class.ClassId)
	assert.EqualValues(t, "churchill", class.LocationId)
	assert.EqualValues(t, "3 pm - 5 pm", class.Times)
	assert.EqualValues(t, now, class.StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*60), class.EndDate)
}

// Test: Create 1 Class, Update it, GetByClassId()
func Test_UpdateClass(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	// Create 1 Class
	now := time.Now().UTC()
	class1 := createClass(
		"program1",
		"2020_spring",
		sql.NullString{String: "class1", Valid: true},
		"program1_2020_spring_class1",
		"churchill",
		"3 pm - 5 pm",
		now,
		now.Add(time.Hour*24*60),
	)
	body1 := createJsonBody(class1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedClass := createClass(
		"program2",
		"2020_summer",
		sql.NullString{},
		"program2_2020_summer",
		"churchill",
		"5 pm - 7 pm",
		now,
		now.Add(time.Hour*24*30),
	)
	updatedBody := createJsonBody(updatedClass)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/class/program1_2020_spring_class1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/program2_2020_summer", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var class domains.Class
	if err := json.Unmarshal(recorder4.Body.Bytes(), &class); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "program2", class.ProgramId)
	assert.EqualValues(t, "2020_summer", class.SemesterId)
	assert.EqualValues(t, sql.NullString{}, class.ClassKey)
	assert.EqualValues(t, "program2_2020_summer", class.ClassId)
	assert.EqualValues(t, "churchill", class.LocationId)
	assert.EqualValues(t, "5 pm - 7 pm", class.Times)
	assert.EqualValues(t, now, class.StartDate)
	assert.EqualValues(t, now.Add(time.Hour*24*30), class.EndDate)
}

// Test: Create 1 Class, Delete it, GetByClassId()
func Test_DeleteClass(t *testing.T) {
	resetTable(t, domains.TABLE_CLASSES)

	// Create
	now := time.Now().UTC()
	class1 := createClass(
		"program1",
		"2020_spring",
		sql.NullString{String: "class1", Valid: true},
		"program1_2020_spring_class1",
		"churchill",
		"3 pm - 5 pm",
		now,
		now.Add(time.Hour*24*60),
	)
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
func createClass(programId, semesterId string, classKey sql.NullString, classId, locationId, times string, startDate, endDate time.Time) domains.Class {
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
