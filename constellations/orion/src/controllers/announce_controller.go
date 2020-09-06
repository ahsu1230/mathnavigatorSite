package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllAnnouncements(c *gin.Context) {
	utils.LogControllerMethod(c, "announceController.GetAllAnnouncements")
	ctx := utils.RetrieveContext(c)
	announceList, err := repos.AnnounceRepo.SelectAll(ctx)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, announceList)
}

func GetAnnouncementById(c *gin.Context) {
	utils.LogControllerMethod(c, "announceController.GetAnnouncementById")
	// Incoming parameters
	id, _ := utils.ParseParamId(c, "id")

	ctx := utils.RetrieveContext(c)
	announce, err := repos.AnnounceRepo.SelectByAnnounceId(ctx, id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, announce)
}

func CreateAnnouncement(c *gin.Context) {
	utils.LogControllerMethod(c, "announceController.CreateAnnouncement")
	// Incoming JSON
	var announceJson domains.Announce
	if err := c.ShouldBindJSON(&announceJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := announceJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	err := repos.AnnounceRepo.Insert(ctx, announceJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateAnnouncement(c *gin.Context) {
	utils.LogControllerMethod(c, "announceController.UpdateAnnouncement")
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamId(c, "id")
	var announceJson domains.Announce
	if err := c.ShouldBindJSON(&announceJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := announceJson.Validate(); err != nil {
		err = appErrors.WrapInvalidDomain(err.Error())
		c.Error(err)
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	err := repos.AnnounceRepo.Update(ctx, id, announceJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteAnnouncement(c *gin.Context) {
	utils.LogControllerMethod(c, "announceController.DeleteAnnouncement")
	// Incoming Parameters
	id, _ := utils.ParseParamId(c, "id")

	ctx := utils.RetrieveContext(c)
	err := repos.AnnounceRepo.Delete(ctx, id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
