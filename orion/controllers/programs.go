package controllers
import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "orion/models"
)

func GetPrograms(c *gin.Context) {
  results := []models.Program{}
  models.GetDb().Find(&results)

  // somehow remove fields (id, createdAt, updatedAt, deletedAt)
  c.JSON(http.StatusOK, results)
  return;
}

func GetProgram(c *gin.Context) {
  programId := c.Param("programId")

  var foundProgram models.Program
  query := models.GetDb().Where(&models.Program{ProgramId: programId}).First(&foundProgram)
  if query.RecordNotFound() {
    c.String(http.StatusNotFound, "No Program " + programId)
  } else {
    // somehow remove fields (id, createdAt, updatedAt, deletedAt)
    c.JSON(http.StatusOK, foundProgram)
  }
  return;
}

func CreateProgram(c *gin.Context) {
  var foundProgram models.Program
  var newProgram models.Program
  c.BindJSON(&newProgram)

  db := models.GetDb()
  programName := newProgram.Name
  query := db.Where(&models.Program{Name: programName}).First(&foundProgram)
  if query.RecordNotFound() {
    db.Create(&newProgram)
    fmt.Println("New Program created")
    c.JSON(http.StatusOK, gin.H{
      "program": newProgram,
    })
  } else {
    c.String(http.StatusBadRequest, "Program already exists " + programName)
  }
  return;
}

func UpdateProgram(c *gin.Context) {
  var foundProgram models.Program
  var updatedProgram models.Program
  programId := c.Param("programId")
  c.BindJSON(&updatedProgram)

  db := models.GetDb()
  query := db.Where(&models.Program{ProgramId: programId}).First(&foundProgram)
  if query.RecordNotFound() {
    c.String(http.StatusNotFound, "No Program " + programId)
  } else {
    foundProgram.ProgramId = updatedProgram.ProgramId
    foundProgram.Name = updatedProgram.Name
    foundProgram.Grade1 = updatedProgram.Grade1
    foundProgram.Grade2 = updatedProgram.Grade2
    foundProgram.Description = updatedProgram.Description
    db.Save(&foundProgram)
    c.String(http.StatusOK, "Updated Program " + programId)
  }
  return;
}

func DeleteProgram(c *gin.Context) {
  var foundProgram models.Program
  programId := c.Param("programId")
  db := models.GetDb()
  query := db.Where(&models.Program{ProgramId: programId}).First(&foundProgram)
  if query.RecordNotFound() {
    c.String(http.StatusNotFound, "No Program " + programId)
  } else {
    db.Delete(&foundProgram)
    c.String(http.StatusOK, "Deleted Program " + programId)
  }
  return;
}
