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

func initUserClassTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.UserClassRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestUserClassRepo(db)
	return db, mock, repo
}

//
// Select Many by Class Id
//
func TestSelectUsersByClassId(t *testing.T) {
	db, mock, repo := initUserClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM userclass WHERE class_id=?").
		ExpectQuery().
		WithArgs("abcd").
		WillReturnRows(rows)
	got, err := repo.SelectByClassId("abcd")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.UserClass{getUserClass()}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select Many By User ID
//
func TestSelectClassesByUserId(t *testing.T) {
	db, mock, repo := initUserClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM userclass WHERE user_id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectByUserId(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.UserClass{getUserClass()}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select One By User ID and Class ID
//
func TestSelectByUserAndClass(t *testing.T) {
	db, mock, repo := initUserClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserClassRows()
	mock.ExpectPrepare(`^SELECT (.+) FROM userclass WHERE user_id=\? AND class_id=?`).
		ExpectQuery().
		WithArgs(1, "abcd").
		WillReturnRows(rows)
	got, err := repo.SelectByUserAndClass(1, "abcd")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getUserClass()

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
func TestInsertUserClass(t *testing.T) {
	db, mock, repo := initUserClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO userclass").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			1,
			"abcd",
			1,
			1,
		).WillReturnResult(result)
	userClass := getUserClass()
	err := repo.Insert(userClass)
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
func TestUpdateUserClass(t *testing.T) {
	db, mock, repo := initUserClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE userclass SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			1,
			"abcd",
			1,
			1,
			1,
		).WillReturnResult(result)
	userClass := domains.UserClass{
		Id:        1,
		CreatedAt: testUtils.TimeNow,
		UpdatedAt: testUtils.TimeNow,
		DeletedAt: sql.NullTime{},
		UserId:    1,
		ClassId:   "abcd",
		AccountId: 1,
		State:     1,
	}
	err := repo.Update(1, userClass)
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
func TestDeleteUserClass(t *testing.T) {
	db, mock, repo := initUserClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM userclass WHERE id=?").
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
func getUserClassRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"UserId",
		"ClassId",
		"AccountId",
		"State",
	}).AddRow(
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		1,
		"abcd",
		1,
		1,
	)
}

func getUserClass() domains.UserClass {
	return domains.UserClass{
		Id:        1,
		CreatedAt: testUtils.TimeNow,
		UpdatedAt: testUtils.TimeNow,
		DeletedAt: sql.NullTime{},
		UserId:    1,
		ClassId:   "abcd",
		AccountId: 1,
		State:     1,
	}
}
