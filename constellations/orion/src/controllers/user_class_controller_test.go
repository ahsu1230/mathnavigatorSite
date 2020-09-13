package controllers_test

import (
	"bytes"
	"context"
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

func TestGetUsersByClassIdSuccess(t *testing.T) {
	testUtils.UserClassRepo.MockSelectByClassId = func(context.Context, string) ([]domains.UserClass, error) {
		return []domains.UserClass{
			testUtils.CreateMockUserClass(
				1,
				1,
				"abcd",
				1,
				domains.USER_CLASS_ACCEPTED,
			),
			testUtils.CreateMockUserClass(
				2,
				2,
				"abcd",
				2,
				domains.USER_CLASS_TRIAL,
			),
		}, nil
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/class/abcd", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userClass []domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, userClass[0].Id)
	assert.EqualValues(t, 1, userClass[0].UserId)
	assert.EqualValues(t, "abcd", userClass[0].ClassId)
	assert.EqualValues(t, 1, userClass[0].AccountId)
	assert.EqualValues(t, domains.USER_CLASS_ACCEPTED, userClass[0].State)

	assert.EqualValues(t, 2, userClass[1].Id)
	assert.EqualValues(t, 2, userClass[1].UserId)
	assert.EqualValues(t, "abcd", userClass[1].ClassId)
	assert.EqualValues(t, 2, userClass[1].AccountId)
	assert.EqualValues(t, domains.USER_CLASS_TRIAL, userClass[1].State)

	assert.EqualValues(t, 2, len(userClass))
}

func TestGetClassesByUserIdSuccess(t *testing.T) {
	testUtils.UserClassRepo.MockSelectByUserId = func(context.Context, uint) ([]domains.UserClass, error) {
		return []domains.UserClass{
			testUtils.CreateMockUserClass(
				1,
				1,
				"abcd",
				1,
				domains.USER_CLASS_ACCEPTED,
			),
			testUtils.CreateMockUserClass(
				2,
				1,
				"abce",
				1,
				domains.USER_CLASS_ACCEPTED,
			),
		}, nil
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userClass []domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, userClass[0].Id)
	assert.EqualValues(t, 1, userClass[0].UserId)
	assert.EqualValues(t, "abcd", userClass[0].ClassId)
	assert.EqualValues(t, 1, userClass[0].AccountId)
	assert.EqualValues(t, domains.USER_CLASS_ACCEPTED, userClass[0].State)

	assert.EqualValues(t, 2, userClass[1].Id)
	assert.EqualValues(t, 1, userClass[1].UserId)
	assert.EqualValues(t, "abce", userClass[1].ClassId)
	assert.EqualValues(t, 1, userClass[1].AccountId)
	assert.EqualValues(t, domains.USER_CLASS_ACCEPTED, userClass[1].State)

	assert.EqualValues(t, 2, len(userClass))
}

func TestGetUserClassByUserAndClassSuccess(t *testing.T) {
	testUtils.UserClassRepo.MockSelectByUserAndClass = func(context.Context, uint, string) (domains.UserClass, error) {
		userClass := testUtils.CreateMockUserClass(
			1,
			1,
			"abcd",
			1,
			domains.USER_CLASS_ACCEPTED,
		)
		return userClass, nil
	}

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/class/abcd/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userClass domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, userClass.Id)
	assert.EqualValues(t, 1, userClass.UserId)
	assert.EqualValues(t, "abcd", userClass.ClassId)
	assert.EqualValues(t, 1, userClass.AccountId)
	assert.EqualValues(t, domains.USER_CLASS_ACCEPTED, userClass.State)

}

func TestGetUsersByNew(t *testing.T) {
	testUtils.UserClassRepo.MockSelectByNew = func(context.Context) ([]domains.UserClass, error) {
		return []domains.UserClass{
			testUtils.CreateMockUserClass(
				1,
				1,
				"abcd",
				1,
				domains.USER_CLASS_ACCEPTED,
			),
			testUtils.CreateMockUserClass(
				2,
				2,
				"abcd",
				2,
				domains.USER_CLASS_TRIAL,
			),
		}, nil
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/new", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userClass []domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, userClass[0].Id)
	assert.EqualValues(t, 1, userClass[0].UserId)
	assert.EqualValues(t, "abcd", userClass[0].ClassId)
	assert.EqualValues(t, 1, userClass[0].AccountId)
	assert.EqualValues(t, domains.USER_CLASS_ACCEPTED, userClass[0].State)

	assert.EqualValues(t, 2, userClass[1].Id)
	assert.EqualValues(t, 2, userClass[1].UserId)
	assert.EqualValues(t, "abcd", userClass[1].ClassId)
	assert.EqualValues(t, 2, userClass[1].AccountId)
	assert.EqualValues(t, domains.USER_CLASS_TRIAL, userClass[1].State)

	assert.EqualValues(t, 2, len(userClass))
}

//
// Test Create
//
func TestCreateUserClassSuccess(t *testing.T) {
	testUtils.UserClassRepo.MockInsert = func(context.Context, domains.UserClass) (uint, error) {
		return 42, nil
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClass(
		1,
		1,
		"abcd",
		1,
		domains.USER_CLASS_ACCEPTED,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateUserClassFailure(t *testing.T) {
	testUtils.UserClassRepo.MockInsert = func(context.Context, domains.UserClass) (uint, error) {
		return 0, appErrors.MockMySQLDuplicateEntryError()
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClass(
		1,
		0,
		"",
		0,
		domains.USER_CLASS_PENDING,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateUserClassSuccess(t *testing.T) {
	testUtils.UserClassRepo.MockUpdate = func(context.Context, uint, domains.UserClass) error {
		return nil // Successful update
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClass(
		1,
		1,
		"abcd",
		1,
		domains.USER_CLASS_ACCEPTED,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/user-class/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateUserClassInvalid(t *testing.T) {
	testUtils.UserClassRepo.MockUpdate = func(context.Context, uint, domains.UserClass) error {
		return appErrors.MockDbNoRowsError()
	}
	// no mock needed
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClass(
		1,
		0,
		"",
		0,
		domains.USER_CLASS_PENDING,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/user-class/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

func TestUpdateUserClassFailure(t *testing.T) {
	testUtils.UserClassRepo.MockUpdate = func(context.Context, uint, domains.UserClass) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClass(
		1,
		1,
		"abcd",
		1,
		domains.USER_CLASS_ACCEPTED,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/user-class/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteUserClassSuccess(t *testing.T) {
	testUtils.UserClassRepo.MockDelete = func(context.Context, uint) error {
		return nil // Return no error, successful delete!
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/user-classes/user-class/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteUserClassFailure(t *testing.T) {
	testUtils.UserClassRepo.MockDelete = func(context.Context, uint) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/user-classes/user-class/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

func TestStateValuesSuccess(t *testing.T) {
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/states", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

//
// Helper Methods
//

func createMockUserClass(id uint, userId uint, classId string, accountId uint, state uint) domains.UserClass {
	return domains.UserClass{
		Id:        id,
		UserId:    userId,
		ClassId:   classId,
		AccountId: accountId,
		State:     state,
	}
}

func createBodyFromUserClass(userClass domains.UserClass) io.Reader {
	marshal, err := json.Marshal(&userClass)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
