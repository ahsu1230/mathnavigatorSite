package domains

import (
	"errors"
	"regexp"
	"time"
)

var TABLE_PROGRAMS = "programs"

type Program struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	UpdatedAt   time.Time `json:"-" db:"updated_at"`
	DeletedAt   NullTime  `json:"-" db:"deleted_at"`
	ProgramId   string    `json:"programId" db:"program_id"`
	Name        string    `json:"name"`
	Grade1      uint      `json:"grade1"`
	Grade2      uint      `json:"grade2"`
	Description string    `json:"description"`
	Featured    uint      `json:"featured"`
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
	if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, programId); !matches {
		return errors.New("invalid program id")
	}

	// Name validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, name); !matches {
		return errors.New("invalid program name")
	}

	// Grade validation
	if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
		return errors.New("invalid grades")
	}

	// Description validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, description); !matches {
		return errors.New("invalid description")
	}

	return nil
}
