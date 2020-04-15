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
	SelectAll(bool) ([]domains.Semester, error)
	SelectAllUnpublished() ([]domains.Semester, error)
	SelectBySemesterId(string) (domains.Semester, error)
	Insert(domains.Semester) error
	Update(string, domains.Semester) error
	Publish([]string) error
	Delete(string) error
}

func (sr *semesterRepo) Initialize(db *sql.DB) {
	sr.db = db
}

func (sr *semesterRepo) SelectAll(publishedOnly bool) ([]domains.Semester, error) {
	results := make([]domains.Semester, 0)

	var query string
	if publishedOnly {
		query = "SELECT * FROM semesters WHERE published_at IS NOT NULL"
	} else {
		query = "SELECT * FROM semesters"
	}
	stmt, err := sr.db.Prepare(query)
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
			&semester.PublishedAt,
			&semester.SemesterId,
			&semester.Title); errScan != nil {
			return results, errScan
		}
		results = append(results, semester)
	}
	return results, nil
}

func (sr *semesterRepo) SelectAllUnpublished() ([]domains.Semester, error) {
	results := make([]domains.Semester, 0)

	stmt, err := sr.db.Prepare("SELECT * FROM semesters WHERE published_at IS NULL")
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
			&semester.PublishedAt,
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
		&semester.PublishedAt,
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

func (sr *semesterRepo) Publish(semesterIds []string) error {
	errorList := make([]domains.PublishErrorBody, 0)

	// Begin Transaction
	tx, err := sr.db.Begin()
	stmt, err := tx.Prepare("UPDATE semesters SET published_at=? WHERE id=? AND published_at IS NULL")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	for _, semesterId := range semesterIds {
		_, err := stmt.Exec(now, semesterId)
		if err != nil {
			errorList = append(errorList, domains.PublishErrorBody{StringId: semesterId, Error: err})
		}
	}

	// End Transaction
	tx.Commit()

	return domains.ConcatErrors(errorList)
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
