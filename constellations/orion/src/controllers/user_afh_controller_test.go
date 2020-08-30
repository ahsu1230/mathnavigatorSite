package controllers_test

import (
	"bytes"
	"encoding/json"

	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

// Test Get UserAfh By UserId
func TestGetUserAfhByUserIdSuccess(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByUserId = func(userId uint) ([]domains.UserAfh, error) {
		return []domains.UserAfh{
			testUtils.CreateMockUserAfh(2, 3, 1),
			testUtils.CreateMockUserAfh(2, 4, 1),
		}, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userAfh[0].UserId)
	assert.EqualValues(t, 3, userAfh[0].AfhId)
	assert.EqualValues(t, 1, userAfh[0].AccountId)
	assert.EqualValues(t, 2, userAfh[1].UserId)
	assert.EqualValues(t, 4, userAfh[1].AfhId)
	assert.EqualValues(t, 1, userAfh[1].AccountId)
}

func TestGetUserAfhByUserIdFailure(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByUserId = func(userId uint) ([]domains.UserAfh, error) {
		return []domains.UserAfh{}, appErrors.MockDbNoRowsError()
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/3", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Get UserAfh By AfhId
func TestGetUserAfhByAfhIdSuccess(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByAfhId = func(afhId uint) ([]domains.UserAfh, error) {
		return []domains.UserAfh{
			testUtils.CreateMockUserAfh(2, 4, 1),
			testUtils.CreateMockUserAfh(3, 4, 1),
		}, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/afh/4", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userAfh[0].UserId)
	assert.EqualValues(t, 4, userAfh[0].AfhId)
	assert.EqualValues(t, 1, userAfh[0].AccountId)
	assert.EqualValues(t, 3, userAfh[1].UserId)
	assert.EqualValues(t, 4, userAfh[1].AfhId)
	assert.EqualValues(t, 1, userAfh[1].AccountId)
}

func TestGetUserAfhByAfhIdFailure(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByAfhId = func(afhId uint) ([]domains.UserAfh, error) {
		return []domains.UserAfh{}, appErrors.MockDbNoRowsError()
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/afh/5", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Get UserAfh by UserId and AfhId
func TestGetUserAfhByBothIdsSuccess(t *testing.T) {
	testUtils.UserAfhRepo.MockSelectByBothIds = func(userId, afhId uint) (domains.UserAfh, error) {
		userAfh := testUtils.CreateMockUserAfh(2, 3, 1)
		return userAfh, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/2/afh/3", nil)

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
			testUtils.CreateMockUserAfh(2, 3, 1),
			testUtils.CreateMockUserAfh(2, 4, 1),
		}, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/new", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userAfh[0].UserId)
	assert.EqualValues(t, 3, userAfh[0].AfhId)
	assert.EqualValues(t, 1, userAfh[0].AccountId)
	assert.EqualValues(t, 2, userAfh[1].UserId)
	assert.EqualValues(t, 4, userAfh[1].AfhId)
	assert.EqualValues(t, 1, userAfh[1].AccountId)
}

// Test Create
func TestCreateUserAfhSuccess(t *testing.T) {
	testUtils.UserAfhRepo.MockInsert = func(userAfh domains.UserAfh) error {
		return nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	userAfh := testUtils.CreateMockUserAfh(2, 3, 1)
	body := createBodyFromUserAfh(userAfh)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-afhs/create", body)

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
func TestUpdateUserAfhSuccess(t *testing.T) {
	testUtils.UserAfhRepo.MockUpdate = func(id uint, userAfh domains.UserAfh) error {
		return nil // Successful update
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	userAfh := testUtils.CreateMockUserAfh(2, 3, 1)
	body := createBodyFromUserAfh(userAfh)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-afhs/user-afh/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateUserAfhFailure(t *testing.T) {
	testUtils.UserAfhRepo.MockUpdate = func(id uint, userAfh domains.UserAfh) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	userAfh := testUtils.CreateMockUserAfh(2, 3, 1)
	body := createBodyFromUserAfh(userAfh)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-afhs/user-afh/2", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Delete
func TestDeleteUserAfhSuccess(t *testing.T) {
	testUtils.UserAfhRepo.MockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/user-afhs/user-afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteUserAfhFailure(t *testing.T) {
	testUtils.UserAfhRepo.MockDelete = func(id uint) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/user-afhs/user-afh/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}
