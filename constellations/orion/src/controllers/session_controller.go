package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllSessionsByClassId(c *gin.Context) {
	classId := c.Param("classId")

	sessionList, err := repos.SessionRepo.SelectAllByClassId(classId)
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
	id, _ := utils.ParseParamId(c, "id")

	session, err := repos.SessionRepo.SelectBySessionId(id)
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
	var sessionsJson []domains.Session
	c.BindJSON(&sessionsJson)

	errs := repos.SessionRepo.Insert(sessionsJson)
	if len(errs) > 0 {
		for _, err := range errs {
			c.Error(err)
		}
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateSession(c *gin.Context) {
	// Incoming JSON & Parameters
	id, _ := utils.ParseParamId(c, "id")
	var sessionJson domains.Session
	c.BindJSON(&sessionJson)

	if err := sessionJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.SessionRepo.Update(id, sessionJson)
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

	errs := repos.SessionRepo.Delete(idsJson)
	if len(errs) > 0 {
		for _, err := range errs {
			c.Error(err)
		}
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}
