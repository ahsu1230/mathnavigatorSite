package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

type UserSearchBody struct {
	Query string `json:"query"`
}

func SearchUsers(c *gin.Context) {
	utils.LogControllerMethod(c, "userController.SearchUsers")

	// Incoming parameters
	var body UserSearchBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	query := body.Query

	user, err := repos.UserRepo.SearchUsers(query)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetUserById(c *gin.Context) {
	utils.LogControllerMethod(c, "userController.GetUserById")

	// Incoming parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	user, err := repos.UserRepo.SelectById(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetUsersByAccountId(c *gin.Context) {
	utils.LogControllerMethod(c, "userController.GetUsersByAccountId")

	// Incoming parameters
	accountId, err := utils.ParseParamId(c, "accountId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("accountId")))
		c.Abort()
		return
	}

	user, err := repos.UserRepo.SelectByAccountId(accountId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetNewUsers(c *gin.Context) {
	utils.LogControllerMethod(c, "userController.GetNewUsers")
	users, err := repos.UserRepo.SelectByNew()
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, &users)
	}
}

func CreateUser(c *gin.Context) {
	utils.LogControllerMethod(c, "userController.CreateUser")

	// Incoming JSON
	var userJson domains.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := userJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	if err := repos.UserRepo.Insert(userJson); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateUser(c *gin.Context) {
	utils.LogControllerMethod(c, "userController.UpdateUser")

	// Incoming JSON & Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	var userJson domains.User
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := userJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	if err := repos.UserRepo.Update(id, userJson); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteUser(c *gin.Context) {
	utils.LogControllerMethod(c, "userController.DeleteUser")

	// Incoming Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	if err := repos.UserRepo.Delete(id); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
