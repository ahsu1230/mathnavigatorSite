package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get User
//
func TestGetUserSuccess(t *testing.T) {
	testUtils.UserRepo.MockSelectById = func(context.Context, uint) (domains.User, error) {
		user := testUtils.CreateMockUser(
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
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)

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
	assert.EqualValues(t, 0, user.AccountId)
	assert.EqualValues(t, "notes1", user.Notes.String)

}

func TestGetUsersByAccountSuccess(t *testing.T) {
	testUtils.UserRepo.MockSelectByAccountId = func(context.Context, uint) ([]domains.User, error) {
		return []domains.User{
			testUtils.CreateMockUser(
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
			testUtils.CreateMockUser(
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
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/users/account/2", nil)

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
	assert.EqualValues(t, 2, users[0].AccountId)
	assert.EqualValues(t, "notes1", users[0].Notes.String)

	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "Bob", users[1].FirstName)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, "Middle", users[1].MiddleName.String)
	assert.EqualValues(t, "bob_joe@example.com", users[1].Email)
	assert.EqualValues(t, "555-555-0199", users[1].Phone)
	assert.EqualValues(t, false, users[1].IsGuardian)
	assert.EqualValues(t, 2, users[1].AccountId)
	assert.EqualValues(t, "notes2", users[1].Notes.String)

	assert.EqualValues(t, 2, len(users))
}

func TestGetUserFailure(t *testing.T) {
	testUtils.UserRepo.MockSelectById = func(context.Context, uint) (domains.User, error) {
		return domains.User{}, appErrors.MockDbNoRowsError()
	}
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Get New
//
func TestGetNewUsers(t *testing.T) {
	testUtils.UserRepo.MockSelectByNew = func(context.Context) ([]domains.User, error) {
		return []domains.User{
			testUtils.CreateMockUser(
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
			testUtils.CreateMockUser(
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
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/users/new", nil)

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
	assert.EqualValues(t, 2, users[0].AccountId)
	assert.EqualValues(t, "notes1", users[0].Notes.String)

	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "Bob", users[1].FirstName)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, "Middle", users[1].MiddleName.String)
	assert.EqualValues(t, "bob_joe@example.com", users[1].Email)
	assert.EqualValues(t, "555-555-0199", users[1].Phone)
	assert.EqualValues(t, false, users[1].IsGuardian)
	assert.EqualValues(t, 2, users[1].AccountId)
	assert.EqualValues(t, "notes2", users[1].Notes.String)

	assert.EqualValues(t, 2, len(users))
}

//
// Test Create
//
func TestCreateUserSuccess(t *testing.T) {
	testUtils.UserRepo.MockInsert = func(context.Context, domains.User) (uint, error) {
		return 42, nil
	}
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	user := testUtils.CreateMockUser(
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
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/users/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateUserFailure(t *testing.T) {
	// no mock needed
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	user := testUtils.CreateMockUser(
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
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/users/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestSearchUsersSuccess(t *testing.T) {
	testUtils.UserRepo.MockSearchUsers = func(context.Context, string) ([]domains.User, error) {
		return []domains.User{
			testUtils.CreateMockUser(
				1,
				"Ben",
				"Smith",
				"",
				"ben_smith@example.com",
				"555-555-0199",
				false,
				2,
				"notes1",
			),
			testUtils.CreateMockUser(
				2,
				"B",
				"Joe",
				"Middle",
				"benjamin_joe@example.com",
				"555-555-0199",
				false,
				2,
				"notes2",
			),
			testUtils.CreateMockUser(
				3,
				"Bob",
				"Ben",
				"Middle",
				"bob@example.com",
				"555-555-0199",
				false,
				3,
				"notes3",
			),
			testUtils.CreateMockUser(
				4,
				"A",
				"J",
				"Ben",
				"a_j@example.com",
				"555-555-0199",
				false,
				4,
				"notes4",
			),
		}, nil
	}
	repos.UserRepo = &testUtils.UserRepo

	// Create search body for HTTP request
	userSearchBody := controllers.UserSearchBody{
		"Ben",
	}
	marshal, err := json.Marshal(&userSearchBody)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(marshal)

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/users/search", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, users[0].Id)
	assert.EqualValues(t, "Ben", users[0].FirstName)
	assert.EqualValues(t, "", users[0].MiddleName.String)
	assert.EqualValues(t, "Smith", users[0].LastName)
	assert.EqualValues(t, "ben_smith@example.com", users[0].Email)

	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "B", users[1].FirstName)
	assert.EqualValues(t, "Middle", users[1].MiddleName.String)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, "benjamin_joe@example.com", users[1].Email)

	assert.EqualValues(t, 3, users[2].Id)
	assert.EqualValues(t, "Bob", users[2].FirstName)
	assert.EqualValues(t, "Middle", users[2].MiddleName.String)
	assert.EqualValues(t, "Ben", users[2].LastName)
	assert.EqualValues(t, "bob@example.com", users[2].Email)

	assert.EqualValues(t, 4, users[3].Id)
	assert.EqualValues(t, "A", users[3].FirstName)
	assert.EqualValues(t, "Ben", users[3].MiddleName.String)
	assert.EqualValues(t, "J", users[3].LastName)
	assert.EqualValues(t, "a_j@example.com", users[3].Email)

}

func TestSearchUsersFailure(t *testing.T) {
	testUtils.UserRepo.MockSearchUsers = func(context.Context, string) ([]domains.User, error) {
		return []domains.User{}, appErrors.MockDbNoRowsError()
	}
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/users/search", nil)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateUserSuccess(t *testing.T) {
	testUtils.UserRepo.MockUpdate = func(context.Context, uint, domains.User) error {
		return nil // Successful update
	}
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	user := testUtils.CreateMockUser(
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
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/users/user/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateUserInvalid(t *testing.T) {
	// no mock needed
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	user := testUtils.CreateMockUser(
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
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/users/user/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateUserFailure(t *testing.T) {
	testUtils.UserRepo.MockUpdate = func(context.Context, uint, domains.User) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	user := testUtils.CreateMockUser(
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
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/users/user/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteUserSuccess(t *testing.T) {
	testUtils.UserRepo.MockDelete = func(context.Context, uint) error {
		return nil // Return no error, successful delete!
	}
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteUserFailure(t *testing.T) {
	testUtils.UserRepo.MockDelete = func(context.Context, uint) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.UserRepo = &testUtils.UserRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Helper Methods
//

func createMockUser(id uint, firstName, lastName, middleName, email, phone string, isGuardian bool, accountId uint, notes string) domains.User {
	return domains.User{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: domains.NewNullString(middleName),
		Email:      email,
		Phone:      phone,
		IsGuardian: isGuardian,
		AccountId:  accountId,
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
