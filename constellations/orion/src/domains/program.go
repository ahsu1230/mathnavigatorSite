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
		return appErrors.WrapInvalidDomain("Invalid program ID")
	}

	// Name validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, name); !matches {
		return appErrors.WrapInvalidDomain("Invalid program name")
	}

	// Grade validation
	if !(grade1 <= grade2 && grade1 >= 1 && grade2 <= 12) {
		return appErrors.WrapInvalidDomain("Invalid grades (must be between 1 and 12) and grade1 <= grade2")
	}

	// Description validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, description); !matches {
		return appErrors.WrapInvalidDomain("Invalid description entry")
	}

	return nil
}
