package domains_test

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"testing"
)

func TestValidYear(t *testing.T) {
	// Checks for valid years
	achieve := domains.Achieve{Year: 2020, Message: "This is a message"}
	if err := achieve.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid years
	achieve.Year = 100
	if err := achieve.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid year")
	}
}

func TestValidAchievementMessage(t *testing.T) {
	// Checks for valid messages
	achieve := domains.Achieve{Year: 2050, Message: "Hello World!"}
	if err := achieve.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid messages
	achieve.Message = ""
	if err := achieve.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid message")
	}
}
