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
	"time"
)

//
// Test Get All
//
func TestGetAllAnnouncements_Success(t *testing.T) {
	now := time.Now().UTC()
	mas.mockGetAll = func() ([]domains.Announce, error) {
		return []domains.Announce{
			{
				Id:       1,
				PostedAt: now,
				Author:   "Author Name",
				Message:  "Valid Message",
			},
			{
				Id:       2,
				PostedAt: now,
				Author:   "Author Name 2",
				Message:  "Valid Message 2",
			},
		}, nil
	}
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/all", nil)

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
	mas.mockGetByAnnounceId = func(id uint) (domains.Announce, error) {
		announce := createMockAnnounce(1, now, "Author Name", "Valid Message")
		return announce, nil
	}
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/announce/1", nil)

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
	mas.mockGetByAnnounceId = func(id uint) (domains.Announce, error) {
		return domains.Announce{}, errors.New("Not Found")
	}
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/announce/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateAnnounce_Success(t *testing.T) {
	mas.mockCreate = func(announce domains.Announce) error {
		return nil
	}
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "Author Name", "Valid Message")
	marshal, _ := json.Marshal(announce)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAnnounce_Failure(t *testing.T) {
	// no mock needed
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "", "Valid Message")
	marshal, _ := json.Marshal(announce)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateAnnounce_Success(t *testing.T) {
	mas.mockUpdate = func(id uint, announce domains.Announce) error {
		return nil // Successful update
	}
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "Author Name", "Valid Message")
	body := createBodyFromAnnounce(announce)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/announce/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAnnounce_Invalid(t *testing.T) {
	// no mock needed
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "", "Valid Message")
	body := createBodyFromAnnounce(announce)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/announce/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAnnounce_Failure(t *testing.T) {
	mas.mockUpdate = func(id uint, announce domains.Announce) error {
		return errors.New("not found")
	}
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	announce := createMockAnnounce(1, now, "Author Name", "Valid Message")
	body := createBodyFromAnnounce(announce)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/announce/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteAnnounce_Success(t *testing.T) {
	mas.mockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/announcements/v1/announce/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteAnnounce_Failure(t *testing.T) {
	mas.mockDelete = func(id uint) error {
		return errors.New("not found")
	}
	services.AnnounceService = &mas

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/announcements/v1/announce/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockAnnounce(id uint, postedAt time.Time, author string, message string) domains.Announce {
	return domains.Announce{
		Id:       id,
		PostedAt: postedAt,
		Author:   author,
		Message:  message,
	}
}

func createBodyFromAnnounce(announce domains.Announce) io.Reader {
	marshal, err := json.Marshal(announce)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
