package repos

import (
	"database/sql"
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"time"
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
	SelectAll(bool) ([]domains.Program, error)
	SelectAllUnpublished() ([]domains.Program, error)
	SelectByProgramId(string) (domains.Program, error)
	Insert(domains.Program) error
	Publish([]string) error
	Update(string, domains.Program) error
	Delete(string) error
}

func (pr *programRepo) Initialize(db *sql.DB) {
	pr.db = db
}

func (pr *programRepo) SelectAll(publishedOnly bool) ([]domains.Program, error) {
	results := make([]domains.Program, 0)

	statement := "SELECT * FROM programs"
	if publishedOnly {
		statement += " WHERE published_at IS NOT NULL"
	}
	stmt, err := pr.db.Prepare(statement)
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
		var program domains.Program
		if errScan := rows.Scan(
			&program.Id,
			&program.CreatedAt,
			&program.UpdatedAt,
			&program.DeletedAt,
			&program.PublishedAt,
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

func (pr *programRepo) SelectAllUnpublished() ([]domains.Program, error) {
	results := make([]domains.Program, 0)

	stmt, err := pr.db.Prepare("SELECT * FROM programs WHERE published_at IS NULL")
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
		var program domains.Program
		if errScan := rows.Scan(
			&program.Id,
			&program.CreatedAt,
			&program.UpdatedAt,
			&program.DeletedAt,
			&program.PublishedAt,
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
		return domains.Program{}, err
	}
	defer stmt.Close()

	var program domains.Program
	row := stmt.QueryRow(programId)
	errScan := row.Scan(
		&program.Id,
		&program.CreatedAt,
		&program.UpdatedAt,
		&program.DeletedAt,
		&program.PublishedAt,
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
		return err
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
		return err
	}
	return handleSqlExecResult(execResult, 1, "program was not inserted")
}

func (pr *programRepo) Publish(programIds []string) error {
	var errorString string

	tx, err := pr.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("UPDATE programs SET published_at=? WHERE program_id=? AND published_at IS NULL")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	for _, programId := range programIds {
		execResult, err := stmt.Exec(now, programId)
		if err != nil {
			appendError(errorString, programId, 0, err)
			continue
		}
		err1 := handleSqlExecResult(execResult, 0, "program was not published") // program is already published, 0 rows affected
		err2 := handleSqlExecResult(execResult, 1, "program was not published") // program was not published, 1 row affected
		if err1 != nil && err2 != nil {
			appendError(errorString, programId, 0, err1)
		}
	}
	appendError(errorString, "", 0, tx.Commit())

	if len(errorString) == 0 {
		return nil
	}
	return errors.New(errorString)
}

func (pr *programRepo) Update(programId string, program domains.Program) error {
	statement := "UPDATE programs SET " +
		"updated_at=?, " +
		"published_at=?, " +
		"program_id=?, " +
		"name=?, " +
		"grade1=?, " +
		"grade2=?, " +
		"description=?, " +
		"featured=? " +
		"WHERE program_id=?"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		program.PublishedAt,
		program.ProgramId,
		program.Name,
		program.Grade1,
		program.Grade2,
		program.Description,
		program.Featured,
		programId)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "program was not updated")
}

func (pr *programRepo) Delete(programId string) error {
	statement := "DELETE FROM programs WHERE program_id=?"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(programId)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "program was not deleted")
}

// For Tests Only
func CreateTestProgramRepo(db *sql.DB) ProgramRepoInterface {
	pr := &programRepo{}
	pr.Initialize(db)
	return pr
}
