package repos_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
	"reflect"
	"testing"
	"time"
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
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
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
		"GuardianId",
	}).AddRow(
		1,
		now,
		now,
		sql.NullTime{},
		"John",
		"Smith",
		"Middle",
		"john.smith@example.com",
		"555-555-0100",
		false,
		sql.NullUint,
	)
	mock.ExpectPrepare("^SELECT (.+) FROM users").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.User{
		{
			Id:         1,
			CreatedAt:  now,
			UpdatedAt:  now,
			DeletedAt:  sql.NullTime{},
			FirstName:  "John",
			LastName:   "Smith",
			MiddleName: "Middle",
			Email:      "john.smith@example.com",
			Phone:      "555-555-0100",
			IsGuardian: false,
			GuardianId: 0,
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
// Select One
//
func TestSelectUser(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
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
		"GuardianId",
	}).AddRow(
		1,
		now,
		now,
		sql.NullTime{},
		"John",
		"Smith",
		"Middle",
		"john.smith@example.com",
		"555-555-0100",
		false,
		0,
	)
	mock.ExpectPrepare("^SELECT (.+) FROM users WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectByUserId(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.User{
		Id:         1,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  sql.NullTime{},
		FirstName:  "John",
		LastName:   "Smith",
		MiddleName: "Middle",
		Email:      "john.smith@example.com",
		Phone:      "555-555-0100",
		IsGuardian: false,
		GuardianId: 0,
	}

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
			"Middle",
			"john.smith@example.com",
			"555-555-0100",
			false,
			0,
		).WillReturnResult(result)
	user := domains.User{
		FirstName:  "John",
		LastName:   "Smith",
		MiddleName: "Middle",
		Email:      "john.smith@example.com",
		Phone:      "555-555-0100",
		IsGuardian: false,
		GuardianId: 0,
	}
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
		WithArgs(sqlmock.AnyArg(), "Bob", "Joe", "", "bob.joe@example.com", "555-555-0199", false, 0).
		WillReturnResult(result)
	user := domains.User{
		FirstName:  "Bob",
		LastName:   "Joe",
		MiddleName: "",
		Email:      "bob.joe@example.com",
		Phone:      "555-555-0199",
		IsGuardian: false,
		GuardianId: 0,
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
