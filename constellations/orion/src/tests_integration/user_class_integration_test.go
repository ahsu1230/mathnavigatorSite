package tests_integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create UserClass and Get by GetClassesByUserId()
func TestCreateUserClass(t *testing.T) {
	createAllUserClasses(t)
	targetUserId := uint(3)
	targetAccountId := uint(1)
	apiGetUrl := fmt.Sprintf("/api/user-classes/user/%d", targetUserId)

	// Call GetClassesByUserId!
	recorder := utils.SendHttpRequest(t, http.MethodGet, apiGetUrl, nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var userClasses []domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClasses); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUserClasses(t, userClasses[0], domains.UserClass{
		ClassId:   "program1_2020_spring_classA",
		UserId:    targetUserId,
		AccountId: targetAccountId,
	})
	assertUserClasses(t, userClasses[1], domains.UserClass{
		ClassId:   "program1_2020_spring_classB",
		UserId:    targetUserId,
		AccountId: targetAccountId,
	})
	assert.EqualValues(t, 2, len(userClasses))

	resetUserClassTables(t)
}

// Test: Create UserClass and GetUserByClassId
func TestGetUsersByClassId(t *testing.T) {
	createAllUserClasses(t)
	targetClassId := "program1_2020_spring_classA"
	apiGetUrl := fmt.Sprintf("/api/user-classes/class/%s", targetClassId)

	// Call GetUserByClassId
	recorder := utils.SendHttpRequest(t, http.MethodGet, apiGetUrl, nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var userClasses []domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClasses); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUserClasses(t, userClasses[0], domains.UserClass{
		ClassId:   targetClassId,
		UserId:    3,
		AccountId: 1,
	})
	assertUserClasses(t, userClasses[1], domains.UserClass{
		ClassId:   targetClassId,
		UserId:    5,
		AccountId: 3,
	})
	assert.EqualValues(t, 2, len(userClasses))

	resetUserClassTables(t)
}

//Test: Create UserClass and GetUserByUserAndClass
func TestGetUserClassByUserAndClass(t *testing.T) {
	createAllUserClasses(t)
	targetClassId := "program1_2020_spring_classA"
	targetUserId := uint(3)
	targetAccountId := uint(1)
	apiGetUrl := fmt.Sprintf("/api/user-classes/class/%s/user/%d", targetClassId, targetUserId)

	// Call Get GetUserByUserAndClass
	recorder := utils.SendHttpRequest(t, http.MethodGet, apiGetUrl, nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var userClass domains.UserClass
	if err := json.Unmarshal(recorder.Body.Bytes(), &userClass); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUserClasses(t, userClass, domains.UserClass{
		ClassId:   targetClassId,
		UserId:    targetUserId,
		AccountId: targetAccountId,
	})

	resetUserClassTables(t)
}

// Test: Update UserClass, GetByUserId()
func TestUpdateUserClass(t *testing.T) {
	createAllUserClasses(t)

	// Update (Harry Potter will now attend classB as a trial)
	updatedUserClass := domains.UserClass{
		Id:        uint(3),
		ClassId:   "program1_2020_spring_classB", // Changed
		UserId:    uint(5),
		AccountId: uint(3),
		State:     domains.USER_CLASS_TRIAL, // Changed
	}
	updatedBody := utils.CreateJsonBody(&updatedUserClass)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/user-classes/user-class/3", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/user-classes/user/5", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var userClasses []domains.UserClass
	if err := json.Unmarshal(recorder3.Body.Bytes(), &userClasses); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertUserClasses(t, updatedUserClass, userClasses[0])

	resetUserClassTables(t)
}

// Test: Create 1 User, Delete it, GetByUserId()
func TestDeleteUserClass(t *testing.T) {
	createAllUserClasses(t)

	// Get
	apiGetUrl := "/api/user-classes/class/program1_2020_spring_classA/user/5"
	recorder1 := utils.SendHttpRequest(t, http.MethodGet, apiGetUrl, nil)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/user-classes/user-class/3", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, apiGetUrl, nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	resetUserClassTables(t)
}

// Helper methods
func createAllUserClasses(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	studentPP := utils.UserPeterParker // Peter Parker (userId 3, accountId 1)
	studentPP.Id = uint(3)
	studentHP := utils.UserPotter // Harry Potter (userId 5, accountId 3)
	studentHP.Id = uint(5)

	// Create Classes
	createClasses(t)
	classId1 := "program1_2020_spring_classA"
	classId2 := "program1_2020_spring_classB"

	// Classes attended by student1
	utils.SendCreateUserClass(t, true, classId1, studentPP.Id, studentPP.AccountId) // userClassId 1
	utils.SendCreateUserClass(t, true, classId2, studentPP.Id, studentPP.AccountId) // userClassId 2

	// Classes attended by student2
	utils.SendCreateUserClass(t, true, classId1, studentHP.Id, studentHP.AccountId) // userClassId 3
}

func assertUserClasses(t *testing.T, expectedUserClass domains.UserClass, userClass domains.UserClass) {
	assert.EqualValues(t, expectedUserClass.UserId, userClass.UserId)
	assert.EqualValues(t, expectedUserClass.ClassId, userClass.ClassId)
	assert.EqualValues(t, expectedUserClass.AccountId, userClass.AccountId)
	assert.EqualValues(t, expectedUserClass.State, userClass.State)
}

func createClasses(t *testing.T) {
	utils.SendCreateProgram(t, true, "program1", "Program1", 1, 3, domains.SUBJECT_MATH, "description1", domains.FEATURED_NONE)
	utils.SendCreateLocationWCHS(t)
	utils.SendCreateSemester(t, true, domains.SPRING, 2020)
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
	utils.SendCreateClass(
		t,
		true,
		"program1",
		"2020_spring",
		"classB",
		"wchs",
		"2:00pm - 3:00pm",
		0,
		600,
	)
}

func resetUserClassTables(t *testing.T) {
	utils.ResetTable(t, domains.TABLE_USER_CLASSES)
	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
	utils.ResetTable(t, domains.TABLE_CLASSES)
	utils.ResetTable(t, domains.TABLE_PROGRAMS)
	utils.ResetTable(t, domains.TABLE_SEMESTERS)
	utils.ResetTable(t, domains.TABLE_LOCATIONS)
}
