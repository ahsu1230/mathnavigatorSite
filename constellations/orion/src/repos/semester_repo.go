package repos

import (
	"context"
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

// Global variable
var SemesterRepo SemesterRepoInterface = &semesterRepo{}

// Implements interface semesterRepoInterface
type semesterRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type SemesterRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectAll(context.Context) ([]domains.Semester, error)
	SelectBySemesterId(context.Context, string) (domains.Semester, error)
	Insert(context.Context, domains.Semester) error
	Update(context.Context, string, domains.Semester) error
	Delete(context.Context, string) error
}

func (sr *semesterRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "semesterRepo.Initialize", logger.Fields{})
	sr.db = db
}

func (sr *semesterRepo) SelectAll(ctx context.Context) ([]domains.Semester, error) {
	utils.LogWithContext(ctx, "semesterRepo.SelectAll", logger.Fields{})
	results := make([]domains.Semester, 0)

	query := "SELECT * FROM semesters ORDER BY year ASC, " +
		"CASE WHEN SEASON='winter' THEN '0' " +
		"WHEN SEASON='spring' THEN '1' " +
		"WHEN SEASON='summer' THEN '2' " +
		"WHEN SEASON='fall' THEN '3' " +
		"END ASC"

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
			&semester.Season,
			&semester.Year,
			&semester.Title); errScan != nil {
			return results, errScan
		}
		results = append(results, semester)
	}
	return results, nil
}

func (sr *semesterRepo) SelectBySemesterId(ctx context.Context, semesterId string) (domains.Semester, error) {
	utils.LogWithContext(ctx, "semesterRepo.SelectBySemesterId", logger.Fields{"semesterId": semesterId})
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
		&semester.Season,
		&semester.Year,
		&semester.Title); err != nil {
		return domains.Semester{}, appErrors.WrapDbExec(err, statement, semesterId)
	}
	return semester, nil
}

func (sr *semesterRepo) Insert(ctx context.Context, semester domains.Semester) error {
	utils.LogWithContext(ctx, "semesterRepo.Insert", logger.Fields{"semester": semester})
	statement := "INSERT INTO semesters (" +
		"created_at, " +
		"updated_at, " +
		"semester_id, " +
		"season, " +
		"year, " +
		"title" +
		") VALUES (?, ?, ?, ?, ?, ?)"

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
		semester.Season,
		semester.Year,
		semester.Title)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, semester)
	}
	return appErrors.ValidateDbResult(execResult, 1, "semester was not inserted")
}

func (sr *semesterRepo) Update(ctx context.Context, semesterId string, semester domains.Semester) error {
	utils.LogWithContext(ctx, "semesterRepo.Update", logger.Fields{"semester": semester})
	statement := "UPDATE semesters SET " +
		"updated_at=?, " +
		"semester_id=?, " +
		"season=?, " +
		"year=?, " +
		"title=? " +
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
		semester.Season,
		semester.Year,
		semester.Title,
		semesterId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, semester, semesterId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "semester was not updated")
}

func (sr *semesterRepo) Delete(ctx context.Context, semesterId string) error {
	utils.LogWithContext(ctx, "semesterRepo.Delete", logger.Fields{"semesterId": semesterId})
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
func CreateTestSemesterRepo(ctx context.Context, db *sql.DB) SemesterRepoInterface {
	sr := &semesterRepo{}
	sr.Initialize(ctx, db)
	return sr
}
