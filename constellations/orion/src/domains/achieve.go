package domains

import (
	"fmt"
	"regexp"
	"time"
)

var TABLE_ACHIEVEMENTS = "achievements"

type Achieve struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	DeletedAt NullTime  `json:"-" db:"deleted_at"`
	Year      uint      `json:"year"`
	Message   string    `json:"message"`
	Position  uint      `json:"position"`
}

type AchieveYearGroup struct {
	Year         uint      `json:"year"`
	Achievements []Achieve `json:"achievements"`
}

// Class Methods

func (achieve *Achieve) Validate() error {
	messageFmt := "Invalid Achievement: %s"
	// Retrieves the inputted values
	year := achieve.Year
	message := achieve.Message

	// Year validation
	if year < 2000 {
		return fmt.Errorf(messageFmt, "Year must be later than 2000")
	}

	// Message validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, message); !matches {
		return fmt.Errorf(messageFmt, "Message must contain at least an alphabetic letter")
	}

	return nil
}
