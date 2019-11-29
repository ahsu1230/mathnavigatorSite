package controllers
import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetPrograms(c *gin.Context) {
  c.String(http.StatusOK, "Get All Programs")
}

func CreateProgram(c *gin.Context) {
  c.String(http.StatusOK, "Create new Program")
}

func UpdateProgram(c *gin.Context) {
  programId := c.Param("programId")
  c.String(http.StatusOK, "Update Program " + programId)
}

func DeleteProgram(c *gin.Context) {
  programId := c.Param("programId")
  c.String(http.StatusOK, "Delete Program " + programId)
}
