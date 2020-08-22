package controllers_test

import (
	"bytes"
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
func TestGetAllClassesSuccess(t *testing.T) {
	testUtils.ClassRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Class, error) {
		return createMockClasses(1, 2, 3, 4), nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classes/all", nil)

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
// Test Get Published
//
func TestGetPublishedClassesSuccess(t *testing.T) {
	testUtils.ClassRepo.MockSelectAll = func(publishedOnly bool) ([]domains.Class, error) {
		return createMockClasses(2, 3), nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classes/all?published=true", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var classes []domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &classes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertMockClasses(t, 2, classes[0])
	assertMockClasses(t, 3, classes[1])
	assert.EqualValues(t, 2, len(classes))
}

//
// Test Get Class
//
func TestGetClassSuccess(t *testing.T) {
	testUtils.ClassRepo.MockSelectByClassId = func(classId string) (domains.Class, error) {
		return createMockClasses(1)[0], nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classes/class/program1_2020_spring_class1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var class domains.Class
	if err := json.Unmarshal(recorder.Body.Bytes(), &class); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertMockClasses(t, 1, class)
}

func TestGetClassFailure(t *testing.T) {
	testUtils.ClassRepo.MockSelectByClassId = func(classId string) (domains.Class, error) {
		return domains.Class{}, appErrors.MockDbNoRowsError()
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classes/class/program2_2020_summer", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Get Classes by other properties
//
func TestGetClassesByProgramSuccess(t *testing.T) {
	testUtils.ClassRepo.MockSelectByProgramId = func(programId string) ([]domains.Class, error) {
		return createMockClasses(1, 2, 3), nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classes/classes/program/program1", nil)

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

func TestGetClassesBySemesterSuccess(t *testing.T) {
	testUtils.ClassRepo.MockSelectBySemesterId = func(semesterId string) ([]domains.Class, error) {
		return createMockClasses(3, 4), nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classes/classes/semester/2020_summer", nil)

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

func TestGetClassesByProgramAndSemesterSuccess(t *testing.T) {
	testUtils.ClassRepo.MockSelectByProgramAndSemesterId = func(programId, semesterId string) ([]domains.Class, error) {
		return createMockClasses(1, 2), nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/classes/classes/program/program1/semester/2020_spring", nil)

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
func TestCreateClassSuccess(t *testing.T) {
	testUtils.ClassRepo.MockInsert = func(class domains.Class) error {
		return nil
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	class := createMockClasses(1)[0]
	body := createBodyFromClass(class)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateClassFailure(t *testing.T) {
	// no mock needed
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	class := testUtils.CreateMockClass("", "", "", "", "", 0, 0) // Empty fields
	body := createBodyFromClass(class)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateClassSuccess(t *testing.T) {
	testUtils.ClassRepo.MockUpdate = func(classId string, class domains.Class) error {
		return nil // Successful update
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	class := createMockClasses(2)[0]
	body := createBodyFromClass(class)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/classes/class/program1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateClassInvalid(t *testing.T) {
	// no mock needed
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	class := testUtils.CreateMockClass("", "", "", "", "", 0, 0) // Empty fields
	body := createBodyFromClass(class)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/classes/class/program1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateClassFailure(t *testing.T) {
	testUtils.ClassRepo.MockUpdate = func(classId string, class domains.Class) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	class := createMockClasses(2)[0]
	body := createBodyFromClass(class)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/classes/class/program1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Publish
//
func TestPublishClassesSuccess(t *testing.T) {
	testUtils.ClassRepo.MockPublish = func(classIds []string) []error {
		return nil // Return no error, successful publish!
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	classIds := []string{"program1_2020_spring_class1"}
	marshal, err := json.Marshal(classIds)
	if err != nil {
		t.Fatal(err)
	}
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/classes/publish", bytes.NewBuffer(marshal))

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestPublishClassesFailure(t *testing.T) {
	testUtils.ClassRepo.MockPublish = func(classIds []string) []error {
		return []error{appErrors.MockDbNoRowsError()}
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	classIds := []string{"program1_2020_spring_class1"}
	marshal, err := json.Marshal(classIds)
	if err != nil {
		panic(err)
	}
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/classes/publish", bytes.NewBuffer(marshal))

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteClassSuccess(t *testing.T) {
	testUtils.ClassRepo.MockDelete = func(classId string) error {
		return nil // Return no error, successful delete!
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/classes/class/some_class", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteClassFailure(t *testing.T) {
	testUtils.ClassRepo.MockDelete = func(classId string) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.ClassRepo = &testUtils.ClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/classes/class/some_class", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Helper Methods
//

func createMockClasses(ids ...int) []domains.Class {
	classes := make([]domains.Class, len(ids))

	for i, id := range ids {
		switch id {
		case 1:
			classes[i] = testUtils.CreateMockClass(
				"program1",
				"2020_spring",
				"class1",
				"churchill",
				"3 pm - 5 pm",
				50,
				0,
			)
		case 2:
			classes[i] = testUtils.CreateMockClass(
				"program1",
				"2020_spring",
				"class2",
				"churchill",
				"5 pm - 7 pm",
				50,
				0,
			)
		case 3:
			classes[i] = testUtils.CreateMockClass(
				"program1",
				"2020_summer",
				"final_review",
				"churchill",
				"5 pm - 8 pm",
				60,
				0,
			)
		case 4:
			classes[i] = testUtils.CreateMockClass(
				"program2",
				"2020_summer",
				"",
				"churchill",
				"4 pm - 6 pm",
				60,
				0,
			)
		default:
			classes[i] = domains.Class{}
		}
	}
	return classes
}

func assertMockClasses(t *testing.T, id int, class domains.Class) {
	switch id {
	case 1:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_spring", class.SemesterId)
		assert.EqualValues(t, "class1", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_spring_class1", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "3 pm - 5 pm", class.Times)
		assert.EqualValues(t, domains.NewNullUint(50), class.PricePerSession)
		assert.EqualValues(t, domains.NewNullUint(0), class.PriceLump)
	case 2:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_spring", class.SemesterId)
		assert.EqualValues(t, "class2", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_spring_class2", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "5 pm - 7 pm", class.Times)
		assert.EqualValues(t, domains.NewNullUint(50), class.PricePerSession)
		assert.EqualValues(t, domains.NewNullUint(0), class.PriceLump)
	case 3:
		assert.EqualValues(t, "program1", class.ProgramId)
		assert.EqualValues(t, "2020_summer", class.SemesterId)
		assert.EqualValues(t, "final_review", class.ClassKey.String)
		assert.EqualValues(t, "program1_2020_summer_final_review", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "5 pm - 8 pm", class.Times)
		assert.EqualValues(t, domains.NewNullUint(60), class.PricePerSession)
		assert.EqualValues(t, domains.NewNullUint(0), class.PriceLump)
	case 4:
		assert.EqualValues(t, "program2", class.ProgramId)
		assert.EqualValues(t, "2020_summer", class.SemesterId)
		assert.EqualValues(t, "", class.ClassKey.String)
		assert.EqualValues(t, "program2_2020_summer", class.ClassId)
		assert.EqualValues(t, "churchill", class.LocationId)
		assert.EqualValues(t, "4 pm - 6 pm", class.Times)
		assert.EqualValues(t, domains.NewNullUint(60), class.PricePerSession)
		assert.EqualValues(t, domains.NewNullUint(0), class.PriceLump)
	}
}

func createBodyFromClass(class domains.Class) io.Reader {
	marshal, err := json.Marshal(&class)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
