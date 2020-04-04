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

type Domains struct {
	Programs  []domains.Program  `json:"programs"`
	Classes   []domains.Class    `json:"classes"`
	Locations []domains.Location `json:"locations"`
	Achieves  []domains.Achieve  `json:"achieves"`
	Semesters []domains.Semester `json:"semesters"`
	Sessions  []domains.Session  `json:"sessions"`
}

func callServices() (Domains, error) {
	// programList, err := services.ProgramService.GetUnpublished()
	// if err != nil {
	// 	return Domains{}, err
	// }
	// classList, err := services.ClassService.GetUnpublished()
	// if err != nil {
	// 	return Domains{}, err
	// }
	// locationList, err := services.LocationService.GetUnpublished()
	// if err != nil {
	// 	return Domains{}, err
	// }
	achieveList, err := services.AchieveService.GetUnpublished()
	if err != nil {
		return Domains{}, err
	}
	// semesterList, err := services.SemesterService.GetUnpublished()
	// if err != nil {
	// 	return Domains{}, err
	// }
	// sessionList, err := services.SessionService.GetUnpublished()
	// if err != nil {
	// 	return Domains{}, err
	// }

	domains := Domains{
		// Programs: programList,
		// Classes:   classList,
		// Locations: locationList,
		Achieves: achieveList,
		// Semesters: semesterList,
		// Sessions:  sessionList,
	}

	return domains, nil
}

func GetAllUnpublished(c *gin.Context) {
	domains, err := callServices()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, domains)
	}
}
