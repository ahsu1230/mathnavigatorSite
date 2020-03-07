package repos_test

import (
	"database/sql"
	"reflect"
    "testing"
    "time"
    sqlmock "github.com/DATA-DOG/go-sqlmock"
    "github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
    "github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
)

func initAnnounceTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.AnnounceRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestAnnounceRepo(db)
	return db, mock, repo
}

func TestSelectAllAnnouncements(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()
	
	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "PostedAt", "Author", "Message"}).AddRow(1, now, now, sql.NullTime{}, now, "Author Name", "Valid Message")
	mock.ExpectPrepare("^SELECT (.+) FROM announcements").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	
	// Validate results
	want := []domains.Announce {
		{
			Id:			1,
			CreatedAt:	now,
			UpdatedAt:	now,
			DeletedAt:	sql.NullTime{},
			PostedAt:	now,
			Author:		"Author Name",
			Message:	"Valid Message",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestSelectAnnouncement(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()
	
	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "PostedAt", "Author", "Message"}).AddRow(1, now, now, sql.NullTime{}, now, "Author Name", "Valid Message")
	mock.ExpectPrepare("^SELECT (.+) FROM announcements WHERE id=?").ExpectQuery().WithArgs(1).WillReturnRows(rows)
	got, err := repo.SelectByAnnounceId(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	
	// Validate results
	want := domains.Announce {
		Id:			1,
		CreatedAt:	now,
		UpdatedAt:	now,
		DeletedAt:	sql.NullTime{},
		PostedAt:	now,
		Author:		"Author Name",
		Message:	"Valid Message",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Values not equal: got = %v, want = %v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestInsertAnnouncement(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()
	
	// Mock DB statements and execute
	now := time.Now().UTC()
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO announcements").ExpectExec().WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), now, "Author Name", "Valid Message").WillReturnResult(result)
	announce := domains.Announce {
		PostedAt:	now,
		Author:		"Author Name",
		Message:	"Valid Message",
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

func TestUpdateAnnouncement(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()
	
	// Mock DB statements and execute
	now := time.Now().UTC()
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE announcements SET (.*) WHERE id=?").ExpectExec().WithArgs(sqlmock.AnyArg(), now, "Author Name", "Valid Message").WillReturnResult(result)
	announce := domains.Announce {
		PostedAt:	now,
		Author:		"Author Name",
		Message:	"Valid Message",
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

func TestDeleteAnnouncement(t *testing.T) {
	db, mock, repo := initAnnounceTest(t)
	defer db.Close()
	
	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM announcements WHERE id=?").ExpectExec().WithArgs(1).WillReturnResult(result)
	err := repo.Delete(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	
	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
