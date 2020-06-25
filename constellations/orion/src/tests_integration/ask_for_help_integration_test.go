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
func Test_CreateAskForHelps(t *testing.T) {
	createLocations(t)

	var date1 = now.Add(time.Hour * 24 * 30)
	var date2 = now.Add(time.Hour * 24 * 31)
	var date3 = now.Add(time.Hour * 24 * 60)
	afh1 := createAFH(1, "AP Calculus Help", date1, "1:00-3:00PM", "Calculus", "wchs")
	afh2 := createAFH(2, "AP Statistics Help", date2, "2:00-4:00PM", "Statistics", "room12")
	afh3 := createAFH(3, "AP CS Help", date3, "3:00-5:00PM", "CS", "room101")
	body1 := utils.CreateJsonBody(&afh1)
	body2 := utils.CreateJsonBody(&afh2)
	body3 := utils.CreateJsonBody(&afh3)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

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
	assert.EqualValues(t, date1, askForHelps[0].Date)
	assert.EqualValues(t, "1:00-3:00PM", askForHelps[0].TimeString)
	assert.EqualValues(t, "Calculus", askForHelps[0].Subject)
	assert.EqualValues(t, "wchs", askForHelps[0].LocationId)
	assert.EqualValues(t, 2, askForHelps[1].Id)
	assert.EqualValues(t, "AP Statistics Help", askForHelps[1].Title)
	assert.EqualValues(t, date2, askForHelps[1].Date)
	assert.EqualValues(t, "2:00-4:00PM", askForHelps[1].TimeString)
	assert.EqualValues(t, "Statistics", askForHelps[1].Subject)
	assert.EqualValues(t, "room12", askForHelps[1].LocationId)
	assert.EqualValues(t, 3, askForHelps[2].Id)
	assert.EqualValues(t, "AP CS Help", askForHelps[2].Title)
	assert.EqualValues(t, date3, askForHelps[2].Date)
	assert.EqualValues(t, "3:00-5:00PM", askForHelps[2].TimeString)
	assert.EqualValues(t, "CS", askForHelps[2].Subject)
	assert.EqualValues(t, "room101", askForHelps[2].LocationId)

	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 1 Ask For Help, Update, Get By ID
func Test_UpdateAFH(t *testing.T) {
	createLocations(t)

	//Create 1 AFH
	var date1 = now.Add(time.Hour * 24 * 30)
	afh1 := createAFH(1, "AP Calculus Help", date1, "1:00-3:00PM", "Calculus", "wchs")
	body1 := utils.CreateJsonBody(&afh1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	var date2 = now.Add(time.Hour * 24 * 31)
	updatedAFH := createAFH(1, "AP Statistics Help", date2, "2:00-4:00PM", "Statistics", "room12")
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
	assert.EqualValues(t, date2, askForHelp.Date)
	assert.EqualValues(t, "2:00-4:00PM", askForHelp.TimeString)
	assert.EqualValues(t, "Statistics", askForHelp.Subject)
	assert.EqualValues(t, "room12", askForHelp.LocationId)

	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 1 AFH, Delete it, GetById()
func Test_DeleteAFH(t *testing.T) {
	createLocations(t)

	// Create
	var date1 = now.Add(time.Hour * 24 * 30)
	afh1 := createAFH(1, "AP Calculus Help", date1, "1:00-3:00PM", "Calculus", "wchs")
	body1 := utils.CreateJsonBody(&afh1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/askforhelp/afh/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/afh/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Helper methods
func createLocations(t *testing.T) {
	location1 := createLocation("wchs", "11300 Gainsborough Road", "Potomac", "MD", "20854", "Room 100")
	location2 := createLocation("room12", "123 Sesame St", "Rockville", "MD", "20814", "Room 8")
	location3 := createLocation("room101", "101 Rainbow Avenue", "Gaithersburg", "MD", "23456", "Room 101")

	body1 := utils.CreateJsonBody(&location1)
	body2 := utils.CreateJsonBody(&location2)
	body3 := utils.CreateJsonBody(&location3)

	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body3)

	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
}

func createAFH(id uint, title string, date time.Time, timeString string, subject string, locationId string) domains.AskForHelp {
	return domains.AskForHelp{
		Id:         id,
		Title:      title,
		Date:       date,
		TimeString: timeString,
		Subject:    subject,
		LocationId: locationId,
	}
}
