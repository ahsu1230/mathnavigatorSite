package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSemesters(c *gin.Context) {
	semesterList, err := services.SemesterService.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, semesterList)
	}
	return
}

func GetSemesterById(c *gin.Context) {
	// Incoming parameters
	semesterId := c.Param("semesterId")

	semester, err := services.SemesterService.GetBySemesterId(semesterId)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, semester)
	}
	return
}

func CreateSemester(c *gin.Context) {
	// Incoming JSON
	var semesterJson domains.Semester
	c.BindJSON(&semesterJson)

	if err := semesterJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.SemesterService.Create(semesterJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func UpdateSemester(c *gin.Context) {
	// Incoming JSON & Parameters
	semesterId := c.Param("semesterId")
	var semesterJson domains.Semester
	c.BindJSON(&semesterJson)

	if err := semesterJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.SemesterService.Update(semesterId, semesterJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func DeleteSemester(c *gin.Context) {
	// Incoming Parameters
	semesterId := c.Param("semesterId")

	err := services.SemesterService.Delete(semesterId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
