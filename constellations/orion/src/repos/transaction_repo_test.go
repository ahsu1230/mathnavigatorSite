package repos_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/testUtils"
)

func initTransactionTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.TransactionRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestTransactionRepo(testUtils.Context, db)
	return db, mock, repo
}

// Test Select By Account Id
func TestSelectByAccountId(t *testing.T) {
	db, mock, repo := initTransactionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getTransactionRows()
	mock.ExpectPrepare("^SELECT (.+) FROM transactions WHERE account_id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectByAccountId(testUtils.Context, 1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate Results
	want := []domains.Transaction{getTransaction()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Test Select By ID
func TestSelectTransaction(t *testing.T) {
	db, mock, repo := initTransactionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getTransactionRows()
	mock.ExpectPrepare("^SELECT (.+) FROM transactions WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectById(testUtils.Context, 1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getTransaction()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Test Insert Transaction
func TestInsertTransaction(t *testing.T) {
	db, mock, repo := initTransactionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO transactions").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			1,
			domains.PAY_PAYPAL,
			100,
			domains.NewNullString("note1"),
		).WillReturnResult(result)
	transaction := domains.Transaction{
		AccountId: 1,
		Type:      domains.PAY_PAYPAL,
		Amount:    100,
		Notes:     domains.NewNullString("note1"),
	}
	_, err := repo.Insert(testUtils.Context, transaction)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Test Update Transaction
func TestUpdateTransaction(t *testing.T) {
	db, mock, repo := initTransactionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE transactions SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			2,
			domains.PAY_PAYPAL,
			101,
			domains.NewNullString("note2"),
			1,
		).WillReturnResult(result)
	transaction := domains.Transaction{
		AccountId: 2,
		Type:      domains.PAY_PAYPAL,
		Amount:    101,
		Notes:     domains.NewNullString("note2"),
	}
	err := repo.Update(testUtils.Context, 1, transaction)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Test Delete Transaction
func TestDeleteTransaction(t *testing.T) {
	db, mock, repo := initTransactionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM transactions WHERE id=?").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(result)

	err := repo.Delete(testUtils.Context, 1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func getTransactionRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"AccountId",
		"Type",
		"Amount",
		"Notes",
	}).
		AddRow(
			1,
			testUtils.TimeNow,
			testUtils.TimeNow,
			domains.NullTime{},
			1,
			domains.PAY_PAYPAL,
			100,
			domains.NewNullString("note1"),
		)
}

func getTransaction() domains.Transaction {
	return domains.Transaction{
		Id:        1,
		CreatedAt: testUtils.TimeNow,
		UpdatedAt: testUtils.TimeNow,
		DeletedAt: domains.NullTime{},
		AccountId: 1,
		Type:      domains.PAY_PAYPAL,
		Amount:    100,
		Notes:     domains.NewNullString("note1"),
	}
}
