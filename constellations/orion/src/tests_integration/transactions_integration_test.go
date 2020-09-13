package tests_integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
)

// Test: Create 3 Transactions and GetAll()
func TestCreateTransactions(t *testing.T) {
	utils.CreateAllTestAccountsAndUsers(t)

	// Create Transactions
	utils.SendCreateTransaction(t, true, 1, domains.PAY_PAYPAL, 100, "notes1")
	utils.SendCreateTransaction(t, true, 2, domains.PAY_CASH, 200, "notes2")
	utils.SendCreateTransaction(t, true, 3, domains.PAY_CHECK, 300, "notes3")
	utils.SendCreateTransaction(t, true, 1, domains.PAY_CHECK, 400, "notes4")

	// Call Get All
	recorder5 := utils.SendHttpRequest(t, http.MethodGet, "/api/transactions/account/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder5.Code)

	// Validate results
	var transactions []domains.Transaction
	if err := json.Unmarshal(recorder5.Body.Bytes(), &transactions); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, transactions[0].Id)
	assert.EqualValues(t, 100, transactions[0].Amount)
	assert.EqualValues(t, domains.PAY_PAYPAL, transactions[0].Type)
	assert.EqualValues(t, "notes1", transactions[0].Notes.String)
	assert.EqualValues(t, 1, transactions[0].AccountId)

	assert.EqualValues(t, 4, transactions[1].Id)
	assert.EqualValues(t, 400, transactions[1].Amount)
	assert.EqualValues(t, domains.PAY_CHECK, transactions[1].Type)
	assert.EqualValues(t, "notes4", transactions[1].Notes.String)
	assert.EqualValues(t, 1, transactions[1].AccountId)

	utils.ResetTable(t, domains.TABLE_TRANSACTIONS)
	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 Transaction, Update, Get By ID
func TestUpdateTransaction(t *testing.T) {
	// Create Account
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)

	// Create 1 Transaction
	utils.SendCreateTransaction(t, true, 1, domains.PAY_PAYPAL, 100, "notes1")

	// Update
	transUpdated := domains.Transaction{
		AccountId: 1,
		Type:      domains.PAY_CASH,
		Amount:    100,
		Notes:     domains.NewNullString("notes2"),
	}

	updatedBody := utils.CreateJsonBody(&transUpdated)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/transactions/transaction/1", updatedBody)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/transactions/transaction/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)

	// Validate results
	var transaction domains.Transaction
	if err := json.Unmarshal(recorder3.Body.Bytes(), &transaction); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, transaction.Id)
	assert.EqualValues(t, 100, transaction.Amount)
	assert.EqualValues(t, domains.PAY_CASH, transaction.Type)
	assert.EqualValues(t, "notes2", transaction.Notes.String)
	assert.EqualValues(t, 1, transaction.AccountId)

	utils.ResetTable(t, domains.TABLE_TRANSACTIONS)
	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}

// Test: Create 1 AFH, Delete it, GetById()
func TestDeleteTransaction(t *testing.T) {
	// Create Account
	utils.SendCreateAccountUser(t, true, utils.AccountTonyStark, utils.UserTonyStark)

	// Create
	utils.SendCreateTransaction(t, true, 1, domains.PAY_PAYPAL, 100, "notes1")

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/transactions/transaction/1", nil)
	assert.EqualValues(t, http.StatusNoContent, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/transactions/transaction/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_TRANSACTIONS)
	utils.ResetTable(t, domains.TABLE_USERS)
	utils.ResetTable(t, domains.TABLE_ACCOUNTS)
}
