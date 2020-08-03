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
	if matches, _ := regexp.MatchString(REGEX_SEMESTER_ID, semesterId); !matches || len(semesterId) > 64 {
		return errors.New("invalid semester id")
	}

	// Title validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, title); !matches || len(title) > 64 {
		return errors.New("invalid semester title")
	}

	return nil
}
