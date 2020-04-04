package repos_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
	"reflect"
	"testing"
	"time"
)

func initProgramTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.ProgramRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestProgramRepo(db)
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
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "ProgramId", "Name", "Grade1", "Grade2", "Description", "PublishedAt"}).
		AddRow(1, now, now, sql.NullTime{}, "prog1", "Program1", 2, 3, "descript1", sql.NullTime{})
	mock.ExpectPrepare("^SELECT (.+) FROM programs").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Program{
		{
			Id:          1,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedAt:   sql.NullTime{},
			ProgramId:   "prog1",
			Name:        "Program1",
			Grade1:      2,
			Grade2:      3,
			Description: "descript1",
			PublishedAt: sql.NullTime{},
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
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "ProgramId", "Name", "Grade1", "Grade2", "Description", "PublishedAt"}).
		AddRow(1, now, now, sql.NullTime{}, "prog1", "Program1", 2, 3, "descript1", sql.NullTime{})
	mock.ExpectPrepare("^SELECT (.+) FROM programs WHERE program_id=?").
		ExpectQuery().
		WithArgs("prog1").
		WillReturnRows(rows)
	got, err := repo.SelectByProgramId("prog1") // Correct programId
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.Program{
		Id:          1,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   sql.NullTime{},
		ProgramId:   "prog1",
		Name:        "Program1",
		Grade1:      2,
		Grade2:      3,
		Description: "descript1",
		PublishedAt: sql.NullTime{},
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
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "prog1", "Program1", 2, 3, "Descript1").
		WillReturnResult(result)
	program := domains.Program{
		ProgramId:   "prog1",
		Name:        "Program1",
		Grade1:      2,
		Grade2:      3,
		Description: "Descript1",
	}
	err := repo.Insert(program)
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
		WithArgs(sqlmock.AnyArg(), "prog2", "Program2", 2, 3, "Descript2", sqlmock.AnyArg(), "prog1").
		WillReturnResult(result)
	program := domains.Program{
		ProgramId:   "prog2",
		Name:        "Program2",
		Grade1:      2,
		Grade2:      3,
		Description: "Descript2",
	}
	err := repo.Update("prog1", program)
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
	err := repo.Delete("prog1")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

//
// Select All Unpublished
//
func TestSelectAllUnpublishedPrograms(t *testing.T) {
	db, mock, repo := initProgramTest(t)
	defer db.Close()

	// Mock DB statements and execute
	rows := sqlmock.NewRows([]string{"ProgramId"}).AddRow("prog1").AddRow("prog2")
	mock.ExpectPrepare("^SELECT program_id FROM programs WHERE published_at IS NULL").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAllUnpublished()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []string{"prog1", "prog2"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
