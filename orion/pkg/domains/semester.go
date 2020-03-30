package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_SEMESTERS = "semesters"

type Semester struct {
	Id         uint
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	SemesterId string       `db:"semester_id" json:"semesterId"`
	Title      string       `json:"title"`
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
