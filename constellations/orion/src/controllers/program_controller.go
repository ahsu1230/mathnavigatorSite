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
	ctx := utils.RetrieveContext(c)
	programList, err := repos.ProgramRepo.SelectAll(ctx)
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

	ctx := utils.RetrieveContext(c)
	program, err := repos.ProgramRepo.SelectByProgramId(ctx, programId)
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

	ctx := utils.RetrieveContext(c)
	id, err := repos.ProgramRepo.Insert(ctx, programJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
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

	ctx := utils.RetrieveContext(c)
	err := repos.ProgramRepo.Update(ctx, programId, programJson)
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

	ctx := utils.RetrieveContext(c)
	err := repos.ProgramRepo.Delete(ctx, programId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func GetAllProgramFeatured(c *gin.Context) {
	utils.LogControllerMethod(c, "programController.GetAllProgramFeatured")
	c.JSON(http.StatusOK, domains.ALL_PROGRAM_FEATURED)
}
