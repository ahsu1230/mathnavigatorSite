package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllAchievements(c *gin.Context) {
	utils.LogControllerMethod(c, "achieveController.GetAllAchievements")
	achieveList, err := repos.AchieveRepo.SelectAll()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, achieveList)
}

func GetAchievementById(c *gin.Context) {
	utils.LogControllerMethod(c, "achieveController.GetAchievementById")
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
	utils.LogControllerMethod(c, "achieveController.GetAllAchievementsGroupedByYear")
	achieveYearGroup, err := repos.AchieveRepo.SelectAllGroupedByYear()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, achieveYearGroup)
}

func CreateAchievement(c *gin.Context) {
	utils.LogControllerMethod(c, "achieveController.CreateAchievement")
	// Incoming JSON
	var achieveJson domains.Achieve
	if err := c.ShouldBindJSON(&achieveJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := achieveJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.AchieveRepo.Insert(achieveJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateAchievement(c *gin.Context) {
	utils.LogControllerMethod(c, "achieveController.UpdateAchievement")
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamId(c, "id")
	var achieveJson domains.Achieve
	if err := c.ShouldBindJSON(&achieveJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := achieveJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.AchieveRepo.Update(id, achieveJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteAchievement(c *gin.Context) {
	utils.LogControllerMethod(c, "achieveController.DeleteAchievement")
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
