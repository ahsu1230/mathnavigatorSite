package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
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
func TestGetAccountSuccess(t *testing.T) {
	testUtils.AccountRepo.MockSelectById = func(context.Context, uint) (domains.Account, error) {
		account := testUtils.CreateMockAccount(
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

func TestSearchAccountSuccess(t *testing.T) {
	testUtils.AccountRepo.MockSelectByPrimaryEmail = func(context.Context, string) (domains.Account, error) {
		return testUtils.CreateMockAccount(
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

func TestSearchAccountFailure(t *testing.T) {
	testUtils.AccountRepo.MockSelectByPrimaryEmail = func(context.Context, string) (domains.Account, error) {
		return domains.Account{}, appErrors.MockDbNoRowsError()
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/search", nil)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestGetAccountsFailure(t *testing.T) {
	testUtils.AccountRepo.MockSelectById = func(context.Context, uint) (domains.Account, error) {
		return domains.Account{}, appErrors.MockDbNoRowsError()
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Get Negative Balances
//
func TestGetNegativeBalanceAccountsSuccess(t *testing.T) {
	testUtils.AccountRepo.MockSelectAllNegativeBalances = func(context.Context) ([]domains.AccountSum, error) {
		return []domains.AccountSum{
			{
				Account: domains.Account{
					Id:           1,
					PrimaryEmail: "test@gmail.com",
					Password:     "password",
				},
				Balance: -300,
			},
			{
				Account: domains.Account{
					Id:           2,
					PrimaryEmail: "test2@gmail.com",
					Password:     "password2",
				},
				Balance: -200,
			},
		}, nil
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/accounts/unpaid", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var accountSums []domains.AccountSum
	if err := json.Unmarshal(recorder.Body.Bytes(), &accountSums); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, accountSums[0].Account.Id)
	assert.EqualValues(t, "test@gmail.com", accountSums[0].Account.PrimaryEmail)
	assert.EqualValues(t, "password", accountSums[0].Account.Password)
	assert.EqualValues(t, -300, accountSums[0].Balance)
	assert.EqualValues(t, 2, accountSums[1].Account.Id)
	assert.EqualValues(t, "test2@gmail.com", accountSums[1].Account.PrimaryEmail)
	assert.EqualValues(t, "password2", accountSums[1].Account.Password)
	assert.EqualValues(t, -200, accountSums[1].Balance)
}

//
// Test Create
//
func TestCreateAccountWithUserSuccess(t *testing.T) {
	testUtils.AccountRepo.MockInsertWithUser = func(context.Context, domains.Account, domains.User) error {
		return nil
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	accountUser := domains.AccountUser{
		Account: testUtils.CreateMockAccount(
			1,
			"john_smith@example.com",
			"password",
		),
		User: testUtils.CreateMockUser(
			1,
			"John",
			"Smith",
			"",
			"john_smith@example.com",
			"555-555-0199",
			true,
			0,
			"notes1",
		),
	}
	body := createBodyFromAccountUser(accountUser)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateAccountWithUserFailure(t *testing.T) {
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	accountUser := domains.AccountUser{
		Account: testUtils.CreateMockAccount(
			1,
			"john_smith@example.com",
			"password",
		),
		User: testUtils.CreateMockUser(
			1,
			"",
			"",
			"",
			"",
			"",
			false,
			0,
			"",
		),
	}
	body := createBodyFromAccountUser(accountUser)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

//
// Test Update
//
func TestUpdateAccountSuccess(t *testing.T) {
	testUtils.AccountRepo.MockUpdate = func(context.Context, uint, domains.Account) error {
		return nil // Successful update
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	account := testUtils.CreateMockAccount(
		1,
		"john_smith@example.com",
		"password",
	)
	body := createBodyFromAccount(account)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/account/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateAccountInvalid(t *testing.T) {
	// no mock needed
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	account := testUtils.CreateMockAccount(
		1,
		"",
		"",
	)
	body := createBodyFromAccount(account)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/account/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateAccountFailure(t *testing.T) {
	t.Log("Testing failure case, expected error...")
	testUtils.AccountRepo.MockUpdate = func(context.Context, uint, domains.Account) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	account := testUtils.CreateMockAccount(
		1,
		"john_smith@example.com",
		"password",
	)
	body := createBodyFromAccount(account)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/accounts/account/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Test Delete
//
func TestDeleteAccountSuccess(t *testing.T) {
	testUtils.AccountRepo.MockDelete = func(context.Context, uint) error {
		return nil // Return no error, successful delete!
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/accounts/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteAccountFailure(t *testing.T) {
	testUtils.AccountRepo.MockDelete = func(context.Context, uint) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.AccountRepo = &testUtils.AccountRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/accounts/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

//
// Helper Methods
//
func createBodyFromAccount(account domains.Account) io.Reader {
	marshal, err := json.Marshal(&account)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}

func createBodyFromAccountUser(accountUser domains.AccountUser) io.Reader {
	marshal, err := json.Marshal(&accountUser)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
