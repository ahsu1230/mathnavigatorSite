package domains

import (
	"fmt"
	"regexp"
	"time"
)

var TABLE_ASKFORHELP = "ask_for_help"

const (
	SUBJECT_MATH        = "math"
	SUBJECT_ENGLISH     = "english"
	SUBJECT_PROGRAMMING = "programming"
)

type AskForHelp struct {
	Id         uint       `json:"id"`
	CreatedAt  time.Time  `json:"-" db:"created_at"`
	UpdatedAt  time.Time  `json:"-" db:"updated_at"`
	DeletedAt  NullTime   `json:"-" db:"deleted_at"`
	Title      string     `json:"title"`
	Date       time.Time  `json:"date"`
	TimeString string     `json:"timeString" db:"time_string"`
	Subject    string     `json:"subject"`
	LocationId string     `json:"locationId" db:"location_id"`
	Notes      NullString `json:"notes"`
}

func (askForHelp *AskForHelp) Validate() error {
	messageFmt := "Invalid AskForHelp: %s"
	title := askForHelp.Title
	subject := askForHelp.Subject

	// Title validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, title); !matches {
		return fmt.Errorf(messageFmt, "Invalid Title")
	}

	// Subject validation
	if subject != SUBJECT_MATH && subject != SUBJECT_ENGLISH && subject != SUBJECT_PROGRAMMING {
		return fmt.Errorf(messageFmt, "Unrecognized subject")
	}
	return nil
}
