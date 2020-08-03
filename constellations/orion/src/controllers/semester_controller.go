package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllSemesters(c *gin.Context) {
	// Incoming optional parameter
	publishedOnly := ParseParamPublishedOnly(c)

	semesterList, err := repos.SemesterRepo.SelectAll(publishedOnly)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, semesterList)
	}
}

func GetSemesterById(c *gin.Context) {
	// Incoming parameters
	semesterId := c.Param("semesterId")

	semester, err := repos.SemesterRepo.SelectBySemesterId(semesterId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &semester)
	}
}

func CreateSemester(c *gin.Context) {
	// Incoming JSON
	var semesterJson domains.Semester
	c.BindJSON(&semesterJson)

	if err := semesterJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.SemesterRepo.Insert(semesterJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateSemester(c *gin.Context) {
	// Incoming JSON & Parameters
	semesterId := c.Param("semesterId")
	var semesterJson domains.Semester
	c.BindJSON(&semesterJson)

	if err := semesterJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.SemesterRepo.Update(semesterId, semesterJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteSemester(c *gin.Context) {
	// Incoming Parameters
	semesterId := c.Param("semesterId")

	err := repos.SemesterRepo.Delete(semesterId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
