package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllSemesters(c *gin.Context) {
	utils.LogControllerMethod(c, "semesterController.GetAllSemesters")
	// Incoming optional parameter
	semesterList, err := repos.SemesterRepo.SelectAll()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, semesterList)
}

func GetSemesterById(c *gin.Context) {
	utils.LogControllerMethod(c, "semesterController.GetSemesterById")
	// Incoming parameters
	semesterId := c.Param("semesterId")

	semester, err := repos.SemesterRepo.SelectBySemesterId(semesterId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &semester)
}

func CreateSemester(c *gin.Context) {
	utils.LogControllerMethod(c, "semesterController.CreateSemester")
	// Incoming JSON
	var semesterJson domains.Semester
	if err := c.ShouldBindJSON(&semesterJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := semesterJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.SemesterRepo.Insert(semesterJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateSemester(c *gin.Context) {
	utils.LogControllerMethod(c, "semesterController.UpdateSemester")
	// Incoming JSON & Parameters
	semesterId := c.Param("semesterId")
	var semesterJson domains.Semester
	if err := c.ShouldBindJSON(&semesterJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := semesterJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.SemesterRepo.Update(semesterId, semesterJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteSemester(c *gin.Context) {
	utils.LogControllerMethod(c, "semesterController.DeleteSemester")
	// Incoming Parameters
	semesterId := c.Param("semesterId")

	err := repos.SemesterRepo.Delete(semesterId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
