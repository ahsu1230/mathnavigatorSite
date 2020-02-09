package programs

import (
  "net/http"
  "github.com/gin-gonic/gin"
  
  "errors"
  "regexp"
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

  if err := CheckValidProgram(programJson); err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

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

  if err := CheckValidProgram(programJson); err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

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

func CheckValidProgram(program Program) error {
  // Retrieves the inputted values
  programId := program.ProgramId
  name := program.Name
  grade1 := program.Grade1
  grade2 := program.Grade2
  description := program.Description

  // Checks if variables aren't empty
  if programId == "" || name == "" || description == "" {
    return errors.New("One or more fields are empty.")
  }

  // Name validation
  match, _ := regexp.MatchString(`^\s*[A-Z]\S*(\s[A-Z]\S*)*\s*$`, name)
  if !match {
  	return errors.New("Invalid name.")
  }

  // Checks if the grades are valid
  if grade1 > grade2 || grade1 < 1 || grade2 > 12 {
    return errors.New("Invalid grades.")
  }

  return nil
}
