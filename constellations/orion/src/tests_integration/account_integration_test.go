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

//create 1 account then get by id
func Test_SearchAccountById(t *testing.T) {
	account1 := createAccount(1)
	body1 := utils.CreateJsonBody(&account1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	recorder2 := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	var account domains.Account
	if err := json.Unmarshal(recorder2.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertAccount(t, 1, account)

	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 3 Accounts and search by pagination
func Test_SearchAccountByPrimaryEmail(t *testing.T) {
	account1 := createAccount(1)
	body1 := utils.CreateJsonBody(&account1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	body := strings.NewReader(`{
		"primaryEmail": "john_smith@example.com"
	}`)

	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/search", body)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	var account domains.Account
	if err := json.Unmarshal(recorder2.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assertAccount(t, 1, account)

	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 3 Accounts and GetAccountById
func Test_GetAccountById(t *testing.T) {
	account1 := createAccount(1)
	body1 := utils.CreateJsonBody(&account1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Call Get All!
	recorder := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var account domains.Account
	if err := json.Unmarshal(recorder.Body.Bytes(), &account); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "john_smith@example.com", account.PrimaryEmail)
	assert.EqualValues(t, "password1", account.Password)

	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 Account, Update it, GetAccountById()
func Test_UpdateAccount(t *testing.T) {
	// Create 1 Account
	account1 := createAccount(1)
	body1 := utils.CreateJsonBody(&account1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	updatedAccount := createAccount(2)
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
	assertAccount(t, 2, account)

	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 Account, Delete it, GetByAccountId()
func Test_DeleteAccount(t *testing.T) {
	// Create
	account1 := createAccount(1)
	body1 := utils.CreateJsonBody(&account1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/accounts/account/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Helper methods
func createAccount(id int) domains.Account {
	switch id {
	case 1:
		return domains.Account{
			Id:           1,
			PrimaryEmail: "john_smith@example.com",
			Password:     "password1",
		}
	case 2:
		return domains.Account{
			Id:           2,
			PrimaryEmail: "bob_smith@example.com",
			Password:     "password2",
		}
	case 3:
		return domains.Account{
			Id:           3,
			PrimaryEmail: "foobar@example.com",
			Password:     "password3",
		}
	default:
		return domains.Account{}
	}
}

func createAllAccounts(t *testing.T) {
	for i := 1; i < 4; i++ {
		account := createAccount(i)
		body := utils.CreateJsonBody(&account)
		recorder := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
	}
}

func assertAccount(t *testing.T, id int, account domains.Account) {
	switch id {
	case 1:
		assert.EqualValues(t, "john_smith@example.com", account.PrimaryEmail)
		assert.EqualValues(t, "password1", account.Password)
	case 2:
		assert.EqualValues(t, "bob_smith@example.com", account.PrimaryEmail)
		assert.EqualValues(t, "password2", account.Password)
	case 3:
		assert.EqualValues(t, "foobar@example.com", account.PrimaryEmail)
		assert.EqualValues(t, "password2", account.Password)
	}
}
