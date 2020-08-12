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
		return errors.New("invalid location id")
	}

	// Street validation
	if matches, _ := regexp.MatchString(REGEX_STREET, street); !matches {
		return errors.New("invalid street")
	}

	// City validation
	if matches, _ := regexp.MatchString(REGEX_CITY, city); !matches {
		return errors.New("invalid city")
	}

	// State validation
	if matches, _ := regexp.MatchString(REGEX_STATE, state); !matches {
		return errors.New("invalid state")
	}

	// Zipcode validation
	if matches, _ := regexp.MatchString(REGEX_ZIPCODE, zipcode); !matches {
		return errors.New("invalid zipcode")
	}

	// Room validation
	if room.Valid {
		if matches, _ := regexp.MatchString(REGEX_ALPHA, room.String); !matches {
			return errors.New("invalid room")
		}
	}

	return nil
}
