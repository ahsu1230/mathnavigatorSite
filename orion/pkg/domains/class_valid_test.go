package domains_test

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"testing"
	"time"
)

func TestValidClassKey(t *testing.T) {
	// Checks for valid class keys
	now := time.Now().UTC()
	later := now.Add(time.Hour * 24 * 100)
	class := domains.Class{
		ClassKey:  "final_review",
		Times:     "3 pm - 5 pm",
		StartDate: now,
		EndDate:   later,
	}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.ClassKey = ""
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.ClassKey = "Valid_Class_100"
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid class keys
	class.ClassKey = "a__a"
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class key")
	}

	class.ClassKey = "_"
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class key")
	}
}

func TestValidTimes(t *testing.T) {
	// Checks for valid times
	now := time.Now().UTC()
	later := now.Add(time.Hour * 24 * 100)
	class := domains.Class{
		ClassKey:  "final_review",
		Times:     "6 pm - 8 pm",
		StartDate: now,
		EndDate:   later,
	}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.Times = "8"
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid times
	class.Times = ""
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class name")
	}

	class.Times = "Hi"
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class name")
	}
}

func TestValidDates(t *testing.T) {
	// Checks for valid dates
	now := time.Now().UTC()
	later := now.Add(time.Hour * 24 * 30)
	class := domains.Class{
		ClassKey:  "final_review",
		Times:     "3 pm - 5 pm",
		StartDate: now,
		EndDate:   later,
	}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.StartDate = now
	class.EndDate = now.Add(time.Hour * 24 * 365 * 100)
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid dates
	class.StartDate = time.Unix(0, 0)
	class.EndDate = now
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}

	class.StartDate = now.Add(time.Hour * 24)
	class.EndDate = now
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}
}