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
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, achieveList)
}

func GetAchievementById(c *gin.Context) {
	// Incoming parameters
	id, _ := utils.ParseParamId(c, "id")

	achieve, err := repos.AchieveRepo.SelectById(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &achieve)
}

func GetAllAchievementsGroupedByYear(c *gin.Context) {
	achieveYearGroup, err := repos.AchieveRepo.SelectAllGroupedByYear()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, achieveYearGroup)
}

func CreateAchievement(c *gin.Context) {
	// Incoming JSON
	var achieveJson domains.Achieve
	if err := c.ShouldBindJSON(&achieveJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := achieveJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err, "Invalid Achievement"))
		c.Abort()
		return
	}

	err := repos.AchieveRepo.Insert(achieveJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateAchievement(c *gin.Context) {
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamId(c, "id")
	var achieveJson domains.Achieve
	if err := c.ShouldBindJSON(&achieveJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := achieveJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err, "Invalid Achievement"))
		c.Abort()
		return
	}

	err := repos.AchieveRepo.Update(id, achieveJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteAchievement(c *gin.Context) {
	// Incoming Parameters
	id, _ := utils.ParseParamId(c, "id")

	err := repos.AchieveRepo.Delete(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
