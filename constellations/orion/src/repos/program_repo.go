package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

// Global variable
var ProgramRepo ProgramRepoInterface = &programRepo{}

// Implements interface programRepoInterface
type programRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type ProgramRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.Program, error)
	SelectByProgramId(string) (domains.Program, error)
	Insert(domains.Program) error
	Update(string, domains.Program) error
	Delete(string) error
}

func (pr *programRepo) Initialize(db *sql.DB) {
	logger.Debug("Initialize ProgramRepo", logger.Fields{})
	pr.db = db
}

func (pr *programRepo) SelectAll() ([]domains.Program, error) {
	results := make([]domains.Program, 0)

	statement := "SELECT * FROM programs"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement)
	}
	defer rows.Close()

	for rows.Next() {
		var program domains.Program
		if errScan := rows.Scan(
			&program.Id,
			&program.CreatedAt,
			&program.UpdatedAt,
			&program.DeletedAt,
			&program.ProgramId,
			&program.Name,
			&program.Grade1,
			&program.Grade2,
			&program.Description,
			&program.Featured); errScan != nil {
			return results, errScan
		}
		results = append(results, program)
	}
	return results, nil
}

func (pr *programRepo) SelectByProgramId(programId string) (domains.Program, error) {
	statement := "SELECT * FROM programs WHERE program_id=?"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return domains.Program{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var program domains.Program
	row := stmt.QueryRow(programId)
	errScan := row.Scan(
		&program.Id,
		&program.CreatedAt,
		&program.UpdatedAt,
		&program.DeletedAt,
		&program.ProgramId,
		&program.Name,
		&program.Grade1,
		&program.Grade2,
		&program.Description,
		&program.Featured)
	return program, errScan
}

func (pr *programRepo) Insert(program domains.Program) error {
	statement := "INSERT INTO programs (" +
		"created_at, " +
		"updated_at, " +
		"program_id, " +
		"name, " +
		"grade1, " +
		"grade2, " +
		"description, " +
		"featured" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		program.ProgramId,
		program.Name,
		program.Grade1,
		program.Grade2,
		program.Description,
		program.Featured)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, program)
	}
	return appErrors.ValidateDbResult(execResult, 1, "program was not inserted")
}

func (pr *programRepo) Update(programId string, program domains.Program) error {
	statement := "UPDATE programs SET " +
		"updated_at=?, " +
		"program_id=?, " +
		"name=?, " +
		"grade1=?, " +
		"grade2=?, " +
		"description=?, " +
		"featured=? " +
		"WHERE program_id=?"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		program.ProgramId,
		program.Name,
		program.Grade1,
		program.Grade2,
		program.Description,
		program.Featured,
		programId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, program, programId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "program was not updated")
}

func (pr *programRepo) Delete(programId string) error {
	statement := "DELETE FROM programs WHERE program_id=?"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(programId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, programId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "program was not deleted")
}

// For Tests Only
func CreateTestProgramRepo(db *sql.DB) ProgramRepoInterface {
	pr := &programRepo{}
	pr.Initialize(db)
	return pr
}
