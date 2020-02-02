package programs

import (
  "database/sql"
  "net/http"

  _ "github.com/lib/pq"
  "github.com/gin-gonic/gin"

  "orion/controllers/utils"
  "orion/database"
)

func GetPrograms(c *gin.Context) {
  // Query DB
  programList := []Program{}
  database.DbConn.Select(&programList, "SELECT * FROM programs")

  // JSON Response
  c.JSON(http.StatusOK, programList)
  return
}

func GetProgram(c *gin.Context) {
  // Incoming parameters
  programId := c.Param("programId")

  // Query DB
  program := Program{}
  sqlStatement := "SELECT * FROM programs WHERE program_id=?"
  err := database.DbConn.Get(&program, sqlStatement, programId)
  if err == sql.ErrNoRows {
    c.String(http.StatusNotFound, "No Program " + programId)
  } else if err != nil {
    panic(err)
  } else {
    c.JSON(http.StatusOK, program)
  }
  return
}

func CreateProgram(c *gin.Context) {
  // Incoming JSON
  var programJson Program
  c.BindJSON(&programJson)
  programId := programJson.ProgramId
  now := utils.TimestampNow()

  // TODO: implement for real!
  // isValid := CheckValidProgram(newProgram);
  // if (isValid) {
  //
  // } else {
  //
  // }

  // Query DB (INSERT & SELECT)
  db := database.DbConn
  sqlStatement := "INSERT INTO programs (created_at, updated_at, deleted_at, program_id, name, grade1, grade2, description) VALUES (:createdAt, :updatedAt, :deletedAt, :programId, :name, :grade1, :grade2, :description)"
  parameters := map[string]interface{}{
    "createdAt": now,
    "updatedAt": now,
    "deletedAt": nil,
    "programId": programId,
    "name": programJson.Name,
    "grade1": programJson.Grade1,
    "grade2": programJson.Grade2,
    "description": programJson.Description,
  }
  _, err := db.NamedExec(sqlStatement, parameters)
  if err != nil {
    panic(err)
  }

  addedProgram := Program{}
  sqlStatement = "SELECT * FROM programs WHERE program_id=?"
  res := db.Get(&addedProgram, sqlStatement, programId)
  if res != nil {
    panic(res)
  } else {
    c.JSON(http.StatusOK, addedProgram)
  }
  return;
}

func UpdateProgram(c *gin.Context) {
  // Incoming JSON & Parameters
  programId := c.Param("programId")
  var programJson Program
  c.BindJSON(&programJson)
  now := utils.TimestampNow()

  // Query DB (UPDATE & SELECT)
  db := database.DbConn
  tx := db.MustBegin()
  sqlStatement := "UPDATE programs SET updated_at=:updatedAt, program_id=:programId, name=:name, grade1=:grade1, grade2=:grade2, description=:description WHERE program_id=:programId"
  parameters := map[string]interface{}{
    "updatedAt": now,
    "programId": programId,
    "name": programJson.Name,
    "grade1": programJson.Grade1,
    "grade2": programJson.Grade2,
    "description": programJson.Description,
  }
  _, err := tx.NamedExec(sqlStatement, parameters)
  if err != nil {
    panic(err)
  }

  updatedProgram := Program{}
  sqlStatement = "SELECT * FROM programs WHERE program_id=?"
  res := tx.Get(&updatedProgram, sqlStatement, programId)
  tx.Commit()
  
  if res != nil {
    panic(res)
  } else {
    c.JSON(http.StatusOK, updatedProgram)
  }
  return
}

func DeleteProgram(c *gin.Context) {
  // Incoming Parameters
  programId := c.Param("programId")

  // Query DB (DELETE)
  sqlStatement := "DELETE FROM programs WHERE program_id=:programId"
  parameters := map[string]interface{}{
    "programId": programId,
  }
  _, err := database.DbConn.NamedExec(sqlStatement, parameters)
  if err != nil {
    panic(err)
  } else {
    c.String(http.StatusOK, "Deleted Program " + programId)
  }
  return
}

func CheckValidProgram() bool {
  return true
}
