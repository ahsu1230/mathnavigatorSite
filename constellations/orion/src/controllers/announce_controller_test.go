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

//
// Test Get All
//
func TestGetAllAnnouncementsSuccess(t *testing.T) {
	now := time.Now().UTC()
	testUtils.AnnounceRepo.MockSelectAll = func(context.Context) ([]domains.Announce, error) {
		return []domains.Announce{
			{
				Id:         1,
				PostedAt:   now,
				Author:     "Author Name",
				Message:    "Valid Message",
				OnHomePage: false,
			},
			{
				Id:         2,
				PostedAt:   now,
				Author:     "Author Name 2",
				Message:    "Valid Message 2",
				OnHomePage: true,
			},
		}, nil
	}
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/announcements/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var announces []domains.Announce
	if err := json.Unmarshal(recorder.Body.Bytes(), &announces); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, announces[0].Id)
	assert.EqualValues(t, "Author Name", announces[0].Author)
	assert.EqualValues(t, "Valid Message", announces[0].Message)
	assert.EqualValues(t, 2, announces[1].Id)
	assert.EqualValues(t, "Author Name 2", announces[1].Author)
	assert.EqualValues(t, "Valid Message 2", announces[1].Message)
	assert.EqualValues(t, 2, len(announces))
}

//
// Test Get Announce
//
func TestGetAnnouncementSuccess(t *testing.T) {
	now := time.Now().UTC()
	testUtils.AnnounceRepo.MockSelectByAnnounceId = func(context.Context, uint) (domains.Announce, error) {
		announce := testUtils.CreateMockAnnounce(1, now, "Author Name", "Valid Message", false)
		return announce, nil
	}
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/announcements/announcement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var announce domains.Announce
	if err := json.Unmarshal(recorder.Body.Bytes(), &announce); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, announce.Id)
	assert.EqualValues(t, "Author Name", announce.Author)
	assert.EqualValues(t, "Valid Message", announce.Message)
}

func TestGetAnnounceFailure(t *testing.T) {
	testUtils.AnnounceRepo.MockSelectByAnnounceId = func(context.Context, uint) (domains.Announce, error) {
		return domains.Announce{}, appErrors.MockDbNoRowsError()
	}
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/announcements/announcement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateAnnounceSuccess(t *testing.T) {
	testUtils.AnnounceRepo.MockInsert = func(context.Context, domains.Announce) error {
		return nil
	}
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := testUtils.CreateMockAnnounce(1, now, "Author Name", "Valid Message", false)
	marshal, _ := json.Marshal(announce)
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/announcements/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAnnounceFailure(t *testing.T) {
	// no mock needed
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := testUtils.CreateMockAnnounce(1, now, "", "Valid Message", false)
	marshal, _ := json.Marshal(announce)
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/announcements/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateAnnounceSuccess(t *testing.T) {
	testUtils.AnnounceRepo.MockUpdate = func(context.Context, uint, domains.Announce) error {
		return nil // Successful update
	}
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := testUtils.CreateMockAnnounce(1, now, "Author Name", "Valid Message", false)
	body := createBodyFromAnnounce(announce)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/announcements/announcement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAnnounceInvalid(t *testing.T) {
	// no mock needed
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := testUtils.CreateMockAnnounce(1, now, "", "Valid Message", false)
	body := createBodyFromAnnounce(announce)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/announcements/announcement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAnnounceFailure(t *testing.T) {
	testUtils.AnnounceRepo.MockUpdate = func(context.Context, uint, domains.Announce) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := testUtils.CreateMockAnnounce(1, now, "Author Name", "Valid Message", false)
	body := createBodyFromAnnounce(announce)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/announcements/announcement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteAnnounceSuccess(t *testing.T) {
	testUtils.AnnounceRepo.MockDelete = func(context.Context, uint) error {
		return nil // Return no error, successful delete!
	}
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/announcements/announcement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteAnnounceFailure(t *testing.T) {
	testUtils.AnnounceRepo.MockDelete = func(context.Context, uint) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.AnnounceRepo = &testUtils.AnnounceRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/announcements/announcement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Helper Methods
//

func createBodyFromAnnounce(announce domains.Announce) io.Reader {
	marshal, err := json.Marshal(announce)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
