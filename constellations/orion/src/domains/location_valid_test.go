package domains_test

import (
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidLocationId(t *testing.T) {
	// Checks for valid location IDs
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
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

func TestValidLocationTitle(t *testing.T) {
	// Checks for valid title
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Check for invalid title
	location.Title = ""
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}
}

func TestValidLocationStreet(t *testing.T) {
	// Checks for valid streets
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Street = domains.NewNullString("12345 Great Terrace Drive")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Street = domains.NewNullString("78 Palace Blvd W")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Street = domains.NewNullString("54234 Nowhere Pl")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Street = domains.NewNullString("")
	location.IsOnline = true
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid streets
	location.Street = domains.NewNullString("48 Place")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid street")
	}

	location.Street = domains.NewNullString("Imposter Ave")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid street")
	}

	location.Street = domains.NewNullString("11285")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid street")
	}
}

func TestValidLocationCity(t *testing.T) {
	// Checks for valid cities
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.City = domains.NewNullString("Rockville")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.City = domains.NewNullString("Driftveil City")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.City = domains.NewNullString("Sky World")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.City = domains.NewNullString("")
	location.IsOnline = true
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	// Checks for invalid cities
	location.City = domains.NewNullString("ss Potomac")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid city")
	}

	location.City = domains.NewNullString("1234")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid city")
	}

	location.City = domains.NewNullString("G0th@m City")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid city")
	}
}

func TestValidLocationState(t *testing.T) {
	// Checks for valid states
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.State = domains.NewNullString("VA")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid states
	location.State = domains.NewNullString("MVA")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid state")
	}

	location.State = domains.NewNullString("md")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid state")
	}
}

func TestValidLocationZipcode(t *testing.T) {
	// Checks for valid zipcodes
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	location.Zipcode = domains.NewNullString("09801-2391")
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid zipcodes
	location.Zipcode = domains.NewNullString("481234")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid zipcode")
	}

	location.Zipcode = domains.NewNullString("12341-1233-1")
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid zipcode")
	}
}

func TestValidLocationRoom(t *testing.T) {
	// Checks for valid rooms
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
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

func TestValidLocationIsOnline(t *testing.T) {
	// Checks for valid online location
	location := domains.Location{
		LocationId: "xkcd",
		Title:      "Online Zoom",
		Street:     domains.NewNullString(""),
		City:       domains.NewNullString(""),
		State:      domains.NewNullString(""),
		Zipcode:    domains.NewNullString(""),
		IsOnline:   true,
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for valid physical location
	location = domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString("4040 Cherry Rd"),
		City:       domains.NewNullString("Potomac"),
		State:      domains.NewNullString("MD"),
		Zipcode:    domains.NewNullString("20854"),
		IsOnline:   false,
	}
	if err := location.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Check for invalid location (marked not-online, but no physical address)
	location = domains.Location{
		LocationId: "xkcd",
		Title:      "School1",
		Street:     domains.NewNullString(""),
		City:       domains.NewNullString(""),
		State:      domains.NewNullString(""),
		Zipcode:    domains.NewNullString(""),
		IsOnline:   false,
	}
	if err := location.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}
}
