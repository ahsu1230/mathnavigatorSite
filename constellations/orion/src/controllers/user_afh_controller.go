package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetUserAfhByUserId(c *gin.Context) {
	// Incoming parameters
	userId := ParseParamUserId(c)

	userAfh, err := repos.UserAfhRepo.SelectByUserId(userId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &userAfh)
	}
}

func GetUserAfhByAfhId(c *gin.Context) {
	// Incoming parameters
	afhId := ParseParamAfhId(c)

	userAfh, err := repos.UserAfhRepo.SelectByAfhId(afhId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &userAfh)
	}
}

func GetUserAfhByBothIds(c *gin.Context) {
	// Incoming parameters
	userId := ParseParamUserId(c)
	afhId := ParseParamAfhId(c)

	userAfh, err := repos.UserAfhRepo.SelectByBothIds(userId, afhId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &userAfh)
	}
}

func GetUserAfhByNew(c *gin.Context) {
	userAfh, err := repos.UserAfhRepo.SelectByNew()
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &userAfh)
	}
}

func CreateUserAfh(c *gin.Context) {
	// Incoming JSON
	var userAfhJson domains.UserAfh
	c.BindJSON(&userAfhJson)

	err := repos.UserAfhRepo.Insert(userAfhJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateUserAfh(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var userAfhJson domains.UserAfh
	c.BindJSON(&userAfhJson)

	err := repos.UserAfhRepo.Update(id, userAfhJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteUserAfh(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := repos.UserAfhRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
