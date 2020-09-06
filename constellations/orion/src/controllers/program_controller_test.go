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
				Name:        "Program1",
				Grade1:      2,
				Grade2:      3,
				Description: "Description1",
				Featured:    0,
			},
			{
				Id:          2,
				ProgramId:   "prog2",
				Name:        "Program2",
				Grade1:      2,
				Grade2:      3,
				Description: "Description2",
				Featured:    1,
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
	assert.EqualValues(t, "Program1", programs[0].Name)
	assert.EqualValues(t, "prog1", programs[0].ProgramId)
	assert.EqualValues(t, "Program2", programs[1].Name)
	assert.EqualValues(t, "prog2", programs[1].ProgramId)
	assert.EqualValues(t, 2, len(programs))
}

//
// Test Get Program
//
func TestGetProgramSuccess(t *testing.T) {
	testUtils.ProgramRepo.MockSelectByProgramId = func(context.Context, string) (domains.Program, error) {
		program := testUtils.CreateMockProgram("prog1", "Program1", 2, 3, "descript1", 0)
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
	assert.EqualValues(t, "Program1", program.Name)
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
	program := testUtils.CreateMockProgram("prog1", "Program1", 2, 3, "descript1", 0)
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
	program := testUtils.CreateMockProgram("prog1", "", 2, 3, "descript1", 0) // Empty Name!
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
	program := testUtils.CreateMockProgram("prog2", "Program2", 2, 3, "descript2", 0)
	body := createBodyFromProgram(program)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/programs/program/prog1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateProgramInvalid(t *testing.T) {
	// no mock needed
	repos.ProgramRepo = &testUtils.ProgramRepo

	// Create new HTTP request to endpoint
	program := testUtils.CreateMockProgram("prog2", "", 2, 3, "descript2", 0) // Empty Name!
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
	program := testUtils.CreateMockProgram("prog2", "Program2", 2, 3, "descript2", 0)
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

func TestGetAllProgramStates(t *testing.T) {
	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/programs/states", nil)

	//Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var programStates []string
	if err := json.Unmarshal(recorder.Body.Bytes(), &programStates); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "normal", programStates[0])
	assert.EqualValues(t, "popular", programStates[1])
	assert.EqualValues(t, "new", programStates[2])
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
