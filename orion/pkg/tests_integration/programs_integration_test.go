package integration_tests

import (
	"encoding/json"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 3 Programs and GetAll()
func Test_CreatePrograms(t *testing.T) {
	refreshTable(t, domains.TABLE_PROGRAMS)

	program1 := createProgram("prog1", "Program1", 2, 3, "descript1")
	program2 := createProgram("prog2", "Program2", 2, 3, "descript2")
	program3 := createProgram("prog3", "Program3", 2, 3, "descript3")
	body1 := createJsonBody(program1)
	body2 := createJsonBody(program2)
	body3 := createJsonBody(program3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var programs []domains.Program
	if err := json.Unmarshal(recorder4.Body.Bytes(), &programs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "Program1", programs[0].Name)
	assert.EqualValues(t, "prog1", programs[0].ProgramId)
	assert.EqualValues(t, "Program2", programs[1].Name)
	assert.EqualValues(t, "prog2", programs[1].ProgramId)
	assert.EqualValues(t, "Program3", programs[2].Name)
	assert.EqualValues(t, "prog3", programs[2].ProgramId)
	assert.EqualValues(t, 3, len(programs))
}

// Test: Create 2 Programs with same programId. Then GetByProgramId()
func Test_UniqueProgramId(t *testing.T) {
	refreshTable(t, domains.TABLE_PROGRAMS)

	program1 := createProgram("prog1", "Program1", 2, 3, "descript1")
	program2 := createProgram("prog1", "Program2", 2, 3, "descript2") // Same programId
	body1 := createJsonBody(program1)
	body2 := createJsonBody(program2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusInternalServerError, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "Duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/program/prog1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var program domains.Program
	if err := json.Unmarshal(recorder3.Body.Bytes(), &program); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog1", program.ProgramId)
	assert.EqualValues(t, "Program1", program.Name)
}

// Test: Create 1 Program, Update it, GetByProgramId()
func Test_UpdateProgram(t *testing.T) {
	refreshTable(t, domains.TABLE_PROGRAMS)

	// Create 1 Program
	program1 := createProgram("prog1", "Program1", 2, 3, "descript1")
	body1 := createJsonBody(program1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedProgram := createProgram("prog2", "Program2a", 2, 3, "Description123")
	updatedBody := createJsonBody(updatedProgram)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/program/prog1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/program/prog1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/program/prog2", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var program domains.Program
	if err := json.Unmarshal(recorder4.Body.Bytes(), &program); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog2", program.ProgramId)
	assert.EqualValues(t, "Program2a", program.Name)
}

// Test: Create 1 Program, Delete it, GetByProgramId()
func Test_DeleteProgram(t *testing.T) {
	refreshTable(t, domains.TABLE_PROGRAMS)

	// Create
	program1 := createProgram("prog1", "Program1", 2, 3, "descript1")
	body1 := createJsonBody(program1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/programs/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/programs/v1/program/prog1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/programs/v1/program/prog1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
}

// Helper methods
func createProgram(programId string, name string, grade1 uint, grade2 uint, description string) domains.Program {
	return domains.Program{
		ProgramId:   programId,
		Name:        name,
		Grade1:      grade1,
		Grade2:      grade2,
		Description: description,
	}
}
