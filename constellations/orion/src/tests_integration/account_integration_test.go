package tests_integration

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Create account with user
func TestCreateAccountAndUserSuccess(t *testing.T) {
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)

	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var account domains.Account
	if err := json.Unmarshal(recorder.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assertAccounts(t, utils.AccountTonyStark, account)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Create 1 account then get by id
func TestSearchAccountById(t *testing.T) {
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)

	// Retrieve account (by id)
	recorder2 := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	var account domains.Account
	if err := json.Unmarshal(recorder2.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertAccounts(t, utils.AccountTonyStark, account)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 3 Accounts and search by pagination
func TestSearchAccountByPrimaryEmail(t *testing.T) {
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)

	// Retrieve account (by email)
	body := strings.NewReader(`{
		"primaryEmail": "tony@stark.com"
	}`)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/search", body)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	var searchedAccount domains.Account
	if err := json.Unmarshal(recorder2.Body.Bytes(), &searchedAccount); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertAccounts(t, utils.AccountTonyStark, searchedAccount)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 3 Accounts and check each Account
func TestGetAccountById(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)
	var account domains.Account

	// Get Account1 by Id and validate
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	if err := json.Unmarshal(recorder.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertAccounts(t, utils.AccountTonyStark, account)

	// Get Account2 by Id and validate
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/2", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	if err := json.Unmarshal(recorder.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertAccounts(t, utils.AccountNatasha, account)

	// Get Account3 by Id and validate
	recorder = utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/3", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	if err := json.Unmarshal(recorder.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertAccounts(t, utils.AccountPotter, account)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 2 accounts with the same email
func TestCreateSameEmailAccountsFailure(t *testing.T) {
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)

	// Create account with same email
	accountUser2 := domains.AccountUser{
		Account: utils.AccountTonyStark, // account email already being used!
		User:    domains.User{},
	}
	body2 := utils.CreateJsonBody(&accountUser2)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body2)
	assert.EqualValues(t, http.StatusBadRequest, recorder2.Code)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 3 Accounts, GetNegativeBalances()
func TestGetNegativeBalanceAccounts(t *testing.T) {
	// Create 3 accounts and 5 transactions, accounts 1 and 2 are negative
	utils.CreateAllTestAccountsAndUsers(t)
	utils.SendCreateTransaction(t, true, 1, domains.PAY_PAYPAL, 100, "notes1")
	utils.SendCreateTransaction(t, true, 2, domains.PAY_CASH, 200, "notes2")
	utils.SendCreateTransaction(t, true, 3, domains.PAY_CHECK, 300, "notes3")
	utils.SendCreateTransaction(t, true, 1, domains.CHARGE, -400, "notes4")
	utils.SendCreateTransaction(t, true, 2, domains.CHARGE, -300, "")

	// Call endpoint
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/unpaid", nil)
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	// Validate results
	var accountBalances []domains.AccountBalance
	if err := json.Unmarshal(recorder.Body.Bytes(), &accountBalances); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, accountBalances[0].Account.Id)
	assert.EqualValues(t, "tony@stark.com", accountBalances[0].Account.PrimaryEmail)
	assert.EqualValues(t, -300, accountBalances[0].Balance)
	assert.EqualValues(t, 2, accountBalances[1].Account.Id)
	assert.EqualValues(t, "natasha@shield.com", accountBalances[1].Account.PrimaryEmail)
	assert.EqualValues(t, -100, accountBalances[1].Balance)
	assert.EqualValues(t, 2, len(accountBalances))

	utils.ResetTable(t, domains.TABLE_TRANSACTIONS)
	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 Account, Update it, GetAccountById()
func TestUpdateAccount(t *testing.T) {
	// Create 1 Account
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)

	// Update
	updatedAccount := domains.Account{
		PrimaryEmail: "ultron@tonysux.com",
		Password:     "password13",
	}
	updatedBody := utils.CreateJsonBody(&updatedAccount)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/account/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var account domains.Account
	if err := json.Unmarshal(recorder3.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertAccounts(t, updatedAccount, account)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 Account, Delete it, GetByAccountId()
func TestDeleteAccount(t *testing.T) {
	// Create
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)

	// Delete
	recorder4 := utils.SendHttpRequest(t, http.MethodDelete, "/api/users/user/1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder4.Code)
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Helper methods
func assertAccounts(
	t *testing.T,
	expectedAccount domains.Account,
	actualAccount domains.Account,
) {
	assert.EqualValues(t, expectedAccount.PrimaryEmail, actualAccount.PrimaryEmail)
	assert.EqualValues(t, expectedAccount.Password, actualAccount.Password)
}
