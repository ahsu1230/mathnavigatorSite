package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// Test: Create 3 Sessions, 2 With Same Class Id, and GetAllByClassId()
func Test_CreateSessions(t *testing.T) {
	resetTable(t, domains.TABLE_SESSIONS)

	now := time.Now().UTC()
	class1 := createClass("fast_track", "sem_1", "class_A", "loc_1", "5:00, 7:00", now, now)
	class2 := createClass("slow_track", "sem_2", "class_B", "loc_1", "3:00, 7:00", now, now)
	session1 := createSession("class_A", now, now, false, "special lecture from guest")
	session2 := createSession("class_A", now, now, true, "May 5th regular meeting")
	session3 := createSession("class_B", now, now, false, "May 5th regular meeting")
	body1 := createJsonBody(class1)
	body2 := createJsonBody(class2)
	body3 := createJsonBody(session1)
	body4 := createJsonBody(session2)
	body5 := createJsonBody(session3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/classes/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body3)
	recorder4 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body4)
	recorder5 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body5)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Call Get All!
	recorder6 := sendHttpRequest(t, http.MethodGet, "/api/classes/v1/class/class_A/sessions", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder6.Code)
	var sessions []domains.Session
	if err := json.Unmarshal(recorder6.Body.Bytes(), &sessions); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, sessions[0].Id)
	assert.EqualValues(t, "class_A", sessions[0].ClassId)
	assert.EqualValues(t, 2, sessions[1].Id)
	assert.EqualValues(t, "class_A", sessions[1].ClassId)
	assert.EqualValues(t, 2, len(sessions))
}

// Test: Create 1 Session, Update it, GetBySessionId()
func Test_UpdateSession(t *testing.T) {
	resetTable(t, domains.TABLE_SESSIONS)

	// Create 1 Session
	now := time.Now().UTC()
	session1 := createSession("class_A", now, now, false, "special lecture from guest")
	body1 := createJsonBody(session1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedSession := createSession("class_B", now, now, true, "May 5th regular meeting")
	updatedBody := createJsonBody(updatedSession)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/session/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/session/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/session/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Validate results
	var session domains.Session
	if err := json.Unmarshal(recorder4.Body.Bytes(), &session); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, session.Id)
	assert.EqualValues(t, "class_B", session.ClassId)
}

// Test: Create 1 Session, Delete it, GetBySessionId()
func Test_DeleteSession(t *testing.T) {
	resetTable(t, domains.TABLE_SESSIONS)

	// Create
	now := time.Now().UTC()
	session1 := createSession("class_A", now, now, false, "special lecture from guest")
	body1 := createJsonBody(session1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/sessions/v1/session/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/session/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
}

// Helper methods
func createSession(classId string, startsAt time.Time, endsAt time.Time, canceled bool, notes string) domains.Session {
	return domains.Session{
		ClassId:  classId,
		StartsAt: startsAt,
		EndsAt:   endsAt,
		Canceled: canceled,
		Notes:    notes,
	}
}
