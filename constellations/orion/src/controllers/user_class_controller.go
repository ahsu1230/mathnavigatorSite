package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
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
	utils.LogControllerMethod(c, "userClassController.GetUsersByClassId")
	// Incoming parameters
	classId := c.Param("classId")

	userClasses, err := repos.UserClassesRepo.SelectByClassId(classId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func GetClassesByUserId(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.GetClassesByUserId")
	// Incoming parameters
	id, err := utils.ParseParamId(c, "userId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("userId")))
		c.Abort()
		return
	}

	userClasses, err := repos.UserClassesRepo.SelectByUserId(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func GetUserClassByUserAndClass(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.GetUserClassByUserAndClass")
	// Incoming parameters
	id, err := utils.ParseParamId(c, "userId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	classId := c.Param("classId")

	userClasses, err := repos.UserClassesRepo.SelectByUserAndClass(id, classId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func GetNewClasses(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.GetNewClasses")
	userClasses, err := repos.UserClassesRepo.SelectByNew()
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func CreateUserClass(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.CreateUserClass")
	// Incoming JSON
	var userClassesJson domains.UserClasses
	if err := c.ShouldBindJSON(&userClassesJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := userClassesJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.UserClassesRepo.Insert(userClassesJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateUserClass(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.UpdateUserClass")
	// Incoming JSON & Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	var userClassesJson domains.UserClasses
	if err := c.ShouldBindJSON(&userClassesJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := userClassesJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	if err := repos.UserClassesRepo.Update(id, userClassesJson); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteUserClass(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.DeleteUserClass")
	// Incoming Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	if err := repos.UserClassesRepo.Delete(id); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func GetStateValues(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.GetStateValues")
	arr := make([]StateValues, 0)
	// create 3 stateValues variables and push them to arr
	stateValue1 := StateValues{"pending", domains.USER_CLASS_PENDING}
	stateValue2 := StateValues{"accepted", domains.USER_CLASS_ACCEPTED}
	stateValue3 := StateValues{"trial", domains.USER_CLASS_TRIAL}

	arr = append(arr, stateValue1, stateValue2, stateValue3)
	c.JSON(http.StatusOK, arr)
}
