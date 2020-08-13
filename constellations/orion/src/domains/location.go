package domains

import (
	"errors"
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
	Street     string     `json:"street"`
	City       string     `json:"city"`
	State      string     `json:"state"`
	Zipcode    string     `json:"zipcode"`
	Room       NullString `json:"room"`
}

func (location *Location) Validate() error {
	// Retrieves the inputted values
	locationId := location.LocationId
	street := location.Street
	city := location.City
	state := location.State
	zipcode := location.Zipcode
	room := location.Room

	// Location ID validation
	if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, locationId); !matches {
		return appErrors.WrapInvalidDomain("Invalid ID")
	}

	// Street validation
	if matches, _ := regexp.MatchString(REGEX_STREET, street); !matches {
		return appErrors.WrapInvalidDomain("Invalid Street format")
	}

	// City validation
	if matches, _ := regexp.MatchString(REGEX_CITY, city); !matches {
		return appErrors.WrapInvalidDomain("Invalid City name")
	}

	// State validation
	if matches, _ := regexp.MatchString(REGEX_STATE, state); !matches {
		return appErrors.WrapInvalidDomain("Invalid State - must be two capitalized letters")
	}

	// Zipcode validation
	if matches, _ := regexp.MatchString(REGEX_ZIPCODE, zipcode); !matches {
		return appErrors.WrapInvalidDomain("Invalid Zipcode format")
	}

	// Room validation
	if room.Valid {
		if matches, _ := regexp.MatchString(REGEX_ALPHA, room.String); !matches {
			return appErrors.WrapInvalidDomain("Invalid room entry")
		}
	}

	return nil
}
