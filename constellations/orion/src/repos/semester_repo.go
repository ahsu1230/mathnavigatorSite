package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
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
	utils.LogWithContext("semesterRepo.Initialize", logger.Fields{})
	sr.db = db
}

func (sr *semesterRepo) SelectAll() ([]domains.Semester, error) {
	utils.LogWithContext("semesterRepo.SelectAll", logger.Fields{})
	results := make([]domains.Semester, 0)

	query := "SELECT * FROM semesters ORDER BY ordering ASC"

	stmt, err := sr.db.Prepare(query)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, query)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, query)
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
			&semester.Title,
			&semester.Ordering); errScan != nil {
			return results, errScan
		}
		results = append(results, semester)
	}
	return results, nil
}

func (sr *semesterRepo) SelectBySemesterId(semesterId string) (domains.Semester, error) {
	utils.LogWithContext("semesterRepo.SelectBySemesterId", logger.Fields{"semesterId": semesterId})
	statement := "SELECT * FROM semesters WHERE semester_id=?"
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return domains.Semester{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var semester domains.Semester
	row := stmt.QueryRow(semesterId)
	if err := row.Scan(
		&semester.Id,
		&semester.CreatedAt,
		&semester.UpdatedAt,
		&semester.DeletedAt,
		&semester.SemesterId,
		&semester.Title,
		&semester.Ordering); err != nil {
		return domains.Semester{}, appErrors.WrapDbExec(err, statement, semesterId)
	}
	return semester, nil
}

func (sr *semesterRepo) Insert(semester domains.Semester) error {
	utils.LogWithContext("semesterRepo.Insert", logger.Fields{"semester": semester})
	statement := "INSERT INTO semesters (" +
		"created_at, " +
		"updated_at, " +
		"semester_id, " +
		"title, " +
		"ordering" +
		") VALUES (?, ?, ?, ?, ?)"

	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		semester.SemesterId,
		semester.Title,
		semester.Ordering)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, semester)
	}
	return appErrors.ValidateDbResult(execResult, 1, "semester was not inserted")
}

func (sr *semesterRepo) Update(semesterId string, semester domains.Semester) error {
	utils.LogWithContext("semesterRepo.Update", logger.Fields{"semester": semester})
	statement := "UPDATE semesters SET " +
		"updated_at=?, " +
		"semester_id=?, " +
		"title=?, " +
		"ordering=? " +
		"WHERE semester_id=?"
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		semester.SemesterId,
		semester.Title,
		semester.Ordering,
		semesterId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, semester, semesterId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "semester was not updated")
}

func (sr *semesterRepo) Delete(semesterId string) error {
	utils.LogWithContext("semesterRepo.Delete", logger.Fields{"semesterId": semesterId})
	statement := "DELETE FROM semesters WHERE semester_id=?"
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(semesterId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, semesterId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "semester was not deleted")
}

// For Tests Only
func CreateTestSemesterRepo(db *sql.DB) SemesterRepoInterface {
	sr := &semesterRepo{}
	sr.Initialize(db)
	return sr
}
