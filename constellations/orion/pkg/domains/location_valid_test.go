package domains_test

import (
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
)

func TestValidLocationId(t *testing.T) {
	// Checks for valid location IDs
	location := domains.Location{
		LocationId: "xkcd",
		Street:     "4040 Cherry Rd",
		City:       "Potomac",
		State:      "MD",
		Zipcode:    "20854",
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.LocationId = "I_am_a_legitimate_id"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.LocationId = "a"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.LocationId = "23T"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid location IDs
	location.LocationId = "spaces are not allowed"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid location id")
	}

	location.LocationId = "112@_asdf"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid location id")
	}

	location.LocationId = "a_A_id_9__"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid location id")
	}

	location.LocationId = "__k__"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid location id")
	}
}

func TestValidLocationStreet(t *testing.T) {
	// Checks for valid streets
	location := domains.Location{
		LocationId: "xkcd",
		Street:     "4040 Cherry Rd",
		City:       "Potomac",
		State:      "MD",
		Zipcode:    "20854",
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Street = "12345 Great Terrace Drive"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Street = "78 Palace Blvd W"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Street = "54234 Nowhere Pl"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid streets
	location.Street = "48 Place"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid street")
	}

	location.Street = "Imposter Ave"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid street")
	}

	location.Street = "11285"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid street")
	}

	location.Street = ""
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid street")
	}
}

func TestValidLocationCity(t *testing.T) {
	// Checks for valid cities
	location := domains.Location{
		LocationId: "xkcd",
		Street:     "4040 Cherry Rd",
		City:       "Potomac",
		State:      "MD",
		Zipcode:    "20854",
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.City = "Rockville"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.City = "Driftveil City"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.City = "Sky World"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid cities
	location.City = "ss Potomac"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid city")
	}

	location.City = "1234"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid city")
	}

	location.City = "G0th@m City"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid city")
	}

	location.City = ""
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid city")
	}
}

func TestValidLocationState(t *testing.T) {
	// Checks for valid states
	location := domains.Location{
		LocationId: "xkcd",
		Street:     "4040 Cherry Rd",
		City:       "Potomac",
		State:      "MD",
		Zipcode:    "20854",
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.State = "VA"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid states
	location.State = "MVA"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid state")
	}

	location.State = "md"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid state")
	}
}

func TestValidLocationZipcode(t *testing.T) {
	// Checks for valid zipcodes
	location := domains.Location{
		LocationId: "xkcd",
		Street:     "4040 Cherry Rd",
		City:       "Potomac",
		State:      "MD",
		Zipcode:    "20854",
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Zipcode = "09801-2391"
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid zipcodes
	location.Zipcode = "481234"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid zipcode")
	}

	location.Zipcode = "12341-1233-1"
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid zipcode")
	}
}

func TestValidLocationRoom(t *testing.T) {
	// Checks for valid rooms
	location := domains.Location{
		LocationId: "xkcd",
		Street:     "4040 Cherry Rd",
		City:       "Potomac",
		State:      "MD",
		Zipcode:    "20854",
		Room:       domains.NewNullString("Room 2"),
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Room = domains.NewNullString("124")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Room = domains.NewNullString("Auditorium")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Room = domains.NewNullString("")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid rooms
	location.Room = domains.NewNullString("@@@")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid room")
	}

	location.Room = domains.NewNullString("#@!*")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid room")
	}
}
