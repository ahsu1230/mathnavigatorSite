package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"time"
)

// Global variable
var AchieveRepo AchieveRepoInterface = &achieveRepo{}

// Implements interface achieveRepoInterface
type achieveRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type AchieveRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.Achieve, error)
	SelectById(uint) (domains.Achieve, error)
	SelectUnpublished() ([]domains.Achieve, error)
	Insert(domains.Achieve) error
	Update(uint, domains.Achieve) error
	Delete(uint) error
	Publish(uint) error
}

func (ar *achieveRepo) Initialize(db *sql.DB) {
	ar.db = db
}

func (ar *achieveRepo) SelectAll() ([]domains.Achieve, error) {
	results := make([]domains.Achieve, 0)

	stmt, err := ar.db.Prepare("SELECT * FROM achievements")
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
		var achieve domains.Achieve
		if errScan := rows.Scan(
			&achieve.Id,
			&achieve.CreatedAt,
			&achieve.UpdatedAt,
			&achieve.DeletedAt,
			&achieve.PublishedAt,
			&achieve.Year,
			&achieve.Message); errScan != nil {
			return results, errScan
		}
		results = append(results, achieve)
	}
	return results, nil
}

func (ar *achieveRepo) SelectById(id uint) (domains.Achieve, error) {
	statement := "SELECT * FROM achievements WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return domains.Achieve{}, err
	}
	defer stmt.Close()

	var achieve domains.Achieve
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&achieve.Id,
		&achieve.CreatedAt,
		&achieve.UpdatedAt,
		&achieve.DeletedAt,
		&achieve.PublishedAt,
		&achieve.Year,
		&achieve.Message)
	return achieve, errScan
}

func (ar *achieveRepo) SelectUnpublished() ([]domains.Achieve, error) {
	results := make([]domains.Achieve, 0)

	stmt, err := ar.db.Prepare("SELECT * FROM achievements WHERE published_at IS NULL")
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
		var achieve domains.Achieve
		if errScan := rows.Scan(
			&achieve.Id,
			&achieve.CreatedAt,
			&achieve.UpdatedAt,
			&achieve.DeletedAt,
			&achieve.PublishedAt,
			&achieve.Year,
			&achieve.Message); errScan != nil {
			return results, errScan
		}
		results = append(results, achieve)
	}
	return results, nil
}

func (ar *achieveRepo) Insert(achieve domains.Achieve) error {
	statement := "INSERT INTO achievements (" +
		"created_at, " +
		"updated_at, " +
		"year, " +
		"message" +
		") VALUES (?, ?, ?, ?)"

	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		achieve.Year,
		achieve.Message)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "achievement was not inserted")
}

func (ar *achieveRepo) Update(id uint, achieve domains.Achieve) error {
	statement := "UPDATE achievements SET " +
		"updated_at=?, " +
		"year=?, " +
		"message=? " +
		"WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		achieve.Year,
		achieve.Message,
		id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "achievement was not updated")
}

func (ar *achieveRepo) Delete(id uint) error {
	statement := "DELETE FROM achievements WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "achievement was not deleted")
}

func (ar *achieveRepo) Publish(id uint) error {
	statement := "UPDATE achievements SET updated_at=? WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(now, id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "achievement was not published")
}

// For Tests Only
func CreateTestAchieveRepo(db *sql.DB) AchieveRepoInterface {
	ar := &achieveRepo{}
	ar.Initialize(db)
	return ar
}
