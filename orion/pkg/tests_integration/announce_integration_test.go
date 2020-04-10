package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// Test: Create 3 Announcements and GetAll()
func Test_CreateAnnouncements(t *testing.T) {
	resetAllTables(t)

	early := time.Unix(0, 0)
	mid := time.Unix(55, 123)
	now := time.Now().UTC()
	announce1 := createAnnouncement(early, "Author 1", "Message 1")
	announce2 := createAnnouncement(mid, "Author 2", "Message 2")
	announce3 := createAnnouncement(now, "Author 3", "Message 3")
	body1 := createJsonBody(announce1)
	body2 := createJsonBody(announce2)
	body3 := createJsonBody(announce3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var announces []domains.Announce
	if err := json.Unmarshal(recorder4.Body.Bytes(), &announces); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 3, announces[0].Id)
	assert.EqualValues(t, "Author 3", announces[0].Author)
	assert.EqualValues(t, "Message 3", announces[0].Message)
	assert.EqualValues(t, 2, announces[1].Id)
	assert.EqualValues(t, "Author 2", announces[1].Author)
	assert.EqualValues(t, "Message 2", announces[1].Message)
	assert.EqualValues(t, 3, len(announces))
	assert.EqualValues(t, 1, announces[2].Id)
	assert.EqualValues(t, "Author 1", announces[2].Author)
	assert.EqualValues(t, "Message 1", announces[2].Message)
}

// Test: Create 1 Announcement, Update it, GetByAnnounceId()
func Test_UpdateAnnouncement(t *testing.T) {
	resetAllTables(t)

	// Create 1 Announcement
	now := time.Now().UTC()
	announce1 := createAnnouncement(now, "Author 1", "Message 1")
	body1 := createJsonBody(announce1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedAnnounce := createAnnouncement(now, "Author 2", "Message 2")
	updatedBody := createJsonBody(updatedAnnounce)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/announcement/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/announcement/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var announce domains.Announce
	if err := json.Unmarshal(recorder3.Body.Bytes(), &announce); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, announce.Id)
	assert.EqualValues(t, "Author 2", announce.Author)
	assert.EqualValues(t, "Message 2", announce.Message)
}

// Test: Create 1 Announcement, Delete it, GetByAnnounceId()
func Test_DeleteAnnouncement(t *testing.T) {
	resetAllTables(t)

	// Create
	now := time.Now().UTC()
	announce1 := createAnnouncement(now, "Author 1", "Message 1")
	body1 := createJsonBody(announce1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/announcements/v1/announcement/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/announcement/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
}

// Helper methods
func createAnnouncement(postedAt time.Time, author string, message string) domains.Announce {
	return domains.Announce{
		PostedAt: postedAt,
		Author:   author,
		Message:  message,
	}
}
