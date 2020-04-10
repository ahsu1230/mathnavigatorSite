package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUnpublished(c *gin.Context) {
	unpublishedDomains, err := callGetUnpublishedServices()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, unpublishedDomains)
	}
}

func callGetUnpublishedServices() (domains.UnpublishedDomains, error) {
	// programList, err := services.ProgramService.GetUnpublished()
	// if err != nil {
	// 	return domains.UnpublishedDomains{}, err
	// }
	classList, err := services.ClassService.GetUnpublished()
	if err != nil {
		return domains.UnpublishedDomains{}, err
	}
	// locationList, err := services.LocationService.GetUnpublished()
	// if err != nil {
	// 	return domains.UnpublishedDomains{}, err
	// }
	achieveList, err := services.AchieveService.GetUnpublished()
	if err != nil {
		return domains.UnpublishedDomains{}, err
	}
	// semesterList, err := services.SemesterService.GetUnpublished()
	// if err != nil {
	// 	return domains.UnpublishedDomains{}, err
	// }
	// sessionList, err := services.SessionService.GetUnpublished()
	// if err != nil {
	// 	return domains.UnpublishedDomains{}, err
	// }

	unpublishedDomains := domains.UnpublishedDomains{
		// Programs: programList,
		Classes: classList,
		// Locations: locationList,
		Achieves: achieveList,
		// Semesters: semesterList,
		// Sessions:  sessionList,
	}

	return unpublishedDomains, nil
}
