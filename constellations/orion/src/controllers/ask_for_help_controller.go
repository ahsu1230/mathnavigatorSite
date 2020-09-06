package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllAFH(c *gin.Context) {
	utils.LogControllerMethod(c, "askForHelpController.GetAllAFH")
	ctx := utils.RetrieveContext(c)
	afhList, err := repos.AskForHelpRepo.SelectAll(ctx)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, afhList)
}

func GetAFHById(c *gin.Context) {
	utils.LogControllerMethod(c, "askForHelpController.GetAFHById")
	// Incoming parameters
	id, _ := utils.ParseParamId(c, "id")

	ctx := utils.RetrieveContext(c)
	askForHelp, err := repos.AskForHelpRepo.SelectById(ctx, id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &askForHelp)
}

func CreateAFH(c *gin.Context) {
	utils.LogControllerMethod(c, "askForHelpController.CreateAFH")
	// Incoming JSON
	var afhJson domains.AskForHelp
	if err := c.ShouldBindJSON(&afhJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := afhJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	id, err := repos.AskForHelpRepo.Insert(ctx, afhJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateAFH(c *gin.Context) {
	utils.LogControllerMethod(c, "askForHelpController.UpdateAFH")
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamId(c, "id")
	var afhJson domains.AskForHelp
	if err := c.ShouldBindJSON(&afhJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := afhJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	err := repos.AskForHelpRepo.Update(ctx, id, afhJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteAFH(c *gin.Context) {
	utils.LogControllerMethod(c, "askForHelpController.DeleteAFH")
	// Incoming Parameters
	id, _ := utils.ParseParamId(c, "id")

	ctx := utils.RetrieveContext(c)
	err := repos.AskForHelpRepo.Delete(ctx, id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func GetAllAFHSubjects(c *gin.Context) {
	utils.LogControllerMethod(c, "askForHelpController.GetAllAFHSubjects")
	c.JSON(http.StatusOK, domains.ALL_AFH_SUBJECTS)
}
