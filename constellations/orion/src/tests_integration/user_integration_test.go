package tests_integration

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

// Test: Create 3 Users and GetAll()
func Test_CreateUsers(t *testing.T) {
	account1 := createAccount(1)
	body1 := utils.CreateJsonBody(&account1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	account2 := createAccount(2)
	body2 := utils.CreateJsonBody(&account2)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	createAllUsers(t)

	// Call Get All!
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/users/account/1", nil)

	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/users/account/2", nil)
	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, 1, users[0])
	assert.EqualValues(t, 1, len(users))

	if err := json.Unmarshal(recorder3.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertUser(t, 2, users[0])
	assertUser(t, 3, users[1])
	assert.EqualValues(t, 2, len(users))

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)

}

// Test: Create 3 Users and search by pagination
func Test_SearchUsers(t *testing.T) {
	account1 := createAccount(1)
	body1 := utils.CreateJsonBody(&account1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	account2 := createAccount(2)
	body2 := utils.CreateJsonBody(&account2)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	createAllUsers(t)

	body := strings.NewReader(`{
		"query": "smith"	
	}`)

	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/users/search", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertUser(t, 1, users[0])
	assertUser(t, 2, users[1])

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 3 Users and GetUserByAccountId
func Test_GetUsersByAccountId(t *testing.T) {
	account1 := createAccount(1)
	body1 := utils.CreateJsonBody(&account1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	account2 := createAccount(2)
	body2 := utils.CreateJsonBody(&account2)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	createAllUsers(t)
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/users/account/1", nil)

	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/users/account/2", nil)
	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, 1, users[0])
	assert.EqualValues(t, 1, len(users))

	if err := json.Unmarshal(recorder3.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertUser(t, 2, users[0])
	assertUser(t, 3, users[1])
	assert.EqualValues(t, 2, len(users))

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)

}

// Test: Create 1 Account, 1 User, Update it, GetUserById()
func Test_UpdateUser(t *testing.T) {
	account := createAccount(1)
	body := utils.CreateJsonBody(&account)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Create 1 User
	user1 := createUser(1)
	body1 := utils.CreateJsonBody(&user1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/users/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedUser := createUser(4)
	updatedBody := utils.CreateJsonBody(&updatedUser)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/users/user/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var user domains.User
	if err := json.Unmarshal(recorder3.Body.Bytes(), &user); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, 4, user)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 User, Delete it, GetByUserId()
func Test_DeleteUser(t *testing.T) {
	account := createAccount(1)
	body := utils.CreateJsonBody(&account)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Create
	user1 := createUser(1)
	body1 := utils.CreateJsonBody(&user1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/users/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)

}

// Helper methods
func createUser(id int) domains.User {
	switch id {
	case 1:
		return domains.User{
			FirstName:      "John",
			LastName:       "Smith",
			MiddleName:     domains.NewNullString("Middle"),
			Email:          "john_smith@example.com",
			Phone:          "555-555-0100",
			IsGuardian:     true,
			AccountId:      1,
			Notes:          domains.NewNullString("notes1"),
			School:         domains.NewNullString("schoolone"),
			GraduationYear: domains.NewNullUint(2001),
		}
	case 2:
		return domains.User{
			FirstName:      "Bob",
			LastName:       "Smith",
			MiddleName:     domains.NewNullString("Middle"),
			Email:          "bob_smith@example.com",
			Phone:          "555-555-0101",
			IsGuardian:     false,
			AccountId:      2,
			Notes:          domains.NewNullString("notes2"),
			School:         domains.NewNullString("schooltwo"),
			GraduationYear: domains.NewNullUint(2002),
		}
	case 3:
		return domains.User{
			FirstName:      "Foo",
			LastName:       "Bar",
			MiddleName:     domains.NewNullString("Smith"),
			Email:          "foobar@example.com",
			Phone:          "555-555-0102",
			IsGuardian:     false,
			AccountId:      2,
			Notes:          domains.NewNullString("notes3"),
			School:         domains.NewNullString(""),
			GraduationYear: domains.NewNullUint(0),
		}
	case 4:
		return domains.User{
			FirstName:      "Austin",
			LastName:       "Hsu",
			MiddleName:     domains.NewNullString(""),
			Email:          "austinhsu@example.com",
			Phone:          "555-555-0103",
			IsGuardian:     false,
			AccountId:      1,
			Notes:          domains.NewNullString("notes4"),
			School:         domains.NewNullString("schoolfour"),
			GraduationYear: domains.NewNullUint(2004),
		}
	default:
		return domains.User{}
	}
}

func createAllUsers(t *testing.T) {
	for i := 1; i < 4; i++ {
		user := createUser(i)
		body := utils.CreateJsonBody(&user)
		recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/users/create", body)
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
		assert.EqualValues(t, 1, user.AccountId)
		assert.EqualValues(t, "notes1", user.Notes.String)
		assert.EqualValues(t, "schoolone", user.School.String)
		assert.EqualValues(t, 2001, user.GraduationYear.Uint)
	case 2:
		assert.EqualValues(t, "Bob", user.FirstName)
		assert.EqualValues(t, "Smith", user.LastName)
		assert.EqualValues(t, "Middle", user.MiddleName.String)
		assert.EqualValues(t, "bob_smith@example.com", user.Email)
		assert.EqualValues(t, "555-555-0101", user.Phone)
		assert.EqualValues(t, false, user.IsGuardian)
		assert.EqualValues(t, 2, user.AccountId)
		assert.EqualValues(t, "notes2", user.Notes.String)
		assert.EqualValues(t, "schooltwo", user.School.String)
		assert.EqualValues(t, 2002, user.GraduationYear.Uint)
	case 3:
		assert.EqualValues(t, "Foo", user.FirstName)
		assert.EqualValues(t, "Bar", user.LastName)
		assert.EqualValues(t, "Smith", user.MiddleName.String)
		assert.EqualValues(t, "foobar@example.com", user.Email)
		assert.EqualValues(t, "555-555-0102", user.Phone)
		assert.EqualValues(t, false, user.IsGuardian)
		assert.EqualValues(t, 2, user.AccountId)
		assert.EqualValues(t, "notes3", user.Notes.String)
		assert.EqualValues(t, "", user.School.String)
		assert.EqualValues(t, 0, user.GraduationYear.Uint)

	}

}
