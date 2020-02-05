package programs

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetPrograms(c *gin.Context) {
  // Query Repo
  programList := GetAllPrograms()

  // JSON Response
  c.JSON(http.StatusOK, programList)
  return
}

func GetProgram(c *gin.Context) {
  // Incoming parameters
  programId := c.Param("programId")

  // Query Repo
  program, err := GetProgramById(programId)
  if err != nil {
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

  // TODO: implement for real!
  // isValid := CheckValidProgram(newProgram);
  // if (isValid) {
  //
  // } else {
  //
  // }

  // Query Repo (INSERT & SELECT)
  err := InsertProgram(programJson)
  if err != nil {
    panic(err)
  } else {
    c.JSON(http.StatusOK, nil)
  }
  return
}

func UpdateProgram(c *gin.Context) {
  // Incoming JSON & Parameters
  programId := c.Param("programId")
  var programJson Program
  c.BindJSON(&programJson)

  // Query Repo (UPDATE & SELECT)
  err := UpdateProgramById(programId, programJson)
  if err != nil {
    panic(err)
  } else {
    c.JSON(http.StatusOK, nil)
  }
  return
}

func DeleteProgram(c *gin.Context) {
  // Incoming Parameters
  programId := c.Param("programId")

  // Query Repo (DELETE)
  err := DeleteProgramById(programId)
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
