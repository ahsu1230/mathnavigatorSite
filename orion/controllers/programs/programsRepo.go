package programs

import (
  "orion/controllers/utils"
  "orion/database"
)

func GetAllPrograms() []Program {
  programList := []Program{}
  database.DbConn.Select(&programList, "SELECT * FROM programs")
  return programList
}

func GetProgramById(programId string) (Program, error) {
  program := Program{}
  sqlStatement := "SELECT * FROM programs WHERE program_id=?"
  err := database.DbConn.Get(&program, sqlStatement, programId)
  return program, err
}

func InsertProgram(program Program) (error) {
  programId := program.ProgramId
  now := utils.TimestampNow()
  db := database.DbConn
  sqlStatement := "INSERT INTO programs " +
        "(created_at, updated_at, deleted_at, program_id, name, " +
          "grade1, grade2, description) " +
        "VALUES (:createdAt, :updatedAt, :deletedAt, :programId, :name, " +
          ":grade1, :grade2, :description)"
  parameters := map[string]interface{}{
    "createdAt": now,
    "updatedAt": now,
    "deletedAt": nil,
    "programId": programId,
    "name": program.Name,
    "grade1": program.Grade1,
    "grade2": program.Grade2,
    "description": program.Description,
  }
  _, err := db.NamedExec(sqlStatement, parameters)
  return err
}

func UpdateProgramById(oldProgramId string, program Program) (error) {
  now := utils.TimestampNow()
  db := database.DbConn

  sqlStatement := "UPDATE programs SET " +
        "updated_at=:updatedAt, " +
        "name=:name, " +
        "program_id=:programId, " +
        "grade1=:grade1, " +
        "grade2=:grade2, " +
        "description=:description " +
        "WHERE program_id=:oldProgramId"
  parameters := map[string]interface{}{
    "updatedAt": now,
    "programId": program.ProgramId,
    "name": program.Name,
    "grade1": program.Grade1,
    "grade2": program.Grade2,
    "description": program.Description,
    "oldProgramId": oldProgramId,
  }
  _, err := db.NamedExec(sqlStatement, parameters)
  return err
}

func DeleteProgramById(programId string) error {
  sqlStatement := "DELETE FROM programs WHERE program_id=:programId"
  parameters := map[string]interface{}{
    "programId": programId,
  }
  _, err := database.DbConn.NamedExec(sqlStatement, parameters)
  return err
}
