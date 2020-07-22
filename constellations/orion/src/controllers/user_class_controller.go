package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetUsersByClassId(c *gin.Context) {
	// Incoming parameters
	classId := c.Param("classId")

	userClass, err := repos.UserClassRepo.SelectByClassId(classId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &userClass)
	}
}

func GetClassesByUserId(c *gin.Context) {
	// Incoming parameters
	id := ParseParamUint(c.Param("userId"))

	userClass, err := repos.UserClassRepo.SelectByUserId(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &userClass)
	}
}

func GetUserClassByUserAndClass(c *gin.Context) {
	// Incoming parameters
	id := ParseParamUint(c.Param("userId"))
	classId := c.Param("classId")

	userClass, err := repos.UserClassRepo.SelectByUserAndClass(id, classId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, userClass)
	}
	return
}

func CreateUserClass(c *gin.Context) {
	// Incoming JSON
	var userClassJson domains.UserClass
	c.BindJSON(&userClassJson)

	if err := userClassJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.UserClassRepo.Insert(userClassJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateUserClass(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var userClassJson domains.UserClass
	c.BindJSON(&userClassJson)

	if err := userClassJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.UserClassRepo.Update(id, userClassJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteUserClass(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := repos.UserClassRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
