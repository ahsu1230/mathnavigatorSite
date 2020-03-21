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
	DeletedAt sql.NullTime `db:"deleted_at"`
	ClassId   string       `db:"class_id" json:"classId"`
	StartsAt  time.Time    `db:"starts_at"`
	EndsAt    time.Time    `db:"ends_at"`
	Canceled  bool         `json:"canceled"`
	Notes     string       `json:"notes"`
}

func (session *Session) Validate() error {
	// Retrieves the inputted values
	classId := session.ClassId
	notes := session.Notes

	// Class ID validation
	if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, classId); !matches {
		return errors.New("invalid class id")
	}

	// Notes validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, notes); !matches {
		return errors.New("invalid notes")
	}

	return nil
}
