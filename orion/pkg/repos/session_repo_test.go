package repos_test

import (
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
	"reflect"
	"testing"
	"time"
)

func initSessionTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.SessionRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestSessionRepo(db)
	return db, mock, repo
}

//
// Test Select All By Class Id
//
func TestSelectAllSessionsByClassId(t *testing.T) {
	db, mock, repo := initSessionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"ClassId",
		"StartsAt",
		"EndsAt",
		"Canceled",
		"Notes",
		"PublishedAt"}).
		AddRow(1, now, now, domains.NullTime{}, "id_1", now, now, false, domains.NewNullString("special lecture from guest"), domains.NullTime{})
	mock.ExpectPrepare("^SELECT (.+) FROM sessions WHERE class_id=?").
		ExpectQuery().
		WithArgs("id_1").
		WillReturnRows(rows)
	got, err := repo.SelectAllByClassId("id_1")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Session{
		{
			Id:          1,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedAt:   domains.NullTime{},
			ClassId:     "id_1",
			StartsAt:    now,
			EndsAt:      now,
			Canceled:    false,
			Notes:       domains.NewNullString("special lecture from guest"),
			PublishedAt: domains.NullTime{},
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
func TestSelectSession(t *testing.T) {
	db, mock, repo := initSessionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"ClassId",
		"StartsAt",
		"EndsAt",
		"Canceled",
		"Notes",
		"PublishedAt"}).
		AddRow(1, now, now, domains.NullTime{}, "id_1", now, now, false, domains.NewNullString("special lecture from guest"), domains.NullTime{})
	mock.ExpectPrepare("^SELECT (.+) FROM sessions WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectBySessionId(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.Session{
		Id:          1,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   domains.NullTime{},
		ClassId:     "id_1",
		StartsAt:    now,
		EndsAt:      now,
		Canceled:    false,
		Notes:       domains.NewNullString("special lecture from guest"),
		PublishedAt: domains.NullTime{},
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
func TestInsertSession(t *testing.T) {
	db, mock, repo := initSessionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO sessions").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "id_1", now, now, false, domains.NewNullString("special lecture from guest")).
		WillReturnResult(result)
	session := domains.Session{
		ClassId:  "id_1",
		StartsAt: now,
		EndsAt:   now,
		Canceled: false,
		Notes:    domains.NewNullString("special lecture from guest"),
	}
	err := repo.Insert(session)
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
func TestUpdateSession(t *testing.T) {
	db, mock, repo := initSessionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE sessions SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), "id_1", now, now, false, domains.NewNullString("special lecture from guest"), sqlmock.AnyArg(), 1).
		WillReturnResult(result)
	session := domains.Session{
		ClassId:  "id_1",
		StartsAt: now,
		EndsAt:   now,
		Canceled: false,
		Notes:    domains.NewNullString("special lecture from guest"),
	}
	err := repo.Update(1, session)
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
func TestDeleteSession(t *testing.T) {
	db, mock, repo := initSessionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM sessions WHERE id=?").
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
// Select All Unpublished
//
func TestSelectAllUnpublishedSessions(t *testing.T) {
	db, mock, repo := initSessionTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"ClassId",
		"StartsAt",
		"EndsAt",
		"Canceled",
		"Notes",
		"PublishedAt"}).
		AddRow(1, now, now, domains.NullTime{}, "id_1", now, now, false, domains.NewNullString("special lecture from guest"), domains.NewNullTime(now))
	mock.ExpectPrepare("^SELECT (.+) FROM sessions WHERE class_id=?").
		ExpectQuery().
		WithArgs("id_1").
		WillReturnRows(rows)
	got, err := repo.SelectAllByClassId("id_1")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Session{
		{
			Id:          1,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedAt:   domains.NullTime{},
			ClassId:     "id_1",
			StartsAt:    now,
			EndsAt:      now,
			Canceled:    false,
			Notes:       domains.NewNullString("special lecture from guest"),
			PublishedAt: domains.NewNullTime(now),
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
