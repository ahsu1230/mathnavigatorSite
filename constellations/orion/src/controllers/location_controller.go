package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	utils.LogControllerMethod(c, "locationController.GetAllLocations")
	locationList, err := repos.LocationRepo.SelectAll()
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, locationList)
}

func GetLocationById(c *gin.Context) {
	utils.LogControllerMethod(c, "locationController.GetLocationById")
	// Incoming parameters
	locationId := c.Param("locationId")

	location, err := repos.LocationRepo.SelectByLocationId(locationId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, &location)
}

func CreateLocation(c *gin.Context) {
	utils.LogControllerMethod(c, "locationController.CreateLocation")
	// Incoming JSON
	var locationJson domains.Location
	if err := c.ShouldBindJSON(&locationJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := locationJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.LocationRepo.Insert(locationJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func UpdateLocation(c *gin.Context) {
	utils.LogControllerMethod(c, "locationController.UpdateLocation")
	// Incoming JSON & Parameters
	locationId := c.Param("locationId")
	var locationJson domains.Location
	if err := c.ShouldBindJSON(&locationJson); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	if err := locationJson.Validate(); err != nil {
		c.Error(appErrors.WrapInvalidDomain(err.Error()))
		c.Abort()
		return
	}

	err := repos.LocationRepo.Update(locationId, locationJson)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusOK)
}

func DeleteLocation(c *gin.Context) {
	utils.LogControllerMethod(c, "locationController.DeleteLocation")
	// Incoming Parameters
	locationId := c.Param("locationId")

	err := repos.LocationRepo.Delete(locationId)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}
	c.Status(http.StatusNoContent)
}
