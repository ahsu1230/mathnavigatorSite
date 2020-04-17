package controllers

import (
	"errors"
	"fmt"
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

func CreateSessions(c *gin.Context) {
	// Incoming JSON
	var sessionsJson, validatedSessions []domains.Session
	c.BindJSON(&sessionsJson)

	var errorString string
	for _, session := range sessionsJson {
		if err := session.Validate(); err != nil {
			errorString = appendError(errorString, fmt.Sprint(session.Id), err)
		} else {
			validatedSessions = append(validatedSessions, session)
		}
	}

	err := services.SessionService.Create(validatedSessions)
	if err != nil {
		errorString = err.Error() + "\n" + errorString
	}

	if len(errorString) > 0 {
		err := errors.New(errorString)
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func PublishSessions(c *gin.Context) {
	// Incoming JSON
	var idsJson []uint
	c.BindJSON(&idsJson)

	errorList := services.SessionService.Publish(idsJson)
	if len(errorList) > 0 {
		err := domains.Concatenate("one or more sessions failed to publish", errorList, false)
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

func DeleteSessions(c *gin.Context) {
	// Incoming Parameters
	var idsJson []uint
	c.BindJSON(&idsJson)

	err := services.SessionService.Delete(idsJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
