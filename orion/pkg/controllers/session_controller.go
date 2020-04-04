package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSessionsByClassId(c *gin.Context) {
	classId := c.Param("classId")

	sessionList, err := services.SessionService.GetAllByClassId(classId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, sessionList)
	}
	return
}

func GetSessionById(c *gin.Context) {
	// Incoming parameters
	id := ParseParamId(c)

	session, err := services.SessionService.GetBySessionId(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &session)
	}
	return
}

func CreateSession(c *gin.Context) {
	// Incoming JSON
	var sessionJson domains.Session
	c.BindJSON(&sessionJson)

	if err := sessionJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.SessionService.Create(sessionJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func UpdateSession(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var sessionJson domains.Session
	c.BindJSON(&sessionJson)

	if err := sessionJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.SessionService.Update(id, sessionJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func DeleteSession(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := services.SessionService.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func GetAllUnpublishedSessions(c *gin.Context) {
	ids, err := services.SessionService.GetAllUnpublished()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, ids)
	}
	return
}

func PublishSessions(c *gin.Context) {
	// Incoming JSON
	var idsJson []uint
	c.BindJSON(&idsJson)

	err := services.SessionService.Publish(idsJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
