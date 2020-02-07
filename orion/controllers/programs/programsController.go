package programs

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
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

  // TODO: test if this works
  if CheckValidProgram(c) == false {
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

func CheckValidProgram(c *gin.Context) bool {
  // Retrieves the inputted values
  name := c.Param("programName")
  grade1 := c.Param("grade1")
  grade2 := c.Param("grade2")

  // Checks if program name is empty
  if name == "" {
    c.String(http.StatusBadRequest, "Invalid name " + name)
    return false
  }

  // Checks if the program name is a string only
  for _, i := range name {
    if (i < 'a' || i > 'z') && (i < 'A' || i > 'Z') {
      c.String(http.StatusBadRequest, "Invalid name " + name)
      return false
    }
  }

  // Checks if the grades are integers
  _, err1 := strconv.Atoi(grade1)
  _, err2 := strconv.Atoi(grade2)

  if err1 != nil || err2 != nil {
    c.String(http.StatusBadRequest, "Invalid grades " + grade1 + " to " + grade2)
    return false
  }

  // Checks if the grades are valid
  if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
    c.String(http.StatusBadRequest, "Invalid grades " + grade1 + " to " + grade2)
    return false
  }

  return true
}
