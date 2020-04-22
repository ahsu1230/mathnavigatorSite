package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get All
//
func TestGetAllLocations_Success(t *testing.T) {
	locationRepo.mockSelectAll = func(publishedOnly bool) ([]domains.Location, error) {
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
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/locations/all", nil)

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
func TestGetLocation_Success(t *testing.T) {
	locationRepo.mockSelectByLocationId = func(LocationId string) (domains.Location, error) {
		location := createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
		return location, nil
	}
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/locations/location/loc1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var location domains.Location
	if err := json.Unmarshal(recorder.Body.Bytes(), &location); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", location.LocationId)
	assert.EqualValues(t, "4040 Location Rd", location.Street)
}

func TestGetLocation_Failure(t *testing.T) {
	locationRepo.mockSelectByLocationId = func(LocationId string) (domains.Location, error) {
		return domains.Location{}, errors.New("Not Found")
	}
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/locations/location/loc2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateLocation_Success(t *testing.T) {
	locationRepo.mockInsert = func(location domains.Location) error {
		return nil
	}
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	marshal, _ := json.Marshal(&location)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateLocation_Failure(t *testing.T) {
	// no mock needed
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "Location Rd", "City", "MA", "77294", "Room 1") // Invalid street
	marshal, _ := json.Marshal(&location)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateLocation_Success(t *testing.T) {
	locationRepo.mockUpdate = func(LocationId string, location domains.Location) error {
		return nil // Successful update
	}
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body := createBodyFromLocation(location)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/location/loc1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateLocation_Invalid(t *testing.T) {
	// no mock needed
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "Location Rd", "City", "MA", "77294", "Room 1") // Invalid street
	body := createBodyFromLocation(location)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/location/loc1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateLocation_Failure(t *testing.T) {
	locationRepo.mockUpdate = func(LocationId string, location domains.Location) error {
		return errors.New("not found")
	}
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	location := createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1")
	body := createBodyFromLocation(location)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/location/loc2", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Publish
//
func TestPublishLocations_Success(t *testing.T) {
	locationRepo.mockPublish = func(LocationIds []string) error {
		return nil // Successful publish
	}
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	LocationIds := []string{"loc1", "loc2"}
	marshal, err := json.Marshal(LocationIds)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/locations/publish", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

//
// Test Delete
//
func TestDeleteLocation_Success(t *testing.T) {
	locationRepo.mockDelete = func(LocationId string) error {
		return nil // Return no error, successful delete!
	}
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/locations/location/some_location", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteLocation_Failure(t *testing.T) {
	locationRepo.mockDelete = func(LocationId string) error {
		return errors.New("not found")
	}
	repos.LocationRepo = &locationRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/locations/location/some_location", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockLocation(LocationId string, street string, city string, state string, zipcode string, room string) domains.Location {
	return domains.Location{
		LocationId: LocationId,
		Street:     street,
		City:       city,
		State:      state,
		Zipcode:    zipcode,
		Room:       domains.NewNullString(room),
	}
}

func createBodyFromLocation(location domains.Location) io.Reader {
	marshal, err := json.Marshal(&location)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
