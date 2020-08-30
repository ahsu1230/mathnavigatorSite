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

func initUserAfhTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.UserAfhRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestUserAfhRepo(db)
	return db, mock, repo
}

// Test Select By User Id
func TestSelectByUserId(t *testing.T) {
	db, mock, repo := initUserAfhTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"UserId",
		"AfhId",
		"AccountId",
	}).AddRow(
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		2,
		3,
		1,
	)
	mock.ExpectPrepare("^SELECT (.+) FROM user_afh WHERE user_id=?").
		ExpectQuery().
		WithArgs(2).
		WillReturnRows(rows)
	got, err := repo.SelectByUserId(2)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.UserAfh{
		{
			Id:        1,
			CreatedAt: testUtils.TimeNow,
			UpdatedAt: testUtils.TimeNow,
			DeletedAt: sql.NullTime{},
			UserId:    2,
			AfhId:     3,
			AccountId: 1,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Test Select By Afh Id
func TestSelectByAfhId(t *testing.T) {
	db, mock, repo := initUserAfhTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"UserId",
		"AfhId",
		"AccountId",
	}).AddRow(
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		2,
		3,
		1,
	)
	mock.ExpectPrepare("^SELECT (.+) FROM user_afh WHERE afh_id=?").
		ExpectQuery().
		WithArgs(3).
		WillReturnRows(rows)
	got, err := repo.SelectByAfhId(3)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.UserAfh{
		{
			Id:        1,
			CreatedAt: testUtils.TimeNow,
			UpdatedAt: testUtils.TimeNow,
			DeletedAt: sql.NullTime{},
			UserId:    2,
			AfhId:     3,
			AccountId: 1,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Test Select By User Id and Afh Id
func TestSelectByBothIds(t *testing.T) {
	db, mock, repo := initUserAfhTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"UserId",
		"AfhId",
		"AccountId",
	}).AddRow(
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		2,
		3,
		1,
	)
	mock.ExpectPrepare(`^SELECT (.+) FROM user_afh WHERE user_id=\? AND afh_id=?`).
		ExpectQuery().
		WithArgs(2, 3).
		WillReturnRows(rows)
	got, err := repo.SelectByBothIds(2, 3)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.UserAfh{
		Id:        1,
		CreatedAt: testUtils.TimeNow,
		UpdatedAt: testUtils.TimeNow,
		DeletedAt: sql.NullTime{},
		UserId:    2,
		AfhId:     3,
		AccountId: 1,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Select by new
func TestSelectByNew(t *testing.T) {
	db, mock, repo := initUserAfhTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"UserId",
		"AfhId",
		"AccountId",
	}).AddRow(
		1,
		testUtils.TimeNow,
		testUtils.TimeNow,
		sql.NullTime{},
		2,
		3,
		1,
	)
	mock.ExpectPrepare("^SELECT (.+) FROM user_afh WHERE created_at>=*").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectByNew()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.UserAfh{
		{
			Id:        1,
			CreatedAt: testUtils.TimeNow,
			UpdatedAt: testUtils.TimeNow,
			DeletedAt: sql.NullTime{},
			UserId:    2,
			AfhId:     3,
			AccountId: 1,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Create
func TestInsertUserAfh(t *testing.T) {
	db, mock, repo := initUserAfhTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO user_afh").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), 2, 3, 1).
		WillReturnResult(result)
	userAfh := domains.UserAfh{
		UserId: 2,
		AfhId:  3,
		AccountId: 1,
	}
	err := repo.Insert(userAfh)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Update
func TestUpdateUserAfh(t *testing.T) {
	db, mock, repo := initUserAfhTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE user_afh SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(3, 3, 2, sqlmock.AnyArg(), 1).
		WillReturnResult(result)
	userAfh := domains.UserAfh{
		Id:     1,
		UserId: 3,
		AfhId:  3,
		AccountId: 2,
	}
	err := repo.Update(1, userAfh)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Delete
func TestDeleteUserAfh(t *testing.T) {
	db, mock, repo := initUserAfhTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM user_afh WHERE id=?").
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
