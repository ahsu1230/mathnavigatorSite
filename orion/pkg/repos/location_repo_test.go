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

func initLocationTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.LocationRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestLocationRepo(db)
	return db, mock, repo
}

//
// Test Select All
//
func TestSelectAllLocations(t *testing.T) {
	db, mock, repo := initLocationTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"LocId",
		"Street",
		"City",
		"State",
		"Zipcode",
		"Room"}).
		AddRow(1, now, now, sql.NullTime{}, "xkcd", "4040 Cherry Rd", "Potomac", "MD", "20854", sql.NullString{String: "Room 2", Valid: true})
	mock.ExpectPrepare("^SELECT (.+) FROM locations").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Location{
		{
			Id:        1,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: sql.NullTime{},
			LocId:     "xkcd",
			Street:    "4040 Cherry Rd",
			City:      "Potomac",
			State:     "MD",
			Zipcode:   "20854",
			Room:      "Room 2",
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
func TestSelectLocation(t *testing.T) {
	db, mock, repo := initLocationTest(t)
	defer db.Close()

	// Mock DB statements and execute
	now := time.Now().UTC()
	rows := sqlmock.NewRows([]string{
		"Id",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"LocId",
		"Street",
		"City",
		"State",
		"Zipcode",
		"Room"}).
		AddRow(1, now, now, sql.NullTime{}, "xkcd", "4040 Cherry Rd", "Potomac", "MD", "20854", sql.NullString{String: "Room 2", Valid: true})
	mock.ExpectPrepare("^SELECT (.+) FROM locations WHERE loc_id=?").
		ExpectQuery().
		WithArgs("xkcd").
		WillReturnRows(rows)
	got, err := repo.SelectByLocationId("xkcd")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.Location{
		Id:        1,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: sql.NullTime{},
		LocId:     "xkcd",
		Street:    "4040 Cherry Rd",
		City:      "Potomac",
		State:     "MD",
		Zipcode:   "20854",
		Room:      "Room 2",
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
func TestInsertLocation(t *testing.T) {
	db, mock, repo := initLocationTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^INSERT INTO locations").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "xkcd", "4040 Cherry Rd", "Potomac", "MD", "20854", sql.NullString{String: "Room 2", Valid: true}).
		WillReturnResult(result)
	location := domains.Location{
		LocId:   "xkcd",
		Street:  "4040 Cherry Rd",
		City:    "Potomac",
		State:   "MD",
		Zipcode: "20854",
		Room:    "Room 2",
	}
	err := repo.Insert(location)
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
func TestUpdateLocation(t *testing.T) {
	db, mock, repo := initLocationTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^UPDATE locations SET (.*) WHERE loc_id=?").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), "www", "4041 Cherry Rd", "San Francisco", "CA", "94016", sql.NullString{String: "Room 41", Valid: true}, "xkcd").
		WillReturnResult(result)
	location := domains.Location{
		LocId:   "www",
		Street:  "4041 Cherry Rd",
		City:    "San Francisco",
		State:   "CA",
		Zipcode: "94016",
		Room:    "Room 41",
	}
	err := repo.Update("xkcd", location)
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
func TestDeleteLocation(t *testing.T) {
	db, mock, repo := initLocationTest(t)
	defer db.Close()

	// Mock DB statements and execute
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("^DELETE FROM locations WHERE loc_id=?").
		ExpectExec().
		WithArgs("xkcd").
		WillReturnResult(result)
	err := repo.Delete("xkcd")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
