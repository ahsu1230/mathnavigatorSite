package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get All
//
func TestGetAllAnnouncements_Success(t *testing.T) {
	now := time.Now().UTC()
	announceRepo.mockSelectAll = func() ([]domains.Announce, error) {
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
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/announcements/all", nil)

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
func TestGetAnnouncement_Success(t *testing.T) {
	now := time.Now().UTC()
	announceRepo.mockSelectByAnnounceId = func(id uint) (domains.Announce, error) {
		announce := createMockAnnounce(1, now, "Author Name", "Valid Message", false)
		return announce, nil
	}
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/announcements/announcement/1", nil)

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

func TestGetAnnounce_Failure(t *testing.T) {
	announceRepo.mockSelectByAnnounceId = func(id uint) (domains.Announce, error) {
		return domains.Announce{}, errors.New("not found")
	}
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/announcements/announcement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateAnnounce_Success(t *testing.T) {
	announceRepo.mockInsert = func(announce domains.Announce) error {
		return nil
	}
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "Author Name", "Valid Message", false)
	marshal, _ := json.Marshal(announce)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAnnounce_Failure(t *testing.T) {
	// no mock needed
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "", "Valid Message", false)
	marshal, _ := json.Marshal(announce)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateAnnounce_Success(t *testing.T) {
	announceRepo.mockUpdate = func(id uint, announce domains.Announce) error {
		return nil // Successful update
	}
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "Author Name", "Valid Message", false)
	body := createBodyFromAnnounce(announce)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/announcement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAnnounce_Invalid(t *testing.T) {
	// no mock needed
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "", "Valid Message", false)
	body := createBodyFromAnnounce(announce)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/announcement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAnnounce_Failure(t *testing.T) {
	announceRepo.mockUpdate = func(id uint, announce domains.Announce) error {
		return errors.New("not found")
	}
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "Author Name", "Valid Message", false)
	body := createBodyFromAnnounce(announce)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/announcement/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteAnnounce_Success(t *testing.T) {
	announceRepo.mockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/announcements/announcement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteAnnounce_Failure(t *testing.T) {
	announceRepo.mockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.AnnounceRepo = &announceRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/announcements/announcement/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockAnnounce(id uint, postedAt time.Time, author string, message string, onHomePage bool) domains.Announce {
	return domains.Announce{
		Id:         id,
		PostedAt:   postedAt,
		Author:     author,
		Message:    message,
		OnHomePage: onHomePage,
	}
}

func createBodyFromAnnounce(announce domains.Announce) io.Reader {
	marshal, err := json.Marshal(announce)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
