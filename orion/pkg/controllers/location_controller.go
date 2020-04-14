package controllers

import (
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	locId := c.Param("locId")

	location, err := services.LocationService.GetByLocationId(locId)
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

func PublishLocations(c *gin.Context) {
	// Incoming JSON
	var locIdsJson []string
	c.BindJSON(&locIdsJson)

	errorList := services.LocationService.Publish(locIdsJson)
	if len(errorList) > 0 {
		var errorStrings []string
		errorStrings = append(errorStrings, "one or more locations failed to publish")
		for _, errorBody := range errorList {
			errorStrings = append(errorStrings, errorBody.StringId+": "+errorBody.Error.Error())
		}
		err := strings.Join(errorStrings, "\n")
		c.Error(errors.New(err))
		c.String(http.StatusInternalServerError, err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func UpdateLocation(c *gin.Context) {
	// Incoming JSON & Parameters
	locId := c.Param("locId")
	var locationJson domains.Location
	c.BindJSON(&locationJson)

	if err := locationJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := services.LocationService.Update(locId, locationJson)
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
	locId := c.Param("locId")

	err := services.LocationService.Delete(locId)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
	return
}
