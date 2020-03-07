package repos

import (
	"database/sql"
	"time"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
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
	SelectById(string) (domains.Achieve, error)
	Insert(domains.Achieve) error
	Update(string, domains.Achieve) error
	Delete(string) error
}

func (pr *achieveRepo) Initialize(db *sql.DB) {
	pr.db = db
}

func (pr *achieveRepo) SelectAll() ([]domains.Achieve, error) {
	results := make([]domains.Achieve, 0)

	stmt, err := pr.db.Prepare("SELECT * FROM achievements")
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
			&achieve.Year,
			&achieve.Message);
			errScan != nil {
			return results, errScan
		}
		results = append(results, achieve)
	}
	return results, nil
}

func (pr *achieveRepo) SelectById(id string) (domains.Achieve, error) {
	statement := "SELECT * FROM achievements WHERE id=?"
	stmt, err := pr.db.Prepare(statement)
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
		&achieve.Year,
		&achieve.Message)
	return achieve, errScan
}

func (pr *achieveRepo) Insert(achieve domains.Achieve) error {
	statement := "INSERT INTO achievements (" +
		"created_at, " +
		"updated_at, " +
		"year, " +
		"message" +
		") VALUES (?, ?, ?, ?)"

	stmt, err := pr.db.Prepare(statement)
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

func (pr *achieveRepo) Update(id string, achieve domains.Achieve) error {
	statement := "UPDATE achievements SET " +
		"updated_at=?, " +
		"id=?, " +
		"year=?, " +
		"message=? " +
		"WHERE id=?"
	stmt, err := pr.db.Prepare(statement)
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

func (pr *achieveRepo) Delete(id string) error {
	statement := "DELETE FROM achievements WHERE id=?"
	stmt, err := pr.db.Prepare(statement)
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

// For Tests Only
func CreateTestAchieveRepo(db *sql.DB) AchieveRepoInterface {
	pr := &achieveRepo{}
	pr.Initialize(db)
	return pr
}
