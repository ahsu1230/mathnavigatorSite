package tests_integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

func TestE2ERegisterClassStudentAndGuardianDoNotExist(t *testing.T) {
	utils.CreateFullClassAndAfhEnvironment(t)
	// No account/user to begin with

	body := createBodyForRegister(utils.UserMorganStark, utils.UserTonyStark)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/register/class/program1_2020_spring_classA", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	validateAccountUsersAreCorrect(t)
	validateUserClassIsCorrect(t)
	utils.ResetAllTables(t)
}

func TestE2ERegisterClassStudentAndGuardianBothExist(t *testing.T) {
	utils.CreateFullClassAndAfhEnvironment(t)
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark) // userId 1, accountId 1
	utils.SendCreateUser(t, true, utils.UserMorganStark)                              // userId 2, accountId 1

	body := createBodyForRegister(utils.UserMorganStark, utils.UserTonyStark)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/register/class/program1_2020_spring_classA", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	validateAccountUsersAreCorrect(t)
	validateUserClassIsCorrect(t)
	utils.ResetAllTables(t)
}

func TestE2ERegisterClassOnlyStudentExists(t *testing.T) {
	utils.CreateFullClassAndAfhEnvironment(t)

	// Create an account with guardian and student
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark) // userId 1, accountId 1
	utils.SendCreateUser(t, true, utils.UserMorganStark)                              // userId 2, accountId 1

	// Register using an unrecognized guardian
	body := createBodyForRegister(utils.UserMorganStark, utils.UserPepperPotts) // userId 3, accountId 1
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/register/class/program1_2020_spring_classA", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	// Validate account & users are correct (3 users, all with same accountId)
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/users/new", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "tony@stark.com", users[0].Email)
	assert.EqualValues(t, "morgan@stark.com", users[1].Email)
	assert.EqualValues(t, "pepper@stark.com", users[2].Email)
	assert.EqualValues(t, 1, users[0].AccountId)
	assert.EqualValues(t, 1, users[1].AccountId)
	assert.EqualValues(t, 1, users[2].AccountId)
	assert.EqualValues(t, 3, len(users))
	validateUserClassIsCorrect(t)
	utils.ResetAllTables(t)
}

func TestE2ERegisterClassOnlyGuardianExists(t *testing.T) {
	utils.CreateFullClassAndAfhEnvironment(t)
	// Create an account with guardian and student
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark) // userId 1, accountId 1

	// Register using an unrecorgnized student
	body := createBodyForRegister(utils.UserMorganStark, utils.UserTonyStark)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/register/class/program1_2020_spring_classA", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	validateAccountUsersAreCorrect(t)
	validateUserClassIsCorrect(t)
	utils.ResetAllTables(t)
}

func TestE2ERegisterAfhStudentExists(t *testing.T) {
	utils.CreateFullClassAndAfhEnvironment(t)

	// Create an account with guardian and student
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark) // userId 1, accountId 1
	utils.SendCreateUser(t, true, utils.UserMorganStark)                              // userId 2, accountId 1

	// Register using a recognized student
	body := createBodyForRegister(utils.UserMorganStark, domains.User{})
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/register/afh/1", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	validateAccountUsersAreCorrect(t)
	validateUserAfhIsCorrect(t)
	utils.ResetAllTables(t)
}

func TestE2ERegisterAfhStudentDoesNotExist(t *testing.T) {
	utils.CreateFullClassAndAfhEnvironment(t)
	// Create an account with guardian and student
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark) // userId 1, accountId 1

	// Register using an unrecognized student
	body := createBodyForRegister(utils.UserMorganStark, domains.User{}) // userId 2, accountId 2
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/register/afh/1", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	// Validate account & users are correct (3 users, all with same accountId)
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/users/new", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "tony@stark.com", users[0].Email)
	assert.EqualValues(t, "morgan@stark.com", users[1].Email)
	assert.EqualValues(t, 1, users[0].AccountId)
	assert.EqualValues(t, 2, users[1].AccountId)
	assert.EqualValues(t, 2, len(users))

	validateUserAfhIsCorrect(t)
	utils.ResetAllTables(t)
}

func validateAccountUsersAreCorrect(t *testing.T) {
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/users/new", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "tony@stark.com", users[0].Email)
	assert.EqualValues(t, "morgan@stark.com", users[1].Email)
	assert.EqualValues(t, users[0].AccountId, users[1].AccountId)
	assert.EqualValues(t, 2, len(users))
}

func validateUserClassIsCorrect(t *testing.T) {
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userClasses []domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClasses); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userClasses[0].UserId)
	assert.EqualValues(t, 1, userClasses[0].AccountId)
	assert.EqualValues(t, "program1_2020_spring_classA", userClasses[0].ClassId)
	assert.EqualValues(t, 1, len(userClasses))
}

func validateUserAfhIsCorrect(t *testing.T) {
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-afhs/users/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userAfhs []domains.UserAfh
	if err := json.Unmarshal(recorder.Body.Bytes(), &userAfhs); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userAfhs[0].UserId)
	assert.EqualValues(t, 1, userAfhs[0].AfhId)
	assert.EqualValues(t, 1, len(userAfhs))
}

func createBodyForRegister(student domains.User, guardian domains.User) io.Reader {
	fields := domains.RegisterBody{
		Student:  student,
		Guardian: guardian,
	}
	marshal, err := json.Marshal(&fields)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
