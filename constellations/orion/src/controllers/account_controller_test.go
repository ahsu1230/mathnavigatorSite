package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
)

//
// Test Get Account
//
func TestGetAccount_Success(t *testing.T) {
	testUtils.AccountRepo.MockSelectById = func(id uint) (domains.Account, error) {
		account := createMockAccount(
			1,
			"john_smith@example.com",
			"password",
		)
		return account, nil
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var account domains.Account
	if err := json.Unmarshal(recorder.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, account.Id)
	assert.EqualValues(t, "john_smith@example.com", account.PrimaryEmail)
	assert.EqualValues(t, "password", account.Password)
}

func TestSearchAccount_Success(t *testing.T) {
	testUtils.AccountRepo.MockSelectByPrimaryEmail = func(primaryEmail string) (domains.Account, error) {
		return createMockAccount(
			1,
			"john_smith@example.com",
			"password",
		), nil
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create search body for HTTP request
	accountSearchBody := controllers.AccountSearchBody{
		"john_smith@example.com",
	}
	marshal, err := json.Marshal(&accountSearchBody)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(marshal)

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/search", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var accounts domains.Account
	if err := json.Unmarshal(recorder.Body.Bytes(), &accounts); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, accounts.Id)
	assert.EqualValues(t, "john_smith@example.com", accounts.PrimaryEmail)
	assert.EqualValues(t, "password", accounts.Password)
}

func TestSearchAccount_Failure(t *testing.T) {
	testUtils.AccountRepo.MockSelectByPrimaryEmail = func(primaryEmail string) (domains.Account, error) {
		return domains.Account{}, appErrors.TestDbNoRowsError()
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/search", nil)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestGetAccounts_Failure(t *testing.T) {
	testUtils.AccountRepo.MockSelectById = func(id uint) (domains.Account, error) {
		return domains.Account{}, appErrors.TestDbNoRowsError()
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Create
//
func TestCreateAccount_Success(t *testing.T) {
	testUtils.AccountRepo.MockInsert = func(account domains.Account) error {
		return nil
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	account := createMockAccount(
		1,
		"john_smith@example.com",
		"password",
	)
	body := createBodyFromAccount(account)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAccount_Failure(t *testing.T) {
	// no mock needed
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	account := createMockAccount(
		1,
		"",
		"",
	)
	body := createBodyFromAccount(account)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateAccount_Success(t *testing.T) {
	testUtils.AccountRepo.MockUpdate = func(id uint, account domains.Account) error {
		return nil // Successful update
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	account := createMockAccount(
		1,
		"john_smith@example.com",
		"password",
	)
	body := createBodyFromAccount(account)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/account/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAccount_Invalid(t *testing.T) {
	// no mock needed
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	account := createMockAccount(
		1,
		"",
		"",
	)
	body := createBodyFromAccount(account)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/account/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAccount_Failure(t *testing.T) {
	testUtils.AccountRepo.MockUpdate = func(id uint, account domains.Account) error {
		return errors.New("not found")
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	account := createMockAccount(
		1,
		"john_smith@example.com",
		"password",
	)
	body := createBodyFromAccount(account)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/account/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Test Delete
//
func TestDeleteAccount_Success(t *testing.T) {
	testUtils.AccountRepo.MockDelete = func(id uint) error {
		return nil // Return no error, successful delete!
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/accounts/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteAccount_Failure(t *testing.T) {
	testUtils.AccountRepo.MockDelete = func(id uint) error {
		return errors.New("not found")
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/accounts/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
}

//
// Helper Methods
//
func createMockAccount(id uint, primary_email string, password string) domains.Account {
	return domains.Account{
		Id:           id,
		PrimaryEmail: primary_email,
		Password:     password,
	}
}

func createBodyFromAccount(account domains.Account) io.Reader {
	marshal, err := json.Marshal(&account)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
