package announce

import (
	"errors"
	"net/http"
	"regexp"
	
	"github.com/gin-gonic/gin"
)

const REGEX = `[A-Za-z]+`

func GetAnnouncements(c *gin.Context) {
	// Query Repo
	announceList := GetAllAnnouncements()

	// JSON Response
	c.JSON(http.StatusOK, announceList)
	return
}

func GetAnnouncement(c *gin.Context) {
	// Incoming parameters
	id := c.Param("id")

	// Query Repo
	announce, err := GetAnnouncementById(id)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, announce)
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
	err := InsertAnnouncement(announceJson)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func UpdateAnnouncement(c *gin.Context) {
	// Incoming JSON & Parameters
	id := c.Param("id")
	var announceJson Announce
	c.BindJSON(&announceJson)

	if err := CheckValidAnnouncement(announceJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (UPDATE & SELECT)
	err := UpdateAnnouncementById(id, announceJson)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func DeleteAnnouncement(c *gin.Context) {
	// Incoming Parameters
	id := c.Param("id")

	// Query Repo (DELETE)
	err := DeleteAnnouncementById(id)
	if err != nil {
		panic(err)
	} else {
		c.String(http.StatusOK, "Deleted Announcement " + id)
	}
	return
}

func CheckValidAnnouncement(announce Announce) error {
	// Retrieves the inputted values
	author := announce.Author
	message := announce.Message
	
	// Author validation
	if match, _ := regexp.MatchString(REGEX, author); !match {
		return errors.New("invalid author")
	}
	
	// Message validation
	if match, _ := regexp.MatchString(REGEX, message); !match {
		return errors.New("invalid message")
	}
	
	return nil
}