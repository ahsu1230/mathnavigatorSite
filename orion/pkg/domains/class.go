package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_CLASSES = "classes"

type Class struct {
	Id         uint
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	ProgramId  string       `db:"program_id" json:"programId"`
	SemesterId string       `db:"semester_id" json:"semesterId"`
	ClassKey   string       `db:"class_key" json:"classKey"`
	ClassId    string       `db:"class_id" json:"classId"`
	LocationId string       `db:"location_id" json:"locationId"`
	Times      string       `json:"times"`
	StartDate  time.Time    `json:"startDate"`
	EndDate    time.Time    `json:"endDate"`
}

// Class Methods

func (class *Class) Validate() error {
	// Retrieves the inputted values
	classKey := class.ClassKey
	times := class.Times
	startDate := class.StartDate
	endDate := class.EndDate

	// Class Key validation
	if len(classKey) != 0 {
		if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, classKey); !matches || len(classKey) > 64 {
			return errors.New("invalid class key")
		}
	}

	// Times validation
	if matches, _ := regexp.MatchString(REGEX_NUMBER, times); !matches || len(times) > 64 {
		return errors.New("invalid times")
	}

	// Start Date validation
	if startDate.Year() < 2000 {
		return errors.New("invalid start date")
	}

	// End Date validation
	if !endDate.After(startDate) {
		return errors.New("invalid end date")
	}

	return nil
}
