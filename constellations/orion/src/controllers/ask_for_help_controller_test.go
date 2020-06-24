package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

// Test Get All
func TestGetAllAFH_Success(t *testing.T) {
	testUtils.AskForHelpRepo.MockSelectAll = func() ([]domains.AskForHelp, error) {
		return []domains.AskForHelp{
			testUtils.CreateMockAFH(1, "AP Calculus Help", "December 25, 2020", "2:00-4:00PM", "AP Calculus", "wchs"),
			testUtils.CreateMockAFH(2, "AP Statistics Help", "February 14, 2020", "3:00-5:00PM", "AP Statistics", "room12"),
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
	assert.EqualValues(t, 2, askForHelps[1].Id)
	assert.EqualValues(t, "February 14, 2020", askForHelps[1].Date)
	assert.EqualValues(t, 2, len(askForHelps))
}

// Test Get Ask For Help
func TestGetAFH_Success(t *testing.T) {
	testUtils.AskForHelpRepo.MockSelectById = func(id uint) (domains.AskForHelp, error) {
		askForHelp := testUtils.CreateMockAFH(1, "AP Calculus Help", "December 25, 2020", "2:00-4:00PM", "AP Calculus", "wchs")
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
	assert.EqualValues(t, "2:00-4:00PM", askForHelp.TimeString)
	assert.EqualValues(t, "wchs", askForHelp.LocationId)
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
	askForHelp := testUtils.CreateMockAFH(1, "AP Calculus Help", "December 25, 2020", "2:00-4:00PM", "AP Calculus", "wchs")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAFH_Failure(t *testing.T) {
	// no mock needed
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	askForHelp := testUtils.CreateMockAFH(1, "", "December 25, 2020", "2:00-4:00PM", "AP Calculus", "wchs")
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
	askForHelp := testUtils.CreateMockAFH(1, "AP Calculus Help", "December 25, 2020", "2:00-4:00PM", "AP Calculus", "wchs")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/afh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAFH_Invalid(t *testing.T) {
	// no mock needed
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	askForHelp := testUtils.CreateMockAFH(1, "AP Calculus Help", "December 25, 2020", "2:00-4:00PM", "", "wchs")
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
	askForHelp := testUtils.CreateMockAFH(1, "AP Calculus Help", "December 25, 2020", "2:00-4:00PM", "AP Calculus", "wchs")
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
