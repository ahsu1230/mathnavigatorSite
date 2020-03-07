package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

type Announce struct {
	Id			uint			`db:"id"`
	CreatedAt	time.Time		`db:"created_at"`
	UpdatedAt	time.Time		`db:"updated_at"`
	DeletedAt	sql.NullTime	`db:"deleted_at"`
	PostedAt	time.Time		`db:"posted_at"`
	Author		string			`json:"author"`
	Message		string			`json:"message"`
}

// Ensures at least one uppercase or lowercase letter
const REGEX_LETTER = `[A-Za-z]+`

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
