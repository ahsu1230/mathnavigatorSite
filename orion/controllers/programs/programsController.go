package programs

import (
  "errors"
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

  // Checks if program name or program ID is empty
  if programId == "" {
    return errors.New("empty Program ID")
  }

  // TODO: use regex
  // Checks if the program ID is alphanumeric
  for _, i := range name {
    if (i < 'a' || i > 'z') && (i < 'A' || i > 'Z') && (i < '1' || i > '0') && i != '_' {
      return errors.New("invalid program name")
    }
  }

  if name == "" {
    return errors.New("empty program name")
  }

  // TODO: Check if program name is alphanumeric and can include spaces, underscores, and ampersands
  // Checks if the program name is alphanumeric
  for _, i := range name {
    if (i < 'a' || i > 'z') && (i < 'A' || i > 'Z') && (i < '1' || i > '0') && i != '_' {
      return errors.New("invalid program name")
    }
  }

  // Checks if the grades are valid
  if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
    return errors.New("invalid grades")
  }

  return nil
}
