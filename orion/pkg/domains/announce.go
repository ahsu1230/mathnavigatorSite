package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_ANNOUNCEMENTS = "announcements"

type Announce struct {
	Id        uint         `db:"id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"-"`
	PostedAt  time.Time    `db:"posted_at"`
	Author    string       `json:"author"`
	Message   string       `json:"message"`
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
