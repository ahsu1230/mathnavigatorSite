package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_PROGRAMS = "programs"

type Program struct {
	Id          uint
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
	ProgramId   string       `db:"program_id" json:"programId"`
	Name        string       `json:"name"`
	Grade1      uint         `json:"grade1"`
	Grade2      uint         `json:"grade2"`
	Description string       `json:"description"`
}

// Class Methods

func (program *Program) Validate() error {
	// Retrieves the inputted values
	programId := program.ProgramId
	name := program.Name
	grade1 := program.Grade1
	grade2 := program.Grade2
	description := program.Description

	// Program ID validation
	if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, programId); !matches || len(programId) > 64 {
		return errors.New("invalid program id")
	}

	// Name validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, name); !matches || len(name) > 255 {
		return errors.New("invalid program name")
	}

	// Grade validation
	if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
		return errors.New("invalid grades")
	}

	// Description validation
	if matches, _ := regexp.MatchString(REGEX_AT_LEAST_ONE_LETTER, description); !matches {
		return errors.New("invalid description")
	}

	return nil
}
