package repos_test

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
)

var now time.Time

func initSemesterTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.SemesterRepoInterface) {
	now = time.Now().UTC()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestSemesterRepo(db)
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
	got, err := repo.SelectAll(false)
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
// Test Select Published
//
func TestSelectPublishedSemesters(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getSemesterRows()
	mock.ExpectPrepare("^SELECT (.+) FROM semesters WHERE published_at IS NOT NULL").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll(true)
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
// Select Unpublished
//
func TestSelectAllUnpublishedSemesters(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := getSemesterRows()
	mock.ExpectPrepare("^SELECT (.+) FROM semesters WHERE published_at IS NULL").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAllUnpublished()
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
	got, err := repo.SelectBySemesterId("2020_fall") // Correct semesterId
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
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "2020_fall", "Fall 2020").
		WillReturnResult(result)
	semester := domains.Semester{
		SemesterId: "2020_fall",
		Title:      "Fall 2020",
	}
	err := repo.Insert(semester)
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
		WithArgs(sqlmock.AnyArg(), "2021_spring", "Spring 2021", "2020_fall").
		WillReturnResult(result)
	semester := domains.Semester{
		SemesterId: "2021_spring",
		Title:      "Spring 2021",
	}
	err := repo.Update("2020_fall", semester)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Publish
//
func TestPublishSemesters(t *testing.T) {
	db, mock, repo := initSemesterTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectBegin()
	mock.ExpectPrepare(`^UPDATE semesters SET published_at=\? WHERE semester_id=\? AND published_at IS NULL`).
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), "2020_fall").
		WillReturnResult(result)
	mock.ExpectCommit()
	err := repo.Publish([]string{"2020_fall"})
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
	err := repo.Delete("2020_fall")
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
	return sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "PublishedAt", "SemesterId", "Title"}).
		AddRow(
			1,
			now,
			now,
			domains.NullTime{},
			domains.NullTime{},
			"2020_fall",
			"Fall 2020",
		)
}

func getSemester() domains.Semester {
	return domains.Semester{
		Id:          1,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   domains.NullTime{},
		PublishedAt: domains.NullTime{},
		SemesterId:  "2020_fall",
		Title:       "Fall 2020",
	}
}
