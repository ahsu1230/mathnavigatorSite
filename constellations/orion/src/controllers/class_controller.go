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
	utils.LogControllerMethod(c, "classController.GetAllClasses")
	publishedOnly := utils.ParseParamPublishedOnly(c)

	ctx := utils.RetrieveContext(c)
	classList, err := repos.ClassRepo.SelectAll(ctx, publishedOnly)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classList)
}

func GetClassById(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.GetClassById")
	classId := c.Param("classId")

	ctx := utils.RetrieveContext(c)
	class, err := repos.ClassRepo.SelectByClassId(ctx, classId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &class)
}

func GetClassesByProgram(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.GetClassesByProgram")
	programId := c.Param("programId")

	ctx := utils.RetrieveContext(c)
	classes, err := repos.ClassRepo.SelectByProgramId(ctx, programId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classes)
}

func GetClassesBySemester(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.GetClassesBySemester")
	semesterId := c.Param("semesterId")

	ctx := utils.RetrieveContext(c)
	classes, err := repos.ClassRepo.SelectBySemesterId(ctx, semesterId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classes)
}

func GetClassesByProgramAndSemester(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.GetClassesByProgramAndSemester")
	programId := c.Param("programId")
	semesterId := c.Param("semesterId")

	ctx := utils.RetrieveContext(c)
	classes, err := repos.ClassRepo.SelectByProgramAndSemesterId(ctx, programId, semesterId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classes)
}

func GetUnpublishedClasses(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.GetUnpublishedClasses")
	ctx := utils.RetrieveContext(c)
	classList, err := repos.ClassRepo.SelectAllUnpublished(ctx)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, classList)
}

func CreateClass(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.CreateClass")
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

	ctx := utils.RetrieveContext(c)
	id, err := repos.ClassRepo.Insert(ctx, classJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateClass(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.UpdateClass")
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

	ctx := utils.RetrieveContext(c)
	err := repos.ClassRepo.Update(ctx, classId, classJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func PublishClasses(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.PublishClass")
	var classIds []string
	if err := c.ShouldBindJSON(&classIds); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	ctx := utils.RetrieveContext(c)
	errs := repos.ClassRepo.Publish(ctx, classIds)
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
	utils.LogControllerMethod(c, "classController.DeleteClass")
	classId := c.Param("classId")

	ctx := utils.RetrieveContext(c)
	err := repos.ClassRepo.Delete(ctx, classId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}

func GetFullStates(c *gin.Context) {
	utils.LogControllerMethod(c, "classController.getFullStates")
	states := []string{
		domains.NOT_FULL_DISPLAY_NAME,
		domains.ALMOST_FULL_DISPLAY_NAME,
		domains.FULL_DISPLAY_NAME,
	}
	c.JSON(http.StatusOK, states)
}
