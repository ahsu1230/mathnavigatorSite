package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get All
//
func TestGetAllSessionsByClassId_Success(t *testing.T) {
	now := time.Now().UTC()
	testUtils.SessionRepo.MockSelectAllByClassId = func(classId string, publishedOnly bool) ([]domains.Session, error) {
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
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/sessions/class/id_1", nil)

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
	testUtils.SessionRepo.MockSelectBySessionId = func(id uint) (domains.Session, error) {
		session := createMockSession(1, "id_1", now, now, true, "special lecture from guest")
		return session, nil
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/sessions/session/1", nil)

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
	testUtils.SessionRepo.MockSelectBySessionId = func(id uint) (domains.Session, error) {
		return domains.Session{}, errors.New("Not Found")
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/sessions/session/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateSessions_Success(t *testing.T) {
	testUtils.SessionRepo.MockInsert = func(session []domains.Session) error {
		return nil
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	sessions := []domains.Session{createMockSession(1, "id_1", now, now, true, "special lecture from guest")}
	marshal, _ := json.Marshal(&sessions)
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/sessions/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateSessions_Failure(t *testing.T) {
	testUtils.SessionRepo.MockInsert = func(session []domains.Session) error {
		return errors.New("invalid notes")
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	sessions := []domains.Session{createMockSession(1, "id_1", now, now, true, "@@")}
	marshal, _ := json.Marshal(&sessions)
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/sessions/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateSession_Success(t *testing.T) {
	testUtils.SessionRepo.MockUpdate = func(id uint, session domains.Session) error {
		return nil // Successful update
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	session := createMockSession(1, "id_1", now, now, true, "special lecture from guest")
	body := createBodyFromSession(session)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/sessions/session/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateSession_Invalid(t *testing.T) {
	// no mock needed
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	session := createMockSession(1, "id_1", now, now, true, "@@")
	body := createBodyFromSession(session)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/sessions/session/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateSession_Failure(t *testing.T) {
	testUtils.SessionRepo.MockUpdate = func(id uint, session domains.Session) error {
		return errors.New("not found")
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	now := time.Now().UTC()
	session := createMockSession(1, "id_1", now, now, true, "special lecture from guest")
	body := createBodyFromSession(session)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/sessions/session/2", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Publish
//
func TestPublishSessions_Success(t *testing.T) {
	testUtils.SessionRepo.MockPublish = func(ids []uint) error {
		return nil // Successful publish
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	ids := []uint{1, 2}
	marshal, err := json.Marshal(ids)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/sessions/publish", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

//
// Test Delete
//
func TestDeleteSessions_Success(t *testing.T) {
	testUtils.SessionRepo.MockDelete = func(ids []uint) error {
		return nil // Return no error, successful delete!
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	ids := []uint{1, 2, 3}
	marshal, err := json.Marshal(ids)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer(marshal)
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/sessions/delete", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteSessions_Failure(t *testing.T) {
	testUtils.SessionRepo.MockDelete = func(id []uint) error {
		return errors.New("not found")
	}
	repos.SessionRepo = &testUtils.SessionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/sessions/delete", nil)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
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
