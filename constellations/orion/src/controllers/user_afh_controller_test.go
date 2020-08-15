package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

// Test Get UserAfh By UserId
func TestGetUserAfhByUserId_Success(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByUserId = func(userId uint) ([]domains.UserAfh, error) {
		return []domains.UserAfh{
			testUtils.CreateMockUserAfh(2, 3),
			testUtils.CreateMockUserAfh(2, 4),
		}, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/users/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userAfh[0].UserId)
	assert.EqualValues(t, 3, userAfh[0].AfhId)
	assert.EqualValues(t, 2, userAfh[1].UserId)
	assert.EqualValues(t, 4, userAfh[1].AfhId)
}

func TestGetUserAfhByUserId_Failure(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByUserId = func(userId uint) ([]domains.UserAfh, error) {
		return []domains.UserAfh{}, errors.New("Not Found")
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/users/3", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Get UserAfh By AfhId
func TestGetUserAfhByAfhId_Success(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByAfhId = func(afhId uint) ([]domains.UserAfh, error) {
		return []domains.UserAfh{
			testUtils.CreateMockUserAfh(2, 4),
			testUtils.CreateMockUserAfh(3, 4),
		}, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/afh/4", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userAfh[0].UserId)
	assert.EqualValues(t, 4, userAfh[0].AfhId)
	assert.EqualValues(t, 3, userAfh[1].UserId)
	assert.EqualValues(t, 4, userAfh[1].AfhId)
}

func TestGetUserAfhByAfhId_Failure(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByAfhId = func(afhId uint) ([]domains.UserAfh, error) {
		return []domains.UserAfh{}, errors.New("Not Found")
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/afh/5", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Get UserAfh by UserId and AfhId
func TestGetUserAfhByBothIds_Success(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByBothIds = func(userId, afhId uint) (domains.UserAfh, error) {
		userAfh := testUtils.CreateMockUserAfh(2, 3)
		return userAfh, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/users/2/afh/3", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
}

func TestGetUserAfhByNew(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByNew = func() ([]domains.UserAfh, error) {
		return []domains.UserAfh{
			testUtils.CreateMockUserAfh(2, 3),
			testUtils.CreateMockUserAfh(2, 4),
		}, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/new", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userAfh[0].UserId)
	assert.EqualValues(t, 3, userAfh[0].AfhId)
	assert.EqualValues(t, 2, userAfh[1].UserId)
	assert.EqualValues(t, 4, userAfh[1].AfhId)
}

// Test Create
func TestCreateUserAfh_Success(t *testing.T) {
	testUtils.UserAfhRepo.MockInsert = func(userAfh domains.UserAfh) error {
		return nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	userAfh := testUtils.CreateMockUserAfh(2, 3)
	body := createBodyFromUserAfh(userAfh)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/userafhs/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func createBodyFromUserAfh(userAfh domains.UserAfh) io.Reader {
	marshal, err := json.Marshal(&userAfh)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}

// Test Update
func TestUpdateUserAfh_Success(t *testing.T) {
	testUtils.UserAfhRepo.MockUpdate = func(id uint, userAfh domains.UserAfh) error {
		return nil // Successful update
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	userAfh := testUtils.CreateMockUserAfh(2, 3)
	body := createBodyFromUserAfh(userAfh)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/userafhs/userafh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateUserAfh_Failure(t *testing.T) {
	testUtils.UserAfhRepo.MockUpdate = func(id uint, userAfh domains.UserAfh) error {
		return errors.New("not found")
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	userAfh := testUtils.CreateMockUserAfh(2, 3)
	body := createBodyFromUserAfh(userAfh)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/userafhs/userafh/2", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

// Test Delete
func TestDeleteUserAfh_Success(t *testing.T) {
	testUtils.UserAfhRepo.MockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/userafhs/userafh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteUserAfh_Failure(t *testing.T) {
	testUtils.UserAfhRepo.MockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/userafhs/userafh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}
