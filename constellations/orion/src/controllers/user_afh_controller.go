package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetUserAfhByUserId(c *gin.Context) {
	utils.LogControllerMethod(c, "userAfhController.GetUserAfhByUserId")
	userId, err := utils.ParseParamId(c, "userId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("userId")))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	userAfh, err := repos.UserAfhRepo.SelectByUserId(ctx, userId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userAfh)
}

func GetUserAfhByAfhId(c *gin.Context) {
	utils.LogControllerMethod(c, "userAfhController.GetUserAfhByAfhId")
	afhId, err := utils.ParseParamId(c, "afhId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("afhId")))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	userAfh, err := repos.UserAfhRepo.SelectByAfhId(ctx, afhId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userAfh)
}

func GetUserAfhByBothIds(c *gin.Context) {
	utils.LogControllerMethod(c, "userAfhController.GetUserAfhByBothIds")
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

	ctx := utils.RetrieveContext(c)
	userAfh, err := repos.UserAfhRepo.SelectByBothIds(ctx, userId, afhId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &userAfh)
}

func GetUserAfhByNew(c *gin.Context) {
	utils.LogControllerMethod(c, "userAfhController.GetUserAfhByNew")
	ctx := utils.RetrieveContext(c)
	userAfh, err := repos.UserAfhRepo.SelectByNew(ctx)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &userAfh)
	}
}

func CreateUserAfh(c *gin.Context) {
	utils.LogControllerMethod(c, "userAfhController.CreateUserAfh")
	var userAfhJson domains.UserAfh
	if err := c.ShouldBindJSON(&userAfhJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	err := repos.UserAfhRepo.Insert(ctx, userAfhJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateUserAfh(c *gin.Context) {
	utils.LogControllerMethod(c, "userAfhController.UpdateUserAfh")
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

	ctx := utils.RetrieveContext(c)
	if err := repos.UserAfhRepo.Update(ctx, id, userAfhJson); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteUserAfh(c *gin.Context) {
	utils.LogControllerMethod(c, "userAfhController.DeleteUserAfh")
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	if err := repos.UserAfhRepo.Delete(ctx, id); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
