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

func initAchieveTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.AchieveRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestAchieveRepo(db)
	return db, mock, repo
}

//
// Test Select All
//
func TestSelectAllAchieves(t *testing.T) {
	db, mock, repo := initAchieveTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "PublishedAt", "Year", "Message"}).
		AddRow(1, now, now, sql.NullTime{}, sql.NullTime{}, 2020, "message1")
	mock.ExpectPrepare("^SELECT (.+) FROM achievements").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAll()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Achieve{
		{
			Id:          1,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedAt:   sql.NullTime{},
			PublishedAt: sql.NullTime{},
			Year:        2020,
			Message:     "message1",
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
// Select All By Year
//
func TestSelectAllGroupedByYear(t *testing.T) {
	db, mock, repo := initAchieveTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "PublishedAt", "Year", "Message"}).
		AddRow(2, now, now, sql.NullTime{}, sql.NullTime{}, 2021, "1600 on SAT").
		AddRow(1, now, now, sql.NullTime{}, sql.NullTime{}, 2020, "message1")
	mock.ExpectPrepare("^SELECT (.+) FROM achievements ORDER BY year DESC").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectAllGroupedByYear()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.AchieveYearGroup{
		{
			Year: 2021,
			Achievements: []domains.Achieve{
				{
					Id:          2,
					CreatedAt:   now,
					UpdatedAt:   now,
					DeletedAt:   sql.NullTime{},
					PublishedAt: sql.NullTime{},
					Year:        2021,
					Message:     "1600 on SAT",
				},
			},
		},
		{
			Year: 2020,
			Achievements: []domains.Achieve{
				{
					Id:          1,
					CreatedAt:   now,
					UpdatedAt:   now,
					DeletedAt:   sql.NullTime{},
					PublishedAt: sql.NullTime{},
					Year:        2020,
					Message:     "message1",
				},
			},
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
func TestSelectAchieve(t *testing.T) {
	db, mock, repo := initAchieveTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "DeletedAt", "Year", "Message"}).
		AddRow(1, now, now, sql.NullTime{}, sql.NullTime{}, 2020, "message1")
	mock.ExpectPrepare("^SELECT (.+) FROM achievements WHERE id=?").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)
	got, err := repo.SelectById(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.Achieve{
		Id:          1,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   sql.NullTime{},
		PublishedAt: sql.NullTime{},
		Year:        2020,
		Message:     "message1",
	}
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
func TestSelectUnpublishedAchieves(t *testing.T) {
	db, mock, repo := initAchieveTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt", "PublishedAt", "Year", "Message"}).
		AddRow(1, now, now, sql.NullTime{}, sql.NullTime{}, 2020, "message1")
	mock.ExpectPrepare("^SELECT (.+) FROM achievements WHERE published_at IS NULL").ExpectQuery().WillReturnRows(rows)
	got, err := repo.SelectUnpublished()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Achieve{
		{
			Id:          1,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedAt:   sql.NullTime{},
			PublishedAt: sql.NullTime{},
			Year:        2020,
			Message:     "message1",
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
// Create
//
func TestInsertAchieve(t *testing.T) {
	db, mock, repo := initAchieveTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO achievements").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), 2020, "message1").
		WillReturnResult(result)
	achieve := domains.Achieve{
		Year:    2020,
		Message: "message1",
	}
	err := repo.Insert(achieve)
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
func TestUpdateAchieve(t *testing.T) {
	db, mock, repo := initAchieveTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE achievements SET (.*) WHERE id=?").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), 2021, "message2", 1).
		WillReturnResult(result)
	achieve := domains.Achieve{
		Year:    2021,
		Message: "message2",
	}
	err := repo.Update(1, achieve)
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
func TestDeleteAchieve(t *testing.T) {
	db, mock, repo := initAchieveTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM achievements WHERE id=?").
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
// Publish
//
func TestPublishAchieves(t *testing.T) {
	db, mock, repo := initAchieveTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare(`^UPDATE achievements SET published_at=\? WHERE id=\?`).
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnResult(result)
	err := repo.Publish(1)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
