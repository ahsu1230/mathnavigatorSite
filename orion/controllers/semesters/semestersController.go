package semesters

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

const REGEX_SEMESTER_ID = `^[1-9]\d{3,}_((spring)|(summer)|(fall)|(winter))$`
const REGEX_TITLE = `^[A-Z0-9][[:alnum:]-]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`

func GetSemesters(c *gin.Context) {
	// Query Repo
	semesterList := GetAllSemesters()

	// JSON Response
	c.JSON(http.StatusOK, semesterList)
	return
}

func GetSemester(c *gin.Context) {
	// Incoming parameters
	semesterId := c.Param("semesterId")

	// Query Repo
	semester, err := GetSemesterById(semesterId)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, semester)
	}
	return
}

func CreateSemester(c *gin.Context) {
	// Incoming JSON
	var semesterJson Semester
	c.BindJSON(&semesterJson)

	if err := CheckValidSemester(semesterJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (INSERT & SELECT)
	err := InsertSemester(semesterJson)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func UpdateSemester(c *gin.Context) {
	// Incoming JSON & Parameters
	semesterId := c.Param("semesterId")
	var semesterJson Semester
	c.BindJSON(&semesterJson)

	if err := CheckValidSemester(semesterJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (UPDATE & SELECT)
	err := UpdateSemesterById(semesterId, semesterJson)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
	return
}

func DeleteSemester(c *gin.Context) {
	// Incoming Parameters
	semesterId := c.Param("semesterId")

	// Query Repo (DELETE)
	err := DeleteSemesterById(semesterId)
	if err != nil {
		panic(err)
	} else {
		c.Status(http.StatusNoContent)
	}
	return
}

func CheckValidSemester(semester Semester) error {
	// Retrieves the inputted values
	semesterId := semester.SemesterId
	title := semester.Title

	// Checks if the semester ID is in the form year_season
	if matches, _ := regexp.MatchString(REGEX_SEMESTER_ID, semesterId); !matches {
		return errors.New("invalid semester id")
	}

	// Title validation
	match, _ := regexp.MatchString(REGEX_TITLE, title)
	if !match || len(title) > 64 {
		return errors.New("invalid semester title")
	}

	return nil
}
