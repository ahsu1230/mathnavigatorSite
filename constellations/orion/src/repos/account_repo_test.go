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

func initAccountTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.AccountRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestAccountRepo(db)
	return db, mock, repo
}

//
// Test Search
//
func Test_SearchAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAccountRows()
	mock.ExpectPrepare(`^SELECT (.+) FROM accounts WHERE (.+)`).
		ExpectQuery().
		WillReturnRows(rows)

	got, err := repo.SelectById(2)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getAccount()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select One
//
func Test_SelectAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAccountRows()
	mock.ExpectPrepare("^SELECT (.+) FROM accounts WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectById(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getAccount()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select One By Primary Email
//
func Test_SelectAccountByPrimaryEmail(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAccountRows()
	mock.ExpectPrepare("^SELECT (.+) FROM accounts WHERE primary_email=?").
		ExpectQuery().
		WithArgs("john_smith@example.com").
		WillReturnRows(rows)
	got, err := repo.SelectByPrimaryEmail("john_smith@example.com")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getAccount()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Create
//
func Test_InsertAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO accounts").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			"john_smith@example.com",
			"password",
		).WillReturnResult(result)
	account := getAccount()
	err := repo.Insert(account)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Update
//
func Test_UpdateAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE accounts SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			"bob_joe@example.com",
			"password2",
			1,
		).WillReturnResult(result)
	account := domains.Account{
		Id:           1,
		CreatedAt:    testUtils.TimeNow,
		UpdatedAt:    testUtils.TimeNow,
		DeletedAt:    sql.NullTime{},
		PrimaryEmail: "bob_joe@example.com",
		Password:     "password2",
	}
	err := repo.Update(1, account)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Delete
//
func Test_DeleteAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM accounts WHERE id=?").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(result)
	err := repo.Delete(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Helper Methods
//
func getAccountRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"PrimaryEmail",
		"Password",
	}).AddRow(
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		"john_smith@example.com",
		"password",
	)
}

func getAccount() domains.Account {
	return domains.Account{
		Id:           1,
		CreatedAt:    testUtils.TimeNow,
		UpdatedAt:    testUtils.TimeNow,
		DeletedAt:    sql.NullTime{},
		PrimaryEmail: "john_smith@example.com",
		Password:     "password",
	}
}