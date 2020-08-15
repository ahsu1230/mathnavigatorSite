package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllUnpublished(c *gin.Context) {
	unpublishedDomains, err := callGetUnpublishedRepos()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, unpublishedDomains)
	}
}

func callGetUnpublishedRepos() (domains.UnpublishedDomains, error) {
	classList, err := repos.ClassRepo.SelectAllUnpublished()
	if err != nil {
		return domains.UnpublishedDomains{}, err
	}

	unpublishedDomains := domains.UnpublishedDomains{
		Classes: classList,
	}

	return unpublishedDomains, nil
}
