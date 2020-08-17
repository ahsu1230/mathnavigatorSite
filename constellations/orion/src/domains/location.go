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
	Street     string     `json:"street"`
	City       string     `json:"city"`
	State      string     `json:"state"`
	Zipcode    string     `json:"zipcode"`
	Room       NullString `json:"room"`
}

func (location *Location) Validate() error {
	messageFmt := "Invalid Location: %s"

	// Retrieves the inputted values
	locationId := location.LocationId
	street := location.Street
	city := location.City
	state := location.State
	zipcode := location.Zipcode
	room := location.Room

	// Location ID validation
	if matches, _ := regexp.MatchString(REGEX_GENERIC_ID, locationId); !matches {
		return fmt.Errorf(messageFmt, "Invalid ID")
	}

	// Street validation
	if matches, _ := regexp.MatchString(REGEX_STREET, street); !matches {
		return fmt.Errorf(messageFmt, "Invalid Street format")
	}

	// City validation
	if matches, _ := regexp.MatchString(REGEX_CITY, city); !matches {
		return fmt.Errorf(messageFmt, "Invalid City name")
	}

	// State validation
	if matches, _ := regexp.MatchString(REGEX_STATE, state); !matches {
		return fmt.Errorf(messageFmt, "Invalid State - must be two capitalized letters")
	}

	// Zipcode validation
	if matches, _ := regexp.MatchString(REGEX_ZIPCODE, zipcode); !matches {
		return fmt.Errorf(messageFmt, "Invalid Zipcode format")
	}

	// Room validation
	if room.Valid {
		if matches, _ := regexp.MatchString(REGEX_ALPHA, room.String); !matches {
			return fmt.Errorf(messageFmt, "Invalid room entry")
		}
	}

	return nil
}
