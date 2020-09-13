package repos_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/testUtils"
	"reflect"
	"testing"
)

func initUserTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.UserRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestUserRepo(testUtils.Context, db)
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
	got, err := repo.SelectAll(testUtils.Context, "", 100, 0)
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
	got, err := repo.SelectAll(testUtils.Context, "Smith", 2, 0)
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
	got, err := repo.SelectById(testUtils.Context, 1)
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
	got, err := repo.SelectByAccountId(testUtils.Context, 2)
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
// Select New
//
func TestSelectNewUsers(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserRows()
	mock.ExpectPrepare("^SELECT (.+) FROM users WHERE created_at>=*").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectByNew(testUtils.Context)
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
			2,
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
	user := getUser()
	_, err := repo.Insert(testUtils.Context, user)
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
			1,
			"Bob",
			domains.NewNullString("Oliver"),
			"Joe",
			"bob_joe@example.com",
			"555-555-0199",
			false,
			true,
			domains.NewNullString("schoolone"),
			domains.NewNullUint(2004),
			domains.NewNullString("notes"),
			1,
		).WillReturnResult(result)
	user := domains.User{
		Id:             1,
		CreatedAt:      testUtils.TimeNow,
		UpdatedAt:      testUtils.TimeNow,
		DeletedAt:      sql.NullTime{},
		AccountId:      1,
		FirstName:      "Bob",
		MiddleName:     domains.NewNullString("Oliver"),
		LastName:       "Joe",
		Email:          "bob_joe@example.com",
		Phone:          "555-555-0199",
		IsAdminCreated: false,
		IsGuardian:     true,
		School:         domains.NewNullString("schoolone"),
		GraduationYear: domains.NewNullUint(2004),
		Notes:          domains.NewNullString("notes"),
	}
	err := repo.Update(testUtils.Context, 1, user)
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
	err := repo.Delete(testUtils.Context, 1)
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
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		2,
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
	)
}

func getUser() domains.User {
	return domains.User{
		Id:             1,
		CreatedAt:      testUtils.TimeNow,
		UpdatedAt:      testUtils.TimeNow,
		DeletedAt:      sql.NullTime{},
		AccountId:      2,
		FirstName:      "John",
		MiddleName:     domains.NewNullString(""),
		LastName:       "Smith",
		Email:          "john_smith@example.com",
		Phone:          "555-555-0100",
		IsAdminCreated: false,
		IsGuardian:     false,
		School:         domains.NewNullString("schoolone"),
		GraduationYear: domains.NewNullUint(2004),
		Notes:          domains.NewNullString(""),
	}
}
