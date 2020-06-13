package domains

import (
	"errors"
	"regexp"
	"time"
)

var TABLE_ANNOUNCEMENTS = "announcements"

type Announce struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	DeletedAt  NullTime  `json:"-" db:"deleted_at"`
	PostedAt   time.Time `json:"postedAt" db:"posted_at"`
	Author     string    `json:"author"`
	Message    string    `json:"message"`
	OnHomePage bool      `json:"onHomePage"`
}

func (announce *Announce) Validate() error {
	// Retrieves the inputted values
	author := announce.Author
	message := announce.Message

	// Author validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, author); !matches {
		return errors.New("invalid author")
	}

	// Message validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, message); !matches {
		return errors.New("invalid message")
	}

	return nil
}
