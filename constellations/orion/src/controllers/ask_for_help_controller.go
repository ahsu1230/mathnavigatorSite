package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllAFH(c *gin.Context) {
	afhList, err := repos.AskForHelpRepo.SelectAll()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, afhList)
}

func GetAFHById(c *gin.Context) {
	// Incoming parameters
	id, _ := utils.ParseParamId(c, "id")

	askForHelp, err := repos.AskForHelpRepo.SelectById(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &askForHelp)
}

func CreateAFH(c *gin.Context) {
	// Incoming JSON
	var afhJson domains.AskForHelp
	if err := c.ShouldBindJSON(&afhJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := afhJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err, "Invalid AFH"))
		c.Abort()
		return
	}

	err := repos.AskForHelpRepo.Insert(afhJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateAFH(c *gin.Context) {
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamId(c, "id")
	var afhJson domains.AskForHelp
	if err := c.ShouldBindJSON(&afhJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := afhJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err, "Invalid AFH"))
		c.Abort()
		return
	}

	err := repos.AskForHelpRepo.Update(id, afhJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteAFH(c *gin.Context) {
	// Incoming Parameters
	id, _ := utils.ParseParamId(c, "id")

	err := repos.AskForHelpRepo.Delete(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
