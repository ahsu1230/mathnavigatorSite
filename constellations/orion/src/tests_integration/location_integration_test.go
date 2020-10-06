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

// Test: Create 3 Locations and GetAll()
func TestCreateLocations(t *testing.T) {
	utils.SendCreateLocation(t, true, "loc1", "High School C", "4040 Location Rd", "City", "MA", "77294", "Room 1", false)
	utils.SendCreateLocation(t, true, "loc2", "High School D", "4040 Location Ave", "Dity", "MD", "77294-1243", "Room 2", false)
	utils.SendCreateLocation(t, true, "zoom", "Zoom Conference", "", "", "", "", "", true)

	// Call Get All!
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/locations/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var locations []domains.Location
	if err := json.Unmarshal(recorder4.Body.Bytes(), &locations); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", locations[0].LocationId)
	assert.EqualValues(t, "High School C", locations[0].Title)
	assert.EqualValues(t, "4040 Location Rd", locations[0].Street.String)
	assert.EqualValues(t, false, locations[0].IsOnline)
	assert.EqualValues(t, "loc2", locations[1].LocationId)
	assert.EqualValues(t, "High School D", locations[1].Title)
	assert.EqualValues(t, "4040 Location Ave", locations[1].Street.String)
	assert.EqualValues(t, false, locations[1].IsOnline)
	assert.EqualValues(t, "zoom", locations[2].LocationId)
	assert.EqualValues(t, "Zoom Conference", locations[2].Title)
	assert.EqualValues(t, "", locations[2].Street.String)
	assert.EqualValues(t, true, locations[2].IsOnline)
	assert.EqualValues(t, 3, len(locations))

	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 2 Locations with same locationId. Then GetByLocationId()
func TestUniqueLocationId(t *testing.T) {
	utils.SendCreateLocation(
		t,
		true,
		"loc1",
		"Potomac School",
		"4040 Location Rd",
		"City",
		"MA",
		"77294",
		"Room 1",
		false,
	)

	_, recorder2 := utils.SendCreateLocation(
		t,
		false,
		"loc1", // Same locationId as first one
		"Glen School",
		"89 South Glen Rd",
		"City",
		"MA",
		"77294",
		"Room 43",
		false,
	)
	assert.EqualValues(t, http.StatusBadRequest, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/locations/location/loc1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var location domains.Location
	if err := json.Unmarshal(recorder3.Body.Bytes(), &location); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", location.LocationId)
	assert.EqualValues(t, "Potomac School", location.Title)
	assert.EqualValues(t, "4040 Location Rd", location.Street.String)

	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 1 Location, Update it, GetByLocationId()
func TestUpdateLocation(t *testing.T) {
	// Create 1 Location
	utils.SendCreateLocation(
		t,
		true,
		"loc1",
		"Potomac School",
		"4040 Location Rd",
		"City",
		"MA",
		"77294",
		"Room 1",
		false,
	)

	// Update
	updatedLocation := domains.Location{
		LocationId: "loc2",
		Title:      "Potomac School",
		Street:     domains.NewNullString("4040 Location Ave"),
		City:       domains.NewNullString("Dity"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("77294-1243"),
		Room:       domains.NewNullString("Room 2"),
	}
	updatedBody := utils.CreateJsonBody(&updatedLocation)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/location/loc1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/locations/location/loc1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/locations/location/loc2", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var location domains.Location
	if err := json.Unmarshal(recorder4.Body.Bytes(), &location); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc2", location.LocationId)
	assert.EqualValues(t, "4040 Location Ave", location.Street.String)
	assert.EqualValues(t, "Dity", location.City.String)
	assert.EqualValues(t, "MD", location.State.String)
	assert.EqualValues(t, "77294-1243", location.Zipcode.String)
	assert.EqualValues(t, "Room 2", location.Room.String)

	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 1 Location, Delete it, GetByLocationId()
func TestDeleteLocation(t *testing.T) {
	// Create
	utils.SendCreateLocation(
		t,
		true,
		"loc1",
		"Potomac School",
		"4040 Location Rd",
		"City",
		"MA",
		"77294",
		"Room 1",
		false,
	)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/locations/location/loc1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/locations/location/loc1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}
