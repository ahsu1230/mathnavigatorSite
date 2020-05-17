package domains_test

import (
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidSemesterId(t *testing.T) {
	// Checks for valid semester IDs
	semester := domains.Semester{SemesterId: "2019_fall", Title: "Fall 2019"}
	if err := semester.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	semester.SemesterId = "2050_winter"
	if err := semester.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid semester IDs
	semester.SemesterId = "0999_spring"
	if err := semester.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid semester id")
	}

	semester.SemesterId = "2019_wrong"
	if err := semester.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid semester id")
	}
}

func TestValidTitle(t *testing.T) {
	// Checks for valid titles
	semester := domains.Semester{SemesterId: "2020_spring", Title: "Test title"}
	if err := semester.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	semester.Title = "Winter 2222 (#1 class)"
	if err := semester.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid titles
	semester.Title = "This title is too long: " + strings.Repeat("A", 255)
	if err := semester.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid semester title")
	}

	semester.Title = "Fall_ _2020"
	if err := semester.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid semester title")
	}
}
