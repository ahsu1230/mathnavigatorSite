package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get User
//
func TestGetUser_Success(t *testing.T) {
	userRepo.mockSelectById = func(id uint) (domains.User, error) {
		user := createMockUser(
			1,
			"John",
			"Smith",
			"",
			"john_smith@example.com",
			"555-555-0199",
			true,
			0,
			"notes1",
		)
		return user, nil
	}
	repos.UserRepo = &userRepo

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
	assert.EqualValues(t, 0, user.FamilyId)
	assert.EqualValues(t, "notes1", user.Notes.String)

}

func TestGetUsersByFamily_Success(t *testing.T) {
	userRepo.mockSelectByFamilyId = func(familyId uint) ([]domains.User, error) {
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
				"notes1",
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
				"notes2",
			),
		}, nil
	}
	repos.UserRepo = &userRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/family/2", nil)

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
	assert.EqualValues(t, 2, users[0].FamilyId)
	assert.EqualValues(t, "notes1", users[0].Notes.String)

	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "Bob", users[1].FirstName)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, "Middle", users[1].MiddleName.String)
	assert.EqualValues(t, "bob_joe@example.com", users[1].Email)
	assert.EqualValues(t, "555-555-0199", users[1].Phone)
	assert.EqualValues(t, false, users[1].IsGuardian)
	assert.EqualValues(t, 2, users[1].FamilyId)
	assert.EqualValues(t, "notes2", users[1].Notes.String)

	assert.EqualValues(t, 2, len(users))
}

func TestGetUser_Failure(t *testing.T) {
	userRepo.mockSelectById = func(id uint) (domains.User, error) {
		return domains.User{}, errors.New("not found")
	}
	repos.UserRepo = &userRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateUser_Success(t *testing.T) {
	userRepo.mockInsert = func(user domains.User) error {
		return nil
	}
	repos.UserRepo = &userRepo

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
		"notes1",
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateUser_Failure(t *testing.T) {
	// no mock needed
	repos.UserRepo = &userRepo

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
		"",
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
	userRepo.mockUpdate = func(id uint, user domains.User) error {
		return nil // Successful update
	}
	repos.UserRepo = &userRepo

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
		"notes1",
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/user/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateUser_Invalid(t *testing.T) {
	// no mock needed
	repos.UserRepo = &userRepo

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
		"",
	)
	body := createBodyFromUser(user)
	recorder := sendHttpRequest(t, http.MethodPost, "/api/users/user/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateUser_Failure(t *testing.T) {
	userRepo.mockUpdate = func(id uint, user domains.User) error {
		return errors.New("not found")
	}
	repos.UserRepo = &userRepo

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
		"notes1",
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
	userRepo.mockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.UserRepo = &userRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteUser_Failure(t *testing.T) {
	userRepo.mockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.UserRepo = &userRepo

	// Create new HTTP request to endpoint
	recorder := sendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//

// func createMockFamily(id uint, email string, password string) domains.Family{
// 	return domains.Family{
// 		Id:	id,
// 		PrimaryEmail: email,
// 		Password: password,
// 	}
// }

func createMockUser(id uint, firstName, lastName, middleName, email, phone string, isGuardian bool, familyId uint, notes string) domains.User {
	return domains.User{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: domains.NewNullString(middleName),
		Email:      email,
		Phone:      phone,
		IsGuardian: isGuardian,
		FamilyId:   familyId,
		Notes:      domains.NewNullString(notes),
	}
}

func createBodyFromUser(user domains.User) io.Reader {
	marshal, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
