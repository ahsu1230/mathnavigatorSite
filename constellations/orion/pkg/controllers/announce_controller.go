package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAnnouncements(c *gin.Context) {
	announceList, err := services.AnnounceService.GetAll()
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

	announce, err := services.AnnounceService.GetByAnnounceId(id)
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

	err := services.AnnounceService.Create(announceJson)
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

	err := services.AnnounceService.Update(id, announceJson)
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

	err := services.AnnounceService.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
