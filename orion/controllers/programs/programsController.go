package programs

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Alphanumeric characters separated by underscores
const REGEX_PROGRAM_ID = `^[[:alnum:]]+(_[[:alnum:]]+)*$`

// Starts with a capital letter or number. Words consist of alphanumeric characters and dashes, spaces, and underscores
// separate words. Words can have parentheses around them and number signs must be followed by numbers.
const REGEX_NAME = `^[A-Z0-9][[:alnum:]]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`

// Ensures at least one uppercase or lowercase letter
const REGEX_ALPHA_ONLY = `[A-Za-z]+`

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
		c.Status(http.StatusNoContent)
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
		c.Status(http.StatusNoContent)
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
		c.Status(http.StatusNoContent)
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

	// Program ID validation
	if matches, _ := regexp.MatchString(REGEX_PROGRAM_ID, programId); !matches || len(programId) > 64 {
		return errors.New("invalid program id")
	}

	// Name validation
	if matches, _ := regexp.MatchString(REGEX_NAME, name); !matches || len(name) > 255 {
		return errors.New("invalid program name")
	}

	// Grade validation
	if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
		return errors.New("invalid grades")
	}

	// Description validation
	if matches, _ := regexp.MatchString(REGEX_ALPHA_ONLY, description); !matches {
		return errors.New("invalid description")
	}

	return nil
}
