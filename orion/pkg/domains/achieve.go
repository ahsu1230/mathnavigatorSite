package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_ACHIEVEMENTS = "achievements"

type Achieve struct {
	Id        uint
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
	Year      uint         `json:"year"`
	Message   string       `json:"message"`
}

type AchieveYearGroup struct {
	Year         uint      `json:"year"`
	Achievements []Achieve `json:"achievements"`
}

// Class Methods

func (achieve *Achieve) Validate() error {
	// Retrieves the inputted values
	year := achieve.Year
	message := achieve.Message

	// Year validation
	if year < 2000 {
		return errors.New("invalid year")
	}

	// Message validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, message); !matches {
		return errors.New("invalid message")
	}

	return nil
}
