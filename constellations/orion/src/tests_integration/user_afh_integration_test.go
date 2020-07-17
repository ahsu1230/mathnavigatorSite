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
func Test_GetUserAfhsByUserId(t *testing.T) {
	createAllUserAfhs(t)

	// Call GetUserAfhsByUserId()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/users/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfhs []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfhs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 2, userAfhs[0].UserId)
	assert.EqualValues(t, 2, userAfhs[0].AfhId)

	resetUserAfhTables(t)
}

// Create 3 UserAfhs. Then update one. Then get by afhId
func Test_GetUserAfhsByAfhId(t *testing.T) {
	createAllUserAfhs(t)

	// Update
	updatedUserAfh := createUserAfh(2, 1)
	updatedBod := utils.CreateJsonBody(&updatedUserAfh)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/userafhs/userafh/1", updatedBod)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Call GetUserAfhsByAfhId()
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/afh/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfhs []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfhs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, userAfhs[0].UserId)
	assert.EqualValues(t, 2, userAfhs[0].AfhId)
	assert.EqualValues(t, 2, userAfhs[1].UserId)
	assert.EqualValues(t, 2, userAfhs[1].AfhId)

	resetUserAfhTables(t)
}

// Create 3 UserAfhs. Then delete, then get that one by userId and afhId
func Test_DeleteUserAfh(t *testing.T) {
	createAllUserAfhs(t)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/userafhs/userafh/3", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get by both ids
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/users/2/afh/2", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	resetUserAfhTables(t)
}

// Create 3 UserAfhs. Then get one by userId and afhId
func Test_GetUserAfhsByBothIds(t *testing.T) {
	createAllUserAfhs(t)

	// Call get userAfh by both ids
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/userafhs/users/2/afh/2", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfh domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfh); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 2, userAfh.UserId)
	assert.EqualValues(t, 2, userAfh.AfhId)

	resetUserAfhTables(t)
}

// Helper methods
func createAllUserAfhs(t *testing.T) {
	createUsersAndAfhs(t)

	userAfh1 := createUserAfh(1, 1)
	userAfh2 := createUserAfh(1, 2)
	userAfh3 := createUserAfh(2, 2)

	body1 := utils.CreateJsonBody(&userAfh1)
	body2 := utils.CreateJsonBody(&userAfh2)
	body3 := utils.CreateJsonBody(&userAfh3)

	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/userafhs/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/userafhs/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/userafhs/create", body3)

	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
}

func createUsersAndAfhs(t *testing.T) {
	var date1 = now.Add(time.Hour * 24 * 30)

	// Create accounts
	account := createAccount(1)
	body := utils.CreateJsonBody(&account)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	account0 := createAccount(2)
	body0 := utils.CreateJsonBody(&account0)
	recorder0 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body0)
	assert.EqualValues(t, http.StatusOK, recorder0.Code)

	// Create locations
	location1 := createLocation("wchs", "11300 Gainsborough Road", "Potomac", "MD", "20854", "Room 100")
	location2 := createLocation("room12", "123 Sesame St", "Rockville", "MD", "20814", "Room 8")
	locBody1 := utils.CreateJsonBody(&location1)
	locBody2 := utils.CreateJsonBody(&location2)
	locRecorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", locBody1)
	locRecorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", locBody2)
	assert.EqualValues(t, http.StatusOK, locRecorder1.Code)
	assert.EqualValues(t, http.StatusOK, locRecorder2.Code)

	// Create users and AFHs
	user1 := createUser(1)
	user2 := createUser(2)
	afh1 := createAFH(
		1,
		"AP Calculus Help",
		date1,
		"2:00-4:00PM",
		"AP Calculus",
		"wchs",
		"test note",
	)
	afh2 := createAFH(
		2,
		"AP Statistics Help",
		date1,
		"3:00-5:00PM",
		"AP Statistics",
		"room12",
		"test note 2",
	)

	body1 := utils.CreateJsonBody(&user1)
	body2 := utils.CreateJsonBody(&user2)
	body3 := utils.CreateJsonBody(&afh1)
	body4 := utils.CreateJsonBody(&afh2)

	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/users/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/users/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body3)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/askforhelp/create", body4)

	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
}

func createUserAfh(userId, afhId uint) domains.UserAfh {
	return domains.UserAfh{
		UserId: userId,
		AfhId:  afhId,
	}
}

func resetUserAfhTables(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_USERAFH)
	utils.ResetTable(t, domains.TABLE_ASKFORHELP)
	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}
