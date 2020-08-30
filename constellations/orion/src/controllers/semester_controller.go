package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllSeasons(c *gin.Context) {
	utils.LogControllerMethod(c, "semesterController.GetAllSeasons")
	c.JSON(http.StatusOK, domains.ALL_SEASONS)
}

func GetAllSemesters(c *gin.Context) {
	utils.LogControllerMethod(c, "semesterController.GetAllSemesters")
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
	var semesterJson domains.Semester
	if err := c.ShouldBindJSON(&semesterJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	// When creating, only need season & year
	// All other fields will be determined for you
	semesterJson = standardizeSemester(semesterJson)
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
	semesterId := c.Param("semesterId")
	var semesterJson domains.Semester
	if err := c.ShouldBindJSON(&semesterJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	// When updating, only need season & year
	// All other fields will be determined for you
	semesterJson = standardizeSemester(semesterJson)
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
	semesterId := c.Param("semesterId")

	err := repos.SemesterRepo.Delete(semesterId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func standardizeSemester(semester domains.Semester) domains.Semester {
	season := semester.Season
	year := semester.Year
	semester.SemesterId = fmt.Sprintf("%d_%s", year, season)
	semester.Title = strings.Title(fmt.Sprintf("%s %d", season, year))
	return semester
}
