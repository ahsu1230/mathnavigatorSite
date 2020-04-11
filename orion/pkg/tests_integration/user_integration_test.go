package integration_tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Users and GetAll()
func Test_CreateUsers(t *testing.T) {
	createAllUsers(t)

	// Call Get All!
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/v1/all", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertUser(t, 1, users[0])
	assertUser(t, 2, users[1])
	assertUser(t, 3, users[2])
	assert.EqualValues(t, 3, len(users))

	resetTable(t, domains.TABLE_USERS)
}

// Test: Create 3 Users and GetUserByGuardianId
func Test_GetUsersByGuardian(t *testing.T) {
	createAllUsers(t)

	// Call Get All!
	recorder := sendHttpRequest(t, http.MethodGet, "/api/users/v1/guardian/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertUser(t, 2, users[0])
	assertUser(t, 3, users[1])
	assert.EqualValues(t, 2, len(users))

	resetTable(t, domains.TABLE_USERS)
}

// Test: Create 1 User, Update it, GetUserById()
func Test_UpdateUser(t *testing.T) {
	// Create 1 User
	user1 := createUser(1)
	body1 := createJsonBody(&user1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedUser := createUser(2)
	updatedBody := createJsonBody(&updatedUser)
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

	assertUser(t, 2, user)

	resetTable(t, domains.TABLE_USERS)
}

// Test: Create 1 User, Delete it, GetByUserId()
func Test_DeleteUser(t *testing.T) {
	// Create
	user1 := createUser(1)
	body1 := createJsonBody(&user1)
	recorder1 := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := sendHttpRequest(t, http.MethodDelete, "/api/users/v1/user/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := sendHttpRequest(t, http.MethodGet, "/api/users/v1/user/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	resetTable(t, domains.TABLE_USERS)
}

// Helper methods
func createUser(id int) domains.User {
	switch id {
	case 1:
		return domains.User{
			FirstName:  "John",
			LastName:   "Smith",
			MiddleName: domains.NewNullString("Middle"),
			Email:      "john_smith@example.com",
			Phone:      "555-555-0100",
			IsGuardian: true,
			GuardianId: domains.NewNullUint(0),
		}
	case 2:
		return domains.User{
			FirstName:  "Bob",
			LastName:   "Joe",
			MiddleName: domains.NewNullString(""),
			Email:      "bob_joe@example.com",
			Phone:      "555-555-0101",
			IsGuardian: false,
			GuardianId: domains.NewNullUint(1),
		}
	case 3:
		return domains.User{
			FirstName:  "Foo",
			LastName:   "Bar",
			MiddleName: domains.NewNullString(""),
			Email:      "foobar@example.com",
			Phone:      "555-555-0102",
			IsGuardian: false,
			GuardianId: domains.NewNullUint(1),
		}
	default:
		return domains.User{}
	}
}

func createAllUsers(t *testing.T) {
	for i := 1; i < 4; i++ {
		user := createUser(i)
		body := createJsonBody(&user)
		recorder := sendHttpRequest(t, http.MethodPost, "/api/users/v1/create", body)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
}

func assertUser(t *testing.T, id int, user domains.User) {
	switch id {
	case 1:
		assert.EqualValues(t, "John", user.FirstName)
		assert.EqualValues(t, "Smith", user.LastName)
		assert.EqualValues(t, "Middle", user.MiddleName.String)
		assert.EqualValues(t, "john_smith@example.com", user.Email)
		assert.EqualValues(t, "555-555-0100", user.Phone)
		assert.EqualValues(t, true, user.IsGuardian)
		assert.EqualValues(t, 0, user.GuardianId.Uint)
	case 2:
		assert.EqualValues(t, "Bob", user.FirstName)
		assert.EqualValues(t, "Joe", user.LastName)
		assert.EqualValues(t, "", user.MiddleName.String)
		assert.EqualValues(t, "bob_joe@example.com", user.Email)
		assert.EqualValues(t, "555-555-0101", user.Phone)
		assert.EqualValues(t, false, user.IsGuardian)
		assert.EqualValues(t, 1, user.GuardianId.Uint)
	case 3:
		assert.EqualValues(t, "Foo", user.FirstName)
		assert.EqualValues(t, "Bar", user.LastName)
		assert.EqualValues(t, "", user.MiddleName.String)
		assert.EqualValues(t, "foobar@example.com", user.Email)
		assert.EqualValues(t, "555-555-0102", user.Phone)
		assert.EqualValues(t, false, user.IsGuardian)
		assert.EqualValues(t, 1, user.GuardianId.Uint)
	}
}
