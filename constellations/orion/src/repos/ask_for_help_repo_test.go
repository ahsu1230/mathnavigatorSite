package repos_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
)

func initAFHTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.AskForHelpRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestAFHRepo(db)
	return db, mock, repo
}

// Test Select All
func TestSelectAllAFH(t *testing.T) {
	db, mock, repo := initAFHTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getAFHRows()
	mock.ExpectPrepare("^SELECT (.+) FROM ask_for_help").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll()
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
	got, err := repo.SelectById(1)
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
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO ask_for_help").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "AP Calculus Help", "December 25, 2020", "2:00-4:00 PM", "AP Calculus", "wchs").
		WillReturnResult(result)
	askForHelp := domains.AskForHelp{
		Title:      "AP Calculus Help",
		Date:       "December 25, 2020",
		TimeString: "2:00-4:00 PM",
		Subject:    "AP Calculus",
		LocationId: "wchs",
	}
	err := repo.Insert(askForHelp)
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
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE ask_for_help SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), 2, "AP Stat Help", "December 25, 2020", "2:00-4:00PM", "AP Stat", "room12", 1).
		WillReturnResult(result)
	askForHelp := domains.AskForHelp{
		Id:         2,
		Title:      "AP Stat Help",
		Date:       "December 25, 2020",
		TimeString: "2:00-4:00PM",
		Subject:    "AP Stat",
		LocationId: "room12",
	}
	err := repo.Update(1, askForHelp)
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
	err := repo.Delete(1)
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
	return sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "Title", "Date", "TimeString", "Subject", "LocationId"}).
		AddRow(
			1,
			now,
			now,
			domains.NullTime{},
			"AP Calculus Help",
			"August 2, 2020",
			"3:00-5:00PM",
			"AP Calculus",
			"wchs",
		)
}

func getAskForHelp() domains.AskForHelp {
	return domains.AskForHelp{
		Id:         1,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  domains.NullTime{},
		Title:      "AP Calculus Help",
		Date:       "August 2, 2020",
		TimeString: "3:00-5:00PM",
		Subject:    "AP Calculus",
		LocationId: "wchs",
	}
}
