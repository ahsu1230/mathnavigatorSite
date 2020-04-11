package controllers_test

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func setupMock() {
	programService.mockGetAllUnpublished = func() ([]domains.Program, error) {
		return []domains.Program{
			createMockProgram("prog1", "Program1", 2, 3, "descript1"),
			createMockProgram("prog2", "Program2", 8, 12, "descript2"),
		}, nil
	}
	locationService.mockGetAllUnpublished = func() ([]domains.Location, error) {
		return []domains.Location{
			createMockLocation("loc1", "4040 Location Rd", "City", "MA", "77294", "Room 1"),
			createMockLocation("loc2", "4040 Sesame St", "City", "MD", "77294", "Room 2"),
		}, nil
	}
	achieveService.mockGetUnpublished = func() ([]domains.Achieve, error) {
		return []domains.Achieve{
			createMockAchievement(1, 2020, "message1"),
			createMockAchievement(2, 2021, "message2"),
		}, nil
	}
	sessionService.mockGetAllUnpublished = func() ([]domains.Session, error) {
		now := time.Now().UTC()
		return []domains.Session{
			createMockSession(1, "id_1", now, now, true, "special lecture from guest"),
			createMockSession(2, "id_2", now, now, false, "daily meeting"),
		}, nil
	}
	services.ProgramService = &programService
	services.LocationService = &locationService
	services.AchieveService = &achieveService
	services.SessionService = &sessionService
}

//
// Test Get Unpublished
//
func TestGetUnpublishedAchievements_Success(t *testing.T) {
	setupMock()

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, unpublishedDomains.Achieves[0].Id)
	assert.EqualValues(t, 2020, unpublishedDomains.Achieves[0].Year)
	assert.EqualValues(t, "message1", unpublishedDomains.Achieves[0].Message)
	assert.EqualValues(t, 2, unpublishedDomains.Achieves[1].Id)
	assert.EqualValues(t, 2021, unpublishedDomains.Achieves[1].Year)
	assert.EqualValues(t, "message2", unpublishedDomains.Achieves[1].Message)
	assert.EqualValues(t, 2, len(unpublishedDomains.Achieves))
}

func TestGetAllUnpublishedPrograms_Success(t *testing.T) {
	setupMock()

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "prog1", unpublishedDomains.Programs[0].ProgramId)
	assert.EqualValues(t, "Program1", unpublishedDomains.Programs[0].Name)
	assert.EqualValues(t, "prog2", unpublishedDomains.Programs[1].ProgramId)
	assert.EqualValues(t, "Program2", unpublishedDomains.Programs[1].Name)
	assert.EqualValues(t, 2, len(unpublishedDomains.Programs))
}

func TestGetAllUnpublishedLocations_Success(t *testing.T) {
	setupMock()

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "loc1", unpublishedDomains.Locations[0].LocId)
	assert.EqualValues(t, "4040 Location Rd", unpublishedDomains.Locations[0].Street)
	assert.EqualValues(t, "MA", unpublishedDomains.Locations[0].State)
	assert.EqualValues(t, "loc2", unpublishedDomains.Locations[1].LocId)
	assert.EqualValues(t, "4040 Sesame St", unpublishedDomains.Locations[1].Street)
	assert.EqualValues(t, "MD", unpublishedDomains.Locations[1].State)
	assert.EqualValues(t, 2, len(unpublishedDomains.Locations))
}

func TestGetAllUnpublishedSessions_Success(t *testing.T) {
	setupMock()

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/v1/unpublished", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var unpublishedDomains domains.UnpublishedDomains
	if err := json.Unmarshal(recorder.Body.Bytes(), &unpublishedDomains); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, unpublishedDomains.Sessions[0].Id)
	assert.EqualValues(t, "id_1", unpublishedDomains.Sessions[0].ClassId)
	assert.EqualValues(t, 2, unpublishedDomains.Sessions[1].Id)
	assert.EqualValues(t, "id_2", unpublishedDomains.Sessions[1].ClassId)
	assert.EqualValues(t, 2, len(unpublishedDomains.Sessions))
}
