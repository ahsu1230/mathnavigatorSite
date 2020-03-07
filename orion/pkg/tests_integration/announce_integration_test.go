package integration_tests

import (
	//"encoding/json"
	//"fmt"
	"net/http"
	"testing"
	"time"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Announcements and GetAll()
func Test_CreateAnnouncements(t *testing.T) {
	refreshTable(t, domains.TABLE_ANNOUNCEMENTS)

	now := time.Now().UTC()
	announce1 := createAnnouncement(1, now, "Author 1", "Message 1")
	announce2 := createAnnouncement(2, now, "Author 2", "Message 2")
	announce3 := createAnnouncement(3, now, "Author 3", "Message 3")
	body1 := createJsonBody(announce1)
	body2 := createJsonBody(announce2)
	body3 := createJsonBody(announce3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body3)
	
	// Dies here, the assert returns false (http.StatusInternalServerError instead of http.StatusOK)
	// "runtime error: invalid memory address or nil pointer dereference"
	// I think something's up with announce repo, but if that was the case it shouldn't pass unit tests?
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

//	// Call Get All!
//	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/all", nil)
//
//	// Validate results
//	assert.EqualValues(t, http.StatusOK, recorder4.Code)
//	var announces []domains.Announce
//	if err := json.Unmarshal(recorder4.Body.Bytes(), &announces); err != nil {
//		t.Errorf("unexpected error: %v\n", err)
//	}
//	assert.EqualValues(t, 1, announces[0].Id)
//	assert.EqualValues(t, 2, announces[1].Id)
//	assert.EqualValues(t, 3, announces[2].Id)
//	assert.EqualValues(t, 3, len(announces))
}

//// Test: Create 2 Announcements with same id. Then GetByAnnounceId()
//func Test_UniqueAnnounceId(t *testing.T) {
//	refreshTable(t, domains.TABLE_ANNOUNCEMENTS)
//
//	now := time.Now().UTC()
//	announce1 := createAnnouncement(1, now, "Author 1", "Message 1")
//	announce2 := createAnnouncement(1, now, "Author 2", "Message 2") // Same announceId
//	body1 := createJsonBody(announce1)
//	body2 := createJsonBody(announce2)
//	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body1)
//	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body2)
//	assert.EqualValues(t, http.StatusOK, recorder1.Code)
//	assert.EqualValues(t, http.StatusInternalServerError, recorder2.Code)
//	errBody := recorder2.Body.String()
//	assert.Contains(t, errBody, "Duplicate entry", fmt.Sprintf("Expected error does not match. Got: %s", errBody))
//
//	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/announce/1", nil)
//	assert.EqualValues(t, http.StatusOK, recorder3.Code)
//
//	// Validate results
//	var announce domains.Announce
//	if err := json.Unmarshal(recorder3.Body.Bytes(), &announce); err != nil {
//		t.Errorf("unexpected error: %v\n", err)
//	}
//	assert.EqualValues(t, 1, announce.Id)
//	assert.EqualValues(t, "Author 2", announce.Author)
//}
//
//// Test: Create 1 Announcement, Update it, GetByAnnounceId()
//func Test_UpdateAnnouncement(t *testing.T) {
//	refreshTable(t, domains.TABLE_ANNOUNCEMENTS)
//
//	// Create 1 Announcement
//	now := time.Now().UTC()
//	announce1 := createAnnouncement(1, now, "Author 1", "Message 1")
//	body1 := createJsonBody(announce1)
//	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body1)
//	assert.EqualValues(t, http.StatusOK, recorder1.Code)
//
//	// Update
//	updatedAnnounce := createAnnouncement(1, now, "Author 2", "Message 2")
//	updatedBody := createJsonBody(updatedAnnounce)
//	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/announce/1", updatedBody)
//	assert.EqualValues(t, http.StatusOK, recorder2.Code)
//
//	// Get
//	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/announce/1", nil)
//	assert.EqualValues(t, http.StatusOK, recorder3.Code)
//
//	// Validate results
//	var announce domains.Announce
//	if err := json.Unmarshal(recorder3.Body.Bytes(), &announce); err != nil {
//		t.Errorf("unexpected error: %v\n", err)
//	}
//	assert.EqualValues(t, 1, announce.Id)
//	assert.EqualValues(t, "Author 2", announce.Author)
//}
//
//// Test: Create 1 Announcement, Delete it, GetByAnnounceId()
//func Test_DeleteAnnouncement(t *testing.T) {
//	refreshTable(t, domains.TABLE_ANNOUNCEMENTS)
//
//	// Create
//	now := time.Now().UTC()
//	announce1 := createAnnouncement(1, now, "Author 1", "Message 1")
//	body1 := createJsonBody(announce1)
//	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/announcements/v1/create", body1)
//	assert.EqualValues(t, http.StatusOK, recorder1.Code)
//
//	// Delete
//	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/announcements/v1/announce/1", nil)
//	assert.EqualValues(t, http.StatusOK, recorder2.Code)
//
//	// Get
//	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/announcements/v1/announce/1", nil)
//	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
//}

// Helper methods
func createAnnouncement(id uint, postedAt time.Time, author string, message string) domains.Announce {
	return domains.Announce{
		Id:			id,
		PostedAt:	postedAt,
		Author:		author,
		Message:	message,
	}
}
