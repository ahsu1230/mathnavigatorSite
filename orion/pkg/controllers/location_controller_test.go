package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

//
// Test Get All
//
func TestGetAllLocations_Success(t *testing.T) {
	mls.mockGetAll = func() ([]domains.Location, error) {
		return []domains.Location{
			{
				Id:      1,
				LocId:   "loc1",
				Street:  "4040 Location Rd",
				City:    "City",
				State:   "MA",
				Zipcode: "77294",
				Room:    "Room 1",
			},
			{
				Id:      2,
				LocId:   "loc2",
				Street:  "4040 Location Ave",
				City:    "Dity",
				State:   "MD",
				Zipcode: "12353",
				Room:    "Room 2",
			},
		}, nil
	}
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var locations []domains.Location
	if err := json.Unmarshal(recorder.Body.Bytes(), &locations); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", locations[0].LocId)
	assert.EqualValues(t, "4040 Location Rd", locations[0].Street)
	assert.EqualValues(t, "loc2", locations[1].LocId)
	assert.EqualValues(t, "4040 Location Ave", locations[1].Street)
	assert.EqualValues(t, 2, len(locations))
}

//
// Test Get Location
//
func TestGetLocation_Success(t *testing.T) {
	mls.mockGetByLocationId = func(locId string) (domains.Location, error) {
		location := createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
		return location, nil
	}
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/location/loc1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var location domains.Location
	if err := json.Unmarshal(recorder.Body.Bytes(), &location); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", location.LocId)
	assert.EqualValues(t, "4040 Location Rd", location.Street)
}

func TestGetLocation_Failure(t *testing.T) {
	mls.mockGetByLocationId = func(locId string) (domains.Location, error) {
		return domains.Location{}, errors.New("Not Found")
	}
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/locations/v1/location/loc2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateLocation_Success(t *testing.T) {
	mls.mockCreate = func(location domains.Location) error {
		return nil
	}
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	marshal, _ := json.Marshal(location)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateLocation_Failure(t *testing.T) {
	// no mock needed
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "Location Rd", "City", "MA", "77294", "Room 1") // Invalid street
	marshal, _ := json.Marshal(location)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateLocation_Success(t *testing.T) {
	mls.mockUpdate = func(locId string, location domains.Location) error {
		return nil // Succesful update
	}
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body := createBodyFromLocation(location)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/location/loc1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateLocation_Invalid(t *testing.T) {
	// no mock needed
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "Location Rd", "City", "MA", "77294", "Room 1") // Invalid street
	body := createBodyFromLocation(location)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/location/loc1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateLocation_Failure(t *testing.T) {
	mls.mockUpdate = func(locId string, location domains.Location) error {
		return errors.New("not found")
	}
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body := createBodyFromLocation(location)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/v1/location/loc2", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteLocation_Success(t *testing.T) {
	mls.mockDelete = func(locId string) error {
		return nil // Return no error, successful delete!
	}
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/locations/v1/location/some_location", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteLocation_Failure(t *testing.T) {
	mls.mockDelete = func(locId string) error {
		return errors.New("not found")
	}
	services.LocationService = &mls

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/locations/v1/location/some_location", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockLocation(locId string, street string, city string, state string, zipcode string, room string) domains.Location {
	return domains.Location{
		LocId:   locId,
		Street:  street,
		City:    city,
		State:   state,
		Zipcode: zipcode,
		Room:    room,
	}
}

func createBodyFromLocation(location domains.Location) io.Reader {
	marshal, err := json.Marshal(location)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
