package announce

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	
	"github.com/gin-gonic/gin"
)

const REGEX_ALPHA_ONLY = `[A-Za-z]+`

func GetAnnouncements(c *gin.Context) {
	// Query Repo
	announceList := GetAllAnnouncements()

	// JSON Response
	c.JSON(http.StatusOK, announceList)
	return
}

func GetAnnouncement(c *gin.Context) {
	// Incoming parameters
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)

	// Query Repo
	if _, err := GetAnnouncementById(id); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func CreateAnnouncement(c *gin.Context) {
	// Incoming JSON
	var announceJson Announce
	c.BindJSON(&announceJson)

	if err := CheckValidAnnouncement(announceJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (INSERT & SELECT)
	if err := InsertAnnouncement(announceJson); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func UpdateAnnouncement(c *gin.Context) {
	// Incoming JSON & Parameters
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	var announceJson Announce
	c.BindJSON(&announceJson)

	if err := CheckValidAnnouncement(announceJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (UPDATE & SELECT)
	if err := UpdateAnnouncementById(id, announceJson); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func DeleteAnnouncement(c *gin.Context) {
	// Incoming Parameters
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)

	// Query Repo (DELETE)
	if err := DeleteAnnouncementById(id); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func CheckValidAnnouncement(announce Announce) error {
	// Retrieves the inputted values
	author := announce.Author
	message := announce.Message
	
	// Author validation
	if match, _ := regexp.MatchString(REGEX_ALPHA_ONLY, author); !match {
		return errors.New("invalid author")
	}
	
	// Message validation
	if match, _ := regexp.MatchString(REGEX_ALPHA_ONLY, message); !match {
		return errors.New("invalid message")
	}
	
	return nil
}