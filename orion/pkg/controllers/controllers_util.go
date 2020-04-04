package controllers

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ParseParamId(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(id)
}

func ParseParamIds(c *gin.Context, key string) []uint {
	ids := make([]uint, 0)

	for _, id := range c.PostFormArray(key) {
		id, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			panic(err)
		}
		ids = append(ids, uint(id))
	}
	return ids
}

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
