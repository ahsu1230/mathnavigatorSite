package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

var date1 = now.Add(time.Hour * 24 * 30)
var date2 = now.Add(time.Hour * 24 * 31)

// Test Get All
func TestGetAllAFH_Success(t *testing.T) {
	testUtils.AskForHelpRepo.MockSelectAll = func() ([]domains.AskForHelp, error) {
		return []domains.AskForHelp{
			testUtils.CreateMockAFH(
				1,
				"AP Calculus Help",
				date1,
				"2:00-4:00PM",
				"AP Calculus",
				"wchs",
				"test note",
			),
			testUtils.CreateMockAFH(
				2,
				"AP Statistics Help",
				date2,
				"3:00-5:00PM",
				"AP Statistics",
				"room12",
				"",
			),
		}, nil
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var askForHelps []domains.AskForHelp
	if err := json.Unmarshal(recorder.Body.Bytes(), &askForHelps); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, askForHelps[0].Id)
	assert.EqualValues(t, "AP Calculus Help", askForHelps[0].Title)
	assert.EqualValues(t, date1, askForHelps[0].Date)
	assert.EqualValues(t, "2:00-4:00PM", askForHelps[0].TimeString)
	assert.EqualValues(t, "AP Calculus", askForHelps[0].Subject)
	assert.EqualValues(t, "wchs", askForHelps[0].LocationId)
	assert.EqualValues(t, "test note", askForHelps[0].Notes)
	assert.EqualValues(t, 2, askForHelps[1].Id)
	assert.EqualValues(t, "AP Statistics Help", askForHelps[1].Title)
	assert.EqualValues(t, date2, askForHelps[1].Date)
	assert.EqualValues(t, "3:00-5:00PM", askForHelps[1].TimeString)
	assert.EqualValues(t, "AP Statistics", askForHelps[1].Subject)
	assert.EqualValues(t, "room12", askForHelps[1].LocationId)
	assert.EqualValues(t, "", askForHelps[1].Notes)
	assert.EqualValues(t, 2, len(askForHelps))
}

// Test Get Ask For Help
func TestGetAFH_Success(t *testing.T) {
	testUtils.AskForHelpRepo.MockSelectById = func(id uint) (domains.AskForHelp, error) {
		askForHelp := testUtils.CreateMockAFH(
			1,
			"AP Calculus Help",
			date1,
			"2:00-4:00PM",
			"AP Calculus",
			"wchs",
			"test note")
		return askForHelp, nil
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var askForHelp domains.AskForHelp
	if err := json.Unmarshal(recorder.Body.Bytes(), &askForHelp); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, askForHelp.Id)
	assert.EqualValues(t, "AP Calculus Help", askForHelp.Title)
	assert.EqualValues(t, date1, askForHelp.Date)
	assert.EqualValues(t, "2:00-4:00PM", askForHelp.TimeString)
	assert.EqualValues(t, "AP Calculus", askForHelp.Subject)
	assert.EqualValues(t, "wchs", askForHelp.LocationId)
	assert.EqualValues(t, "test note", askForHelp.Notes)
}

func TestGetAFH_Failure(t *testing.T) {
	testUtils.AskForHelpRepo.MockSelectById = func(id uint) (domains.AskForHelp, error) {
		return domains.AskForHelp{}, errors.New("not found")
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Create
func TestCreateAFH_Success(t *testing.T) {
	testUtils.AskForHelpRepo.MockInsert = func(askForHelp domains.AskForHelp) error {
		return nil
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	askForHelp := testUtils.CreateMockAFH(
		1,
		"AP Calculus Help",
		date1,
		"2:00-4:00PM",
		"AP Calculus",
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAFH_Failure(t *testing.T) {
	// no mock needed
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	askForHelp := testUtils.CreateMockAFH(
		1,
		"",
		date1,
		"2:00-4:00PM",
		"AP Calculus",
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

// Test Update
func TestUpdateAFH_Success(t *testing.T) {
	testUtils.AskForHelpRepo.MockUpdate = func(id uint, askForHelp domains.AskForHelp) error {
		return nil // Successful update
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	askForHelp := testUtils.CreateMockAFH(
		1,
		"AP Calculus Help",
		date1,
		"2:00-4:00PM",
		"AP Calculus",
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/afh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAFH_Invalid(t *testing.T) {
	// no mock needed
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	askForHelp := testUtils.CreateMockAFH(
		1,
		"AP Calculus Help",
		date1,
		"2:00-4:00PM",
		"",
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/afh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAFH_Failure(t *testing.T) {
	testUtils.AskForHelpRepo.MockUpdate = func(id uint, askForHelp domains.AskForHelp) error {
		return errors.New("not found")
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	askForHelp := testUtils.CreateMockAFH(
		1,
		"AP Calculus Help",
		date1,
		"2:00-4:00PM",
		"AP Calculus",
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/afh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

// Test Delete
func TestDeleteAFH_Success(t *testing.T) {
	testUtils.AskForHelpRepo.MockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/askforhelp/afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteAFH_Failure(t *testing.T) {
	testUtils.AskForHelpRepo.MockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/askforhelp/afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

// Helper Methods
func createBodyFromAFH(askForHelp domains.AskForHelp) io.Reader {
	marshal, err := json.Marshal(&askForHelp)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
