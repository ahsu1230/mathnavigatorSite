package domains_test

import (
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidAFHTitle(t *testing.T) {
	askForHelp := domains.AskForHelp{
		Id:         1,
		Title:      "AP Calculus Help",
		Date:       "August 20, 2020",
		TimeString: "3:00 - 5:00 PM",
		Subject:    "AP Calculus",
		LocationId: "wchs",
	}
	if err := askForHelp.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	askForHelp.Title = ""
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}

	askForHelp.Title = "Too long: " + strings.Repeat("A", 256)
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}
}

func TestValidAFHSubject(t *testing.T) {
	askForHelp := domains.AskForHelp{
		Id:         1,
		Title:      "AP Calculus Help",
		Date:       "August 20, 2020",
		TimeString: "3:00 - 5:00 PM",
		Subject:    "AP Calculus",
		LocationId: "wchs",
	}

	askForHelp.Subject = "AP Calculus 2"
	if err := askForHelp.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	askForHelp.Subject = ""
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid subject")
	}

	askForHelp.Subject = "Too long: " + strings.Repeat("A", 128)
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid subject")
	}
}

func TestValidAFHLocationId(t *testing.T) {
	askForHelp := domains.AskForHelp{
		Id:         1,
		Title:      "AP Calculus Help",
		Date:       "August 20, 2020",
		TimeString: "3:00 - 5:00 PM",
		Subject:    "AP Calculus",
		LocationId: "wchs",
	}

	askForHelp.LocationId = "churchill_high_school"
	if err := askForHelp.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	askForHelp.LocationId = "Room_23"
	if err := askForHelp.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	askForHelp.LocationId = "Room 23 Space"
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid AFH location id")
	}
}
