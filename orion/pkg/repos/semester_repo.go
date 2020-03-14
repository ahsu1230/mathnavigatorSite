package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"time"
)

// Global variable
var SemesterRepo SemesterRepoInterface = &semesterRepo{}

// Implements interface semesterRepoInterface
type semesterRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type SemesterRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.Semester, error)
	SelectBySemesterId(string) (domains.Semester, error)
	Insert(domains.Semester) error
	Update(string, domains.Semester) error
	Delete(string) error
}

func (sr *semesterRepo) Initialize(db *sql.DB) {
	sr.db = db
}

func (sr *semesterRepo) SelectAll() ([]domains.Semester, error) {
	results := make([]domains.Semester, 0)

	stmt, err := sr.db.Prepare("SELECT * FROM semesters")
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
		var semester domains.Semester
		if errScan := rows.Scan(
			&semester.Id,
			&semester.CreatedAt,
			&semester.UpdatedAt,
			&semester.DeletedAt,
			&semester.SemesterId,
			&semester.Title); errScan != nil {
			return results, errScan
		}
		results = append(results, semester)
	}
	return results, nil
}

func (sr *semesterRepo) SelectBySemesterId(semesterId string) (domains.Semester, error) {
	statement := "SELECT * FROM semesters WHERE semester_id=?"
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return domains.Semester{}, err
	}
	defer stmt.Close()

	var semester domains.Semester
	row := stmt.QueryRow(semesterId)
	errScan := row.Scan(
		&semester.Id,
		&semester.CreatedAt,
		&semester.UpdatedAt,
		&semester.DeletedAt,
		&semester.SemesterId,
		&semester.Title)
	return semester, errScan
}

func (sr *semesterRepo) Insert(semester domains.Semester) error {
	statement := "INSERT INTO semesters (" +
		"created_at, " +
		"updated_at, " +
		"semester_id, " +
		"title" +
		") VALUES (?, ?, ?, ?)"

	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		semester.SemesterId,
		semester.Title)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "semester was not inserted")
}

func (sr *semesterRepo) Update(semesterId string, semester domains.Semester) error {
	statement := "UPDATE semesters SET " +
		"updated_at=?, " +
		"semester_id=?, " +
		"name=?, " +
		"title=? " +
		"WHERE semester_id=?"
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		semester.SemesterId,
		semester.Title,
		semesterId)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "semester was not updated")
}

func (sr *semesterRepo) Delete(semesterId string) error {
	statement := "DELETE FROM semesters WHERE semester_id=?"
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(semesterId)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "semester was not deleted")
}

// For Tests Only
func CreateTestSemesterRepo(db *sql.DB) SemesterRepoInterface {
	sr := &semesterRepo{}
	sr.Initialize(db)
	return sr
}