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

// Full delete an account
// Create 1class, 1afh, 1location, 1program, 1semester
// Create 1 account, 2 users (userA & userB)
// userB is registered for class and afh
//
// When full deleting account,
// The class, afh, location, program, semester are still intact
// The account is not found
// userA and userB are not found
func TestE2EFullDeleteAccount(t *testing.T) {
	// Create Environment
	createCrossEnvironment(t)
	// accountId: 1, userId: 1
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)
	// userId: 2
	utils.SendCreateUser(t, true, utils.UserMorganStark)
	// Create UserAfh & UserClass for (userId 2, accountId 1)
	utils.SendCreateUserAfh(t, true, 1, 2, 1)
	utils.SendCreateUserClass(t, true, "program1_2020_spring_classA", 2, 1)

	// Full Delete Account
	recorderDel := utils.SendHttpRequest(t, http.MethodDelete, "/api/accounts/full/account/1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorderDel.Code)

	// Check account is not found
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	// Check userA & userB are not found
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/users/user/2", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	// Check for UserAfh & UserClasses are not found
	var userClasses []domains.UserClass
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClasses); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(userClasses))

	var userAfhs []domains.UserAfh
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfhs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(userAfhs))

	utils.ResetAllTables(t)
}

// Full delete a user
// Create 1class, 1afh, 1location, 1program, 1semester
// Create 1 account, 2 users (userA & userB)
// userB is registered for class and afh
//
// When full deleting userB,
// The class, afh, location, program, semester are still intact
// The account is intact with userA
// userB is not found
func TestE2EFullDeleteUser(t *testing.T) {
	// Create Environment
	createCrossEnvironment(t)
	// accountId: 1, userId: 1
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)
	// userId: 2
	utils.SendCreateUser(t, true, utils.UserMorganStark)
	// Create UserAfh & UserClass for (userId 2, accountId 1)
	utils.SendCreateUserAfh(t, true, 1, 2, 1)
	utils.SendCreateUserClass(t, true, "program1_2020_spring_classA", 2, 1)

	// Full Delete UserB (MorganStark, userId 2)
	recorderDel := utils.SendHttpRequest(t, http.MethodDelete, "/api/users/full/user/2", nil)
	assert.EqualValues(t, http.StatusNoContent, recorderDel.Code)

	// Check account is still found
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Check userA is still found
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Check userB is not found
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/users/user/2", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	// Check for UserAfh & UserClasses are not found
	var userClasses []domains.UserClass
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClasses); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(userClasses))

	var userAfhs []domains.UserAfh
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfhs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 0, len(userAfhs))

	utils.ResetAllTables(t)
}

func createCrossEnvironment(t *testing.T) {
	time1 := time.Now().UTC()
	time2 := time1.Add(time.Hour * 1)

	utils.SendCreateProgram(
		t,
		true,
		"program1",
		"Program1",
		1,
		3,
		domains.SUBJECT_MATH,
		"description1",
		domains.FEATURED_NONE,
	)
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)
	utils.SendCreateLocationWCHS(t)
	utils.SendCreateClass(
		t,
		true,
		"program1",
		"2020_spring",
		"classA",
		"wchs",
		"2:00pm - 3:00pm",
		0,
		600,
	)
	utils.SendCreateAskForHelp(
		t,
		true,
		time1,
		time2,
		"AP Statistics Help",
		domains.SUBJECT_MATH,
		"wchs",
		"test note 2",
	)
}
