package integration_tests

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 3 Users and GetAll()
func Test_CreateUsers(t *testing.T) {
	resetTable(t, domains.TABLE_USERS)

	user1 := createUser(
		"John",
		"Smith",
		"",
		"john_smith@example.com",
		"555-555-0100",
		true,
		0,
	)
	user2 := createUser(
		"Bob",
		"Joe",
		"",
		"bob_joe@example.com",
		"555-555-0101",
		true,
		0,
	)
	user3 := createUser(
		"Foo",
		"Bar",
		"",
		"foobar@example.com",
		"555-555-0102",
		true,
		0,
	)
	body1 := createJsonBody(user1)
	body2 := createJsonBody(user2)
	body3 := createJsonBody(user3)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body1)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body2)
	recorder3 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Call Get All!
	recorder4 := sendHttpRequest(t, http.MethodGet, "/api/users/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder4.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, users[0].Id)
	assert.EqualValues(t, "John", users[0].FirstName)
	assert.EqualValues(t, "Smith", users[0].LastName)
	assert.EqualValues(t, "", users[0].MiddleName)
	assert.EqualValues(t, "john_smith@example.com", users[0].Email)
	assert.EqualValues(t, "555-555-0100", users[0].Phone)
	assert.EqualValues(t, true, users[0].IsGuardian)
	assert.EqualValues(t, "", users[0].GuardianId)
	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "Bob", users[1].FirstName)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, "", users[1].MiddleName)
	assert.EqualValues(t, "bob_joe@example.com", users[1].Email)
	assert.EqualValues(t, "555-555-0101", users[1].Phone)
	assert.EqualValues(t, true, users[1].IsGuardian)
	assert.EqualValues(t, "", users[1].GuardianId)
	assert.EqualValues(t, 3, users[2].Id)
	assert.EqualValues(t, "Foo", users[2].FirstName)
	assert.EqualValues(t, "Bar", users[2].LastName)
	assert.EqualValues(t, "", users[2].MiddleName)
	assert.EqualValues(t, "foobar@example.com", users[2].Email)
	assert.EqualValues(t, "555-555-0102", users[2].Phone)
	assert.EqualValues(t, true, users[2].IsGuardian)
	assert.EqualValues(t, "", users[2].GuardianId)
	assert.EqualValues(t, 3, len(users))
}

// Test: Create 1 User, Update it, GetByUserId()
func Test_UpdateUser(t *testing.T) {
	resetTable(t, domains.TABLE_USERS)

	// Create 1 User
	user1 := createUser(
		"John",
		"Smith",
		"",
		"john_smith@example.com",
		"555-555-0100",
		true,
		0,
	)
	body1 := createJsonBody(user1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedUser := createUser(
		"Bob",
		"Joe",
		"",
		"bob_joe@example.com",
		"555-555-0101",
		true,
		0,
	)
	updatedBody := createJsonBody(updatedUser)
	recorder2 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/user/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/users/v1/user/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var user domains.User
	if err := json.Unmarshal(recorder3.Body.Bytes(), &user); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Bob", user.FirstName)
	assert.EqualValues(t, "Joe", user.LastName)
	assert.EqualValues(t, "", user.MiddleName)
	assert.EqualValues(t, "bob_joe@example.com", user.Email)
	assert.EqualValues(t, "555-555-0101", user.Phone)
	assert.EqualValues(t, true, user.IsGuardian)
	assert.EqualValues(t, 0, user.GuardianId)
}

// Test: Create 1 User, Delete it, GetByUserId()
func Test_DeleteUser(t *testing.T) {
	resetTable(t, domains.TABLE_USERS)

	// Create
	user1 := createUser(
		"John",
		"Smith",
		"",
		"john_smith@example.com",
		"555-555-0100",
		true,
		0,
	)
	body1 := createJsonBody(user1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/users/v1/user/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/users/v1/user/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)
}

// Helper methods
func createUser(firstName, lastName, middleName, email, phone string, isGuardian bool, guardianId uint) domains.User {
	return domains.User{
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
		Email:      email,
		Phone:      phone,
		IsGuardian: isGuardian,
		GuardianId: guardianId,
	}
}
