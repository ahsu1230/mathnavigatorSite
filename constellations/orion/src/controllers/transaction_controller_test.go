package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/testUtils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

// Test Get All
func TestGetTransactionsByAccountIdSuccess(t *testing.T) {
	testUtils.TransactionRepo.MockSelectByAccountId = func(context.Context, uint) ([]domains.Transaction, error) {
		return []domains.Transaction{
			testUtils.CreateMockTransaction(
				1,
				100,
				domains.PAY_PAYPAL,
				"notes1",
				1,
			),
			testUtils.CreateMockTransaction(
				2,
				200,
				domains.PAY_PAYPAL,
				"notes2",
				1,
			),
		}, nil
	}
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/transactions/account/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var transactions []domains.Transaction
	if err := json.Unmarshal(recorder.Body.Bytes(), &transactions); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	assert.EqualValues(t, 1, transactions[0].Id)
	assert.EqualValues(t, 100, transactions[0].Amount)
	assert.EqualValues(t, domains.PAY_PAYPAL, transactions[0].PaymentType)
	assert.EqualValues(t, "notes1", transactions[0].PaymentNotes.String)
	assert.EqualValues(t, 1, transactions[0].AccountId)

	assert.EqualValues(t, 2, transactions[1].Id)
	assert.EqualValues(t, 200, transactions[1].Amount)
	assert.EqualValues(t, domains.PAY_PAYPAL, transactions[1].PaymentType)
	assert.EqualValues(t, "notes2", transactions[1].PaymentNotes.String)
	assert.EqualValues(t, 1, transactions[1].AccountId)

}

// Test Get Transaction
func TestGetTransactionSuccess(t *testing.T) {
	testUtils.TransactionRepo.MockSelectById = func(context.Context, uint) (domains.Transaction, error) {
		transaction := testUtils.CreateMockTransaction(
			1,
			100,
			domains.PAY_PAYPAL,
			"notes1",
			1)
		return transaction, nil
	}
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/transactions/transaction/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	var transaction domains.Transaction
	if err := json.Unmarshal(recorder.Body.Bytes(), &transaction); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, 1, transaction.Id)
	assert.EqualValues(t, 100, transaction.Amount)
	assert.EqualValues(t, domains.PAY_PAYPAL, transaction.PaymentType)
	assert.EqualValues(t, "notes1", transaction.PaymentNotes.String)
	assert.EqualValues(t, 1, transaction.AccountId)
}

func TestGetTransactionFailure(t *testing.T) {
	testUtils.TransactionRepo.MockSelectById = func(context.Context, uint) (domains.Transaction, error) {
		return domains.Transaction{}, appErrors.MockDbNoRowsError()
	}
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/transactions/transaction/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Create
func TestCreateTransactionSuccess(t *testing.T) {
	testUtils.TransactionRepo.MockUpdate = func(context.Context, uint, domains.Transaction) error {
		return nil
	}
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	transaction := testUtils.CreateMockTransaction(
		1,
		100,
		domains.PAY_PAYPAL,
		"notes1",
		1)
	body := createBodyFromTransaction(transaction)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/transactions/transaction/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestCreateTransactionFailure(t *testing.T) {
	// no mock needed
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	transaction := testUtils.CreateMockTransaction(
		1,
		100,
		"",
		"notes1",
		1)
	body := createBodyFromTransaction(transaction)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/transactions/create", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

// Test Update
func TestUpdateTransactionSuccess(t *testing.T) {
	testUtils.TransactionRepo.MockUpdate = func(context.Context, uint, domains.Transaction) error {
		return nil // Successful update
	}
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	transaction := testUtils.CreateMockTransaction(
		1,
		100,
		domains.PAY_PAYPAL,
		"notes1",
		1)
	body := createBodyFromTransaction(transaction)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/transactions/transaction/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestUpdateTransactionInvalid(t *testing.T) {
	// no mock needed
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	transaction := testUtils.CreateMockTransaction(
		1,
		100,
		"",
		"notes1",
		1)
	body := createBodyFromTransaction(transaction)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/transactions/transaction/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
}

func TestUpdateTransactionFailure(t *testing.T) {
	testUtils.TransactionRepo.MockUpdate = func(context.Context, uint, domains.Transaction) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	transaction := testUtils.CreateMockTransaction(
		1,
		100,
		domains.PAY_PAYPAL,
		"notes1",
		1)
	body := createBodyFromTransaction(transaction)
	recorder := testUtils.SendHttpRequest(t, http.MethodPost, "/api/transactions/transaction/1", body)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

// Test Delete
func TestDeleteTransactionSuccess(t *testing.T) {
	testUtils.TransactionRepo.MockDelete = func(context.Context, uint) error {
		return nil // Return no error, successful delete!
	}
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/transactions/transaction/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNoContent, recorder.Code)
}

func TestDeleteTransactionFailure(t *testing.T) {
	testUtils.TransactionRepo.MockDelete = func(context.Context, uint) error {
		return appErrors.MockDbNoRowsError()
	}
	repos.TransactionRepo = &testUtils.TransactionRepo

	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodDelete, "/api/transactions/transaction/1", nil)

	// Validate results
	assert.EqualValues(t, http.StatusNotFound, recorder.Code)
}

func TestGetAllPaymentTypes(t *testing.T) {
	// Create new HTTP request to endpoint
	recorder := testUtils.SendHttpRequest(t, http.MethodGet, "/api/transactions/types", nil)

	//Validate results
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var paymentTypes []string
	if err := json.Unmarshal(recorder.Body.Bytes(), &paymentTypes); err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}
	assert.EqualValues(t, "pay_paypal", paymentTypes[0])
	assert.EqualValues(t, "pay_cash", paymentTypes[1])
	assert.EqualValues(t, "pay_check", paymentTypes[2])
	assert.EqualValues(t, "charge", paymentTypes[3])
	assert.EqualValues(t, "refund", paymentTypes[4])

}

// Helper Methods
func createBodyFromTransaction(transaction domains.Transaction) io.Reader {
	marshal, err := json.Marshal(&transaction)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(marshal)
}
