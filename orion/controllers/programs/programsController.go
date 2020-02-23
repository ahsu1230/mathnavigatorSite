package programs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

/*
Regex checks ensure program IDs only contain alphanumeric characters, names do not
start with lowercase or special characters, and descriptions contain at least one
alphabetical character.
*/
const REGEX_PROGRAM_ID = `^[[:alnum:]]+(_[[:alnum:]]+)*$`
const REGEX_NAME = `^[A-Z0-9][[:alnum:]-]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`
const REGEX_DESCRIPTION = `[a-zA-z]+`

func GetPrograms(c *gin.Context) {
	// Query Repo
	programList := GetAllPrograms()

	// JSON Response
	c.JSON(http.StatusOK, programList)
	return
}

func GetProgram(c *gin.Context) {
	// Incoming parameters
	programId := c.Param("programId")

	// Query Repo
	if _, err := GetProgramById(programId); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func CreateProgram(c *gin.Context) {
	// Incoming JSON
	var programJson Program
	c.BindJSON(&programJson)

	if err := CheckValidProgram(programJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (INSERT & SELECT)
	if err := InsertProgram(programJson); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func UpdateProgram(c *gin.Context) {
	// Incoming JSON & Parameters
	programId := c.Param("programId")
	var programJson Program
	c.BindJSON(&programJson)

	if err := CheckValidProgram(programJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Query Repo (UPDATE & SELECT)
	if err := UpdateProgramById(programId, programJson); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func DeleteProgram(c *gin.Context) {
	// Incoming Parameters
	programId := c.Param("programId")

	// Query Repo (DELETE)
	if err := DeleteProgramById(programId); err != nil {
		panic(err)
	} else {
		c.Status(http.StatusOK)
	}
	return
}

func CheckValidProgram(program Program) error {
	// Retrieves the inputted values
	programId := program.ProgramId
	name := program.Name
	grade1 := program.Grade1
	grade2 := program.Grade2
	description := program.Description

	// Checks if the program ID is in the form of alphanumeric strings separated by underscores
	if match, _ := regexp.MatchString(REGEX_PROGRAM_ID, programId); !match {
		return errors.New("invalid program id")
	}

	// Name validation
	if match, _ := regexp.MatchString(REGEX_NAME, name); !match {
		return errors.New("invalid program name")
	}

	// Checks if the grades are valid
	if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
		return errors.New("invalid grades")
	}

	// Description validation
	if match, _ := regexp.MatchString(REGEX_DESCRIPTION, description); !match {
		return errors.New("invalid description")
	}

	return nil
}
