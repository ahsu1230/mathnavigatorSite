package repos

import (
	"database/sql"
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
	SelectAll() ([]domains.Program, error)
	SelectByProgramId(string) (domains.Program, error)
	Insert(domains.Program) error
	Update(string, domains.Program) error
	Delete(string) error
	SelectAllUnpublished() ([]string, error)
	Publish([]string) error
}

func (pr *programRepo) Initialize(db *sql.DB) {
	pr.db = db
}

func (pr *programRepo) SelectAll() ([]domains.Program, error) {
	results := make([]domains.Program, 0)

	stmt, err := pr.db.Prepare("SELECT * FROM programs")
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
			&program.ProgramId,
			&program.Name,
			&program.Grade1,
			&program.Grade2,
			&program.Description,
			&program.PublishedAt); errScan != nil {
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
		&program.ProgramId,
		&program.Name,
		&program.Grade1,
		&program.Grade2,
		&program.Description,
		&program.PublishedAt)
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
		"description" +
		") VALUES (?, ?, ?, ?, ?, ?, ?)"

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
		program.Description)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "program was not inserted")
}

func (pr *programRepo) Update(programId string, program domains.Program) error {
	statement := "UPDATE programs SET " +
		"updated_at=?, " +
		"program_id=?, " +
		"name=?, " +
		"grade1=?, " +
		"grade2=?, " +
		"description=?, " +
		"published_at=? " +
		"WHERE program_id=?"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return err
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
		program.PublishedAt,
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

func (pr *programRepo) SelectAllUnpublished() ([]string, error) {
	results := make([]string, 0)

	stmt, err := pr.db.Prepare("SELECT program_id FROM programs WHERE published_at IS NULL")
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
		var programId string
		if errScan := rows.Scan(&programId); errScan != nil {
			return results, errScan
		}
		results = append(results, programId)
	}
	return results, nil
}

func (pr *programRepo) Publish(programIds []string) error {
	for _, programId := range programIds {
		program, err := pr.SelectByProgramId(programId)
		if err != nil {
			return err
		}
		if !program.PublishedAt.Valid {
			now := time.Now().UTC()
			program.PublishedAt.Scan(now)
			pr.Update(programId, program)
		}
	}
	return nil
}

// For Tests Only
func CreateTestProgramRepo(db *sql.DB) ProgramRepoInterface {
	pr := &programRepo{}
	pr.Initialize(db)
	return pr
}
