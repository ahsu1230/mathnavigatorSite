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

func initSemesterTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.SemesterRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestSemesterRepo(testUtils.Context, db)
	return db, mock, repo
}

//
// Test Select All
//
func TestSelectAllSemesters(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getSemesterRows()
	mock.ExpectPrepare("^SELECT (.+) FROM semesters").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll(testUtils.Context)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Semester{getSemester()}
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
func TestSelectSemester(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getSemesterRows()
	mock.ExpectPrepare("^SELECT (.+) FROM semesters WHERE semester_id=?").
		ExpectQuery().
		WithArgs("2020_fall").
		WillReturnRows(rows)
	got, err := repo.SelectBySemesterId(testUtils.Context, "2020_fall") // Correct semesterId
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := getSemester()
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
func TestInsertSemester(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO semesters").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "2020_fall", domains.FALL, 2020, "Fall 2020").
		WillReturnResult(result)
	semester := domains.Semester{
		SemesterId: "2020_fall",
		Season:     domains.FALL,
		Year:       2020,
		Title:      "Fall 2020",
	}
	_, err := repo.Insert(testUtils.Context, semester)
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
func TestUpdateSemester(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE semesters SET (.*) WHERE semester_id=?").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), "2021_spring", domains.SPRING, 2021, "Spring 2021", "2020_fall").
		WillReturnResult(result)
	semester := domains.Semester{
		SemesterId: "2021_spring",
		Season:     domains.SPRING,
		Year:       2021,
		Title:      "Spring 2021",
	}
	err := repo.Update(testUtils.Context, "2020_fall", semester)
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
func TestDeleteSemester(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM semesters WHERE semester_id=?").
		ExpectExec().
		WithArgs("2020_fall").
		WillReturnResult(result)
	err := repo.Delete(testUtils.Context, "2020_fall")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestArchiveSemester(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("UPDATE semesters SET deleted_at=(.+) WHERE semester_id=(.+)").
		ExpectExec().
		WithArgs(testUtils.MockAnyTime{}, "2020_fall").
		WillReturnResult(result)
	err := repo.Archive(testUtils.Context, "2020_fall")
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
func getSemesterRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "SemesterId", "Season", "Year", "Title"}).
		AddRow(
			1,
			testUtils.TimeNow,
			testUtils.TimeNow,
			domains.NullTime{},
			"2020_fall",
			"fall",
			2020,
			"Fall 2020",
		)
}

func getSemester() domains.Semester {
	return domains.Semester{
		Id:         1,
		CreatedAt:  testUtils.TimeNow,
		UpdatedAt:  testUtils.TimeNow,
		DeletedAt:  domains.NullTime{},
		SemesterId: "2020_fall",
		Season:     "fall",
		Year:       2020,
		Title:      "Fall 2020",
	}
}
