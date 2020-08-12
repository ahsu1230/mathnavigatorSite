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
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, afhList)
	}
}

func GetAFHById(c *gin.Context) {
	// Incoming parameters
	id, _ := utils.ParseParamIdString(c, "id")

	askForHelp, err := repos.AskForHelpRepo.SelectById(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &askForHelp)
	}
}

func CreateAFH(c *gin.Context) {
	// Incoming JSON
	var afhJson domains.AskForHelp
	c.BindJSON(&afhJson)

	if err := afhJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.AskForHelpRepo.Insert(afhJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateAFH(c *gin.Context) {
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamIdString(c, "id")
	var afhJson domains.AskForHelp
	c.BindJSON(&afhJson)

	if err := afhJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.AskForHelpRepo.Update(id, afhJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteAFH(c *gin.Context) {
	// Incoming Parameters
	id, _ := utils.ParseParamIdString(c, "id")

	err := repos.AskForHelpRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
