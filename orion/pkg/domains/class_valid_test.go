package domains_test

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"testing"
	"time"
)

func TestValidClassKey(t *testing.T) {
	// Checks for valid class keys
	class := domains.Class{
		ClassKey:  sql.NullString{String: "final_review", Valid: true},
		Times:     "3 pm - 5 pm",
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Hour * 24 * 100),
	}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.ClassKey = sql.NullString{}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.ClassKey = sql.NullString{String: "Valid_Class_100", Valid: true}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid class keys
	class.ClassKey = sql.NullString{String: "a__a", Valid: true}
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class key")
	}

	class.ClassKey = sql.NullString{String: "_", Valid: true}
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class key")
	}

	class.ClassKey = sql.NullString{String: "", Valid: true}
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class key")
	}
}

func TestValidTimes(t *testing.T) {
	// Checks for valid times
	class := domains.Class{
		ClassKey:  sql.NullString{String: "final_review", Valid: true},
		Times:     "6 pm - 8 pm",
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Hour * 24 * 100),
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
	class := domains.Class{
		ClassKey:  sql.NullString{String: "final_review", Valid: true},
		Times:     "3 pm - 5 pm",
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Hour * 24 * 30),
	}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.StartDate = time.Now()
	class.EndDate = time.Now().Add(time.Hour * 24 * 365 * 100)
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid dates
	class.StartDate = time.Unix(0, 0)
	class.EndDate = class.StartDate.Add(time.Hour * 24 * 60)
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}

	class.StartDate = time.Now().Add(time.Hour * 24)
	class.EndDate = time.Now()
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}
}
