package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

// Global variable
var LocationRepo LocationRepoInterface = &locationRepo{}

type locationRepo struct {
	db *sql.DB
}

type LocationRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.Location, error)
	SelectByLocationId(string) (domains.Location, error)
	Insert(domains.Location) error
	Update(string, domains.Location) error
	Delete(string) error
}

func (lr *locationRepo) Initialize(db *sql.DB) {
	lr.db = db
}

func (lr *locationRepo) SelectAll() ([]domains.Location, error) {
	results := make([]domains.Location, 0)

	statement := "SELECT * FROM locations"

	stmt, err := lr.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var location domains.Location
		if errScan := rows.Scan(
			&location.Id,
			&location.CreatedAt,
			&location.UpdatedAt,
			&location.DeletedAt,
			&location.LocationId,
			&location.Street,
			&location.City,
			&location.State,
			&location.Zipcode,
			&location.Room); errScan != nil {
			return results, errScan
		}
		results = append(results, location)
	}

	return results, nil
}

func (lr *locationRepo) SelectByLocationId(locationId string) (domains.Location, error) {
	stmt, err := lr.db.Prepare("SELECT * FROM locations WHERE location_id=?")
	if err != nil {
		return domains.Location{}, err
	}
	defer stmt.Close()

	var location domains.Location
	row := stmt.QueryRow(locationId)
	errScan := row.Scan(
		&location.Id,
		&location.CreatedAt,
		&location.UpdatedAt,
		&location.DeletedAt,
		&location.LocationId,
		&location.Street,
		&location.City,
		&location.State,
		&location.Zipcode,
		&location.Room)

	return location, errScan
}

func (lr *locationRepo) Insert(location domains.Location) error {
	stmt, err := lr.db.Prepare("INSERT INTO locations (" +
		"created_at, " +
		"updated_at, " +
		"location_id, " +
		"street, " +
		"city, " +
		"state, " +
		"zipcode, " +
		"room" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		now,
		location.LocationId,
		location.Street,
		location.City,
		location.State,
		location.Zipcode,
		location.Room)
	if err != nil {
		return err
	}

	return utils.HandleSqlExecResult(result, 1, "location was not inserted")
}

func (lr *locationRepo) Update(locationId string, location domains.Location) error {
	stmt, err := lr.db.Prepare("UPDATE locations SET " +
		"updated_at=?, " +
		"location_id=?, " +
		"street=?, " +
		"city=?, " +
		"state=?, " +
		"zipcode=?, " +
		"room=? " +
		"WHERE location_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		location.LocationId,
		location.Street,
		location.City,
		location.State,
		location.Zipcode,
		location.Room,
		locationId)
	if err != nil {
		return err
	}

	return utils.HandleSqlExecResult(result, 1, "location was not updated")
}

func (lr *locationRepo) Delete(locationId string) error {
	stmt, err := lr.db.Prepare("DELETE FROM locations WHERE location_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(locationId)
	if err != nil {
		return err
	}

	return utils.HandleSqlExecResult(result, 1, "location was not deleted")
}

func CreateTestLocationRepo(db *sql.DB) LocationRepoInterface {
	lr := &locationRepo{}
	lr.Initialize(db)
	return lr
}
