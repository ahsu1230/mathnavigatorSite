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
	GoogleClassCode NullString `json:"googleClassCode" db:"google_class_code"`
	FullState       int        `json:"fullState" db:"full_state"`
	PricePerSession NullUint   `json:"pricePerSession" db:"price_per_session"`
	PriceLump       NullUint   `json:"priceLump" db:"price_lump"`
	PaymentNotes    NullString `json:"paymentNotes" db:"payment_notes"`
}

// Class Methods

func (class *Class) Validate() error {
	// Retrieves the inputted values
	classKey := class.ClassKey
	times := class.Times
	pricePerSession := class.PricePerSession
	priceLump := class.PriceLump

	// Class Key validation
	if classKey.Valid {
		if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, classKey.String); !matches {
			return appErrors.WrapInvalidDomain("Invalid Class Key")
		}
	}

	// Times validation
	if matches, _ := regexp.MatchString(REGEX_NUMBER, times); !matches {
		return appErrors.WrapInvalidDomain("Invalid Time Format")
	}

	//Price validation
	//Both valid
	if priceLump.Valid && pricePerSession.Valid {
		return appErrors.WrapInvalidDomain("Cannot have both priceLump and pricePerSession defined")
	}

	return nil
}
