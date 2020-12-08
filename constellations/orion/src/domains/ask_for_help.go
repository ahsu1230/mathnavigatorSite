package domains

import (
	"fmt"
	"regexp"
	"time"
)

var TABLE_ASKFORHELP = "ask_for_help"

type AskForHelp struct {
	Id         uint       `json:"id"`
	CreatedAt  time.Time  `json:"-" db:"created_at"`
	UpdatedAt  time.Time  `json:"-" db:"updated_at"`
	DeletedAt  NullTime   `json:"-" db:"deleted_at"`
	StartsAt   time.Time  `json:"startsAt" db:"starts_at"`
	EndsAt     time.Time  `json:"endsAt" db:"ends_at"`
	Title      string     `json:"title"`
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
	if !validateSubject(subject) {
		return fmt.Errorf(messageFmt, "Unrecognized subject")
	}

	// Time validation
	if askForHelp.StartsAt.After(askForHelp.EndsAt) {
		return fmt.Errorf(messageFmt, "Start time must be before end time")
	}

	return nil
}
