package store

import (
    "database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
    "github.com/ahsu1230/mathnavigatorSite/orion/domains"
    "github.com/ahsu1230/mathnavigatorSite/orion/utils"
)

type ProgramService struct {
    DbSql *sql.DB       // golang native db connection
    DbSqlx *sqlx.DB     // sqlx wrapper over db connection
}

func GetAllPrograms(dbx *sqlx.DB) ([]domains.Program, error) {
    programList := []domains.Program{}
	err := dbx.Select(&programList, "SELECT * FROM programs")
	return programList, err
}

func GetProgramById(dbx *sqlx.DB, programId string) (domains.Program, error) {
    program := domains.Program{}
	sqlStatement := "SELECT * FROM programs WHERE program_id=?"
	err := dbx.Get(&program, sqlStatement, programId)
	return program, err
}

func InsertProgram(dbx *sqlx.DB, newProgram domains.Program) error {
    programId := newProgram.ProgramId
	now := utils.TimestampNow()
	sqlStatement := "INSERT INTO programs " +
		"(created_at, updated_at, program_id, name, grade1, grade2, description) " +
		"VALUES (:createdAt, :updatedAt, :programId, :name, :grade1, :grade2, :description)"
	parameters := map[string]interface{}{
		"createdAt":   now,
		"updatedAt":   now,
		"programId":   programId,
		"name":        newProgram.Name,
		"grade1":      newProgram.Grade1,
		"grade2":      newProgram.Grade2,
		"description": newProgram.Description,
	}
	_, err := dbx.NamedExec(sqlStatement, parameters)
	return err
}

func UpdateProgram(dbx *sqlx.DB, oldProgramId string, newProgram domains.Program) error {
    now := utils.TimestampNow()
	sqlStatement := "UPDATE programs SET " +
		"updated_at=:updatedAt, " +
		"name=:name, " +
		"program_id=:programId, " +
		"grade1=:grade1, " +
		"grade2=:grade2, " +
		"description=:description " +
		"WHERE program_id=:oldProgramId"
	parameters := map[string]interface{}{
		"updatedAt":    now,
		"programId":    newProgram.ProgramId,
		"name":         newProgram.Name,
		"grade1":       newProgram.Grade1,
		"grade2":       newProgram.Grade2,
		"description":  newProgram.Description,
		"oldProgramId": oldProgramId,
	}
	_, err := dbx.NamedExec(sqlStatement, parameters)
	return err
}

func DeleteProgram(dbx *sqlx.DB, programId string) error {
    sqlStatement := "DELETE FROM programs WHERE program_id=:programId"
	parameters := map[string]interface{}{
		"programId": programId,
	}
	_, err := dbx.NamedExec(sqlStatement, parameters)
	return err
}
