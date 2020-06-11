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

func initAnnounceTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.AnnounceRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestAnnounceRepo(db)
	return db, mock, repo
}

//
// Test Select All
//
func TestSelectAllAnnouncements(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	early := time.Unix(0, 0)
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"PostedAt",
		"Author",
		"Message",
		"OnHomePage"}).
		AddRow(1, now, now, domains.NullTime{}, now, "Author Name", "Valid Message", false).
		AddRow(2, early, early, domains.NullTime{}, early, "Author Name 2", "Valid Message 2", true)
	mock.ExpectPrepare("^SELECT (.+) FROM announcements").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Announce{
		{
			Id:        1,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: domains.NullTime{},
			PostedAt:  now,
			Author:    "Author Name",
			Message:   "Valid Message",
			OnHomePage: false,
		},
		{
			Id:        2,
			CreatedAt: early,
			UpdatedAt: early,
			DeletedAt: domains.NullTime{},
			PostedAt:  early,
			Author:    "Author Name 2",
			Message:   "Valid Message 2",
			OnHomePage: true,
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
func TestSelectAnnouncement(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"PostedAt",
		"Author",
		"Message",
		"OnHomePage"}).
		AddRow(1, now, now, domains.NullTime{}, now, "Author Name", "Valid Message", false)
	mock.ExpectPrepare("^SELECT (.+) FROM announcements WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectByAnnounceId(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.Announce{
		Id:        1,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: domains.NullTime{},
		PostedAt:  now,
		Author:    "Author Name",
		Message:   "Valid Message",
		OnHomePage: false,
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
func TestInsertAnnouncement(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO announcements").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), now, "Author Name", "Valid Message").
		WillReturnResult(result)
	announce := domains.Announce{
		PostedAt: now,
		Author:   "Author Name",
		Message:  "Valid Message",
	}
	err := repo.Insert(announce)
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
func TestUpdateAnnouncement(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE announcements SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), now, "Author Name", "Valid Message", 1).
		WillReturnResult(result)
	announce := domains.Announce{
		PostedAt: now,
		Author:   "Author Name",
		Message:  "Valid Message",
	}
	err := repo.Update(1, announce)
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
func TestDeleteAnnouncement(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM announcements WHERE id=?").
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
