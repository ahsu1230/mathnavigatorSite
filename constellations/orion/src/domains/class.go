package domains

import (
	"errors"
	"regexp"
	"time"
)

var TABLE_CLASSES = "classes"

type Class struct {
	Id              uint       `json:"id"`
	CreatedAt       time.Time  `json:"-" db:"created_at"`
	UpdatedAt       time.Time  `json:"-" db:"updated_at"`
	DeletedAt       NullTime   `json:"-" db:"deleted_at"`
	PublishedAt     NullTime   `json:"publishedAt" db:"published_at"`
	ProgramId       string     `json:"programId" db:"program_id"`
	SemesterId      string     `json:"semesterId" db:"semester_id"`
	ClassKey        NullString `json:"classKey" db:"class_key"`
	ClassId         string     `json:"classId" db:"class_id"`
	LocationId      string     `json:"locationId" db:"location_id"`
	Times           string     `json:"times"`
	StartDate       time.Time  `json:"startDate" db:"start_date"`
	EndDate         time.Time  `json:"endDate" db:"end_date"`
	GoogleClassCode NullString `json:"googleClassCode" db:"google_class_code"`
	FullState       int        `json:"fullState" db:"full_state"`
	PricePerSession NullUint   `json:"pricePerSession" db:"price_per_session"`
	PriceLump       NullUint   `json:"priceLump" db:"price_lump"`
}

// Class Methods

func (class *Class) Validate() error {
	// Retrieves the inputted values
	classKey := class.ClassKey
	times := class.Times
	startDate := class.StartDate
	endDate := class.EndDate

	// Class Key validation
	if classKey.Valid {
		if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, classKey.String); !matches || len(classKey.String) > 64 {
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
