package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/services"
	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	publishedOnly := ParseParamPublishedOnly(c)

	locationList, err := services.LocationService.GetAll(publishedOnly)
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

	location, err := services.LocationService.GetByLocationId(locationId)
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

	err := services.LocationService.Create(locationJson)
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

	err := services.LocationService.Update(locationId, locationJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func PublishLocations(c *gin.Context) {
	// Incoming JSON
	var locationIdsJson []string
	c.BindJSON(&locationIdsJson)

	err := services.LocationService.Publish(locationIdsJson)
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

	err := services.LocationService.Delete(locationId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
