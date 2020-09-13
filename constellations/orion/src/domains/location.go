package domains

import (
	"fmt"
	"regexp"
	"time"
)

var TABLE_LOCATIONS = "locations"

type Location struct {
	Id         uint       `json:"id"`
	CreatedAt  time.Time  `json:"-" db:"created_at"`
	UpdatedAt  time.Time  `json:"-" db:"update_at"`
	DeletedAt  NullTime   `json:"-" db:"deleted_at"`
	LocationId string     `json:"locationId" db:"location_id"`
	Title      string     `json:"title"`
	Street     NullString `json:"street"`
	City       NullString `json:"city"`
	State      NullString `json:"state"`
	Zipcode    NullString `json:"zipcode"`
	Room       NullString `json:"room"`
	IsOnline   bool       `json:"isOnline" db:"is_online"`
}

func (location *Location) Validate() error {
	messageFmt := "Invalid Location: %s"

	// Retrieves the inputted values
	locationId := location.LocationId
	title := location.Title
	street := location.Street
	city := location.City
	state := location.State
	zipcode := location.Zipcode
	room := location.Room

	// Location ID validation
	if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, locationId); !matches {
		return fmt.Errorf(messageFmt, "Invalid ID")
	}

	// Title validation
	if matches, _ := regexp.MatchString(REGEX_TITLE, title); !matches {
		return fmt.Errorf(messageFmt, "Invalid Title")
	}

	// IsOnline validation
	// If class is NOT online, street, city, etc. must all be filled.
	if !location.IsOnline && (!street.Valid || !city.Valid || !state.Valid || !zipcode.Valid) {
		return fmt.Errorf(messageFmt, "Non-online class MUST have a physical address")
	}

	// Street validation
	if street.Valid {
		if matches, _ := regexp.MatchString(REGEX_STREET, street.String); !matches {
			return fmt.Errorf(messageFmt, "Invalid Street format")
		}
	}

	// City validation
	if city.Valid {
		if matches, _ := regexp.MatchString(REGEX_CITY, city.String); !matches {
			return fmt.Errorf(messageFmt, "Invalid City name")
		}
	}

	// State validation
	if state.Valid {
		if matches, _ := regexp.MatchString(REGEX_STATE, state.String); !matches {
			return fmt.Errorf(messageFmt, "Invalid State - must be two capitalized letters")
		}
	}

	// Zipcode validation
	if zipcode.Valid {
		if matches, _ := regexp.MatchString(REGEX_ZIPCODE, zipcode.String); !matches {
			return fmt.Errorf(messageFmt, "Invalid Zipcode format")
		}
	}

	// Room validation
	if room.Valid {
		if matches, _ := regexp.MatchString(REGEX_ALPHA, room.String); !matches {
			return fmt.Errorf(messageFmt, "Invalid room entry")
		}
	}

	return nil
}
