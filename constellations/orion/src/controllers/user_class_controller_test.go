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
	testUtils.UserClassesRepo.MockSelectByClassId = func(classId string) ([]domains.UserClasses, error) {
		return []domains.UserClasses{
			testUtils.CreateMockUserClasses(
				1,
				1,
				"abcd",
				1,
				domains.USER_CLASS_ACCEPTED,
			),
			testUtils.CreateMockUserClasses(
				2,
				2,
				"abcd",
				2,
				domains.USER_CLASS_TRIAL,
			),
		}, nil
	}
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/class/abcd", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userClass []domains.UserClasses
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

func TestGetClassesByUserId_Success(t *testing.T) {
	testUtils.UserClassesRepo.MockSelectByUserId = func(id uint) ([]domains.UserClasses, error) {
		return []domains.UserClasses{
			testUtils.CreateMockUserClasses(
				1,
				1,
				"abcd",
				1,
				domains.USER_CLASS_ACCEPTED,
			),
			testUtils.CreateMockUserClasses(
				2,
				1,
				"abce",
				1,
				domains.USER_CLASS_ACCEPTED,
			),
		}, nil
	}
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userClass []domains.UserClasses
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

func TestGetUserClassByUserAndClass_Success(t *testing.T) {
	testUtils.UserClassesRepo.MockSelectByUserAndClass = func(id uint, classId string) (domains.UserClasses, error) {
		userClass := testUtils.CreateMockUserClasses(
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
	var userClass domains.UserClasses
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, userClass.Id)
	assert.EqualValues(t, 1, userClass.UserId)
	assert.EqualValues(t, "abcd", userClass.ClassId)
	assert.EqualValues(t, 1, userClass.AccountId)
	assert.EqualValues(t, domains.USER_CLASS_ACCEPTED, userClass.State)

}

//
// Test Create
//
func TestCreateUserClass_Success(t *testing.T) {
	testUtils.UserClassesRepo.MockInsert = func(userClass domains.UserClasses) error {
		return nil
	}
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClasses(
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

func TestCreateUserClass_Failure(t *testing.T) {
	testUtils.UserClassesRepo.MockInsert = func(userClass domains.UserClasses) error {
		return errors.New("not found")
	}
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClasses(
		1,
		0,
		"",
		0,
		domains.USER_CLASS_PENDING,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Update
//
func TestUpdateUserClass_Success(t *testing.T) {
	testUtils.UserClassesRepo.MockUpdate = func(id uint, userClass domains.UserClasses) error {
		return nil // Successful update
	}
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClasses(
		1,
		1,
		"abcd",
		1,
		domains.USER_CLASS_ACCEPTED,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/userclass/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateUserClass_Invalid(t *testing.T) {
	testUtils.UserClassesRepo.MockUpdate = func(id uint, userClass domains.UserClasses) error {
		return errors.New("not found")
	}
	// no mock needed
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClasses(
		1,
		0,
		"",
		0,
		domains.USER_CLASS_PENDING,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/userclass/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

func TestUpdateUserClass_Failure(t *testing.T) {
	testUtils.UserClassesRepo.MockUpdate = func(id uint, userClass domains.UserClasses) error {
		return errors.New("not found")
	}
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	userClass := testUtils.CreateMockUserClasses(
		1,
		1,
		"abcd",
		1,
		domains.USER_CLASS_ACCEPTED,
	)
	body := createBodyFromUserClass(userClass)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/userclass/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteUserClass_Success(t *testing.T) {
	testUtils.UserClassesRepo.MockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/user-classes/userclass/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteUserClass_Failure(t *testing.T) {
	testUtils.UserClassesRepo.MockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.UserClassesRepo = &testUtils.UserClassesRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/user-classes/userclass/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

func TestStateValues_Success(t *testing.T) {
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/states", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

//
// Helper Methods
//

func createMockUserClass(id uint, userId uint, classId string, accountId uint, state uint) domains.UserClasses {
	return domains.UserClasses{
		Id:        id,
		UserId:    userId,
		ClassId:   classId,
		AccountId: accountId,
		State:     state,
	}
}

func createBodyFromUserClass(userClass domains.UserClasses) io.Reader {
	marshal, err := json.Marshal(&userClass)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
