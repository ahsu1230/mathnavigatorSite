package controllers

import (
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

func GetAllUnpublished(c *gin.Context) {
	unpublishedDomains, err := callGetUnpublishedRepos()
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, unpublishedDomains)
}

func callGetUnpublishedRepos() (domains.UnpublishedDomains, error) {
	classList, err := repos.ClassRepo.SelectAllUnpublished()
	if err != nil {
		return domains.UnpublishedDomains{}, appErrors.WrapRepo(err)
	}
	unpublishedDomains := domains.UnpublishedDomains{
		Classes: classList,
		// Other domains were removed (unused)
	}
	return unpublishedDomains, nil
}
