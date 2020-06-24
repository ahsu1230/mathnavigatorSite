package domains

import (
	"errors"
	"regexp"
	"time"
)

var TABLE_ASKFORHELP = "askForHelp"

type AskForHelp struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	DeletedAt  NullTime  `json:"-" db:"deleted_at"`
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
	TimeString string    `json:"timeString" db:"time_string"`
	Subject    string    `json:"subject"`
	LocationId string    `json:"locationId" db:"location_id"`
}

func (askForHelp *AskForHelp) Validate() error {
	title := askForHelp.Title
	subject := askForHelp.Subject
	locationId := askForHelp.LocationId

	// Title validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, title); !matches || len(title) > 256 {
		return errors.New("invalid title")
	}

	// Subject validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, subject); !matches || len(title) > 128 {
		return errors.New("invalid subject")
	}

	// Location ID validation
	if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, locationId); !matches {
		return errors.New("invalid AFH location id")
	}
	return nil
}