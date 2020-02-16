package announce

import (
	"errors"
	"net/http"
	"regexp"
	
	"github.com/gin-gonic/gin"
)

const REGEX_ANNOUNCE_ID = `^[[:alnum:]]+(_[[:alnum:]]+)*$`
const REGEX_TITLE = `^[A-Z0-9][[:alnum:]-]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`
const REGEX_MESSAGE = `^\s*$`

func GetAnnouncements(c *gin.Context) {
	// Query Repo
	announceList := GetAllAnnouncements()

	// JSON Response
	c.JSON(http.StatusOK, announceList)
	return
}

func GetAnnouncement(c *gin.Context) {
	// Incoming parameters
	announceId := c.Param("announceId")

	// Query Repo
	announce, err := GetAnnouncementById(announceId)
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
	announceId := c.Param("announceId")
	var announceJson Announce
	c.BindJSON(&announceJson)

	if err := CheckValidAnnouncement(announceJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (UPDATE & SELECT)
	err := UpdateAnnouncementById(announceId, announceJson)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func DeleteAnnouncement(c *gin.Context) {
	// Incoming Parameters
	announceId := c.Param("announceId")

	// Query Repo (DELETE)
	err := DeleteAnnouncementById(announceId)
	if err != nil {
		panic(err)
	} else {
		c.String(http.StatusOK, "Deleted Announcement "+announceId)
	}
	return
}

func CheckValidAnnouncement(announce Announce) error {
	// Retrieves the inputted values
	announceId := announce.AnnounceId
	title := announce.Title
	message := announce.Message
	
	// Announcement ID validation
	if match, _ := regexp.MatchString(REGEX_ANNOUNCE_ID, announceId); !match {
		return errors.New("invalid announcement id")
	}
	
	// Title validation
	if match, _ := regexp.MatchString(REGEX_TITLE, title); !match {
		return errors.New("invalid title")
	}
	
	// Message validation
	if match, _ := regexp.MatchString(REGEX_MESSAGE, message); match {
		return errors.New("empty message")
	}
	
	return nil
}