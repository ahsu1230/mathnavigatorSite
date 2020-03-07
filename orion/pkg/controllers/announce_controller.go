package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
)

func parseParamId(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(id)
}

func GetAllAnnouncements(c *gin.Context) {
	announceList, err := services.AnnounceService.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, announceList)
	}
}

func GetAnnouncementById(c *gin.Context) {
	// Incoming parameters
	id := parseParamId(c)

	announce, err := services.AnnounceService.GetByAnnounceId(id)
	if err != nil {
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
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	err := services.AnnounceService.Create(announceJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func UpdateAnnouncement(c *gin.Context) {
	// Incoming JSON & Parameters
	id := parseParamId(c)
	var announceJson domains.Announce
	c.BindJSON(&announceJson)

	if err := announceJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.AnnounceService.Update(id, announceJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteAnnouncement(c *gin.Context) {
	// Incoming Parameters
	id := parseParamId(c)
	
	err := services.AnnounceService.Delete(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}