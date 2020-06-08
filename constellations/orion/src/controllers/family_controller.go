package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetFamilyById(c *gin.Context) {
	// Incoming parameters
	id := ParseParamId(c)

	family, err := repos.FamilyRepo.SelectById(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &family)
	}
}

func GetFamilyByPrimaryEmail(c *gin.Context) {
	// Incoming parameters
	primaryEmail := ParseParamUint(c.Param("primaryEmail"))

	family, err := repos.FamilyRepo.SelectByPrimaryEmail(primaryemail)
	if err != nil {
		c.Error(err)
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, &family)
	}
}

func CreateFamily(c *gin.Context) {
	// Incoming JSON
	var familyJson domains.Family
	c.BindJSON(&familyJson)

	if err := familyJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.FamilyRepo.Insert(familyJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func UpdateFamily(c *gin.Context) {
	// Incoming JSON & Parameters
	id := ParseParamId(c)
	var familyJson domains.Family
	c.BindJSON(&familyJson)

	if err := familyJson.Validate(); err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := repos.FamilyRepo.Update(id, familyJson)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteFamily(c *gin.Context) {
	// Incoming Parameters
	id := ParseParamId(c)

	err := repos.FamilyRepo.Delete(id)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
