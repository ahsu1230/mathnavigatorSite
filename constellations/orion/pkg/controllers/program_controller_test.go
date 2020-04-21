package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/services"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

//
// Test Get All
//
func TestGetAllPrograms_Success(t *testing.T) {
	programService.mockGetAll = func(publishedOnly bool) ([]domains.Program, error) {
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
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/all", nil)

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
func TestGetProgram_Success(t *testing.T) {
	programService.mockGetByProgramId = func(programId string) (domains.Program, error) {
		program := createMockProgram("prog1", "Program1", 2, 3, "descript1", 0)
		return program, nil
	}
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/program/prog1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var program domains.Program
	if err := json.Unmarshal(recorder.Body.Bytes(), &program); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog1", program.ProgramId)
	assert.EqualValues(t, "Program1", program.Name)
}

func TestGetProgram_Failure(t *testing.T) {
	programService.mockGetByProgramId = func(programId string) (domains.Program, error) {
		return domains.Program{}, errors.New("not found")
	}
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/program/prog2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateProgram_Success(t *testing.T) {
	programService.mockCreate = func(program domains.Program) error {
		return nil
	}
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	program := createMockProgram("prog1", "Program1", 2, 3, "descript1", 0)
	marshal, _ := json.Marshal(&program)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateProgram_Failure(t *testing.T) {
	// no mock needed
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	program := createMockProgram("prog1", "", 2, 3, "descript1", 0) // Empty Name!
	marshal, _ := json.Marshal(&program)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateProgram_Success(t *testing.T) {
	programService.mockUpdate = func(programId string, program domains.Program) error {
		return nil // Successful update
	}
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	program := createMockProgram("prog2", "Program2", 2, 3, "descript2", 0)
	body := createBodyFromProgram(program)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/program/prog1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateProgram_Invalid(t *testing.T) {
	// no mock needed
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	program := createMockProgram("prog2", "", 2, 3, "descript2", 0) // Empty Name!
	body := createBodyFromProgram(program)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/program/prog1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateProgram_Failure(t *testing.T) {
	programService.mockUpdate = func(programId string, program domains.Program) error {
		return errors.New("not found")
	}
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	program := createMockProgram("prog2", "Program2", 2, 3, "descript2", 0)
	body := createBodyFromProgram(program)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/program/prog1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Publish
//
func TestPublishPrograms_Success(t *testing.T) {
	programService.mockPublish = func(programIds []string) error {
		return nil // Successful publish
	}
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	programIds := []string{"prog1", "prog2"}
	marshal, err := json.Marshal(programIds)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/publish", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

//
// Test Delete
//
func TestDeleteProgram_Success(t *testing.T) {
	programService.mockDelete = func(programId string) error {
		return nil // Return no error, successful delete!
	}
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/programs/v1/program/some_program", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteProgram_Failure(t *testing.T) {
	programService.mockDelete = func(programId string) error {
		return errors.New("not found")
	}
	services.ProgramService = &programService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/programs/v1/program/some_program", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockProgram(programId string, name string, grade1 uint, grade2 uint, description string, featured uint) domains.Program {
	return domains.Program{
		ProgramId:   programId,
		Name:        name,
		Grade1:      grade1,
		Grade2:      grade2,
		Description: description,
		Featured:    featured,
	}
}

func createBodyFromProgram(program domains.Program) io.Reader {
	marshal, err := json.Marshal(&program)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
