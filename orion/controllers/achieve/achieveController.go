package achieve

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAchievements(c *gin.Context) {
	// Query Repo
	achieveList := GetAllAchievements()

	// JSON Response
	c.JSON(http.StatusOK, achieveList)
	return
}

func GetAchievement(c *gin.Context) {
	id := parseParamId(c)

	// Query Repo
	achieve, err := GetAchievementById(id)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, achieve)
	}
	return
}

func CreateAchievement(c *gin.Context) {
	// Incoming JSON
	var achieveJson Achieve
	c.BindJSON(&achieveJson)

	if err := CheckValidAchievement(achieveJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (INSERT & SELECT)
	err := InsertAchievement(achieveJson)
	if err != nil {
		panic(err)
	} else {
		c.Status(http.StatusNoContent)
	}
	return
}

func UpdateAchievement(c *gin.Context) {
	// Incoming JSON & Parameters
	id := parseParamId(c)
	var achieveJson Achieve
	c.BindJSON(&achieveJson)

	if err := CheckValidAchievement(achieveJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (UPDATE & SELECT)
	err := UpdateAchievementById(id, achieveJson)
	if err != nil {
		panic(err)
	} else {
		c.Status(http.StatusNoContent)
	}
	return
}

func DeleteAchievement(c *gin.Context) {
	id := parseParamId(c)

	// Query Repo (DELETE)
	err := DeleteAchievementById(id)
	if err != nil {
		panic(err)
	} else {
		c.Status(http.StatusNoContent)
	}
	return
}

func parseParamId(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(id)
}

func CheckValidAchievement(achieve Achieve) error {
	// Retrieves the inputted values
	year := achieve.Year
	message := achieve.Message

	// Year validation
	if year < 2000 {
		return errors.New("invalid year")
	}

	// Message validation
	if message == "" {
		return errors.New("invalid message")
	}

	return nil
}
