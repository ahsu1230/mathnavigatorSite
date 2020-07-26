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

func initUserClassesTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.UserClassesRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestUserClassesRepo(db)
	return db, mock, repo
}

//
// Select Many by Class Id
//
func TestSelectUsersByClassId(t *testing.T) {
	db, mock, repo := initUserClassesTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserClassesRows()
	mock.ExpectPrepare("^SELECT (.+) FROM user_classes WHERE class_id=?").
		ExpectQuery().
		WithArgs("abcd").
		WillReturnRows(rows)
	got, err := repo.SelectByClassId("abcd")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.UserClasses{getUserClasses()}

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
	db, mock, repo := initUserClassesTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserClassesRows()
	mock.ExpectPrepare("^SELECT (.+) FROM user_classes WHERE user_id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectByUserId(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.UserClasses{getUserClasses()}

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
	db, mock, repo := initUserClassesTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getUserClassesRows()
	mock.ExpectPrepare(`^SELECT (.+) FROM user_classes WHERE user_id=\? AND class_id=?`).
		ExpectQuery().
		WithArgs(1, "abcd").
		WillReturnRows(rows)
	got, err := repo.SelectByUserAndClass(1, "abcd")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getUserClasses()

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
	db, mock, repo := initUserClassesTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO user_classes").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			1,
			"abcd",
			1,
			domains.USER_CLASS_ACCEPTED,
		).WillReturnResult(result)
	userClasses := getUserClasses()
	err := repo.Insert(userClasses)
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
	db, mock, repo := initUserClassesTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE user_classes SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			1,
			"abcd",
			1,
			domains.USER_CLASS_ACCEPTED,
			1,
		).WillReturnResult(result)
	userClasses := domains.UserClasses{
		Id:        1,
		CreatedAt: testUtils.TimeNow,
		UpdatedAt: testUtils.TimeNow,
		DeletedAt: sql.NullTime{},
		UserId:    1,
		ClassId:   "abcd",
		AccountId: 1,
		State:     domains.USER_CLASS_ACCEPTED,
	}
	err := repo.Update(1, userClasses)
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
	db, mock, repo := initUserClassesTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM user_classes WHERE id=?").
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
func getUserClassesRows() *sqlmock.Rows {
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
		domains.USER_CLASS_ACCEPTED,
	)
}

func getUserClasses() domains.UserClasses {
	return domains.UserClasses{
		Id:        1,
		CreatedAt: testUtils.TimeNow,
		UpdatedAt: testUtils.TimeNow,
		DeletedAt: sql.NullTime{},
		UserId:    1,
		ClassId:   "abcd",
		AccountId: 1,
		State:     domains.USER_CLASS_ACCEPTED,
	}
}
