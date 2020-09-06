package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllSessionsByClassId(c *gin.Context) {
	utils.LogControllerMethod(c, "sessionController.GetAllSessionsByClassId")
	classId := c.Param("classId")

	ctx := utils.RetrieveContext(c)
	sessionList, err := repos.SessionRepo.SelectAllByClassId(ctx, classId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, sessionList)
}

func GetSessionById(c *gin.Context) {
	utils.LogControllerMethod(c, "sessionController.GetSessionById")
	// Incoming parameters
	id, err := utils.ParseParamId(c, "id")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("id")))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	session, err := repos.SessionRepo.SelectBySessionId(ctx, id)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &session)
}

func CreateSessions(c *gin.Context) {
	utils.LogControllerMethod(c, "sessionController.CreateSession")
	// Incoming JSON
	var sessionsJson []domains.Session
	if err := c.ShouldBindJSON(&sessionsJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	errs := repos.SessionRepo.Insert(ctx, sessionsJson)
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
	utils.LogControllerMethod(c, "sessionController.UpdateSession")
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

	ctx := utils.RetrieveContext(c)
	if err := repos.SessionRepo.Update(ctx, id, sessionJson); err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteSessions(c *gin.Context) {
	utils.LogControllerMethod(c, "sessionController.DeleteSessions")
	// Incoming Parameters
	var idsJson []uint
	if err := c.ShouldBindJSON(&idsJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	errs := repos.SessionRepo.Delete(ctx, idsJson)
	if len(errs) > 0 {
		for _, err := range errs {
			c.Error(err)
		}
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
