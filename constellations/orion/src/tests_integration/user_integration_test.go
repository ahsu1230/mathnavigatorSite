package tests_integration

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Users and GetAll (new)
func TestE2ECreateUsers(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	// Call Get All (new)
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/users/new", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Validate results
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, utils.UserTonyStark, users[0])
	assertUser(t, utils.UserMorganStark, users[1])
	assertUser(t, utils.UserPeterParker, users[2])
	assertUser(t, utils.UserNatasha, users[3])
	assertUser(t, utils.UserPotter, users[4])
	assert.EqualValues(t, 5, len(users))

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)

}

// Test: Create 3 Users and search by pagination
func TestE2ESearchUsers(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	body := strings.NewReader(`{
		"query": "stark"	
	}`)

	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/users/search", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertUser(t, utils.UserTonyStark, users[0])
	assertUser(t, utils.UserMorganStark, users[1])

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// TODO (aaron) TEST IGNORED - Need to fix the JSON Unmarshaling for nested structs?
// Test: Create 3 Users and GetUsersByIds
func IgnoreTestE2EGetUsersIds(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	// Validate results (first call)
	body := strings.NewReader(`[1,2]`)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/users/map", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userMap1 map[uint]domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &userMap1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, utils.UserTonyStark, userMap1[1])
	assertUser(t, utils.UserMorganStark, userMap1[2])
	assert.EqualValues(t, 3, len(userMap1))

	// Validate results (second call)
	body = strings.NewReader(`[2,3]`)
	recorder = utils.SendHttpRequest(t, http.MethodPost, "/api/users/map", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userMap2 map[uint]domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &userMap2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, utils.UserMorganStark, userMap2[2])
	assertUser(t, utils.UserPeterParker, userMap2[3])
	assert.EqualValues(t, 2, len(userMap2))

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 3 Users and GetUserByAccountId
func TestE2EGetUsersByAccountId(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	// Validate results (first call)
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/users/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var usersForAccount1 []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &usersForAccount1); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, utils.UserTonyStark, usersForAccount1[0])
	assertUser(t, utils.UserMorganStark, usersForAccount1[1])
	assertUser(t, utils.UserPeterParker, usersForAccount1[2])
	assert.EqualValues(t, 3, len(usersForAccount1))

	// Validate results (second call)
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/users/account/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var usersForAccount2 []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &usersForAccount2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUser(t, utils.UserNatasha, usersForAccount2[0])
	assert.EqualValues(t, 1, len(usersForAccount2))

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 2 Users with the same email
func TestE2ECreateSameEmailUsersFailure(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	// Create account and user with same email
	account1 := utils.AccountTonyStark
	user1 := utils.UserTonyStark
	_, recorder := utils.SendCreateAccountUser(t, false, account1, user1)
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 Account, 1 User, Update it, GetUserById()
func TestE2EUpdateUser(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	// Update
	updatedUser := domains.User{
		AccountId:      1,
		FirstName:      "Tony",
		LastName:       "Stank",                         // changed
		MiddleName:     domains.NewNullString("Edward"), // changed
		Email:          "ironman@stark.com",             // changed
		Phone:          domains.NewNullString("555-555-0101"),
		IsGuardian:     true,
		Notes:          domains.NewNullString("Avengers CEO"),
		School:         domains.NewNullString("MIT"), // changed
		GraduationYear: domains.NewNullUint(0),
	}
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
	assertUser(t, updatedUser, user)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 User, Delete it, GetByUserId()
func TestE2EDeleteUser(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)

}

// Helper methods
func assertUser(t *testing.T, expectedUser domains.User, user domains.User) {
	assert.EqualValues(t, expectedUser.AccountId, user.AccountId)
	assert.EqualValues(t, expectedUser.FirstName, user.FirstName)
	assert.EqualValues(t, expectedUser.LastName, user.LastName)
	assert.EqualValues(t, expectedUser.MiddleName.String, user.MiddleName.String)
	assert.EqualValues(t, expectedUser.Email, user.Email)
	assert.EqualValues(t, expectedUser.Phone.String, user.Phone.String)
	assert.EqualValues(t, expectedUser.IsGuardian, user.IsGuardian)
	assert.EqualValues(t, expectedUser.Notes.String, user.Notes.String)
	assert.EqualValues(t, expectedUser.School.String, user.School.String)
	assert.EqualValues(t, expectedUser.GraduationYear.Uint, user.GraduationYear.Uint)
}
