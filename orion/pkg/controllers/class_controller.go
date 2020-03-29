package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllClasses(c *gin.Context) {
	classList, err := services.ClassService.GetAll()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, classList)
	}
	return
}

func GetClassById(c *gin.Context) {
	// Incoming parameters
	classId := c.Param("classId")

	class, err := services.ClassService.GetByClassId(classId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &class)
	}
	return
}

func GetClassesByProgram(c *gin.Context) {
	// Incoming parameters
	programId := c.Param("programId")

	classes, err := services.ClassService.GetByProgramId(programId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, classes)
	}
	return
}

func GetClassesBySemester(c *gin.Context) {
	// Incoming parameters
	semesterId := c.Param("semesterId")

	classes, err := services.ClassService.GetBySemesterId(semesterId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, classes)
	}
	return
}

func GetClassesByProgramAndSemester(c *gin.Context) {
	// Incoming parameters
	programId := c.Param("programId")
	semesterId := c.Param("semesterId")

	classes, err := services.ClassService.GetByProgramAndSemesterId(programId, semesterId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, classes)
	}
	return
}

func CreateClass(c *gin.Context) {
	// Incoming JSON
	var classJson domains.Class
	c.BindJSON(&classJson)

	if err := classJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.ClassService.Create(classJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func UpdateClass(c *gin.Context) {
	// Incoming JSON & Parameters
	classId := c.Param("classId")
	var classJson domains.Class
	c.BindJSON(&classJson)

	if err := classJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.ClassService.Update(classId, classJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func DeleteClass(c *gin.Context) {
	// Incoming Parameters
	classId := c.Param("classId")

	err := services.ClassService.Delete(classId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
