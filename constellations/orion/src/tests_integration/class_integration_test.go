package tests_integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 4 Classes and GetAll(false)
func TestCreateClasses(t *testing.T) {
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
func TestUniqueClassId(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	class2 := createClass(1)
	body1 := utils.CreateJsonBody(&class1)
	body2 := utils.CreateJsonBody(&class2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusBadRequest, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

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
func TestGetClassesByProgram(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	createAllClasses(t)

	// Call GetClassesByProgram()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/program/program1", nil)

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
func TestGetClassesBySemester(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	createAllClasses(t)

	// Call GetClassesBySemester()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/semester/2020_summer", nil)

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
func TestGetClassesByProgramAndSemester(t *testing.T) {
	createAllProgramsSemestersLocations(t)
	createAllClasses(t)

	// Call GetClassesByProgramAndSemester()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/program/program1/semester/2020_spring", nil)

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
func TestUpdateClass(t *testing.T) {
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

// Test: Create 2 Classes and Publish 1
func TestPublishClasses(t *testing.T) {
	// Create
	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	class2 := createClass(2)
	body1 := utils.CreateJsonBody(&class1)
	body2 := utils.CreateJsonBody(&class2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get All Published
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var classes1 []domains.Class
	if err := json.Unmarshal(recorder3.Body.Bytes(), &classes1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(classes1))

	// Publish
	classIds := []string{"program1_2020_spring_class2"}
	body3 := utils.CreateJsonBody(&classIds)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/publish", body3)
	assert.EqualValues(t, http.StatusNoContent, recorder4.Code)

	// Get All Published
	recorder5 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/all?published=true", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var classes2 []domains.Class
	if err := json.Unmarshal(recorder5.Body.Bytes(), &classes2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 2, classes2[0])
	assert.EqualValues(t, 1, len(classes2))

	// Get All Unpublished
	recorder6 := utils.SendHttpRequest(t, http.MethodGet, "/api/classes/unpublished", nil)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)

	// Validate results
	var unpublishedClasses []domains.Class
	if err := json.Unmarshal(recorder6.Body.Bytes(), &unpublishedClasses); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertClass(t, 1, unpublishedClasses[0])
	assert.EqualValues(t, 1, len(unpublishedClasses))

	resetClassTables(t)
}

// Test: Create 1 Class, Delete it, GetByClassId()
func TestDeleteClass(t *testing.T) {
	// Create
	createAllProgramsSemestersLocations(t)
	class1 := createClass(1)
	body1 := utils.CreateJsonBody(&class1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/classes/class/program1_2020_spring_class1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

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
			ProgramId:       "program1",
			SemesterId:      "2020_spring",
			ClassKey:        domains.NewNullString("class1"),
			ClassId:         "program1_2020_spring_class1",
			LocationId:      "wchs",
			TimesStr:        "3 pm - 5 pm",
			GoogleClassCode: domains.NewNullString("ab12cd34"),
			FullState:       0,
			PricePerSession: domains.NewNullUint(0),
			PriceLumpSum:    domains.NewNullUint(10),
			PaymentNotes:    domains.NewNullString(""),
		}
	case 2:
		return domains.Class{
			ProgramId:       "program1",
			SemesterId:      "2020_spring",
			ClassKey:        domains.NewNullString("class2"),
			ClassId:         "program1_2020_spring_class2",
			LocationId:      "wchs",
			TimesStr:        "5 pm - 7 pm",
			GoogleClassCode: domains.NewNullString("ab12cd35"),
			FullState:       1,
			PricePerSession: domains.NewNullUint(10),
			PriceLumpSum:    domains.NewNullUint(0),
			PaymentNotes:    domains.NewNullString("notes2"),
		}
	case 3:
		return domains.Class{
			ProgramId:       "program1",
			SemesterId:      "2020_summer",
			ClassKey:        domains.NewNullString("final_review"),
			ClassId:         "program1_2020_summer_final_review",
			LocationId:      "wchs",
			TimesStr:        "5 pm - 8 pm",
			GoogleClassCode: domains.NewNullString("ab12cd36"),
			FullState:       2,
			PricePerSession: domains.NewNullUint(0),
			PriceLumpSum:    domains.NewNullUint(20),
			PaymentNotes:    domains.NewNullString("notes3"),
		}
	case 4:
		return domains.Class{
			ProgramId:       "program2",
			SemesterId:      "2020_summer",
			ClassKey:        domains.NewNullString(""),
			ClassId:         "program2_2020_summer",
			LocationId:      "wchs",
			TimesStr:        "4 pm - 6 pm",
			GoogleClassCode: domains.NewNullString("ab12cd37"),
			FullState:       0,
			PricePerSession: domains.NewNullUint(20),
			PriceLumpSum:    domains.NewNullUint(0),
			PaymentNotes:    domains.NewNullString("notes4"),
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
	utils.SendCreateProgram(t, true, "program1", "Program1", 1, 3, "description1", domains.FEATURED_NONE)
	utils.SendCreateProgram(t, true, "program2", "Program2", 6, 8, "description2", domains.FEATURED_POPULAR)
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)
	utils.SendCreateSemester(t, true, domains.SUMMER, 2020)
	utils.SendCreateLocationWCHS(t)
}

func assertClass(t *testing.T, id int, class domains.Class) {
	switch id {
	case 1:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_spring", class.SemesterId)
		assert.EqualValues(t, "class1", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_spring_class1", class.ClassId)
		assert.EqualValues(t, "wchs", class.LocationId)
		assert.EqualValues(t, "3 pm - 5 pm", class.TimesStr)
		assert.EqualValues(t, "ab12cd34", class.GoogleClassCode.String)
		assert.EqualValues(t, 0, class.FullState)
		assert.EqualValues(t, 0, class.PricePerSession.Uint)
		assert.EqualValues(t, 10, class.PriceLumpSum.Uint)
		assert.EqualValues(t, "", class.PaymentNotes.String)
	case 2:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_spring", class.SemesterId)
		assert.EqualValues(t, "class2", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_spring_class2", class.ClassId)
		assert.EqualValues(t, "wchs", class.LocationId)
		assert.EqualValues(t, "5 pm - 7 pm", class.TimesStr)
		assert.EqualValues(t, "ab12cd35", class.GoogleClassCode.String)
		assert.EqualValues(t, 1, class.FullState)
		assert.EqualValues(t, 10, class.PricePerSession.Uint)
		assert.EqualValues(t, 0, class.PriceLumpSum.Uint)
		assert.EqualValues(t, "notes2", class.PaymentNotes.String)
	case 3:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_summer", class.SemesterId)
		assert.EqualValues(t, "final_review", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_summer_final_review", class.ClassId)
		assert.EqualValues(t, "wchs", class.LocationId)
		assert.EqualValues(t, "5 pm - 8 pm", class.TimesStr)
		assert.EqualValues(t, "ab12cd36", class.GoogleClassCode.String)
		assert.EqualValues(t, 2, class.FullState)
		assert.EqualValues(t, 0, class.PricePerSession.Uint)
		assert.EqualValues(t, 20, class.PriceLumpSum.Uint)
		assert.EqualValues(t, "notes3", class.PaymentNotes.String)
	case 4:
		assert.EqualValues(t, "program2", class.ProgramId)
		assert.EqualValues(t, "2020_summer", class.SemesterId)
		assert.EqualValues(t, "", class.ClassKey.String)
		assert.EqualValues(t, "program2_2020_summer", class.ClassId)
		assert.EqualValues(t, "wchs", class.LocationId)
		assert.EqualValues(t, "4 pm - 6 pm", class.TimesStr)
		assert.EqualValues(t, "ab12cd37", class.GoogleClassCode.String)
		assert.EqualValues(t, 0, class.FullState)
		assert.EqualValues(t, 20, class.PricePerSession.Uint)
		assert.EqualValues(t, 0, class.PriceLumpSum.Uint)
		assert.EqualValues(t, "notes4", class.PaymentNotes.String)
	}
}

func resetClassTables(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_CLASSES)
	utils.ResetTable(t, domains.TABLE_PROGRAMS)
	utils.ResetTable(t, domains.TABLE_SEMESTERS)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}
