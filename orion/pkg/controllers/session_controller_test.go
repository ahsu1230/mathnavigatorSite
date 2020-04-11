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
func TestGetAllSessionsByClassId_Success(t *testing.T) {
	now := time.Now().UTC()
	sessionService.mockGetAllByClassId = func(classId string, publishedOnly bool) ([]domains.Session, error) {
		return []domains.Session{
			{
				Id:       1,
				ClassId:  "id_1",
				StartsAt: now,
				EndsAt:   now,
				Canceled: false,
				Notes:    domains.NewNullString("special lecture from guest"),
			},
			{
				Id:       2,
				ClassId:  "id_1",
				StartsAt: now,
				EndsAt:   now,
				Canceled: true,
				Notes:    domains.NewNullString("live demonstration of science experiment"),
			},
		}, nil
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/class/id_1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var sessions []domains.Session
	if err := json.Unmarshal(recorder.Body.Bytes(), &sessions); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, sessions[0].Id)
	assert.EqualValues(t, "id_1", sessions[0].ClassId)
	assert.EqualValues(t, 2, sessions[1].Id)
	assert.EqualValues(t, "id_1", sessions[1].ClassId)
	assert.EqualValues(t, 2, len(sessions))
}

//
// Test Get Session
//
func TestGetSession_Success(t *testing.T) {
	now := time.Now().UTC()
	sessionService.mockGetBySessionId = func(id uint) (domains.Session, error) {
		session := createMockSession(1, "id_1", now, now, true, "special lecture from guest")
		return session, nil
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/session/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var session domains.Session
	if err := json.Unmarshal(recorder.Body.Bytes(), &session); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, session.Id)
	assert.EqualValues(t, "id_1", session.ClassId)
}

func TestGetSession_Failure(t *testing.T) {
	sessionService.mockGetBySessionId = func(id uint) (domains.Session, error) {
		return domains.Session{}, errors.New("Not Found")
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/sessions/v1/session/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateSession_Success(t *testing.T) {
	sessionService.mockCreate = func(session domains.Session) error {
		return nil
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	session := createMockSession(1, "id_1", now, now, true, "special lecture from guest")
	marshal, _ := json.Marshal(&session)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateSession_Failure(t *testing.T) {
	// no mock needed
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	session := createMockSession(1, "id_1", now, now, true, "@@")
	marshal, _ := json.Marshal(&session)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Publish
//
func TestPublishSessions_Success(t *testing.T) {
	sessionService.mockPublish = func(ids []uint) []domains.SessionErrorBody {
		return nil // Successful update
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	ids := []uint{1, 2}
	marshal, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/publish", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

//
// Test Update
//
func TestUpdateSession_Success(t *testing.T) {
	sessionService.mockUpdate = func(id uint, session domains.Session) error {
		return nil // Successful update
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	session := createMockSession(1, "id_1", now, now, true, "special lecture from guest")
	body := createBodyFromSession(session)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/session/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateSession_Invalid(t *testing.T) {
	// no mock needed
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	session := createMockSession(1, "id_1", now, now, true, "@@")
	body := createBodyFromSession(session)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/session/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateSession_Failure(t *testing.T) {
	sessionService.mockUpdate = func(id uint, session domains.Session) error {
		return errors.New("not found")
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	session := createMockSession(1, "id_1", now, now, true, "special lecture from guest")
	body := createBodyFromSession(session)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/sessions/v1/session/2", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteSession_Success(t *testing.T) {
	sessionService.mockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/sessions/v1/session/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteSession_Failure(t *testing.T) {
	sessionService.mockDelete = func(id uint) error {
		return errors.New("not found")
	}
	services.SessionService = &sessionService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/sessions/v1/session/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockSession(id uint, classId string, startsAt time.Time, endsAt time.Time, canceled bool, notes string) domains.Session {
	return domains.Session{
		Id:       id,
		ClassId:  classId,
		StartsAt: startsAt,
		EndsAt:   endsAt,
		Canceled: canceled,
		Notes:    domains.NewNullString(notes),
	}
}

func createBodyFromSession(session domains.Session) io.Reader {
	marshal, err := json.Marshal(&session)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
