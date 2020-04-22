package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/services"
	"github.com/stretchr/testify/assert"
)

//
// Test Get All
//
func TestGetAllUsers_Success(t *testing.T) {
	userService.mockGetAll = func(search string, pageSize, offset int) ([]domains.User, error) {
		return []domains.User{
			createMockUser(
				1,
				"John",
				"Smith",
				"",
				"john_smith@example.com",
				"555-555-0199",
				true,
				0,
			),
			createMockUser(
				2,
				"Bob",
				"Joe",
				"Middle",
				"bob_joe@example.com",
				"555-555-0199",
				false,
				1,
			),
		}, nil
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, users[0].Id)
	assert.EqualValues(t, "John", users[0].FirstName)
	assert.EqualValues(t, "Smith", users[0].LastName)
	assert.EqualValues(t, "", users[0].MiddleName.String)
	assert.EqualValues(t, "john_smith@example.com", users[0].Email)
	assert.EqualValues(t, "555-555-0199", users[0].Phone)
	assert.EqualValues(t, true, users[0].IsGuardian)
	assert.EqualValues(t, 0, users[0].GuardianId.Uint)

	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "Bob", users[1].FirstName)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, "Middle", users[1].MiddleName.String)
	assert.EqualValues(t, "bob_joe@example.com", users[1].Email)
	assert.EqualValues(t, "555-555-0199", users[1].Phone)
	assert.EqualValues(t, false, users[1].IsGuardian)
	assert.EqualValues(t, 1, users[1].GuardianId.Uint)
	assert.EqualValues(t, 2, len(users))
}

//
// Test Get User
//
func TestGetUser_Success(t *testing.T) {
	userService.mockGetById = func(id uint) (domains.User, error) {
		user := createMockUser(
			1,
			"John",
			"Smith",
			"",
			"john_smith@example.com",
			"555-555-0199",
			true,
			0,
		)
		return user, nil
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var user domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &user); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "John", user.FirstName)
	assert.EqualValues(t, "Smith", user.LastName)
	assert.EqualValues(t, "", user.MiddleName.String)
	assert.EqualValues(t, "john_smith@example.com", user.Email)
	assert.EqualValues(t, "555-555-0199", user.Phone)
	assert.EqualValues(t, true, user.IsGuardian)
	assert.EqualValues(t, 0, user.GuardianId.Uint)
}

func TestGetUserByGuardian_Success(t *testing.T) {
	userService.mockGetByGuardianId = func(guardianId uint) ([]domains.User, error) {
		return []domains.User{
			createMockUser(
				1,
				"John",
				"Smith",
				"",
				"john_smith@example.com",
				"555-555-0199",
				false,
				2,
			),
			createMockUser(
				2,
				"Bob",
				"Joe",
				"Middle",
				"bob_joe@example.com",
				"555-555-0199",
				false,
				2,
			),
		}, nil
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/guardian/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, users[0].Id)
	assert.EqualValues(t, "John", users[0].FirstName)
	assert.EqualValues(t, "Smith", users[0].LastName)
	assert.EqualValues(t, "", users[0].MiddleName.String)
	assert.EqualValues(t, "john_smith@example.com", users[0].Email)
	assert.EqualValues(t, "555-555-0199", users[0].Phone)
	assert.EqualValues(t, false, users[0].IsGuardian)
	assert.EqualValues(t, 2, users[0].GuardianId.Uint)

	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "Bob", users[1].FirstName)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, "Middle", users[1].MiddleName.String)
	assert.EqualValues(t, "bob_joe@example.com", users[1].Email)
	assert.EqualValues(t, "555-555-0199", users[1].Phone)
	assert.EqualValues(t, false, users[1].IsGuardian)
	assert.EqualValues(t, 2, users[1].GuardianId.Uint)

	assert.EqualValues(t, 2, len(users))
}

func TestGetUser_Failure(t *testing.T) {
	userService.mockGetById = func(id uint) (domains.User, error) {
		return domains.User{}, errors.New("not found")
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)

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
		"",
		"john_smith@example.com",
		"555-555-0199",
		true,
		0,
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/create", body)

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
		"",
		"",
		"",
		false,
		0,
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/create", body)

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
		"",
		"john_smith@example.com",
		"555-555-0199",
		true,
		0,
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/user/1", body)

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
		"",
		"",
		"",
		false,
		0,
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/user/1", body)

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
		"",
		"john_smith@example.com",
		"555-555-0199",
		true,
		0,
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/user/1", body)

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
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteUser_Failure(t *testing.T) {
	userService.mockDelete = func(id uint) error {
		return errors.New("not found")
	}
	services.UserService = &userService

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockUser(id uint, firstName, lastName, middleName, email, phone string, isGuardian bool, guardianId uint) domains.User {
	return domains.User{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: domains.NewNullString(middleName),
		Email:      email,
		Phone:      phone,
		IsGuardian: isGuardian,
		GuardianId: domains.NewNullUint(guardianId),
	}
}

func createBodyFromUser(user domains.User) io.Reader {
	marshal, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
