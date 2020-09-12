package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

// Test Get All
func TestGetAllAFHSuccess(t *testing.T) {
	start1 := testUtils.TimeNow.Add(time.Hour * 24 * 3)
	end1 := start1.Add(time.Hour * 1)
	start2 := testUtils.TimeNow.Add(time.Hour * 24 * 10)
	end2 := start2.Add(time.Hour * 1)
	testUtils.AskForHelpRepo.MockSelectAll = func(context.Context) ([]domains.AskForHelp, error) {
		return []domains.AskForHelp{
			testUtils.CreateMockAFH(
				1,
				start1,
				end1,
				"AP Calculus Help",
				domains.SUBJECT_MATH,
				"wchs",
				"test note",
			),
			testUtils.CreateMockAFH(
				2,
				start2,
				end2,
				"AP Statistics Help",
				domains.SUBJECT_MATH,
				"room12",
				"test note 2",
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
	assert.EqualValues(t, start1, askForHelps[0].StartsAt)
	assert.EqualValues(t, end1, askForHelps[0].EndsAt)
	assert.EqualValues(t, "AP Calculus Help", askForHelps[0].Title)
	assert.EqualValues(t, domains.SUBJECT_MATH, askForHelps[0].Subject)
	assert.EqualValues(t, "wchs", askForHelps[0].LocationId)
	assert.EqualValues(t, domains.NewNullString("test note"), askForHelps[0].Notes)
	assert.EqualValues(t, 2, askForHelps[1].Id)
	assert.EqualValues(t, start2, askForHelps[1].StartsAt)
	assert.EqualValues(t, end2, askForHelps[1].EndsAt)
	assert.EqualValues(t, "AP Statistics Help", askForHelps[1].Title)
	assert.EqualValues(t, domains.SUBJECT_MATH, askForHelps[1].Subject)
	assert.EqualValues(t, "room12", askForHelps[1].LocationId)
	assert.EqualValues(t, domains.NewNullString("test note 2"), askForHelps[1].Notes)
	assert.EqualValues(t, 2, len(askForHelps))
}

// Test Get Ask For Help
func TestGetAFHSuccess(t *testing.T) {
	start1 := testUtils.TimeNow.Add(time.Hour * 24 * 3)
	end1 := start1.Add(time.Hour * 1)
	testUtils.AskForHelpRepo.MockSelectById = func(context.Context, uint) (domains.AskForHelp, error) {
		askForHelp := testUtils.CreateMockAFH(
			1,
			start1,
			end1,
			"AP Calculus Help",
			domains.SUBJECT_MATH,
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
	assert.EqualValues(t, start1, askForHelp.StartsAt)
	assert.EqualValues(t, end1, askForHelp.EndsAt)
	assert.EqualValues(t, "AP Calculus Help", askForHelp.Title)
	assert.EqualValues(t, domains.SUBJECT_MATH, askForHelp.Subject)
	assert.EqualValues(t, "wchs", askForHelp.LocationId)
	assert.EqualValues(t, domains.NewNullString("test note"), askForHelp.Notes)
}

func TestGetAFHFailure(t *testing.T) {
	testUtils.AskForHelpRepo.MockSelectById = func(context.Context, uint) (domains.AskForHelp, error) {
		return domains.AskForHelp{}, appErrors.MockDbNoRowsError()
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Create
func TestCreateAFHSuccess(t *testing.T) {
	testUtils.AskForHelpRepo.MockInsert = func(context.Context, domains.AskForHelp) (uint, error) {
		return 42, nil
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	start1 := testUtils.TimeNow.Add(time.Hour * 24 * 3)
	end1 := start1.Add(time.Hour * 1)
	askForHelp := testUtils.CreateMockAFH(
		1,
		start1,
		end1,
		"AP Calculus Help",
		domains.SUBJECT_MATH,
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAFHFailure(t *testing.T) {
	// no mock needed
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	start1 := testUtils.TimeNow.Add(time.Hour * 24 * 3)
	end1 := start1.Add(time.Hour * 1)
	askForHelp := testUtils.CreateMockAFH(
		1,
		start1,
		end1,
		"", // empty title
		domains.SUBJECT_MATH,
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

// Test Update
func TestUpdateAFHSuccess(t *testing.T) {
	testUtils.AskForHelpRepo.MockUpdate = func(context.Context, uint, domains.AskForHelp) error {
		return nil // Successful update
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	start1 := testUtils.TimeNow.Add(time.Hour * 24 * 3)
	end1 := start1.Add(time.Hour * 1)
	askForHelp := testUtils.CreateMockAFH(
		1,
		start1,
		end1,
		"AP Calculus Help",
		domains.SUBJECT_MATH,
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/afh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAFHInvalid(t *testing.T) {
	// no mock needed
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	start1 := testUtils.TimeNow.Add(time.Hour * 24 * 3)
	end1 := start1.Add(time.Hour * 1)
	askForHelp := testUtils.CreateMockAFH(
		1,
		start1,
		end1,
		"AP Calculus Help",
		"", // invalid subject
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/afh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAFHFailure(t *testing.T) {
	testUtils.AskForHelpRepo.MockUpdate = func(context.Context, uint, domains.AskForHelp) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	start1 := testUtils.TimeNow.Add(time.Hour * 24 * 3)
	end1 := start1.Add(time.Hour * 1)
	askForHelp := testUtils.CreateMockAFH(
		1,
		start1,
		end1,
		"AP Calculus Help",
		domains.SUBJECT_MATH,
		"wchs",
		"test note")
	body := createBodyFromAFH(askForHelp)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/afh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Delete
func TestDeleteAFHSuccess(t *testing.T) {
	testUtils.AskForHelpRepo.MockDelete = func(context.Context, uint) error {
		return nil // Return no error, successful delete!
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/askforhelp/afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteAFHFailure(t *testing.T) {
	testUtils.AskForHelpRepo.MockDelete = func(context.Context, uint) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.AskForHelpRepo = &testUtils.AskForHelpRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/askforhelp/afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

func TestGetAllAFHSubjects(t *testing.T) {
	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/askforhelp/subjects", nil)

	//Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var afhSubjects []string
	if err := json.Unmarshal(recorder.Body.Bytes(), &afhSubjects); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "math", afhSubjects[0])
	assert.EqualValues(t, "english", afhSubjects[1])
	assert.EqualValues(t, "programming", afhSubjects[2])
}

// Helper Methods
func createBodyFromAFH(askForHelp domains.AskForHelp) io.Reader {
	marshal, err := json.Marshal(&askForHelp)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
