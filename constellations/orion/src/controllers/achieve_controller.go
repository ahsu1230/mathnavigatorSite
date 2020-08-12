package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllAchievements(c *gin.Context) {
	achieveList, err := repos.AchieveRepo.SelectAll()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, achieveList)
	}
}

func GetAchievementById(c *gin.Context) {
	// Incoming parameters
	id, _ := utils.ParseParamId(c, "id")

	achieve, err := repos.AchieveRepo.SelectById(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &achieve)
	}
}

func GetAllAchievementsGroupedByYear(c *gin.Context) {
	achieveYearGroup, err := repos.AchieveRepo.SelectAllGroupedByYear()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, achieveYearGroup)
	}
}

func CreateAchievement(c *gin.Context) {
	// Incoming JSON
	var achieveJson domains.Achieve
	c.BindJSON(&achieveJson)

	if err := achieveJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.AchieveRepo.Insert(achieveJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateAchievement(c *gin.Context) {
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamId(c, "id")
	var achieveJson domains.Achieve
	c.BindJSON(&achieveJson)

	if err := achieveJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.AchieveRepo.Update(id, achieveJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteAchievement(c *gin.Context) {
	// Incoming Parameters
	id, _ := utils.ParseParamId(c, "id")

	err := repos.AchieveRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
