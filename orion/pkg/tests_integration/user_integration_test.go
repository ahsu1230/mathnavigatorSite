package integration_tests

import (
	"database/sql"
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/sql_helper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 3 Users and GetAll()
func Test_CreateUsers(t *testing.T) {
	refreshTable(t, domains.TABLE_USERS)

	user1 := createUser(
		"John",
		"Smith",
		sql.NullString{},
		"john_smith@example.com",
		"555-555-0100",
		true,
		sql_helper.NullUint{},
	)
	user2 := createUser(
		"Bob",
		"Joe",
		sql.NullString{},
		"bob_joe@example.com",
		"555-555-0101",
		true,
		sql_helper.NullUint{},
	)
	user3 := createUser(
		"Foo",
		"Bar",
		sql.NullString{},
		"foobar@example.com",
		"555-555-0102",
		true,
		sql_helper.NullUint{},
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
	assert.EqualValues(t, sql.NullString{}, users[0].MiddleName)
	assert.EqualValues(t, "john_smith@example.com", users[0].Email)
	assert.EqualValues(t, "555-555-0100", users[0].Phone)
	assert.EqualValues(t, true, users[0].IsGuardian)
	assert.EqualValues(t, sql_helper.NullUint{}, users[0].GuardianId)
	assert.EqualValues(t, 2, users[1].Id)
	assert.EqualValues(t, "Bob", users[1].FirstName)
	assert.EqualValues(t, "Joe", users[1].LastName)
	assert.EqualValues(t, sql.NullString{}, users[1].MiddleName)
	assert.EqualValues(t, "bob_joe@example.com", users[1].Email)
	assert.EqualValues(t, "555-555-0101", users[1].Phone)
	assert.EqualValues(t, true, users[1].IsGuardian)
	assert.EqualValues(t, sql_helper.NullUint{}, users[1].GuardianId)
	assert.EqualValues(t, 3, users[2].Id)
	assert.EqualValues(t, "Foo", users[2].FirstName)
	assert.EqualValues(t, "Bar", users[2].LastName)
	assert.EqualValues(t, sql.NullString{}, users[2].MiddleName)
	assert.EqualValues(t, "foobar@example.com", users[2].Email)
	assert.EqualValues(t, "555-555-0102", users[2].Phone)
	assert.EqualValues(t, true, users[2].IsGuardian)
	assert.EqualValues(t, sql_helper.NullUint{}, users[2].GuardianId)
	assert.EqualValues(t, 3, len(users))
}

// Test: Create 1 User, Update it, GetByUserId()
func Test_UpdateUser(t *testing.T) {
	refreshTable(t, domains.TABLE_USERS)

	// Create 1 User
	user1 := createUser(
		"John",
		"Smith",
		sql.NullString{},
		"john_smith@example.com",
		"555-555-0100",
		true,
		sql_helper.NullUint{},
	)
	body1 := createJsonBody(user1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedUser := createUser(
		"Bob",
		"Joe",
		sql.NullString{},
		"bob_joe@example.com",
		"555-555-0101",
		true,
		sql_helper.NullUint{},
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
	assert.EqualValues(t, sql.NullString{}, user.MiddleName)
	assert.EqualValues(t, "bob_joe@example.com", user.Email)
	assert.EqualValues(t, "555-555-0101", user.Phone)
	assert.EqualValues(t, true, user.IsGuardian)
	assert.EqualValues(t, sql_helper.NullUint{}, user.GuardianId)
}

// Test: Create 1 User, Delete it, GetByUserId()
func Test_DeleteUser(t *testing.T) {
	refreshTable(t, domains.TABLE_USERS)

	// Create
	user1 := createUser(
		"John",
		"Smith",
		sql.NullString{},
		"john_smith@example.com",
		"555-555-0100",
		true,
		sql_helper.NullUint{},
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
func createUser(firstName string, lastName string, middleName sql.NullString, email string, phone string, isGuardian bool, guardianId sql_helper.NullUint) domains.User {
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
