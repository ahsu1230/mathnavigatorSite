package tests_integration

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Announcements and GetAll()
func TestE2ECreateAnnouncements(t *testing.T) {
	now := time.Now().UTC()
	early := now.Add(time.Hour * 24 * -20)
	mid := now.Add(time.Hour * 24 * -10)

	utils.SendCreateAnnouncement(t, true, early, "Author 1", "Message 1", false)
	utils.SendCreateAnnouncement(t, true, mid, "Author 2", "Message 2", true)
	utils.SendCreateAnnouncement(t, true, now, "Author 3", "Message 3", false)

	// Call Get All!
	recorder4 := utils.SendHttpRequest(t, http.MethodGet, "/api/announcements/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var announces []domains.Announce
	if err := json.Unmarshal(recorder4.Body.Bytes(), &announces); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 3, announces[0].Id)
	assert.EqualValues(t, "Author 3", announces[0].Author)
	assert.EqualValues(t, "Message 3", announces[0].Message)
	assert.EqualValues(t, false, announces[0].OnHomePage)
	assert.EqualValues(t, 2, announces[1].Id)
	assert.EqualValues(t, "Author 2", announces[1].Author)
	assert.EqualValues(t, "Message 2", announces[1].Message)
	assert.EqualValues(t, true, announces[1].OnHomePage)
	assert.EqualValues(t, 1, announces[2].Id)
	assert.EqualValues(t, "Author 1", announces[2].Author)
	assert.EqualValues(t, "Message 1", announces[2].Message)
	assert.EqualValues(t, false, announces[2].OnHomePage)
	assert.EqualValues(t, 3, len(announces))

	utils.ResetTable(t, domains.TABLE_ANNOUNCEMENTS)
}

// Test: Create 1 Announcement, Update it, GetByAnnounceId()
func TestE2EUpdateAnnouncement(t *testing.T) {
	// Create 1 Announcement
	now := time.Now().UTC()
	utils.SendCreateAnnouncement(t, true, now, "Author 1", "Message 1", true)

	// Update
	updatedAnnounce := domains.Announce{
		PostedAt:   now,
		Author:     "Author 2",
		Message:    "Message 2",
		OnHomePage: false,
	}
	updatedBody := utils.CreateJsonBody(&updatedAnnounce)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/announcements/announcement/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/announcements/announcement/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var announce domains.Announce
	if err := json.Unmarshal(recorder3.Body.Bytes(), &announce); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, announce.Id)
	assert.EqualValues(t, "Author 2", announce.Author)
	assert.EqualValues(t, "Message 2", announce.Message)
	assert.EqualValues(t, false, announce.OnHomePage)

	utils.ResetTable(t, domains.TABLE_ANNOUNCEMENTS)
}

// Test: Create 1 Announcement, Delete it, GetByAnnounceId()
func TestE2EDeleteAnnouncement(t *testing.T) {
	// Create
	now := time.Now().UTC()
	utils.SendCreateAnnouncement(t, true, now, "Author 1", "Message 1", true)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/announcements/announcement/1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/announcements/announcement/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_ANNOUNCEMENTS)
}
