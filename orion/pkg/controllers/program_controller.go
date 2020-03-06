package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
)

func GetAllPrograms(c *gin.Context) {
	programList, err := services.ProgramService.GetAll()
	if err != nil {
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
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, program)
	}
	return
}

func CreateProgram(c *gin.Context) {
	// Incoming JSON
	var programJson domains.Program
	c.BindJSON(&programJson)

	if err := programJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.ProgramService.Create(programJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func UpdateProgram(c *gin.Context) {
	// Incoming JSON & Parameters
	programId := c.Param("programId")
	var programJson domains.Program
	c.BindJSON(&programJson)

	if err := programJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.ProgramService.Update(programId, programJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func DeleteProgram(c *gin.Context) {
	// Incoming Parameters
	programId := c.Param("programId")

	err := services.ProgramService.Delete(programId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
