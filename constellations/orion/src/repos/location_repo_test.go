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

func initLocationTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.LocationRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestLocationRepo(testUtils.Context, db)
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
		"LocationId",
		"Title",
		"Street",
		"City",
		"State",
		"Zipcode",
		"Room",
		"IsOnline"}).
		AddRow(
			1,
			now,
			now,
			domains.NullTime{},
			"xkcd",
			"Potomac High School",
			domains.NewNullString("4040 Cherry Rd"),
			domains.NewNullString("Potomac"),
			domains.NewNullString("MD"),
			domains.NewNullString("20854"),
			domains.NewNullString("Room 2"),
			false,
		)
	mock.ExpectPrepare("^SELECT (.+) FROM locations").
		ExpectQuery().
		WillReturnRows(rows)
	got, err := repo.SelectAll(testUtils.Context)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := []domains.Location{
		{
			Id:         1,
			CreatedAt:  now,
			UpdatedAt:  now,
			DeletedAt:  domains.NullTime{},
			LocationId: "xkcd",
			Title:      "Potomac High School",
			Street:     domains.NewNullString("4040 Cherry Rd"),
			City:       domains.NewNullString("Potomac"),
			State:      domains.NewNullString("MD"),
			Zipcode:    domains.NewNullString("20854"),
			Room:       domains.NewNullString("Room 2"),
			IsOnline:   false,
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
		"LocationId",
		"Title",
		"Street",
		"City",
		"State",
		"Zipcode",
		"Room",
		"IsOnline"}).
		AddRow(
			1,
			now,
			now,
			domains.NullTime{},
			"xkcd",
			"Potomac High School",
			domains.NewNullString("4040 Cherry Rd"),
			domains.NewNullString("Potomac"),
			domains.NewNullString("MD"),
			domains.NewNullString("20854"),
			domains.NewNullString("Room 2"),
			false,
		)
	mock.ExpectPrepare("^SELECT (.+) FROM locations WHERE location_id=?").
		ExpectQuery().
		WithArgs("xkcd").
		WillReturnRows(rows)
	got, err := repo.SelectByLocationId(testUtils.Context, "xkcd")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	want := domains.Location{
		Id:         1,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  domains.NullTime{},
		LocationId: "xkcd",
		Title:      "Potomac High School",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
		Room:       domains.NewNullString("Room 2"),
		IsOnline:   false,
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
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			"xkcd",
			"School1",
			domains.NewNullString("4040 Cherry Rd"),
			domains.NewNullString("Potomac"),
			domains.NewNullString("MD"),
			domains.NewNullString("20854"),
			domains.NewNullString("Room 2"),
		).
		WillReturnResult(result)
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
		Room:       domains.NewNullString("Room 2"),
	}
	_, err := repo.Insert(testUtils.Context, location)
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
	mock.ExpectPrepare("^UPDATE locations SET (.*) WHERE location_id=?").
		ExpectExec().
		WithArgs(
			sqlmock.AnyArg(),
			"www",
			"School1",
			domains.NewNullString("4041 Cherry Rd"),
			domains.NewNullString("San Francisco"),
			domains.NewNullString("CA"),
			domains.NewNullString("94016"),
			domains.NewNullString("Room 41"),
			"xkcd",
		).
		WillReturnResult(result)
	location := domains.Location{
		LocationId: "www",
		Title:      "School1",
		Street:     domains.NewNullString("4041 Cherry Rd"),
		City:       domains.NewNullString("San Francisco"),
		State:      domains.NewNullString("CA"),
		Zipcode:    domains.NewNullString("94016"),
		Room:       domains.NewNullString("Room 41"),
	}
	err := repo.Update(testUtils.Context, "xkcd", location)
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
	mock.ExpectPrepare("^DELETE FROM locations WHERE location_id=?").
		ExpectExec().
		WithArgs("xkcd").
		WillReturnResult(result)
	err := repo.Delete(testUtils.Context, "xkcd")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	// Validate results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
