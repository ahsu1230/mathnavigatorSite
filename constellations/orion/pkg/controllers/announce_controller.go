package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
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
	id := ParseParamId(c)

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
	id := ParseParamId(c)
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
	id := ParseParamId(c)

	err := repos.AnnounceRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
