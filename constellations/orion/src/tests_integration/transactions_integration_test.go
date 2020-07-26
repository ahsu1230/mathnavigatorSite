package tests_integration

import (
	"encoding/json"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/tests_integration/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// Test: Create 3 Transactions and GetAll()
func Test_CreateTransactions(t *testing.T) {
	createAccounts(t)

	trans1 := createTransaction(1, 100, domains.PAY_PAYPAL, "notes1", 1)
	trans2 := createTransaction(2, 200, domains.PAY_CASH, "notes2", 2)
	trans3 := createTransaction(3, 300, domains.PAY_CHECK, "notes3", 3)
	trans4 := createTransaction(4, 400, domains.PAY_CHECK, "notes4", 1)

	body1 := utils.CreateJsonBody(&trans1)
	body2 := utils.CreateJsonBody(&trans2)
	body3 := utils.CreateJsonBody(&trans3)
	body4 := utils.CreateJsonBody(&trans4)

	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/transactions/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/transactions/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/transactions/create", body3)
	recorder4 := utils.SendHttpRequest(t, http.MethodPost, "/api/transactions/create", body4)

	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
	assert.EqualValues(t, http.StatusOK, recorder4.Code)

	// Call Get All
	recorder5 := utils.SendHttpRequest(t, http.MethodGet, "/api/transactions/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder4.Code)
	var transactions []domains.Transaction
	if err := json.Unmarshal(recorder5.Body.Bytes(), &transactions); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, transactions[0].Id)
	assert.EqualValues(t, 100, transactions[0].Amount)
	assert.EqualValues(t, domains.PAY_PAYPAL, transactions[0].PaymentType)
	assert.EqualValues(t, "notes1", transactions[0].PaymentNotes.String)
	assert.EqualValues(t, 1, transactions[0].AccountId)

	assert.EqualValues(t, 4, transactions[1].Id)
	assert.EqualValues(t, 400, transactions[1].Amount)
	assert.EqualValues(t, domains.PAY_CHECK, transactions[1].PaymentType)
	assert.EqualValues(t, "notes4", transactions[1].PaymentNotes.String)
	assert.EqualValues(t, 1, transactions[1].AccountId)

	utils.ResetTable(t, domains.TABLE_TRANSACTIONS)
}

// Test: Create 1 Transaction, Update, Get By ID
func Test_UpdateTransaction(t *testing.T) {

	// Create 1 Transaction
	trans1 := createTransaction(1, 100, domains.PAY_PAYPAL, "notes1", 1)
	body1 := utils.CreateJsonBody(&trans1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/transactions/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Update
	transUpdated := createTransaction(1, 100, domains.PAY_CASH, "notes1", 1)

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
	assert.EqualValues(t, domains.PAY_CASH, transaction.PaymentType)
	assert.EqualValues(t, "notes1", transaction.PaymentNotes.String)
	assert.EqualValues(t, 1, transaction.AccountId)

	utils.ResetTable(t, domains.TABLE_TRANSACTIONS)
}

// Test: Create 1 AFH, Delete it, GetById()
func Test_DeleteTransaction(t *testing.T) {

	// Create
	trans1 := createTransaction(1, 100, domains.PAY_PAYPAL, "notes1", 1)
	body1 := utils.CreateJsonBody(&trans1)
	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/transactions/create", body1)
	assert.EqualValues(t, http.StatusOK, recorder1.Code)

	// Delete
	recorder2 := utils.SendHttpRequest(t, http.MethodDelete, "/api/transactions/transaction/1", nil)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)

	// Get
	recorder3 := utils.SendHttpRequest(t, http.MethodGet, "/api/transactions/transaction/1", nil)
	assert.EqualValues(t, http.StatusNotFound, recorder3.Code)

	utils.ResetTable(t, domains.TABLE_TRANSACTIONS)
}

// Helper methods
func createAccounts(t *testing.T) {
	acc1 := createAccount(1)
	acc2 := createAccount(2)
	acc3 := createAccount(3)

	body1 := utils.CreateJsonBody(&acc1)
	body2 := utils.CreateJsonBody(&acc2)
	body3 := utils.CreateJsonBody(&acc3)

	recorder1 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body1)
	recorder2 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body2)
	recorder3 := utils.SendHttpRequest(t, http.MethodPost, "/api/accounts/create", body3)

	assert.EqualValues(t, http.StatusOK, recorder1.Code)
	assert.EqualValues(t, http.StatusOK, recorder2.Code)
	assert.EqualValues(t, http.StatusOK, recorder3.Code)
}
func createTransaction(id uint, amount int, paymentType string, paymentNotes string, accountId uint) domains.Transaction {
	return domains.Transaction{
		Id:           id,
		Amount:       amount,
		PaymentType:  paymentType,
		PaymentNotes: domains.NewNullString(paymentNotes),
		AccountId:    accountId,
	}
}
