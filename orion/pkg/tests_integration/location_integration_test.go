package integration_tests

import (
	"encoding/json"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 3 Locations and GetAll()
func Test_CreateLocations(t *testing.T) {
	resetTable(t, domains.TABLE_LOCATIONS)

	location1 := createLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	location2 := createLocation("loc2", "4040 Location Ave", "Dity", "MD", "77294-1243", "Room 2")
	location3 := createLocation("loc3", "4040 Location Blvd", "Eity", "ND", "08430-0302", "Room 3")
	body1 := createJsonBody(&location1)
	body2 := createJsonBody(&location2)
	body3 := createJsonBody(&location3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var locations []domains.Location
	if err := json.Unmarshal(recorder4.Body.Bytes(), &locations); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", locations[0].LocId)
	assert.EqualValues(t, "4040 Location Rd", locations[0].Street)
	assert.EqualValues(t, "loc2", locations[1].LocId)
	assert.EqualValues(t, "4040 Location Ave", locations[1].Street)
	assert.EqualValues(t, "loc3", locations[2].LocId)
	assert.EqualValues(t, "4040 Location Blvd", locations[2].Street)
	assert.EqualValues(t, 3, len(locations))
}

// Test: Create 2 Locations with same locId. Then GetByLocationId()
func Test_UniqueLocationId(t *testing.T) {
	resetTable(t, domains.TABLE_LOCATIONS)

	location1 := createLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	location2 := createLocation("loc1", "89 South Glen Rd", "City", "MD", "77294", "Room 43") // Same locId
	body1 := createJsonBody(&location1)
	body2 := createJsonBody(&location2)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusInternalServerError, recorder2.Code)
	errBody := recorder2.Body.String()
	assert.Contains(t, errBody, "Duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))

	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/location/loc1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var location domains.Location
	if err := json.Unmarshal(recorder3.Body.Bytes(), &location); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", location.LocId)
	assert.EqualValues(t, "4040 Location Rd", location.Street)
}

// Test: Create 1 Location, Update it, GetByLocationId()
func Test_UpdateLocation(t *testing.T) {
	resetTable(t, domains.TABLE_LOCATIONS)

	// Create 1 Location
	location1 := createLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body1 := createJsonBody(&location1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedLocation := createLocation("loc2", "4040 Location Ave", "Dity", "MD", "77294-1243", "Room 2")
	updatedBody := createJsonBody(&updatedLocation)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/location/loc1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/location/loc1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/location/loc2", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var location domains.Location
	if err := json.Unmarshal(recorder4.Body.Bytes(), &location); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc2", location.LocId)
	assert.EqualValues(t, "4040 Location Ave", location.Street)
}

// Test: Create 1 Location, Delete it, GetByLocationId()
func Test_DeleteLocation(t *testing.T) {
	resetTable(t, domains.TABLE_LOCATIONS)

	// Create
	location1 := createLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body1 := createJsonBody(&location1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/locations/v1/location/loc1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/location/loc1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
}

// Test: Get All Unpublished Locations, Publish a Few, then Get All Unpublished Again
func Test_PublishLocations(t *testing.T) {
	resetTable(t, domains.TABLE_LOCATIONS)

	// Create
	location1 := createLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	location2 := createLocation("loc2", "4040 Location Ave", "Dity", "MD", "77294-1243", "Room 2")
	location3 := createLocation("loc3", "4040 Location Blvd", "Eity", "ND", "08430-0302", "Room 3")
	body1 := createJsonBody(&location1)
	body2 := createJsonBody(&location2)
	body3 := createJsonBody(&location3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Get All Unpublished
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var locIds []string
	if err := json.Unmarshal(recorder4.Body.Bytes(), &locIds); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", locIds[0])
	assert.EqualValues(t, "loc2", locIds[1])
	assert.EqualValues(t, "loc3", locIds[2])
	assert.EqualValues(t, 3, len(locIds))

	// Publish loc1 and loc3
	publishIds := []string{"loc1", "loc3"}
	publishBody := createJsonBody(publishIds)
	recorder5 := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/publish", publishBody)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Get All Unpublished Again
	recorder6 := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder6.Code)
	if err := json.Unmarshal(recorder6.Body.Bytes(), &locIds); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc2", locIds[0])
	assert.EqualValues(t, 1, len(locIds))
}

// Helper methods
func createLocation(locId string, street string, city string, state string, zipcode string, room string) domains.Location {
	return domains.Location{
		LocId:   locId,
		Street:  street,
		City:    city,
		State:   state,
		Zipcode: zipcode,
		Room:    domains.NewNullString(room),
	}
}
