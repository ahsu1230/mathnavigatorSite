package tests_integration

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 UserAfhs. Then get by userId
func TestGetUserAfhsByUserId(t *testing.T) {
	createAllUserAfhs(t)

	// Call GetUserAfhsByUserId()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfhs []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfhs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 2, userAfhs[0].UserId)
	assert.EqualValues(t, 2, userAfhs[0].AfhId)
	assert.EqualValues(t, 1, userAfhs[0].AccountId)

	resetUserAfhTables(t)
}

// Create 3 UserAfhs. Then update one. Then get by afhId
func TestGetUserAfhsByAfhId(t *testing.T) {
	createAllUserAfhs(t)

	// Update
	updatedUserAfh := domains.UserAfh{
		AfhId:     1,
		UserId:    2,
		AccountId: 1,
	}
	updatedBod := utils.CreateJsonBody(&updatedUserAfh)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/user-afhs/user-afh/1", updatedBod)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Call GetUserAfhsByAfhId()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/afh/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfhs []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfhs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, userAfhs[0].UserId)
	assert.EqualValues(t, 2, userAfhs[0].AfhId)
	assert.EqualValues(t, 1, userAfhs[0].AccountId)

	assert.EqualValues(t, 2, userAfhs[1].UserId)
	assert.EqualValues(t, 2, userAfhs[1].AfhId)
	assert.EqualValues(t, 1, userAfhs[1].AccountId)

	resetUserAfhTables(t)
}

// Create 3 UserAfhs. Then delete, then get that one by userId and afhId
func TestDeleteUserAfh(t *testing.T) {
	createAllUserAfhs(t)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/user-afhs/user-afh/3", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get by both ids
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/2/afh/2", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	resetUserAfhTables(t)
}

// Create 3 UserAfhs. Then get one by userId and afhId
func TestGetUserAfhsByBothIds(t *testing.T) {
	createAllUserAfhs(t)

	// Call get userAfh by both ids
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/2/afh/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 2, userAfh.UserId)
	assert.EqualValues(t, 2, userAfh.AfhId)
	assert.EqualValues(t, 1, userAfh.AccountId)

	resetUserAfhTables(t)
}

// Helper methods
func createAllUserAfhs(t *testing.T) {
	createUsersAndAfhs(t)

	// AfhId(1) attended by User(1) from Account(1)
	utils.SendCreateUserAfh(t, true, 1, 1, 1)

	// AfhId(2) attended by User(1) from Account(1)
	utils.SendCreateUserAfh(t, true, 2, 1, 1)

	// AfhId(2) attended by User(2) from Account(1)
	utils.SendCreateUserAfh(t, true, 2, 2, 1)
}

func createUsersAndAfhs(t *testing.T) {
	// Create 2 Accounts
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)
	utils.SendCreateAccountUser(t, true, utils.AccountNatasha, utils.UserNatasha)

	// Create locations
	utils.SendCreateLocationWCHS(t)
	utils.SendCreateLocation(
		t,
		true,
		"room12",
		"Sesame High School",
		"123 Sesame St",
		"Rockville",
		"MD",
		"20814",
		"Room 8",
		false,
	)

	// Create AFHs
	start1 := time.Now().UTC()
	end1 := start1.Add(time.Hour * 1)
	start2 := start1.Add(time.Hour * 7)
	end2 := start2.Add(time.Hour * 1)
	utils.SendCreateAskForHelp(
		t,
		true,
		start1,
		end1,
		"AP Calculus Help",
		domains.SUBJECT_MATH,
		"wchs",
		"test note",
	)
	utils.SendCreateAskForHelp(
		t,
		true,
		start2,
		end2,
		"AP Statistics Help",
		domains.SUBJECT_MATH,
		"room12",
		"test note 2",
	)
}

func resetUserAfhTables(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_USER_AFHS)
	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}
