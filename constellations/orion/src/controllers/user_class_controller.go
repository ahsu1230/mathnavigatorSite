package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StateValues struct {
	Name  string `json:"name"`
	Value uint   `json:"value"`
}

func GetUsersByClassId(c *gin.Context) {
	// Incoming parameters
	classId := c.Param("classId")

	userClasses, err := repos.UserClassesRepo.SelectByClassId(classId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func GetClassesByUserId(c *gin.Context) {
	// Incoming parameters
	id := ParseParamUint(c.Param("userId"))

	userClasses, err := repos.UserClassesRepo.SelectByUserId(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func GetUserClassByUserAndClass(c *gin.Context) {
	// Incoming parameters
	id := ParseParamUint(c.Param("userId"))
	classId := c.Param("classId")

	userClasses, err := repos.UserClassesRepo.SelectByUserAndClass(id, classId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func GetNewClasses(c *gin.Context) {
	userClasses, err := repos.UserClassesRepo.SelectByNew()
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func CreateUserClass(c *gin.Context) {
	// Incoming JSON
	var userClassesJson domains.UserClasses
	c.BindJSON(&userClassesJson)

	if err := userClassesJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.UserClassesRepo.Insert(userClassesJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func UpdateUserClass(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var userClassesJson domains.UserClasses
	c.BindJSON(&userClassesJson)

	if err := userClassesJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.UserClassesRepo.Update(id, userClassesJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func DeleteUserClass(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := repos.UserClassesRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func GetStateValues(c *gin.Context) {
	arr := make([]StateValues, 0)
	// create 3 stateValues variables and push them to arr
	stateValue1 := StateValues{"pending", 0}
	stateValue2 := StateValues{"accepted", 1}
	stateValue3 := StateValues{"trial", 2}

	arr = append(arr, stateValue1, stateValue2, stateValue3)
	c.JSON(http.StatusOK, arr)
}
