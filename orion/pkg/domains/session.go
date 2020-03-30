package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_SESSIONS = "sessions"

type Session struct {
	Id        uint
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"-"`
	ClassId   string       `db:"class_id" json:"classId"`
	StartsAt  time.Time    `json:"startsAt"`
	EndsAt    time.Time    `json:"endsAt"`
	Canceled  bool         `json:"canceled"`
	Notes     NullString   `json:"notes"`
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
