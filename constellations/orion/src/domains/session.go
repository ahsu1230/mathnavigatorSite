package domains

import (
	"errors"
	"regexp"
	"time"
)

var TABLE_SESSIONS = "sessions"

type Session struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt time.Time  `json:"-" db:"updated_at"`
	DeletedAt NullTime   `json:"-" db:"deleted_at"`
	ClassId   string     `json:"classId" db:"class_id"`
	StartsAt  time.Time  `json:"startsAt" db:"starts_at"`
	EndsAt    time.Time  `json:"endsAt" db:"ends_at"`
	Canceled  bool       `json:"canceled" db:"canceled"`
	Notes     NullString `json:"notes" db:"notes"`
}

func (session *Session) Validate() error {
	// Retrieves the inputted values
	notes := session.Notes

	// Notes validation
	if notes.Valid {
		if matches, _ := regexp.MatchString(REGEX_LETTER, notes.String); !matches {
			return errors.New("invalid notes")
		}
	}

	return nil
}
