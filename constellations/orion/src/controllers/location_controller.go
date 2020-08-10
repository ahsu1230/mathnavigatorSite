package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	locationList, err := repos.LocationRepo.SelectAll()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, locationList)
	}
	return
}

func GetLocationById(c *gin.Context) {
	// Incoming parameters
	locationId := c.Param("locationId")

	location, err := repos.LocationRepo.SelectByLocationId(locationId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &location)
	}
	return
}

func CreateLocation(c *gin.Context) {
	// Incoming JSON
	var locationJson domains.Location
	c.BindJSON(&locationJson)

	if err := locationJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.LocationRepo.Insert(locationJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func UpdateLocation(c *gin.Context) {
	// Incoming JSON & Parameters
	locationId := c.Param("locationId")
	var locationJson domains.Location
	c.BindJSON(&locationJson)

	if err := locationJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.LocationRepo.Update(locationId, locationJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func DeleteLocation(c *gin.Context) {
	// Incoming Parameters
	locationId := c.Param("locationId")

	err := repos.LocationRepo.Delete(locationId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
