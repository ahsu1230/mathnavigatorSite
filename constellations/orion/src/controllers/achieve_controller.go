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
	ctx := utils.RetrieveContext(c)
	achieveList, err := repos.AchieveRepo.SelectAll(ctx)
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

	ctx := utils.RetrieveContext(c)
	achieve, err := repos.AchieveRepo.SelectById(ctx, id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &achieve)
}

func GetAllAchievementsGroupedByYear(c *gin.Context) {
	utils.LogControllerMethod(c, "achieveController.GetAllAchievementsGroupedByYear")
	ctx := utils.RetrieveContext(c)
	achieveYearGroup, err := repos.AchieveRepo.SelectAllGroupedByYear(ctx)
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

	ctx := utils.RetrieveContext(c)
	err := repos.AchieveRepo.Insert(ctx, achieveJson)
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

	ctx := utils.RetrieveContext(c)
	err := repos.AchieveRepo.Update(ctx, id, achieveJson)
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

	ctx := utils.RetrieveContext(c)
	err := repos.AchieveRepo.Delete(ctx, id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
