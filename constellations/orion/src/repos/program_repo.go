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
var ProgramRepo ProgramRepoInterface = &programRepo{}

// Implements interface programRepoInterface
type programRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type ProgramRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectAll(context.Context) ([]domains.Program, error)
	SelectByProgramId(context.Context, string) (domains.Program, error)
	Insert(context.Context, domains.Program) (uint, error)
	Update(context.Context, string, domains.Program) error
	Archive(context.Context, string) error
	Delete(context.Context, string) error
}

func (pr *programRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "programRepo.Initialize", logger.Fields{})
	pr.db = db
}

func (pr *programRepo) SelectAll(ctx context.Context) ([]domains.Program, error) {
	utils.LogWithContext(ctx, "programRepo.SelectAll", logger.Fields{})
	results := make([]domains.Program, 0)

	statement := "SELECT * FROM programs WHERE deleted_at IS NULL"
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
			&program.Title,
			&program.Grade1,
			&program.Grade2,
			&program.Subject,
			&program.Description,
			&program.Featured); errScan != nil {
			return results, errScan
		}
		results = append(results, program)
	}
	return results, nil
}

func (pr *programRepo) SelectByProgramId(ctx context.Context, programId string) (domains.Program, error) {
	utils.LogWithContext(ctx, "programRepo.SelectByProgramId", logger.Fields{"programId": programId})
	statement := "SELECT * FROM programs WHERE program_id=? AND deleted_at IS NULL"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return domains.Program{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var program domains.Program
	row := stmt.QueryRow(programId)
	if err = row.Scan(
		&program.Id,
		&program.CreatedAt,
		&program.UpdatedAt,
		&program.DeletedAt,
		&program.ProgramId,
		&program.Title,
		&program.Grade1,
		&program.Grade2,
		&program.Subject,
		&program.Description,
		&program.Featured); err != nil {
		return domains.Program{}, appErrors.WrapDbExec(err, statement, programId)
	}
	return program, nil
}

func (pr *programRepo) Insert(ctx context.Context, program domains.Program) (uint, error) {
	utils.LogWithContext(ctx, "programRepo.Insert", logger.Fields{"program": program})
	statement := "INSERT INTO programs (" +
		"created_at, " +
		"updated_at, " +
		"program_id, " +
		"title, " +
		"grade1, " +
		"grade2, " +
		"subject, " +
		"description, " +
		"featured" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return 0, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		program.ProgramId,
		program.Title,
		program.Grade1,
		program.Grade2,
		program.Subject,
		program.Description,
		program.Featured)
	if err != nil {
		return 0, appErrors.WrapDbExec(err, statement, program)
	}

	rowId, err := execResult.LastInsertId()
	if err != nil {
		return 0, appErrors.WrapSQLBadInsertResult(err)
	}

	return uint(rowId), appErrors.ValidateDbResult(execResult, 1, "program was not inserted")
}

func (pr *programRepo) Update(ctx context.Context, programId string, program domains.Program) error {
	utils.LogWithContext(ctx, "programRepo.Update", logger.Fields{"programId": programId, "program": program})
	statement := "UPDATE programs SET " +
		"updated_at=?, " +
		"program_id=?, " +
		"title=?, " +
		"grade1=?, " +
		"grade2=?, " +
		"subject=?, " +
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
		program.Title,
		program.Grade1,
		program.Grade2,
		program.Subject,
		program.Description,
		program.Featured,
		programId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, program, programId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "program was not updated")
}

func (pr *programRepo) Archive(ctx context.Context, programId string) error {
	utils.LogWithContext(ctx, "programRepo.Archive", logger.Fields{"programId": programId})
	statement := "UPDATE programs SET deleted_at=? WHERE program_id=?"
	stmt, err := pr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(now, programId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, programId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "program was not deleted")
}

func (pr *programRepo) Delete(ctx context.Context, programId string) error {
	utils.LogWithContext(ctx, "programRepo.Delete", logger.Fields{"programId": programId})
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
func CreateTestProgramRepo(ctx context.Context, db *sql.DB) ProgramRepoInterface {
	pr := &programRepo{}
	pr.Initialize(ctx, db)
	return pr
}
