package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/repos"
	"github.com/gin-gonic/gin"
)

func GetAllPrograms(c *gin.Context) {
	publishedOnly := ParseParamPublishedOnly(c)

	programList, err := repos.ProgramRepo.SelectAll(publishedOnly)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, programList)
	}
	return
}

func GetProgramById(c *gin.Context) {
	// Incoming parameters
	programId := c.Param("programId")

	program, err := repos.ProgramRepo.SelectByProgramId(programId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &program)
	}
	return
}

func CreateProgram(c *gin.Context) {
	// Incoming JSON
	var programJson domains.Program
	c.BindJSON(&programJson)

	if err := programJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.ProgramRepo.Insert(programJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func UpdateProgram(c *gin.Context) {
	// Incoming JSON & Parameters
	programId := c.Param("programId")
	var programJson domains.Program
	c.BindJSON(&programJson)

	if err := programJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.ProgramRepo.Update(programId, programJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func PublishPrograms(c *gin.Context) {
	// Incoming JSON
	var programIdsJson []string
	c.BindJSON(&programIdsJson)

	err := repos.ProgramRepo.Publish(programIdsJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func DeleteProgram(c *gin.Context) {
	// Incoming Parameters
	programId := c.Param("programId")

	err := repos.ProgramRepo.Delete(programId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
