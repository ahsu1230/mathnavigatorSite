package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllPrograms(c *gin.Context) {
	utils.LogControllerMethod(c, "programController.GetAllPrograms")
	programList, err := repos.ProgramRepo.SelectAll()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, programList)
}

func GetProgramById(c *gin.Context) {
	utils.LogControllerMethod(c, "programController.GetProgramById")
	// Incoming parameters
	programId := c.Param("programId")

	program, err := repos.ProgramRepo.SelectByProgramId(programId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &program)
}

func CreateProgram(c *gin.Context) {
	utils.LogControllerMethod(c, "programController.CreateProgram")
	// Incoming JSON
	var programJson domains.Program
	if err := c.ShouldBindJSON(&programJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := programJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.ProgramRepo.Insert(programJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateProgram(c *gin.Context) {
	utils.LogControllerMethod(c, "programController.UpdateProgram")
	// Incoming JSON & Parameters
	programId := c.Param("programId")
	var programJson domains.Program
	if err := c.ShouldBindJSON(&programJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := programJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.ProgramRepo.Update(programId, programJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteProgram(c *gin.Context) {
	utils.LogControllerMethod(c, "programController.DeleteProgram")
	// Incoming Parameters
	programId := c.Param("programId")

	err := repos.ProgramRepo.Delete(programId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
