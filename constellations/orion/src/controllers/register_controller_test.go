package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

type RegisterRepoCount struct {
	numInsertAccount   int
	numInsertUser      int
	numInsertUserClass int
	numInsertUserAfh   int
}

// Test RegisterClass

func TestRegisterClassStudentAndGuardianBothExist(t *testing.T) {
	counter := newRegisterCounter()
	studentEmail := "student@gmail.com"
	guardianEmail := "guardian@yahoo.com"
	mockStudentAndGuardianEmailsExist(studentEmail, true, guardianEmail, true)

	student := createMockStudent(studentEmail)
	guardian := createMockGuardian(guardianEmail)
	body := createBodyForRegister(student, guardian)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/register/class/classA", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	assert.EqualValues(t, 0, counter.numInsertAccount)
	assert.EqualValues(t, 0, counter.numInsertUser)
	assert.EqualValues(t, 1, counter.numInsertUserClass)
	assert.EqualValues(t, 0, counter.numInsertUserAfh)
}

func TestRegisterClassStudentAndGuardianBothDoNotExist(t *testing.T) {
	counter := newRegisterCounter()
	studentEmail := "student@gmail.com"
	guardianEmail := "guardian@yahoo.com"
	mockStudentAndGuardianEmailsExist(studentEmail, false, guardianEmail, false)

	student := createMockStudent(studentEmail)
	guardian := createMockGuardian(guardianEmail)
	body := createBodyForRegister(student, guardian)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/register/class/classA", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	assert.EqualValues(t, 1, counter.numInsertAccount)
	assert.EqualValues(t, 2, counter.numInsertUser)
	assert.EqualValues(t, 1, counter.numInsertUserClass)
	assert.EqualValues(t, 0, counter.numInsertUserAfh)
}

func TestRegisterClassStudentExists(t *testing.T) {
	counter := newRegisterCounter()
	studentEmail := "student@gmail.com"
	guardianEmail := "guardian@yahoo.com"
	mockStudentAndGuardianEmailsExist(studentEmail, true, guardianEmail, false)

	student := createMockStudent(studentEmail)
	guardian := createMockGuardian(guardianEmail)
	body := createBodyForRegister(student, guardian)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/register/class/classA", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	assert.EqualValues(t, 0, counter.numInsertAccount)
	assert.EqualValues(t, 1, counter.numInsertUser)
	assert.EqualValues(t, 1, counter.numInsertUserClass)
	assert.EqualValues(t, 0, counter.numInsertUserAfh)
}

func TestRegisterClassGuardianExists(t *testing.T) {
	counter := newRegisterCounter()
	studentEmail := "student@gmail.com"
	guardianEmail := "guardian@yahoo.com"
	mockStudentAndGuardianEmailsExist(studentEmail, false, guardianEmail, true)

	student := createMockStudent(studentEmail)
	guardian := createMockGuardian(guardianEmail)
	body := createBodyForRegister(student, guardian)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/register/class/classA", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	assert.EqualValues(t, 0, counter.numInsertAccount)
	assert.EqualValues(t, 1, counter.numInsertUser)
	assert.EqualValues(t, 1, counter.numInsertUserClass)
	assert.EqualValues(t, 0, counter.numInsertUserAfh)
}

// Test RegisterAfh

func TestRegisterAfhStudentExists(t *testing.T) {
	counter := newRegisterCounter()
	studentEmail := "student@gmail.com"
	guardianEmail := "guardian@yahoo.com"
	mockStudentAndGuardianEmailsExist(studentEmail, true, guardianEmail, false)

	student := createMockStudent(studentEmail)
	guardian := createMockGuardian(guardianEmail)
	body := createBodyForRegister(student, guardian)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/register/afh/1", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	assert.EqualValues(t, 0, counter.numInsertAccount)
	assert.EqualValues(t, 0, counter.numInsertUser)
	assert.EqualValues(t, 0, counter.numInsertUserClass)
	assert.EqualValues(t, 1, counter.numInsertUserAfh)
}

func TestRegisterAfhStudentDoesNotExist(t *testing.T) {
	counter := newRegisterCounter()
	studentEmail := "student@gmail.com"
	guardianEmail := "guardian@yahoo.com"
	mockStudentAndGuardianEmailsExist(studentEmail, false, guardianEmail, false)

	student := createMockStudent(studentEmail)
	guardian := createMockGuardian(guardianEmail)
	body := createBodyForRegister(student, guardian)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/register/afh/1", body)
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)

	assert.EqualValues(t, 1, counter.numInsertAccount)
	assert.EqualValues(t, 1, counter.numInsertUser)
	assert.EqualValues(t, 0, counter.numInsertUserClass)
	assert.EqualValues(t, 1, counter.numInsertUserAfh)
}

// Helper methods

func mockStudentAndGuardianEmailsExist(studentEmail string, studentExists bool, guardianEmail string, guardianExists bool) {
	testUtils.UserRepo.MockSelectByEmail = func(ctx context.Context, email string) (domains.User, error) {
		if studentExists && email == studentEmail {
			return domains.User{}, nil
		}
		if guardianExists && email == guardianEmail {
			return domains.User{}, nil
		}
		return domains.User{}, appErrors.MockDbNoRowsError()
	}
	repos.UserRepo = &testUtils.UserRepo
}

func mockAccountRepoInsert(counter *RegisterRepoCount) {
	testUtils.AccountRepo.MockInsertWithUser = func(context.Context, domains.Account, domains.User) (uint, error) {
		counter.numInsertAccount++
		counter.numInsertUser++
		return 1, nil
	}
	repos.AccountRepo = &testUtils.AccountRepo
}

func mockUserRepoInsert(counter *RegisterRepoCount) {
	testUtils.UserRepo.MockInsert = func(context.Context, domains.User) (uint, error) {
		counter.numInsertUser++
		return 1, nil
	}
	repos.UserRepo = &testUtils.UserRepo
}

func mockUserClassRepoInsert(counter *RegisterRepoCount) {
	testUtils.UserClassRepo.MockInsert = func(context.Context, domains.UserClass) (uint, error) {
		counter.numInsertUserClass++
		return 1, nil
	}
	repos.UserClassRepo = &testUtils.UserClassRepo
}

func mockUserAfhRepoInsert(counter *RegisterRepoCount) {
	testUtils.UserAfhRepo.MockInsert = func(context.Context, domains.UserAfh) (uint, error) {
		counter.numInsertUserAfh++
		return 1, nil
	}
	repos.UserAfhRepo = &testUtils.UserAfhRepo
}

func newRegisterCounter() *RegisterRepoCount {
	counter := RegisterRepoCount{
		numInsertAccount:   0,
		numInsertUser:      0,
		numInsertUserClass: 0,
		numInsertUserAfh:   0,
	}
	mockAccountRepoInsert(&counter)
	mockUserRepoInsert(&counter)
	mockUserClassRepoInsert(&counter)
	mockUserAfhRepoInsert(&counter)
	return &counter
}

func createMockStudent(email string) domains.User {
	return domains.User{
		FirstName:  "Student",
		MiddleName: domains.NewNullString(""),
		LastName:   "Chang",
		Email:      email,
		IsGuardian: false,
		Phone:      domains.NewNullString(""),
		Notes:      domains.NewNullString(""),
	}
}

func createMockGuardian(email string) domains.User {
	return domains.User{
		FirstName:  "Parent",
		MiddleName: domains.NewNullString(""),
		LastName:   "Chang",
		Email:      email,
		IsGuardian: true,
		Phone:      domains.NewNullString(""),
		Notes:      domains.NewNullString(""),
	}
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
