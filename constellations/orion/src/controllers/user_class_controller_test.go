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

func TestGetUsersByClassId_Success(t *testing.T) {
	testUtils.UserClassRepo.MockSelectByClassId = func(classId string) ([]domains.UserClass, error) {
		return []domains.UserClass{
			testUtils.CreateMockUserClass(
				1,
				1,
				"abcd",
				1,
				1,
			),
			testUtils.CreateMockUserClass(
				2,
				2,
				"abcd",
				2,
				2,
			),
		}, nil
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userclass/class/abcd", nil)

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
	assert.EqualValues(t, 1, userClass[0].State)

	assert.EqualValues(t, 2, userClass[1].Id)
	assert.EqualValues(t, 2, userClass[1].UserId)
	assert.EqualValues(t, "abcd", userClass[1].ClassId)
	assert.EqualValues(t, 2, userClass[1].AccountId)
	assert.EqualValues(t, 2, userClass[1].State)

	assert.EqualValues(t, 2, len(userClass))
}

func TestGetClassesByUserId_Success(t *testing.T) {
	testUtils.UserClassRepo.MockSelectByUserId = func(id uint) ([]domains.UserClass, error) {
		return []domains.UserClass{
			testUtils.CreateMockUserClass(
				1,
				1,
				"abcd",
				1,
				1,
			),
			testUtils.CreateMockUserClass(
				2,
				1,
				"abce",
				1,
				1,
			),
		}, nil
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userclass/user/1", nil)

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
	assert.EqualValues(t, 1, userClass[0].State)

	assert.EqualValues(t, 2, userClass[1].Id)
	assert.EqualValues(t, 1, userClass[1].UserId)
	assert.EqualValues(t, "abce", userClass[1].ClassId)
	assert.EqualValues(t, 1, userClass[1].AccountId)
	assert.EqualValues(t, 1, userClass[1].State)

	assert.EqualValues(t, 2, len(userClass))
}

func TestGetUserClassByUserAndClass_Success(t *testing.T) {
	testUtils.UserClassRepo.MockSelectByUserAndClass = func(id uint, classId string) (domains.UserClass, error) {
		userClass := testUtils.CreateMockUserClass(
			1,
			1,
			"abcd",
			1,
			1,
		)
		return userClass, nil
	}

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/userclass/class/abcd/user/1", nil)

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
	assert.EqualValues(t, 1, userClass.State)

}

//
// Test Create
//
func TestCreateUserClass_Success(t *testing.T) {
	testUtils.UserClassRepo.MockInsert = func(userClass domains.UserClass) error {
		return nil
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClass(
		1,
		1,
		"abcd",
		1,
		1,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/userclass/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

// func TestCreateUserClass_Failure(t *testing.T) {
// 	// no mock needed
// 	repos.UserClassRepo = &testUtils.UserClassRepo

// 	// Create new HTTP request to endpoint
// 	userClass := testUtils.CreateMockUserClass(
// 		1,
// 		0,
// 		"",
// 		0,
// 		0,
// 	)
// 	body := createBodyFromUserClass(userClass)
// 	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/userclass/create", body)

// 	// Validate results
// 	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
// }

//
// Test Update
//
func TestUpdateUserClass_Success(t *testing.T) {
	testUtils.UserClassRepo.MockUpdate = func(id uint, userClass domains.UserClass) error {
		return nil // Successful update
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClass(
		1,
		1,
		"abcd",
		1,
		1,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/userclass/uc/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

// func TestUpdateUserClass_Invalid(t *testing.T) {
// 	// no mock needed
// 	repos.UserClassRepo = &testUtils.UserClassRepo

// 	// Create new HTTP request to endpoint
// 	userClass := testUtils.CreateMockUserClass(
// 		1,
// 		0,
// 		"",
// 		0,
// 		0,
// 	)
// 	body := createBodyFromUserClass(userClass)
// 	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/userclass/uc/1", body)

// 	// Validate results
// 	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
// }

func TestUpdateUserClass_Failure(t *testing.T) {
	testUtils.UserClassRepo.MockUpdate = func(id uint, userClass domains.UserClass) error {
		return errors.New("not found")
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClass(
		1,
		1,
		"abcd",
		1,
		1,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/userclass/uc/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteUserClass_Success(t *testing.T) {
	testUtils.UserClassRepo.MockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/userclass/uc/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteUserClass_Failure(t *testing.T) {
	testUtils.UserClassRepo.MockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.UserClassRepo = &testUtils.UserClassRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/userclass/uc/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
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
