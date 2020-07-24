package tests_integration

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create UserClasses and Get by GetClassesByUserId()
func Test_CreateUserClasses(t *testing.T) {
	createAccountUser(t)
	createClasses(t)

	createAllUserClasses(t)

	// Call Get All!
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var userClass []domains.UserClasses
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUserClass(t, 1, userClass[0])
	assertUserClass(t, 2, userClass[1])

	assert.EqualValues(t, 2, len(userClass))

	resetAllTables(t)
}

// Test: Create UserClasses and GetUserByClassId
func Test_GetUsersByClassId(t *testing.T) {
	createAccountUser(t)
	createClasses(t)
	createAllUserClasses(t)

	// Call Get All!
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/class/program1_2020_spring_class1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var userClass []domains.UserClasses
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUserClass(t, 1, userClass[0])
	assertUserClass(t, 3, userClass[1])

	assert.EqualValues(t, 2, len(userClass))

	resetAllTables(t)
}

//Test: Create UserClasses and GetUserByUserAndClass
func Test_GetUserClassByUserAndClass(t *testing.T) {
	createAccountUser(t)
	createClasses(t)
	createAllUserClasses(t)

	// Call Get All!
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/class/program1_2020_spring_class1/user/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var userClass domains.UserClasses
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUserClass(t, 1, userClass)

	resetAllTables(t)
}

// Test: Create 1 Account, 1 User, Update it, GetUserById()
func Test_UpdateUserClass(t *testing.T) {
	createAccountUser(t)
	createClasses(t)

	userClass := createUserClass(1)
	body := utils.CreateJsonBody(&userClass)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/create", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Update
	updatedUserClass := createUserClass(4)
	updatedBody := utils.CreateJsonBody(&updatedUserClass)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/user-class/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/3", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var userClass2 []domains.UserClasses
	if err := json.Unmarshal(recorder3.Body.Bytes(), &userClass2); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUserClass(t, 4, userClass2[0])

	resetAllTables(t)
}

// Test: Create 1 User, Delete it, GetByUserId()
func Test_DeleteUserClass(t *testing.T) {
	createAccountUser(t)
	createClasses(t)

	userClass := createUserClass(1)
	body := utils.CreateJsonBody(&userClass)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/create", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Update
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/user-classes/user-class/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user-class/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	resetAllTables(t)
}

// Helper methods
func createUserClass(id int) domains.UserClasses {
	switch id {
	case 1:
		return domains.UserClasses{
			UserId:    1,
			ClassId:   "program1_2020_spring_class1",
			AccountId: 1,
			State:     1,
		}
	case 2:
		return domains.UserClasses{
			UserId:    1,
			ClassId:   "program1_2020_spring_class2",
			AccountId: 1,
			State:     1,
		}
	case 3:
		return domains.UserClasses{
			UserId:    2,
			ClassId:   "program1_2020_spring_class1",
			AccountId: 2,
			State:     1,
		}
	case 4:
		return domains.UserClasses{
			UserId:    3,
			ClassId:   "program1_2020_spring_class2",
			AccountId: 3,
			State:     1,
		}
	default:
		return domains.UserClasses{}
	}
}

func createAllUserClasses(t *testing.T) {
	for i := 1; i < 5; i++ {
		userClass := createUserClass(i)
		body := utils.CreateJsonBody(&userClass)
		recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/create", body)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
}

func assertUserClass(t *testing.T, id int, userClass domains.UserClasses) {
	switch id {
	case 1:
		assert.EqualValues(t, 1, userClass.UserId)
		assert.EqualValues(t, "program1_2020_spring_class1", userClass.ClassId)
		assert.EqualValues(t, 1, userClass.AccountId)
		assert.EqualValues(t, 1, userClass.State)
	case 2:
		assert.EqualValues(t, 1, userClass.UserId)
		assert.EqualValues(t, "program1_2020_spring_class2", userClass.ClassId)
		assert.EqualValues(t, 1, userClass.AccountId)
		assert.EqualValues(t, 1, userClass.State)
	case 3:
		assert.EqualValues(t, 2, userClass.UserId)
		assert.EqualValues(t, "program1_2020_spring_class1", userClass.ClassId)
		assert.EqualValues(t, 2, userClass.AccountId)
		assert.EqualValues(t, 1, userClass.State)
	case 4:
		assert.EqualValues(t, 3, userClass.UserId)
		assert.EqualValues(t, "program1_2020_spring_class2", userClass.ClassId)
		assert.EqualValues(t, 3, userClass.AccountId)
		assert.EqualValues(t, 1, userClass.State)
	}
}

func createAccountUser(t *testing.T) {
	account := createAccount(1)
	body := utils.CreateJsonBody(&account)
	recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	account2 := createAccount(2)
	bodya2 := utils.CreateJsonBody(&account2)
	recordera2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", bodya2)
	assert.EqualValues(t, http.StatusOK, recordera2.Code)

	account3 := createAccount(3)
	bodya3 := utils.CreateJsonBody(&account3)
	recordera3 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", bodya3)
	assert.EqualValues(t, http.StatusOK, recordera3.Code)

	user := createUser(1)
	body6 := utils.CreateJsonBody(&user)
	recorder6 := utils.SendHttpRequest(t, http.MethodPost, "/api/users/create", body6)
	assert.EqualValues(t, http.StatusOK, recorder6.Code)

	user2 := createUser(2)
	bodyu2 := utils.CreateJsonBody(&user2)
	recorderu2 := utils.SendHttpRequest(t, http.MethodPost, "/api/users/create", bodyu2)
	assert.EqualValues(t, http.StatusOK, recorderu2.Code)

	user3 := createUser(3)
	bodyu3 := utils.CreateJsonBody(&user3)
	recorderu3 := utils.SendHttpRequest(t, http.MethodPost, "/api/users/create", bodyu3)
	assert.EqualValues(t, http.StatusOK, recorderu3.Code)
}

func createClasses(t *testing.T) {
	program := createProgram("program1", "Program1", 1, 3, "description1", 0)
	body2 := utils.CreateJsonBody(&program)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/programs/create", body2)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	location := createLocation("churchill", "11300 Gainsborough Road", "Potomac", "MD", "20854", "Room 100")
	body3 := utils.CreateJsonBody(&location)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/locations/create", body3)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	semester := createSemester("2020_spring", "Spring 2020")
	body4 := utils.CreateJsonBody(&semester)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/semesters/create", body4)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	class := createClass(1)
	body5 := utils.CreateJsonBody(&class)
	recorder5 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", body5)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	class2 := createClass(2)
	bodyc2 := utils.CreateJsonBody(&class2)
	recorderc2 := utils.SendHttpRequest(t, http.MethodPost, "/api/classes/create", bodyc2)
	assert.EqualValues(t, http.StatusOK, recorderc2.Code)
}

func resetAllTables(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_USER_CLASSES)
	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
	utils.ResetTable(t, domains.TABLE_CLASSES)
	utils.ResetTable(t, domains.TABLE_PROGRAMS)
	utils.ResetTable(t, domains.TABLE_SEMESTERS)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}
