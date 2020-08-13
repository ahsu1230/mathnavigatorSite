package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllClasses(c *gin.Context) {
	// Incoming optional parameter
	publishedOnly := utils.ParseParamPublishedOnly(c)

	classList, err := repos.ClassRepo.SelectAll(publishedOnly)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classList)
}

func GetClassById(c *gin.Context) {
	// Incoming parameters
	classId := c.Param("classId")

	class, err := repos.ClassRepo.SelectByClassId(classId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &class)
}

func GetClassesByProgram(c *gin.Context) {
	// Incoming parameters
	programId := c.Param("programId")

	classes, err := repos.ClassRepo.SelectByProgramId(programId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classes)
}

func GetClassesBySemester(c *gin.Context) {
	// Incoming parameters
	semesterId := c.Param("semesterId")

	classes, err := repos.ClassRepo.SelectBySemesterId(semesterId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classes)
}

func GetClassesByProgramAndSemester(c *gin.Context) {
	// Incoming parameters
	programId := c.Param("programId")
	semesterId := c.Param("semesterId")

	classes, err := repos.ClassRepo.SelectByProgramAndSemesterId(programId, semesterId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classes)
}

func CreateClass(c *gin.Context) {
	// Incoming JSON
	var classJson domains.Class
	if err := c.ShouldBindJSON(&classJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := classJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.ClassRepo.Insert(classJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateClass(c *gin.Context) {
	// Incoming JSON & Parameters
	classId := c.Param("classId")
	var classJson domains.Class
	if err := c.ShouldBindJSON(&classJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := classJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.ClassRepo.Update(classId, classJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func PublishClasses(c *gin.Context) {
	// Incoming JSON
	var classIds []string
	if err := c.ShouldBindJSON(&classIds); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	errs := repos.ClassRepo.Publish(classIds)
	if len(errs) > 0 {
		for _, err := range errs {
			c.Error(err)
		}
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteClass(c *gin.Context) {
	// Incoming Parameters
	classId := c.Param("classId")

	err := repos.ClassRepo.Delete(classId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
