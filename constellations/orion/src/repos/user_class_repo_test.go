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
	repo := repos.CreateTestUserClassRepo(testUtils.Context, db)
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
	mock.ExpectPrepare("^SELECT (.+) FROM user_classes WHERE class_id=?").
		ExpectQuery().
		WithArgs("abcd").
		WillReturnRows(rows)
	got, err := repo.SelectByClassId(testUtils.Context, "abcd")
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
	mock.ExpectPrepare("^SELECT (.+) FROM user_classes WHERE user_id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectByUserId(testUtils.Context, 1)
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
	mock.ExpectPrepare(`^SELECT (.+) FROM user_classes WHERE user_id=\? AND class_id=?`).
		ExpectQuery().
		WithArgs(1, "abcd").
		WillReturnRows(rows)
	got, err := repo.SelectByUserAndClass(testUtils.Context, 1, "abcd")
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
// Select New Classes
//
func TestSelectByNow(t *testing.T) {
	db, mock, repo := initUserClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserClassRows()
	mock.ExpectPrepare("^SELECT (.+) FROM user_classes WHERE created_at>=*").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectByNew(testUtils.Context)
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
// Create
//
func TestInsertUserClass(t *testing.T) {
	db, mock, repo := initUserClassTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO user_classes").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			"abcd",
			2,
			1,
			domains.USER_CLASS_ENROLLED,
		).WillReturnResult(result)
	userClass := getUserClass()
	_, err := repo.Insert(testUtils.Context, userClass)
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
	mock.ExpectPrepare("^UPDATE user_classes SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			"abcd",
			2,
			1,
			domains.USER_CLASS_ENROLLED,
			1,
		).WillReturnResult(result)
	userClasses := domains.UserClass{
		Id:        1,
		CreatedAt: testUtils.TimeNow,
		UpdatedAt: testUtils.TimeNow,
		DeletedAt: sql.NullTime{},
		ClassId:   "abcd",
		UserId:    2,
		AccountId: 1,
		State:     domains.USER_CLASS_ENROLLED,
	}
	err := repo.Update(testUtils.Context, 1, userClasses)
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
	mock.ExpectPrepare("^DELETE FROM user_classes WHERE id=?").
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
func getUserClassRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"ClassId",
		"UserId",
		"AccountId",
		"State",
	}).AddRow(
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		"abcd",
		2,
		1,
		domains.USER_CLASS_ENROLLED,
	)
}

func getUserClass() domains.UserClass {
	return domains.UserClass{
		Id:        1,
		CreatedAt: testUtils.TimeNow,
		UpdatedAt: testUtils.TimeNow,
		DeletedAt: sql.NullTime{},
		ClassId:   "abcd",
		UserId:    2,
		AccountId: 1,
		State:     domains.USER_CLASS_ENROLLED,
	}
}
