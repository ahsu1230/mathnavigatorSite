package domains

import (
	"fmt"
	"regexp"
	"time"
)

var TABLE_SEMESTERS = "semesters"

const (
	WINTER = "winter"
	SPRING = "spring"
	SUMMER = "summer"
	FALL   = "fall"
)

var ALL_SEASONS = []string{WINTER, SPRING, SUMMER, FALL}

type Semester struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	DeletedAt  NullTime  `json:"-" db:"deleted_at"`
	SemesterId string    `json:"semesterId" db:"semester_id"`
	Season     string    `json:"season"`
	Year       uint      `json:"year"`
	Title      string    `json:"title"`
}

// Class Methods

func (semester *Semester) Validate() error {
	messageFmt := "Invalid Semester: %s"

	// Retrieves the inputted values
	semesterId := semester.SemesterId
	title := semester.Title
	season := semester.Season
	year := semester.Year

	// Year validation
	if year < 2000 {
		return fmt.Errorf(messageFmt, "Invalid semester year")
	}

	// Season validation
	if season != WINTER && season != SPRING && season != SUMMER && season != FALL {
		return fmt.Errorf(messageFmt, "Invalid semester season")
	}

	// Semester ID validation
	if matches, _ := regexp.MatchString(REGEX_SEMESTER_ID, semesterId); !matches {
		return fmt.Errorf(messageFmt, "Invalid Semester ID (should be format `year_season`)")
	}

	// Title validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, title); !matches {
		return fmt.Errorf(messageFmt, "Invalid Semester Title")
	}

	return nil
}
