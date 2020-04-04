package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAchievements(c *gin.Context) {
	// Incoming optional parameter
	published := c.Query("published")

	achieveList, err := services.AchieveService.GetAll(published == "true")
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, achieveList)
	}
}

func GetAchievementById(c *gin.Context) {
	// Incoming parameters
	id := ParseParamId(c)

	achieve, err := services.AchieveService.GetById(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, achieve)
	}
}

func GetAllAchievementsGroupedByYear(c *gin.Context) {
	achieveYearGroup, err := services.AchieveService.GetAllGroupedByYear()
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

	err := services.AchieveService.Create(achieveJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateAchievement(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var achieveJson domains.Achieve
	c.BindJSON(&achieveJson)

	if err := achieveJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.AchieveService.Update(id, achieveJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteAchievement(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := services.AchieveService.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func PublishAchievements(c *gin.Context) {
	// Incoming JSON
	type Ids struct{
		Ids []uint
	}
	var ids Ids
	c.BindJSON(&ids)

	err := services.AchieveService.Publish(ids.Ids)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
