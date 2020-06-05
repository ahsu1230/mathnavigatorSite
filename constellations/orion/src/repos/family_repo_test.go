package repos_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
)

func initFamilyTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.FamilyRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestFamilyRepo(db)
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
	mock.ExpectPrepare("^SELECT (.+) FROM family").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll("", 100, 0)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Family{getFamily()}
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
func TestSearchFamily(t *testing.T) {
	db, mock, repo := initFamilyTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getFamilyRows()
	mock.ExpectPrepare(`^SELECT (.+) FROM family WHERE (.+) LIMIT (.+) OFFSET (.+)`).
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll("Smith", 2, 0)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Family{getUser()}
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
func TestSelectFamily(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getFamilyRows()
	mock.ExpectPrepare("^SELECT (.+) FROM family WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectById(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getFamily()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select One By Guardian ID
//
func TestSelectFamilyByPrimaryEmail(t *testing.T) {
	db, mock, repo := initUserTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getFamilyRows()
	mock.ExpectPrepare("^SELECT (.+) FROM people WHERE email=?").
		ExpectQuery().
		WithArgs(2).
		WillReturnRows(rows)
	got, err := repo.SelectByGuardianId(2)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.Family{getFamily()}

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
func TestInsertFamily(t *testing.T) {
	db, mock, repo := initFamilyTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO family").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			"john_smith@example.com",
			"password",
		).WillReturnResult(result)
	family := getFamily()
	err := repo.Insert(family)
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
	mock.ExpectPrepare("^UPDATE family SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			"bob_joe@example.com",
			"password2",
			1,
		).WillReturnResult(result)
	family := domains.Family{
		Id:         1,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  sql.NullTime{},
		PrimaryEmail:      "bob_joe@example.com",
		Password:   "password2",
	}
	err := repo.Update(1, family)
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
func TestDeleteFamily(t *testing.T) {
	db, mock, repo := initFamilyTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM family WHERE id=?").
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
func getFamilyRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"PrimaryEmail",
		"Password",
	}).AddRow(
		1,
		now,
		now,
		sql.NullTime{},
		"john_smith@example.com",
		"password",
	)
}

func getFamily() domains.Family {
	return domains.Family{
		Id:         1,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  sql.NullTime{},
		PrimaryEmail:      "john_smith@example.com",
		Password:   "password",
	}
}
