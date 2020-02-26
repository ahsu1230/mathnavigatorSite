package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

type Program struct {
	Id          uint
	CreatedAt   time.Time     `db:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at"`
	DeletedAt   sql.NullTime  `db:"deleted_at"`
	ProgramId   string        `db:"program_id" json:"programId"`
	Name        string        `json:"name"`
	Grade1      uint          `json:"grade1"`
	Grade2      uint          `json:"grade2"`
	Description string        `json:"description"`
}

// Class Methods

const REGEX_PROGRAM_ID = `^[[:alnum:]]+(_[[:alnum:]]+)*$`
const REGEX_NAME = `^[A-Z0-9][[:alnum:]-]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`

func (program *Program) Validate() error {
	// Retrieves the inputted values
	programId := program.ProgramId
	name := program.Name
	grade1 := program.Grade1
	grade2 := program.Grade2
	description := program.Description

	// Checks if the program ID is in the form of alphanumeric strings separated by underscores
	if matches, _ := regexp.MatchString(REGEX_PROGRAM_ID, programId); !matches {
		return errors.New("invalid program id")
	}

	// Name validation
	match, _ := regexp.MatchString(REGEX_NAME, name)
	if !match {
		return errors.New("invalid program name")
	}

	// Checks if the grades are valid
	if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
		return errors.New("invalid grades")
	}

	// Description validation
	if description == "" {
		return errors.New("empty description")
	}

	return nil
}