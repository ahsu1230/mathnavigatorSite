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

func initUserTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.UserRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestUserRepo(db)
	return db, mock, repo
}

//
// Test Select All
//
func TestSelectAllUsers(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserRows()
	mock.ExpectPrepare("^SELECT (.+) FROM users").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll("", 100, 0)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.User{getUser()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Test Search
//
func TestSearchUsers(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserRows()
	mock.ExpectPrepare(`^SELECT (.+) FROM users WHERE (.+) LIMIT (.+) OFFSET (.+)`).
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll("Smith", 2, 0)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.User{getUser()}
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
func TestSelectUser(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserRows()
	mock.ExpectPrepare("^SELECT (.+) FROM users WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectById(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getUser()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select Many By Account ID
//
func TestSelectUsersByAccountId(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserRows()
	mock.ExpectPrepare("^SELECT (.+) FROM users WHERE account_id=?").
		ExpectQuery().
		WithArgs(2).
		WillReturnRows(rows)
	got, err := repo.SelectByAccountId(2)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.User{getUser()}

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
func TestInsertUser(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO users").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			"John",
			"Smith",
			domains.NewNullString(""),
			"john_smith@example.com",
			"555-555-0100",
			false,
			2,
			domains.NewNullString(""),
		).WillReturnResult(result)
	user := getUser()
	err := repo.Insert(user)
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
func TestUpdateUser(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE users SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			"Bob",
			"Joe",
			domains.NewNullString("Oliver"),
			"bob_joe@example.com",
			"555-555-0199",
			true,
			0,
			domains.NewNullString("notes"),
			1,
		).WillReturnResult(result)
	user := domains.User{
		Id:         1,
		CreatedAt:  testUtils.TimeNow,
		UpdatedAt:  testUtils.TimeNow,
		DeletedAt:  sql.NullTime{},
		FirstName:  "Bob",
		LastName:   "Joe",
		MiddleName: domains.NewNullString("Oliver"),
		Email:      "bob_joe@example.com",
		Phone:      "555-555-0199",
		IsGuardian: true,
		AccountId:   0,
		Notes:      domains.NewNullString("notes"),
	}
	err := repo.Update(1, user)
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
func TestDeleteUser(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM users WHERE id=?").
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
func getUserRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"FirstName",
		"LastName",
		"MiddleName",
		"Email",
		"Phone",
		"IsGuardian",
		"AccountId",
		"Notes",
	}).AddRow(
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		"John",
		"Smith",
		domains.NewNullString(""),
		"john_smith@example.com",
		"555-555-0100",
		false,
		2,
		domains.NewNullString(""),
	)
}

func getUser() domains.User {
	return domains.User{
		Id:         1,
		CreatedAt:  testUtils.TimeNow,
		UpdatedAt:  testUtils.TimeNow,
		DeletedAt:  sql.NullTime{},
		FirstName:  "John",
		LastName:   "Smith",
		MiddleName: domains.NewNullString(""),
		Email:      "john_smith@example.com",
		Phone:      "555-555-0100",
		IsGuardian: false,
		AccountId:   2,
		Notes:      domains.NewNullString(""),
	}
}
