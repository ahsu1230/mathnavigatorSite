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

func initProgramTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.ProgramRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestProgramRepo(testUtils.Context, db)
	return db, mock, repo
}

//
// Test Select All
//
func TestSelectAllPrograms(t *testing.T) {
	db, mock, repo := initProgramTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"ProgramId",
		"Title",
		"Grade1",
		"Grade2",
		"Subject",
		"Description",
		"Featured",
	}).
		AddRow(
			1,
			now,
			now,
			domains.NullTime{},
			"prog1",
			"Program1",
			2,
			3,
			domains.SUBJECT_MATH,
			"descript1",
			domains.FEATURED_NONE,
		)
	mock.ExpectPrepare("^SELECT (.+) FROM programs").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll(testUtils.Context)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Program{
		{
			Id:          1,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedAt:   domains.NullTime{},
			ProgramId:   "prog1",
			Title:       "Program1",
			Grade1:      2,
			Grade2:      3,
			Subject:     domains.SUBJECT_MATH,
			Description: "descript1",
			Featured:    domains.FEATURED_NONE,
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
func TestSelectProgram(t *testing.T) {
	db, mock, repo := initProgramTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"ProgramId",
		"Title",
		"Grade1",
		"Grade2",
		"Subject",
		"Description",
		"Featured",
	}).
		AddRow(
			1,
			now,
			now,
			domains.NullTime{},
			"prog1",
			"Program1",
			2,
			3,
			domains.SUBJECT_MATH,
			"descript1",
			domains.FEATURED_NONE,
		)
	mock.ExpectPrepare("^SELECT (.+) FROM programs WHERE program_id=?").
		ExpectQuery().
		WithArgs("prog1").
		WillReturnRows(rows)
	got, err := repo.SelectByProgramId(testUtils.Context, "prog1") // Correct programId
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.Program{
		Id:          1,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   domains.NullTime{},
		ProgramId:   "prog1",
		Title:       "Program1",
		Grade1:      2,
		Grade2:      3,
		Subject:     domains.SUBJECT_MATH,
		Description: "descript1",
		Featured:    domains.FEATURED_NONE,
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
func TestInsertProgram(t *testing.T) {
	db, mock, repo := initProgramTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO programs").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			"prog1",
			"Program1",
			2,
			3,
			domains.SUBJECT_MATH,
			"Descript1",
			domains.FEATURED_NONE,
		).WillReturnResult(result)
	program := domains.Program{
		ProgramId:   "prog1",
		Title:       "Program1",
		Grade1:      2,
		Grade2:      3,
		Subject:     domains.SUBJECT_MATH,
		Description: "Descript1",
		Featured:    domains.FEATURED_NONE,
	}
	_, err := repo.Insert(testUtils.Context, program)
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
func TestUpdateProgram(t *testing.T) {
	db, mock, repo := initProgramTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE programs SET (.*) WHERE program_id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			"prog2",
			"Program2",
			2,
			3,
			domains.SUBJECT_MATH,
			"Descript2",
			domains.FEATURED_NONE,
			"prog1",
		).WillReturnResult(result)
	program := domains.Program{
		ProgramId:   "prog2",
		Title:       "Program2",
		Grade1:      2,
		Grade2:      3,
		Subject:     domains.SUBJECT_MATH,
		Description: "Descript2",
		Featured:    domains.FEATURED_NONE,
	}
	err := repo.Update(testUtils.Context, "prog1", program)
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
func TestDeleteProgram(t *testing.T) {
	db, mock, repo := initProgramTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM programs WHERE program_id=?").
		ExpectExec().
		WithArgs("prog1").
		WillReturnResult(result)
	err := repo.Delete(testUtils.Context, "prog1")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Archive
//
func TestArchiveProgram(t *testing.T) {
	db, mock, repo := initProgramTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("UPDATE programs SET deleted_at=(.+) WHERE program_id=(.+)").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), "prog1").
		WillReturnResult(result)
	err := repo.Archive(testUtils.Context, "prog1")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
