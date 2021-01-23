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
func TestE2ECreatePrograms(t *testing.T) {
	utils.SendCreateProgram(t, true, "prog1", "Program1", 2, 3, domains.SUBJECT_MATH, "descript1", domains.FEATURED_NONE)
	utils.SendCreateProgram(t, true, "prog2", "Program2", 2, 3, domains.SUBJECT_MATH, "descript2", domains.FEATURED_POPULAR)
	utils.SendCreateProgram(t, true, "prog3", "Program3", 2, 3, domains.SUBJECT_MATH, "descript3", domains.FEATURED_NONE)

	// Call Get All!
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var programs []domains.Program
	if err := json.Unmarshal(recorder4.Body.Bytes(), &programs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "Program1", programs[0].Title)
	assert.EqualValues(t, "prog1", programs[0].ProgramId)
	assert.EqualValues(t, "Program2", programs[1].Title)
	assert.EqualValues(t, "prog2", programs[1].ProgramId)
	assert.EqualValues(t, "Program3", programs[2].Title)
	assert.EqualValues(t, "prog3", programs[2].ProgramId)
	assert.EqualValues(t, 3, len(programs))

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 2 Programs with same programId. Then GetByProgramId()
func TestE2EUniqueProgramId(t *testing.T) {
	utils.SendCreateProgram(
		t,
		true,
		"prog1",
		"Program1",
		2,
		3,
		domains.SUBJECT_MATH,
		"descript1",
		domains.FEATURED_NONE,
	)
	_, recorder2 := utils.SendCreateProgram(
		t,
		false,
		"prog1", // Same programId
		"Program2",
		2,
		3,
		domains.SUBJECT_MATH,
		"descript2",
		domains.FEATURED_POPULAR,
	)
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
	assert.EqualValues(t, "Program1", program.Title)

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 1 Program, Update it, GetByProgramId()
func TestE2EUpdateProgram(t *testing.T) {
	// Create 1 Program
	utils.SendCreateProgram(t, true, "prog1", "Program1", 2, 3, domains.SUBJECT_MATH, "descript1", domains.FEATURED_NONE)

	// Update
	updatedProgram := domains.Program{
		ProgramId:   "prog2",
		Title:       "Program2a",
		Grade1:      2,
		Grade2:      3,
		Subject:     domains.SUBJECT_MATH,
		Description: "Description123",
		Featured:    domains.FEATURED_POPULAR,
	}
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
	assert.EqualValues(t, "Program2a", program.Title)

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 1 Program, Delete it, GetByProgramId()
func TestE2EDeleteProgram(t *testing.T) {
	// Create
	utils.SendCreateProgram(t, true, "prog1", "Program1", 2, 3, domains.SUBJECT_MATH, "descript1", domains.FEATURED_NONE)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/programs/program/prog1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/program/prog1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}

// Test: Create 1 Program, Archive it, GetByProgramId()
func TestE2EArchiveProgram(t *testing.T) {
	// Create
	utils.SendCreateProgram(t, true, "prog1", "Program1", 2, 3, domains.SUBJECT_MATH, "descript1", domains.FEATURED_NONE)

	// Archive
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/programs/archive/prog1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/programs/program/prog1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_PROGRAMS)
}
