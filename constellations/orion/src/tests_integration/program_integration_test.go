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

// Test: Create 3 Programs and GetAll()
func TestCreatePrograms(t *testing.T) {
	program1 := createProgram("prog1", "Program1", 2, 3, "descript1", 0)
	program2 := createProgram("prog2", "Program2", 2, 3, "descript2", 1)
	program3 := createProgram("prog3", "Program3", 2, 3, "descript3", 0)
	body1 := utils.CreateJsonBody(&program1)
	body2 := utils.CreateJsonBody(&program2)
	body3 := utils.CreateJsonBody(&program3)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/all", nil)

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

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 2 Programs with same programId. Then GetByProgramId()
func TestUniqueProgramId(t *testing.T) {
	program1 := createProgram("prog1", "Program1", 2, 3, "descript1", 0)
	program2 := createProgram("prog1", "Program2", 2, 3, "descript2", 1) // Same programId
	body1 := utils.CreateJsonBody(&program1)
	body2 := utils.CreateJsonBody(&program2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusBadRequest, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/program/prog1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var program domains.Program
	if err := json.Unmarshal(recorder3.Body.Bytes(), &program); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog1", program.ProgramId)
	assert.EqualValues(t, "Program1", program.Name)

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 1 Program, Update it, GetByProgramId()
func TestUpdateProgram(t *testing.T) {
	// Create 1 Program
	program1 := createProgram("prog1", "Program1", 2, 3, "descript1", 0)
	body1 := utils.CreateJsonBody(&program1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedProgram := createProgram("prog2", "Program2a", 2, 3, "Description123", 1)
	updatedBody := utils.CreateJsonBody(&updatedProgram)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/program/prog1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/program/prog1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/program/prog2", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var program domains.Program
	if err := json.Unmarshal(recorder4.Body.Bytes(), &program); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog2", program.ProgramId)
	assert.EqualValues(t, "Program2a", program.Name)

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 1 Program, Delete it, GetByProgramId()
func TestDeleteProgram(t *testing.T) {
	// Create
	program1 := createProgram("prog1", "Program1", 2, 3, "descript1", 0)
	body1 := utils.CreateJsonBody(&program1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/programs/program/prog1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/program/prog1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Helper methods
func createProgram(programId string, name string, grade1 uint, grade2 uint, description string, featured uint) domains.Program {
	return domains.Program{
		ProgramId:   programId,
		Name:        name,
		Grade1:      grade1,
		Grade2:      grade2,
		Description: description,
		Featured:    featured,
	}
}
