package tests_integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

var now = time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
var later1 = now.Add(time.Hour * 24 * 30)
var later2 = now.Add(time.Hour * 24 * 31)
var later3 = now.Add(time.Hour * 24 * 60)

// Test: Create 4 Classes and GetAll(false)
func Test_CreateClasses(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	createAllClasses(t)

	// Call Get All!
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/all", nil)

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

	resetClassTables(t)
}

// Test: Create 2 Classes with same classId. Then GetByClassId()
func Test_UniqueClassId(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	class2 := createClass(1)
	body1 := utils.CreateJsonBody(&class1)
	body2 := utils.CreateJsonBody(&class2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusInternalServerError, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "Duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var class domains.Class
	if err := json.Unmarshal(recorder3.Body.Bytes(), &class); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 1, class)

	resetClassTables(t)
}

// Test: Create 4 Classes and GetClassesByProgram()
func Test_GetClassesByProgram(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	createAllClasses(t)

	// Call GetClassesByProgram()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/classes/program/program1", nil)

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

	resetClassTables(t)
}

// Test: Create 4 Classes and GetClassesBySemester()
func Test_GetClassesBySemester(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	createAllClasses(t)

	// Call GetClassesBySemester()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/classes/semester/2020_summer", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 3, classes[0])
	assertClass(t, 4, classes[1])
	assert.EqualValues(t, 2, len(classes))

	resetClassTables(t)
}

// Test: Create 4 Classes and GetClassesByProgramAndSemester()
func Test_GetClassesByProgramAndSemester(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	createAllClasses(t)

	// Call GetClassesByProgramAndSemester()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/classes/program/program1/semester/2020_spring", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 1, classes[0])
	assertClass(t, 2, classes[1])
	assert.EqualValues(t, 2, len(classes))

	resetClassTables(t)
}

// Test: Create 1 Class, Update it, GetByClassId()
func Test_UpdateClass(t *testing.T) {
	// Create 1 Class
	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	body1 := utils.CreateJsonBody(&class1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedClass := createClass(2)
	updatedBody := utils.CreateJsonBody(&updatedClass)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/class/program1_2020_spring_class1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/class/program1_2020_spring_class2", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var class domains.Class
	if err := json.Unmarshal(recorder4.Body.Bytes(), &class); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 2, class)

	resetClassTables(t)
}

// Test: Create 1 Class, Delete it, GetByClassId()
func Test_DeleteClass(t *testing.T) {
	// Create
	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	body1 := utils.CreateJsonBody(&class1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/classes/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	resetClassTables(t)
}

// Helper methods
func createClass(id int) domains.Class {
	switch id {
	case 1:
		return domains.Class{
			ProgramId:  "program1",
			SemesterId: "2020_spring",
			ClassKey:   domains.NewNullString("class1"),
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
			ClassKey:   domains.NewNullString("class2"),
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
			ClassKey:   domains.NewNullString("final_review"),
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
			ClassKey:   domains.NewNullString(""),
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
		body := utils.CreateJsonBody(&class)
		recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
}

func createAllProgramsSemestersLocations(t *testing.T) {
	program1 := createProgram("program1", "Program1", 1, 3, "description1", 0)
	program2 := createProgram("program2", "Program2", 6, 8, "description2", 1)
	semester1 := createSemester("2020_spring", "Spring 2020")
	semester2 := createSemester("2020_summer", "Summer 2020")
	location1 := createLocation("churchill", "11300 Gainsborough Road", "Potomac", "MD", "20854", "Room 100")

	body1 := utils.CreateJsonBody(&program1)
	body2 := utils.CreateJsonBody(&program2)
	body3 := utils.CreateJsonBody(&semester1)
	body4 := utils.CreateJsonBody(&semester2)
	body5 := utils.CreateJsonBody(&location1)

	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body3)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body4)
	recorder5 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body5)

	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)
}

func assertClass(t *testing.T, id int, class domains.Class) {
	switch id {
	case 1:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_spring", class.SemesterId)
		assert.EqualValues(t, "class1", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_spring_class1", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "3 pm - 5 pm", class.Times)
		assert.EqualValues(t, now, class.StartDate)
		assert.EqualValues(t, later1, class.EndDate)
	case 2:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_spring", class.SemesterId)
		assert.EqualValues(t, "class2", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_spring_class2", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "5 pm - 7 pm", class.Times)
		assert.EqualValues(t, now, class.StartDate)
		assert.EqualValues(t, later1, class.EndDate)
	case 3:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_summer", class.SemesterId)
		assert.EqualValues(t, "final_review", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_summer_final_review", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "5 pm - 8 pm", class.Times)
		assert.EqualValues(t, later1, class.StartDate)
		assert.EqualValues(t, later2, class.EndDate)
	case 4:
		assert.EqualValues(t, "program2", class.ProgramId)
		assert.EqualValues(t, "2020_summer", class.SemesterId)
		assert.EqualValues(t, "", class.ClassKey.String)
		assert.EqualValues(t, "program2_2020_summer", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "4 pm - 6 pm", class.Times)
		assert.EqualValues(t, later2, class.StartDate)
		assert.EqualValues(t, later3, class.EndDate)
	}
}

func resetClassTables(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_CLASSES)
	utils.ResetTable(t, domains.TABLE_PROGRAMS)
	utils.ResetTable(t, domains.TABLE_SEMESTERS)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}
