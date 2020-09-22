package domains

import (
	"fmt"
	"regexp"
	"time"
)

var TABLE_CLASSES = "classes"

const (
	NOT_FULL    = 0
	ALMOST_FULL = 1
	FULL        = 2

	NOT_FULL_DISPLAY_NAME    = "Not full"
	ALMOST_FULL_DISPLAY_NAME = "Almost full"
	FULL_DISPLAY_NAME        = "Full"
)

type Class struct {
	Id              uint       `json:"id"`
	CreatedAt       time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt       NullTime   `json:"-" db:"deleted_at"`
	PublishedAt     NullTime   `json:"publishedAt" db:"published_at"`
	ProgramId       string     `json:"programId" db:"program_id"`
	SemesterId      string     `json:"semesterId" db:"semester_id"`
	ClassKey        NullString `json:"classKey" db:"class_key"`
	ClassId         string     `json:"classId" db:"class_id"`
	LocationId      string     `json:"locationId" db:"location_id"`
	TimesStr        string     `json:"timesStr" db:"times_str"`
	GoogleClassCode NullString `json:"googleClassCode" db:"google_class_code"`
	FullState       uint       `json:"fullState" db:"full_state"`
	PricePerSession NullUint   `json:"pricePerSession" db:"price_per_session"`
	PriceLumpSum    NullUint   `json:"priceLumpSum" db:"price_lump_sum"`
	PaymentNotes    NullString `json:"paymentNotes" db:"payment_notes"`
}

// Class Methods

func (class *Class) Validate() error {
	messageFmt := "Invalid Class: %s"

	// Retrieves the inputted values
	classKey := class.ClassKey
	times := class.TimesStr
	pricePerSession := class.PricePerSession
	priceLump := class.PriceLumpSum

	// Class Key validation
	if classKey.Valid {
		if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, classKey.String); !matches {
			return fmt.Errorf(messageFmt, "Invalid Class Key")
		}
	}

	// Times validation
	if matches, _ := regexp.MatchString(REGEX_NUMBER, times); !matches {
		return fmt.Errorf(messageFmt, "Invalid Time Format")
	}

	// Price validation
	if priceLump.Valid == pricePerSession.Valid {
		return fmt.Errorf(messageFmt, "Only One Price Can Be Defined")
	}

	// Full state validation
	if class.FullState >= 3 || class.FullState < 0 {
		return fmt.Errorf(messageFmt, "Invalid full state")
	}

	return nil
}
