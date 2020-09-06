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
func TestGetAllLocationsSuccess(t *testing.T) {
	testUtils.LocationRepo.MockSelectAll = func(context.Context) ([]domains.Location, error) {
		return []domains.Location{
			{
				Id:         1,
				LocationId: "loc1",
				Street:     "4040 Location Rd",
				City:       "City",
				State:      "MA",
				Zipcode:    "77294",
				Room:       domains.NewNullString("Room 1"),
			},
			{
				Id:         2,
				LocationId: "loc2",
				Street:     "4040 Location Ave",
				City:       "Dity",
				State:      "MD",
				Zipcode:    "12353",
				Room:       domains.NewNullString("Room 2"),
			},
		}, nil
	}
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/locations/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var locations []domains.Location
	if err := json.Unmarshal(recorder.Body.Bytes(), &locations); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", locations[0].LocationId)
	assert.EqualValues(t, "4040 Location Rd", locations[0].Street)
	assert.EqualValues(t, "loc2", locations[1].LocationId)
	assert.EqualValues(t, "4040 Location Ave", locations[1].Street)
	assert.EqualValues(t, 2, len(locations))
}

//
// Test Get Location
//
func TestGetLocationSuccess(t *testing.T) {
	testUtils.LocationRepo.MockSelectByLocationId = func(context.Context, string) (domains.Location, error) {
		location := testUtils.CreateMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
		return location, nil
	}
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/locations/location/loc1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var location domains.Location
	if err := json.Unmarshal(recorder.Body.Bytes(), &location); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", location.LocationId)
	assert.EqualValues(t, "4040 Location Rd", location.Street)
}

func TestGetLocationFailure(t *testing.T) {
	testUtils.LocationRepo.MockSelectByLocationId = func(context.Context, string) (domains.Location, error) {
		return domains.Location{}, appErrors.MockDbNoRowsError()
	}
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/locations/location/loc2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateLocationSuccess(t *testing.T) {
	testUtils.LocationRepo.MockInsert = func(context.Context, domains.Location) (uint, error) {
		return 42, nil
	}
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	location := testUtils.CreateMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	marshal, _ := json.Marshal(&location)
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateLocationFailure(t *testing.T) {
	// no mock needed
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	location := testUtils.CreateMockLocation("loc1", "Location Rd", "City", "MA", "77294", "Room 1") // Invalid street
	marshal, _ := json.Marshal(&location)
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateLocationSuccess(t *testing.T) {
	testUtils.LocationRepo.MockUpdate = func(context.Context, string, domains.Location) error {
		return nil // Successful update
	}
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	location := testUtils.CreateMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body := createBodyFromLocation(location)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/locations/location/loc1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateLocationInvalid(t *testing.T) {
	// no mock needed
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	location := testUtils.CreateMockLocation("loc1", "Location Rd", "City", "MA", "77294", "Room 1") // Invalid street
	body := createBodyFromLocation(location)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/locations/location/loc1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateLocationFailure(t *testing.T) {
	testUtils.LocationRepo.MockUpdate = func(context.Context, string, domains.Location) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	location := testUtils.CreateMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body := createBodyFromLocation(location)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/locations/location/loc2", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteLocationSuccess(t *testing.T) {
	testUtils.LocationRepo.MockDelete = func(context.Context, string) error {
		return nil // Return no error, successful delete!
	}
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/locations/location/some_location", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteLocationFailure(t *testing.T) {
	testUtils.LocationRepo.MockDelete = func(context.Context, string) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.LocationRepo = &testUtils.LocationRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/locations/location/some_location", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Helper Methods
//

func createBodyFromLocation(location domains.Location) io.Reader {
	marshal, err := json.Marshal(&location)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
