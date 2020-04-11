package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllPrograms(c *gin.Context) {
	publishedOnly := ParseParamPublishedOnly(c)

	programList, err := services.ProgramService.GetAll(publishedOnly)
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

	program, err := services.ProgramService.GetByProgramId(programId)
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

	err := services.ProgramService.Create(programJson)
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

	errors := services.ProgramService.Publish(programIdsJson)
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

	err := services.ProgramService.Update(programId, programJson)
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

	err := services.ProgramService.Delete(programId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
