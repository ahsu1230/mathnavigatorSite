package domains

import (
	"errors"
	"regexp"
	"time"
)

var TABLE_ACHIEVEMENTS = "achievements"

type Achieve struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	UpdatedAt   time.Time `json:"-" db:"updated_at"`
	DeletedAt   NullTime  `json:"-" db:"deleted_at"`
	PublishedAt NullTime  `json:"publishedAt" db:"published_at"`
	Year        uint      `json:"year"`
	Message     string    `json:"message"`
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
