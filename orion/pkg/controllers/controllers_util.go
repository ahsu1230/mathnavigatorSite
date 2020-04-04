package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ParseParamId(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(id)
}

func ParseParamIds(c *gin.Context, key string) []uint {
	ids := make([]uint, 0)

	for _, id := range c.PostFormArray(key) {
		id, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			panic(err)
		}
		ids = append(ids, uint(id))
	}
	return ids
}

func checkError(c *gin.Context, list interface{}, err error) {
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, list)
	}
}

func GetAllUnpublished(c *gin.Context) {
	programList, err := services.ProgramService.GetAll()
	checkError(c, programList, err)
	classList, err := services.ClassService.GetAll()
	checkError(c, classList, err)
	locationList, err := services.LocationService.GetAll()
	checkError(c, locationList, err)
	achieveList, err := services.AchieveService.GetAll()
	checkError(c, achieveList, err)
	semesterList, err := services.SemesterService.GetAll()
	checkError(c, semesterList, err)
	sessionList, err := services.SessionService.GetAll()
	checkError(c, sessionList, err)
}