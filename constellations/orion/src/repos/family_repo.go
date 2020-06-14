package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

// Global variable
var FamilyRepo FamilyRepoInterface = &familyRepo{}

// Implements interface userRepoInterface
type familyRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type FamilyRepoInterface interface {
	Initialize(db *sql.DB)
	SelectById(uint) (domains.Family, error)
	SelectByPrimaryEmail(string) (domains.Family, error)
	Insert(domains.Family) error
	Update(uint, domains.Family) error
	Delete(uint) error
}

func (fam *familyRepo) Initialize(db *sql.DB) {
	fam.db = db
}

func (fam *familyRepo) SelectById(id uint) (domains.Family, error) {
	statement := "SELECT * FROM families WHERE id=?"
	stmt, err := fam.db.Prepare(statement)
	if err != nil {
		return domains.Family{}, err
	}
	defer stmt.Close()

	var family domains.Family
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&family.Id,
		&family.CreatedAt,
		&family.UpdatedAt,
		&family.DeletedAt,
		&family.PrimaryEmail,
		&family.Password)
	return family, errScan
}

func (fam *familyRepo) SelectByPrimaryEmail(primary_email string) (domains.Family, error) {
	statement := "SELECT * FROM families WHERE primary_email=?"
	stmt, err := fam.db.Prepare(statement)
	if err != nil {
		return domains.Family{}, err
	}
	defer stmt.Close()

	var family domains.Family
	row := stmt.QueryRow(primary_email)

	errScan := row.Scan(
		&family.Id,
		&family.CreatedAt,
		&family.UpdatedAt,
		&family.DeletedAt,
		&family.PrimaryEmail,
		&family.Password,) 

	return family, errScan
}

func (fam *familyRepo) Insert(family domains.Family) error {
	statement := "INSERT INTO family (" +
		"created_at, " +
		"updated_at, " +
		"primary_email," +
		"password," +
		") VALUES (?, ?, ?, ?)"

	stmt, err := fam.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		family.PrimaryEmail,
		family.Password,
	)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "family was not inserted")
}

func (fam *familyRepo) Update(id uint, family domains.Family) error {
	statement := "UPDATE families SET " +
		"updated_at=?, " +
		"email=?, " +
		"password=? " +
		"WHERE id=?"
	stmt, err := fam.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		family.PrimaryEmail,
		family.Password,
		id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "family was not updated")
}

func (fam *familyRepo) Delete(id uint) error {
	statement := "DELETE FROM family WHERE id=?"
	stmt, err := fam.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "family was not deleted")
}

// For Tests Only
func CreateTestFamilyRepo(db *sql.DB) FamilyRepoInterface {
	fam := &familyRepo{}
	fam.Initialize(db)
	return fam
}
