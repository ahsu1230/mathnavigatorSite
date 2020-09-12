package repos_test

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/testUtils"
)

func initAFHTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.AskForHelpRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestAFHRepo(testUtils.Context, db)
	return db, mock, repo
}

// Test Select All
func TestSelectAllAFH(t *testing.T) {
	db, mock, repo := initAFHTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAFHRows()
	mock.ExpectPrepare("^SELECT (.+) FROM ask_for_help").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll(testUtils.Context)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.AskForHelp{getAskForHelp()}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

// Test Select By ID
func TestSelectAFH(t *testing.T) {
	db, mock, repo := initAFHTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAFHRows()
	mock.ExpectPrepare("^SELECT (.+) FROM ask_for_help WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectById(testUtils.Context, 1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getAskForHelp()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestInsertAFH(t *testing.T) {
	db, mock, repo := initAFHTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	start := now.Add(time.Hour * 24 * 30)
	end := start.Add(time.Hour * 1)
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO ask_for_help").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			start,
			end,
			"AP Calculus Help",
			domains.SUBJECT_MATH,
			"wchs",
			domains.NewNullString("test note"),
		).WillReturnResult(result)
	askForHelp := domains.AskForHelp{
		Title:      "AP Calculus Help",
		StartsAt:   start,
		EndsAt:     end,
		Subject:    domains.SUBJECT_MATH,
		LocationId: "wchs",
		Notes:      domains.NewNullString("test note"),
	}
	_, err := repo.Insert(testUtils.Context, askForHelp)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUpdateAFH(t *testing.T) {
	db, mock, repo := initAFHTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	start := now.Add(time.Hour * 24 * 30)
	end := start.Add(time.Hour * 1)
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE ask_for_help SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			2,
			start,
			end,
			"AP Stat Help",
			domains.SUBJECT_MATH,
			"room12",
			domains.NewNullString("test note"),
			1).
		WillReturnResult(result)
	askForHelp := domains.AskForHelp{
		Id:         2,
		StartsAt:   start,
		EndsAt:     end,
		Title:      "AP Stat Help",
		Subject:    domains.SUBJECT_MATH,
		LocationId: "room12",
		Notes:      domains.NewNullString("test note"),
	}
	err := repo.Update(testUtils.Context, 1, askForHelp)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestDeleteAFH(t *testing.T) {
	db, mock, repo := initAFHTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM ask_for_help WHERE id=?").
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

// Helper Methods
func getAFHRows() *sqlmock.Rows {
	start := testUtils.TimeNow
	end := start.Add(time.Hour * 1)
	return sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"StartsAt",
		"EndsAt",
		"Title",
		"Subject",
		"LocationId",
		"Notes",
	}).
		AddRow(
			1,
			testUtils.TimeNow,
			testUtils.TimeNow,
			domains.NullTime{},
			start,
			end,
			"AP Calculus Help",
			domains.SUBJECT_MATH,
			"wchs",
			domains.NewNullString("test note"),
		)
}

func getAskForHelp() domains.AskForHelp {
	start := testUtils.TimeNow
	end := start.Add(time.Hour * 1)
	return domains.AskForHelp{
		Id:         1,
		CreatedAt:  testUtils.TimeNow,
		UpdatedAt:  testUtils.TimeNow,
		DeletedAt:  domains.NullTime{},
		StartsAt:   start,
		EndsAt:     end,
		Title:      "AP Calculus Help",
		Subject:    domains.SUBJECT_MATH,
		LocationId: "wchs",
		Notes:      domains.NewNullString("test note"),
	}
}
