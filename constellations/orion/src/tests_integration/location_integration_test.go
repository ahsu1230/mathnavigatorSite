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
	location1 := createLocation("loc1", "High School C", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	location2 := createLocation("loc2", "High School D", "4040 Location Ave", "Dity", "MD", "77294-1243", "Room 2")
	location3 := createLocation("loc3", "High School E", "4040 Location Blvd", "Eity", "ND", "08430-0302", "Room 3")
	body1 := utils.CreateJsonBody(&location1)
	body2 := utils.CreateJsonBody(&location2)
	body3 := utils.CreateJsonBody(&location3)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

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
	assert.EqualValues(t, "loc2", locations[1].LocationId)
	assert.EqualValues(t, "High School D", locations[1].Title)
	assert.EqualValues(t, "4040 Location Ave", locations[1].Street.String)
	assert.EqualValues(t, "loc3", locations[2].LocationId)
	assert.EqualValues(t, "High School E", locations[2].Title)
	assert.EqualValues(t, "4040 Location Blvd", locations[2].Street.String)
	assert.EqualValues(t, 3, len(locations))

	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 2 Locations with same locationId. Then GetByLocationId()
func TestUniqueLocationId(t *testing.T) {
	location1 := createLocation("loc1", "Potomac School", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	location2 := createLocation("loc1", "Glen School", "89 South Glen Rd", "City", "MD", "77294", "Room 43") // Same locationId
	body1 := utils.CreateJsonBody(&location1)
	body2 := utils.CreateJsonBody(&location2)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
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
	location1 := createLocation("loc1", "Potomac School", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body1 := utils.CreateJsonBody(&location1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedLocation := createLocation("loc2", "Potomac School", "4040 Location Ave", "Dity", "MD", "77294-1243", "Room 2")
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

	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Test: Create 1 Location, Delete it, GetByLocationId()
func TestDeleteLocation(t *testing.T) {
	// Create
	location1 := createLocation("loc1", "Potomac School", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body1 := utils.CreateJsonBody(&location1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/locations/location/loc1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/locations/location/loc1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}

// Helper methods
func createLocation(
	locationId string,
	title string,
	street string,
	city string,
	state string,
	zipcode string,
	room string,
) domains.Location {
	return domains.Location{
		LocationId: locationId,
		Title:      title,
		Street:     domains.NewNullString(street),
		City:       domains.NewNullString(city),
		State:      domains.NewNullString(state),
		Zipcode:    domains.NewNullString(zipcode),
		Room:       domains.NewNullString(room),
	}
}
