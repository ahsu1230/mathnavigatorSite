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
	// utils.SendCreateAccountUser(t, true, AccountTonyStark, UserTonyStark)	// userId 1, accountId 1
	// utils.SendCreateUser(t, true, UserMorganStark)							// userId 2, accountId 1

	body := createBodyForRegister(utils.UserMorganStark, utils.UserTonyStark)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/register/class/program1_2020_spring_classA", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Validate User & Account info is correct
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/users/new", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var users []domains.User
	if err := json.Unmarshal(recorder.Body.Bytes(), &users); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "tony@stark.com", users[0].Email)
	assert.EqualValues(t, "morgan@stark.com", users[1].Email)
	assert.EqualValues(t, users[0].AccountId, users[1].AccountId)
	assert.EqualValues(t, 2, len(users))

	// Validate UserClass info is correct
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var userClasses []domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClasses); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 2, userClasses[0].UserId)
	assert.EqualValues(t, 1, userClasses[0].AccountId)
	assert.EqualValues(t, "program1_2020_spring_classA", userClasses[0].ClassId)
	assert.EqualValues(t, 2, len(userClasses))

	utils.ResetAllTables(t)
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
