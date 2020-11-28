package tests_integration

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Ask For Helps and GetAll()
func TestCreateAskForHelps(t *testing.T) {
	createLocations(t)

	start1 := time.Now().UTC()
	end1 := start1.Add(time.Hour * 1)
	start2 := start1.Add(time.Hour * 24 * 7)
	end2 := start2.Add(time.Hour * 1)
	start3 := start2.Add(time.Hour * 24 * 7)
	end3 := start3.Add(time.Hour * 1)

	utils.SendCreateAskForHelp(t, true, start1, end1, "AP Calculus Help", domains.SUBJECT_MATH, "wchs", "test note")
	utils.SendCreateAskForHelp(t, true, start2, end2, "AP Statistics Help", domains.SUBJECT_MATH, "room12", "test note 2")
	utils.SendCreateAskForHelp(t, true, start3, end3, "AP CS Help", domains.SUBJECT_PROGRAMMING, "room101", "test note 3")

	// Call Get All
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var askForHelps []domains.AskForHelp
	if err := json.Unmarshal(recorder4.Body.Bytes(), &askForHelps); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, askForHelps[0].Id)
	assert.EqualValues(t, "AP Calculus Help", askForHelps[0].Title)
	assert.EqualValues(t, domains.SUBJECT_MATH, askForHelps[0].Subject)
	assert.EqualValues(t, "wchs", askForHelps[0].LocationId)
	assert.EqualValues(t, "test note", askForHelps[0].Notes.String)
	assert.EqualValues(t, 2, askForHelps[1].Id)
	assert.EqualValues(t, "AP Statistics Help", askForHelps[1].Title)
	assert.EqualValues(t, domains.SUBJECT_MATH, askForHelps[1].Subject)
	assert.EqualValues(t, "room12", askForHelps[1].LocationId)
	assert.EqualValues(t, "test note 2", askForHelps[1].Notes.String)
	assert.EqualValues(t, 3, askForHelps[2].Id)
	assert.EqualValues(t, "AP CS Help", askForHelps[2].Title)
	assert.EqualValues(t, domains.SUBJECT_PROGRAMMING, askForHelps[2].Subject)
	assert.EqualValues(t, "room101", askForHelps[2].LocationId)
	assert.EqualValues(t, "test note 3", askForHelps[2].Notes.String)

	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 1 Ask For Help, Update, Get By ID
func TestUpdateAFH(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
	createLocations(t)

	start1 := time.Now().UTC()
	end1 := start1.Add(time.Hour * 1)

	// Create 1 AFH
	utils.SendCreateAskForHelp(t, true, start1, end1, "AP Calculus Help", domains.SUBJECT_MATH, "wchs", "test note")

	// Update
	updatedAFH := domains.AskForHelp{
		StartsAt:   start1,
		EndsAt:     end1,
		Title:      "AP Statistics Help",
		Subject:    domains.SUBJECT_MATH,
		LocationId: "room12",
		Notes:      domains.NewNullString("test note 2"),
	}
	updatedBody := utils.CreateJsonBody(&updatedAFH)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/afh/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/afh/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var askForHelp domains.AskForHelp
	if err := json.Unmarshal(recorder3.Body.Bytes(), &askForHelp); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, askForHelp.Id)
	assert.EqualValues(t, "AP Statistics Help", askForHelp.Title)
	assert.EqualValues(t, domains.SUBJECT_MATH, askForHelp.Subject)
	assert.EqualValues(t, "room12", askForHelp.LocationId)
	assert.EqualValues(t, "test note 2", askForHelp.Notes.String)

	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 1 AFH, Delete it, GetById()
func TestDeleteAFH(t *testing.T) {
	createLocations(t)

	start1 := time.Now().UTC()
	end1 := start1.Add(time.Hour * 1)

	// Create
	utils.SendCreateAskForHelp(t, true, start1, end1, "AP Calculus Help", domains.SUBJECT_MATH, "wchs", "test note")

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/askforhelp/afh/1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/afh/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 1 AFH, Archive it, GetById()
func TestArchiveAFH(t *testing.T) {
	createLocations(t)

	start1 := time.Now().UTC()
	end1 := start1.Add(time.Hour * 1)

	// Create
	utils.SendCreateAskForHelp(t, true, start1, end1, "AP Calculus Help", domains.SUBJECT_MATH, "wchs", "test note")

	// Archive
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/askforhelp/archive/1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/afh/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Helper methods
func createLocations(t *testing.T) {
	utils.SendCreateLocationWCHS(t)
	utils.SendCreateLocation(t, true, "room12", "Sesame School", "123 Sesame St", "Rockville", "MD", "20814", "Room 8", false)
	utils.SendCreateLocation(t, true, "room101", "Rainbow Academy", "101 Rainbow Avenue", "Gaithersburg", "MD", "23456", "Room 101", false)
}
