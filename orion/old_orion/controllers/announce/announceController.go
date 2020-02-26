package announce

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Ensures at least one uppercase or lowercase letter
const REGEX_LETTER = `[A-Za-z]+`

func parseParamId(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(id)
}

func GetAnnouncements(c *gin.Context) {
	// Query Repo
	announceList := GetAllAnnouncements()

	// JSON Response
	c.JSON(http.StatusOK, announceList)
	return
}

func GetAnnouncement(c *gin.Context) {
	// Incoming parameters
	id := parseParamId(c)

	// Query Repo
	if announcement, err := GetAnnouncementById(id); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, announcement)
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
		c.Status(http.StatusNoContent)
	}
	return
}

func UpdateAnnouncement(c *gin.Context) {
	// Incoming JSON & Parameters
	id := parseParamId(c)
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
		c.Status(http.StatusNoContent)
	}
	return
}

func DeleteAnnouncement(c *gin.Context) {
	// Incoming Parameters
	id := parseParamId(c)

	// Query Repo (DELETE)
	if err := DeleteAnnouncementById(id); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusNoContent)
	}
	return
}

func CheckValidAnnouncement(announce Announce) error {
	// Retrieves the inputted values
	author := announce.Author
	message := announce.Message

	// Author validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, author); !matches {
		return errors.New("invalid author")
	}

	// Message validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, message); !matches {
		return errors.New("invalid message")
	}

	return nil
}
