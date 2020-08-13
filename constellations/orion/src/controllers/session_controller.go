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
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sessionList)
}

func GetSessionById(c *gin.Context) {
	// Incoming parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	session, err := repos.SessionRepo.SelectBySessionId(id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &session)
}

func CreateSessions(c *gin.Context) {
	// Incoming JSON
	var sessionsJson []domains.Session
	if err := c.ShouldBindJSON(&sessionsJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	errs := repos.SessionRepo.Insert(sessionsJson)
	if len(errs) > 0 {
		for _, err := range errs {
			c.Error(err)
		}
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func UpdateSession(c *gin.Context) {
	// Incoming JSON & Parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	var sessionJson domains.Session
	if err := c.ShouldBindJSON(&sessionJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := sessionJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.SessionRepo.Update(id, sessionJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteSessions(c *gin.Context) {
	// Incoming Parameters
	var idsJson []uint
	if err := c.ShouldBindJSON(&idsJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	errs := repos.SessionRepo.Delete(idsJson)
	if len(errs) > 0 {
		for _, err := range errs {
			c.Error(err)
		}
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
