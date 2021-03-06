package domains_test

import (
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidAFHTitle(t *testing.T) {
	now := time.Now().UTC()
	start := now.Add(time.Hour * 24 * 30)
	end := start.Add(time.Hour * 1)

	askForHelp := domains.AskForHelp{
		Id:         1,
		Title:      "AP Calculus Help",
		StartsAt:   start,
		EndsAt:     end,
		Subject:    domains.SUBJECT_MATH,
		LocationId: "wchs",
	}
	if err := askForHelp.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	askForHelp.Title = ""
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}

	askForHelp.Title = "ap calc help"
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}
}

func TestValidAFHSubject(t *testing.T) {
	now := time.Now().UTC()
	start := now.Add(time.Hour * 24 * 30)
	end := start.Add(time.Hour * 1)
	askForHelp := domains.AskForHelp{
		Id:         1,
		Title:      "AP Calculus Help",
		StartsAt:   start,
		EndsAt:     end,
		Subject:    domains.SUBJECT_MATH,
		LocationId: "wchs",
	}
	if err := askForHelp.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	askForHelp.Subject = ""
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid subject")
	}

	askForHelp.Subject = "history" // not a valid subject
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid subject")
	}
}

func TestValidAFHTimes(t *testing.T) {
	now := time.Now().UTC()
	start := now.Add(time.Hour * 24 * 30)
	end := start.Add(time.Hour * 1)
	askForHelp := domains.AskForHelp{
		Id:         1,
		Title:      "AP Calculus Help",
		StartsAt:   start,
		EndsAt:     end,
		Subject:    domains.SUBJECT_MATH,
		LocationId: "wchs",
	}
	if err := askForHelp.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	askForHelp.EndsAt = start.Add(time.Hour * -1)
	if err := askForHelp.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid start/end times")
	}
}
