package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/sql_helper"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

//
// Test Get All
//
func TestGetAllUsers_Success(t *testing.T) {
	userService.mockGetAll = func() ([]domains.User, error) {
		return []domains.User{
			{
				Id:         1,
				FirstName:  "John",
				LastName:   "Smith",
				MiddleName: sql.NullString{},
				Email:      "john_smith@example.com",
				Phone:      "555-555-0199",
				IsGuardian: true,
				GuardianId: sql_helper.NullUint{},
			},
			{
				Id:         2,
				FirstName:  "Bob",
				LastName:   "Joe",
				MiddleName: sql.NullString{String: "Middle", Valid: true},
				Email:      "bob_joe@example.com",
				Phone:      "555-555-0199",
				IsGuardian: false,
				GuardianId: sql_helper.NullUint{Uint: 1, Valid: true},
			},
		}, nil
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, users[0].Id)
	assert.EqualValues(t, "John", users[0].FirstName)
	assert.EqualValues(t, "Smith", users[0].LastName)
	assert.EqualValues(t, sql.NullString{}, users[0].MiddleName)
	assert.EqualValues(t, "john_smith@example.com", users[0].Email)
	assert.EqualValues(t, "555-555-0199", users[0].Phone)
	assert.EqualValues(t, true, users[0].IsGuardian)
	assert.EqualValues(t, sql_helper.NullUint{}, users[0].GuardianId)
	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "Bob", users[1].FirstName)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, sql.NullString{String: "Middle", Valid: true}, users[1].MiddleName)
	assert.EqualValues(t, "bob_joe@example.com", users[1].Email)
	assert.EqualValues(t, "555-555-0199", users[1].Phone)
	assert.EqualValues(t, false, users[1].IsGuardian)
	assert.EqualValues(t, sql_helper.NullUint{Uint: 1, Valid: true}, users[1].GuardianId)
	assert.EqualValues(t, 2, len(users))
}

//
// Test Get User
//
func TestGetUser_Success(t *testing.T) {
	userService.mockGetById = func(id uint) (domains.User, error) {
		user := createMockUser(1, "John", "message1")
		return user, nil
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/v1/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var user domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &user); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, 2020, user.Year)
	assert.EqualValues(t, "message1", user.Message)
}

func TestGetUser_Failure(t *testing.T) {
	userService.mockGetById = func(id uint) (domains.User, error) {
		return domains.User{}, errors.New("not found")
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/v1/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateUser_Success(t *testing.T) {
	userService.mockCreate = func(user domains.User) error {
		return nil
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	user := createMockUser(
		1,
		"John",
		"Smith",
		sql.NullString{},
		"john_smith@example.com",
		"555-555-0199",
		true,
		sql_helper.NullUint{},
	)
	marshal, _ := json.Marshal(user)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateUser_Failure(t *testing.T) {
	// no mock needed
	services.UserService = &userService

	// Create new HTTP request to endpoint
	user := createMockUser(
		1,
		"",
		"",
		sql.NullString{},
		"",
		"",
		false,
		sql_helper.NullUint{},
	)
	marshal, _ := json.Marshal(user)
	body := bytes.NewBuffer(marshal)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateUser_Success(t *testing.T) {
	userService.mockUpdate = func(id uint, user domains.User) error {
		return nil // Successful update
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	user := createMockUser(
		1,
		"John",
		"Smith",
		sql.NullString{},
		"john_smith@example.com",
		"555-555-0199",
		true,
		sql_helper.NullUint{},
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/v1/user/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateUser_Invalid(t *testing.T) {
	// no mock needed
	services.UserService = &userService

	// Create new HTTP request to endpoint
	user := createMockUser(
		1,
		"",
		"",
		sql.NullString{},
		"",
		"",
		false,
		sql_helper.NullUint{},
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/v1/user/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateUser_Failure(t *testing.T) {
	userService.mockUpdate = func(id uint, user domains.User) error {
		return errors.New("not found")
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	user := createMockUser(
		1,
		"John",
		"Smith",
		sql.NullString{},
		"john_smith@example.com",
		"555-555-0199",
		true,
		sql_helper.NullUint{},
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/v1/user/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteUser_Success(t *testing.T) {
	userService.mockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/users/v1/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteUser_Failure(t *testing.T) {
	userService.mockDelete = func(id uint) error {
		return errors.New("not found")
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/users/v1/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockUser(id uint, firstName string, lastName string, middleName sql.NullString, email string, phone string, isGuardian bool, guardianId sql_helper.NullUint) domains.User {
	return domains.User{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
		Email:      email,
		Phone:      phone,
		IsGuardian: isGuardian,
		GuardianId: guardianId,
	}
}

func createBodyFromUser(user domains.User) io.Reader {
	marshal, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
