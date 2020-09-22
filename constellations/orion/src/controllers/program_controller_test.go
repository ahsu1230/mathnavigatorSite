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
func TestGetAllProgramsSuccess(t *testing.T) {
	testUtils.ProgramRepo.MockSelectAll = func(context.Context) ([]domains.Program, error) {
		return []domains.Program{
			{
				Id:          1,
				ProgramId:   "prog1",
				Title:       "Program1",
				Grade1:      2,
				Grade2:      3,
				Description: "Description1",
				Featured:    domains.FEATURED_NONE,
			},
			{
				Id:          2,
				ProgramId:   "prog2",
				Title:       "Program2",
				Grade1:      2,
				Grade2:      3,
				Description: "Description2",
				Featured:    domains.FEATURED_POPULAR,
			},
		}, nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/programs/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var programs []domains.Program
	if err := json.Unmarshal(recorder.Body.Bytes(), &programs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "Program1", programs[0].Title)
	assert.EqualValues(t, "prog1", programs[0].ProgramId)
	assert.EqualValues(t, "Program2", programs[1].Title)
	assert.EqualValues(t, "prog2", programs[1].ProgramId)
	assert.EqualValues(t, 2, len(programs))
}

//
// Test Get Program
//
func TestGetProgramSuccess(t *testing.T) {
	testUtils.ProgramRepo.MockSelectByProgramId = func(context.Context, string) (domains.Program, error) {
		program := testUtils.CreateMockProgram("prog1", "Program1", 2, 3, "descript1", domains.FEATURED_NONE)
		return program, nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/programs/program/prog1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var program domains.Program
	if err := json.Unmarshal(recorder.Body.Bytes(), &program); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog1", program.ProgramId)
	assert.EqualValues(t, "Program1", program.Title)
}

func TestGetProgramFailure(t *testing.T) {
	testUtils.ProgramRepo.MockSelectByProgramId = func(context.Context, string) (domains.Program, error) {
		return domains.Program{}, appErrors.MockDbNoRowsError()
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/programs/program/prog2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateProgramSuccess(t *testing.T) {
	testUtils.ProgramRepo.MockInsert = func(context.Context, domains.Program) (uint, error) {
		return 42, nil
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	program := testUtils.CreateMockProgram("prog1", "Program1", 2, 3, "descript1", domains.FEATURED_NONE)
	marshal, _ := json.Marshal(&program)
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateProgramFailure(t *testing.T) {
	// no mock needed
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	program := testUtils.CreateMockProgram("prog1", "", 2, 3, "descript1", domains.FEATURED_NONE) // Empty Name!
	marshal, _ := json.Marshal(&program)
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateProgramSuccess(t *testing.T) {
	testUtils.ProgramRepo.MockUpdate = func(context.Context, string, domains.Program) error {
		return nil // Successful update
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	program := testUtils.CreateMockProgram("prog2", "Program2", 2, 3, "descript2", domains.FEATURED_NONE)
	body := createBodyFromProgram(program)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/programs/program/prog1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateProgramInvalid(t *testing.T) {
	// no mock needed
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	program := testUtils.CreateMockProgram("prog2", "", 2, 3, "descript2", domains.FEATURED_NONE) // Empty Name!
	body := createBodyFromProgram(program)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/programs/program/prog1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateProgramFailure(t *testing.T) {
	testUtils.ProgramRepo.MockUpdate = func(context.Context, string, domains.Program) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	program := testUtils.CreateMockProgram("prog2", "Program2", 2, 3, "descript2", domains.FEATURED_NONE)
	body := createBodyFromProgram(program)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/programs/program/prog1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteProgramSuccess(t *testing.T) {
	testUtils.ProgramRepo.MockDelete = func(context.Context, string) error {
		return nil // Return no error, successful delete!
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/programs/program/some_program", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteProgramFailure(t *testing.T) {
	testUtils.ProgramRepo.MockDelete = func(context.Context, string) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/programs/program/some_program", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

func TestGetAllProgramFeatured(t *testing.T) {
	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/programs/featured", nil)

	//Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var programFeatured []string
	if err := json.Unmarshal(recorder.Body.Bytes(), &programFeatured); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "none", programFeatured[0])
	assert.EqualValues(t, "popular", programFeatured[1])
	assert.EqualValues(t, "new", programFeatured[2])
}

//
// Helper Methods
//

func createBodyFromProgram(program domains.Program) io.Reader {
	marshal, err := json.Marshal(&program)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
