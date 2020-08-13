package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetUserAfhByUserId(c *gin.Context) {
	// Incoming parameters
	userId, err := utils.ParseParamId(c, "userId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("userId")))
		c.Abort()
		return
	}

	userAfh, err := repos.UserAfhRepo.SelectByUserId(userId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userAfh)
}

func GetUserAfhByAfhId(c *gin.Context) {
	// Incoming parameters
	afhId, err := utils.ParseParamId(c, "afhId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("afhId")))
		c.Abort()
		return
	}

	userAfh, err := repos.UserAfhRepo.SelectByAfhId(afhId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userAfh)
}

func GetUserAfhByBothIds(c *gin.Context) {
	// Incoming parameters
	userId, err := utils.ParseParamId(c, "userId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("userId")))
		c.Abort()
		return
	}

	afhId, err := utils.ParseParamId(c, "afhId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("afhId")))
		c.Abort()
		return
	}

	userAfh, err := repos.UserAfhRepo.SelectByBothIds(userId, afhId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userAfh)
}

func CreateUserAfh(c *gin.Context) {
	// Incoming JSON
	var userAfhJson domains.UserAfh
	if err := c.ShouldBindJSON(&userAfhJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	err := repos.UserAfhRepo.Insert(userAfhJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateUserAfh(c *gin.Context) {
	// Incoming JSON & Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	var userAfhJson domains.UserAfh
	if err := c.ShouldBindJSON(&userAfhJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	err := repos.UserAfhRepo.Update(id, userAfhJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteUserAfh(c *gin.Context) {
	// Incoming Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	err := repos.UserAfhRepo.Delete(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
