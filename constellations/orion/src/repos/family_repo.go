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
	SelectAll(string, int, int) ([]domains.Family, error)
	SelectById(uint) (domains.Family, error)
	SelectByPrimaryEmail(string) (domains.Family, error)
	Insert(domains.Family) error
	Update(uint, domains.Family) error
	Delete(uint) error
}

func (fam *familyRepo) Initialize(db *sql.DB) {
	fam.db = db
}

func (fam *familyRepo) SelectAll(search string, pageSize, offset int) ([]domains.Family, error) {
	results := make([]domains.Family, 0)

	getAll := len(search) == 0
	var query string
	if getAll {
		query = "SELECT * FROM users LIMIT ? OFFSET ?"
	} else {
		query = "SELECT * FROM users WHERE ? IN (first_name,last_name,middle_name) LIMIT ? OFFSET ?"
	}
	stmt, err := fam.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows *sql.Rows
	if getAll {
		rows, err = stmt.Query(pageSize, offset)
	} else {
		rows, err = stmt.Query(search, pageSize, offset)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domains.Family
		if errScan := rows.Scan(
			&family.Id,
			&family.CreatedAt,
			&family.UpdatedAt,
			&family.DeletedAt,
			&family.PrimaryEmail,
			&family.Password
			); errScan != nil {
			return results, errScan
		}
		results = append(results, family)
	}
	return results, nil
}

func (fam *familyRepo) SelectById(id uint) (domains.Family, error) {
	statement := "SELECT * FROM users WHERE id=?"
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
		&family.Password,
	return family, errScan
}

func (fam *familyRepo) SelectByPrimaryEmail(email string) (domains.Family, error) {
	results := make(domains.Family, 0)

	stmt, err := fam.db.Prepare("SELECT * FROM people WHERE email=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(domains.NewNullString(email))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domains.Family
		if errScan := rows.Scan(
			&family.Id,
			&family.CreatedAt,
			&family.UpdatedAt,
			&family.DeletedAt,
			&family.PrimaryEmail,
			&family.Password; errScan != nil {
			return results, errScan
		}
		results = append(results, family)
	}
	return results, nil
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
		family.Password
	)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "family was not inserted")
}

func (fam *familyRepo) Update(id uint, family domains.Family) error {
	statement := "UPDATE users SET " +
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
		family.Password
		id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "family was not updated")
}

func (fam) *familyRepo) Delete(id uint) error {
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
