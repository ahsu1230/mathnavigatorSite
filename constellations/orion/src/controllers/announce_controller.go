package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllAnnouncements(c *gin.Context) {
	announceList, err := repos.AnnounceRepo.SelectAll()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, announceList)
	}
}

func GetAnnouncementById(c *gin.Context) {
	// Incoming parameters
	id, _ := utils.ParseParamIdString(c, "id")

	announce, err := repos.AnnounceRepo.SelectByAnnounceId(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, announce)
	}
}

func CreateAnnouncement(c *gin.Context) {
	// Incoming JSON
	var announceJson domains.Announce
	c.BindJSON(&announceJson)

	if err := announceJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.AnnounceRepo.Insert(announceJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateAnnouncement(c *gin.Context) {
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamIdString(c, "id")
	var announceJson domains.Announce
	c.BindJSON(&announceJson)

	if err := announceJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.AnnounceRepo.Update(id, announceJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteAnnouncement(c *gin.Context) {
	// Incoming Parameters
	id, _ := utils.ParseParamIdString(c, "id")

	err := repos.AnnounceRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
