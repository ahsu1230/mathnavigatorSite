package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllLocations(c *gin.Context) {
	locationList, err := services.LocationService.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, locationList)
	}
	return
}

func GetLocationById(c *gin.Context) {
	// Incoming parameters
	locId := c.Param("locId")

	location, err := services.LocationService.GetByLocationId(locId)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, location)
	}
	return
}

func CreateLocation(c *gin.Context) {
	// Incoming JSON
	var locationJson domains.Location
	c.BindJSON(&locationJson)

	if err := locationJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.LocationService.Create(locationJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func UpdateLocation(c *gin.Context) {
	// Incoming JSON & Parameters
	locId := c.Param("locId")
	var locationJson domains.Location
	c.BindJSON(&locationJson)

	if err := locationJson.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.LocationService.Update(locId, locationJson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func DeleteLocation(c *gin.Context) {
	// Incoming Parameters
	locId := c.Param("locId")

	err := services.LocationService.Delete(locId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}