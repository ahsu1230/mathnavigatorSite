package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAchievements(c *gin.Context) {
	achieveList, err := services.AchieveService.GetAll()
	if err != nil {
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
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, achieve)
	}
}

func CreateAchievement(c *gin.Context) {
	// Incoming JSON
	var achieveJson domains.Achieve
	c.BindJSON(&achieveJson)

	if err := achieveJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.AchieveService.Create(achieveJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func UpdateAchievement(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var achieveJson domains.Achieve
	c.BindJSON(&achieveJson)

	if err := achieveJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.AchieveService.Update(id, achieveJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteAchievement(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := services.AchieveService.Delete(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}