package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsersByClassId(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.GetUsersByClassId")
	// Incoming parameters
	classId := c.Param("classId")

	ctx := utils.RetrieveContext(c)
	userClasses, err := repos.UserClassRepo.SelectByClassId(ctx, classId)
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

	ctx := utils.RetrieveContext(c)
	userClasses, err := repos.UserClassRepo.SelectByUserId(ctx, id)
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

	ctx := utils.RetrieveContext(c)
	classId := c.Param("classId")
	userClasses, err := repos.UserClassRepo.SelectByUserAndClass(ctx, id, classId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userClasses)
}

func GetNewClasses(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.GetNewClasses")
	ctx := utils.RetrieveContext(c)
	userClasses, err := repos.UserClassRepo.SelectByNew(ctx)
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
	var userClassesJson domains.UserClass
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

	ctx := utils.RetrieveContext(c)
	id, err := repos.UserClassRepo.Insert(ctx, userClassesJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
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

	var userClassesJson domains.UserClass
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

	ctx := utils.RetrieveContext(c)
	if err := repos.UserClassRepo.Update(ctx, id, userClassesJson); err != nil {
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

	ctx := utils.RetrieveContext(c)
	if err := repos.UserClassRepo.Delete(ctx, id); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func GetStateValues(c *gin.Context) {
	utils.LogControllerMethod(c, "userClassController.GetStateValues")
	c.JSON(http.StatusOK, domains.ALL_USER_CLASS_STATES)
}
