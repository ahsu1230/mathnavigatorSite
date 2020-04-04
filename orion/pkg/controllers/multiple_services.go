package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func callServices() (domains.Domains, error) {
	// programList, err := services.ProgramService.GetUnpublished()
	// if err != nil {
	// 	return domains.Domains{}, err
	// }
	// classList, err := services.ClassService.GetUnpublished()
	// if err != nil {
	// 	return domains.Domains{}, err
	// }
	// locationList, err := services.LocationService.GetUnpublished()
	// if err != nil {
	// 	return domains.Domains{}, err
	// }
	achieveList, err := services.AchieveService.GetUnpublished()
	if err != nil {
		return domains.Domains{}, err
	}
	// semesterList, err := services.SemesterService.GetUnpublished()
	// if err != nil {
	// 	return domains.Domains{}, err
	// }
	// sessionList, err := services.SessionService.GetUnpublished()
	// if err != nil {
	// 	return domains.Domains{}, err
	// }

	allDomains := domains.Domains{
		// Programs: programList,
		// Classes:   classList,
		// Locations: locationList,
		Achieves: achieveList,
		// Semesters: semesterList,
		// Sessions:  sessionList,
	}

	return allDomains, nil
}

func GetAllUnpublished(c *gin.Context) {
	allDomains, err := callServices()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, allDomains)
	}
}
