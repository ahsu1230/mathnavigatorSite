package domains

import (
	"errors"
	"regexp"
	"time"
)

var TABLE_SEMESTERS = "semesters"

type Semester struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	DeletedAt  NullTime  `json:"-" db:"deleted_at"`
	SemesterId string    `json:"semesterId" db:"semester_id"`
	Title      string    `json:"title"`
	Ordering   uint      `json:"ordering" db:"ordering"`
}

// Class Methods

func (semester *Semester) Validate() error {
	// Retrieves the inputted values
	semesterId := semester.SemesterId
	title := semester.Title

	// Semester ID validation
	if matches, _ := regexp.MatchString(REGEX_SEMESTER_ID, semesterId); !matches {
		return appErrors.WrapInvalidDomain("Invalid Semester ID (should be format `year_season`)")
	}

	// Title validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, title); !matches {
		return appErrors.WrapInvalidDomain("Invalid Semester Title")
	}

	return nil
}
