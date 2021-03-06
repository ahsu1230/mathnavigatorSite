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
	repo := repos.CreateTestAccountRepo(testUtils.Context, db)
	return db, mock, repo
}

//
// Test Search
//
func TestSearchAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAccountRows()
	mock.ExpectPrepare(`^SELECT (.+) FROM accounts WHERE (.+)`).
		ExpectQuery().
		WillReturnRows(rows)

	got, err := repo.SelectById(testUtils.Context, 2)
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
func TestSelectAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAccountRows()
	mock.ExpectPrepare("^SELECT (.+) FROM accounts WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectById(testUtils.Context, 1)
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
func TestSelectAccountByPrimaryEmail(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAccountRows()
	mock.ExpectPrepare("^SELECT (.+) FROM accounts WHERE primary_email=?").
		ExpectQuery().
		WithArgs("john_smith@example.com").
		WillReturnRows(rows)
	got, err := repo.SelectByPrimaryEmail(testUtils.Context, "john_smith@example.com")
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
// Get accounts with negative balances
//
func TestSelectAllNegativeBalances(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := testUtils.TimeNow
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "PrimaryEmail", "Password", "Balance"}).
		AddRow(
			1,
			now,
			now,
			sql.NullTime{},
			"test@gmail.com",
			"password",
			-300,
		).
		AddRow(
			2,
			now,
			now,
			sql.NullTime{},
			"test2@gmail.com",
			"password2",
			-200,
		)
	mock.ExpectPrepare("^SELECT (.*) FROM accounts (.+)").
		ExpectQuery().
		WillReturnRows(rows)

	got, err := repo.SelectAllNegativeBalances(testUtils.Context)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.AccountBalance{
		{
			Account: domains.Account{
				Id:           1,
				CreatedAt:    now,
				UpdatedAt:    now,
				DeletedAt:    sql.NullTime{},
				PrimaryEmail: "test@gmail.com",
				Password:     "password",
			},
			Balance: -300,
		},
		{
			Account: domains.Account{
				Id:           2,
				CreatedAt:    now,
				UpdatedAt:    now,
				DeletedAt:    sql.NullTime{},
				PrimaryEmail: "test2@gmail.com",
				Password:     "password2",
			},
			Balance: -200,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Create User and Account
//
func TestInsertAccountAndUser(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectBegin()
	mock.ExpectPrepare("^INSERT INTO accounts").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			"john_smith@example.com",
			"password",
		).WillReturnResult(result)
	mock.ExpectPrepare("^INSERT INTO users").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			1,
			"John",
			domains.NewNullString(""),
			"Smith",
			"john_smith@example.com",
			"555-555-0100",
			false,
			false,
			domains.NewNullString("schoolone"),
			domains.NewNullUint(2004),
			domains.NewNullString(""),
		).WillReturnResult(result)
	mock.ExpectCommit()
	account := getAccount()
	user := getUser()
	_, err := repo.InsertWithUser(testUtils.Context, account, user)
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
func TestUpdateAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectBegin()
	mock.ExpectPrepare("^UPDATE accounts SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			"bob_joe@example.com",
			"password2",
			1,
		).WillReturnResult(result)
	mock.ExpectCommit()
	account := domains.Account{
		Id:           1,
		CreatedAt:    testUtils.TimeNow,
		UpdatedAt:    testUtils.TimeNow,
		DeletedAt:    sql.NullTime{},
		PrimaryEmail: "bob_joe@example.com",
		Password:     "password2",
	}
	err := repo.Update(testUtils.Context, 1, account)
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
func TestDeleteAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectBegin()
	mock.ExpectPrepare("^DELETE FROM accounts WHERE id=?").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(result)
	mock.ExpectCommit()
	err := repo.Delete(testUtils.Context, 1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFullDeleteAccount(t *testing.T) {
	db, mock, repo := initAccountTest(t)
	defer db.Close()

	// Mock expected user rows for selecting all users in account
	accountId := uint(1)
	userId := uint(1)
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"AccountId",
		"FirstName",
		"MiddleName",
		"LastName",
		"Email",
		"Phone",
		"IsAdminCreated",
		"IsGuardian",
		"School",
		"GraduationYear",
		"Notes",
	}).AddRow(
		userId,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		accountId,
		"John",
		domains.NewNullString(""),
		"Smith",
		"john_smith@example.com",
		domains.NewNullString("555-555-0100"),
		false,
		false,
		domains.NewNullString("schoolone"),
		domains.NewNullUint(2004),
		domains.NewNullString(""),
	)
	result := sqlmock.NewResult(1, 1)

	// Mock DB statements and execute
	mock.ExpectBegin()
	mock.ExpectPrepare("SELECT (.+) FROM users WHERE account_id=?").
		ExpectQuery().
		WithArgs(accountId).
		WillReturnRows(rows)
	mock.ExpectPrepare("^DELETE FROM user_classes WHERE user_id=?").
		ExpectExec().
		WithArgs(userId).
		WillReturnResult(result)
	mock.ExpectPrepare("^DELETE FROM user_afhs WHERE user_id=?").
		ExpectExec().
		WithArgs(userId).
		WillReturnResult(result)
	mock.ExpectPrepare("^DELETE FROM users WHERE id=?").
		ExpectExec().
		WithArgs(userId).
		WillReturnResult(result)
	mock.ExpectPrepare("^DELETE FROM accounts WHERE id=?").
		ExpectExec().
		WithArgs(accountId).
		WillReturnResult(result)
	mock.ExpectCommit()
	err := repo.FullDelete(testUtils.Context, accountId)
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
