package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSessionsByClassId(c *gin.Context) {
	classId := c.Param("classId")
	publishedOnly := ParseParamPublishedOnly(c)

	sessionList, err := services.SessionService.GetAllByClassId(classId, publishedOnly)
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

func PublishSessions(c *gin.Context) {
	// Incoming JSON
	var idsJson []uint
	c.BindJSON(&idsJson)

	errors := services.SessionService.Publish(idsJson)
	if len(errors) > 0 {
		for _, err := range errors {
			c.Error(err.Error)
			c.String(http.StatusInternalServerError, err.Error.Error())
		}
	} else {
		c.Status(http.StatusOK)
	}
	return
}
