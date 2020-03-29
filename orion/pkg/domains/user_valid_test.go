package domains_test

import (
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
)

func TestValidFirstName(t *testing.T) {
	// Checks for valid first names
	user := domains.User{FirstName: "John", LastName: "Smith", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true}
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	user.FirstName = "TestName"
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid first names
	user.FirstName = ""
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid first name")
	}

	user.FirstName = "Too long" + strings.Repeat("A", 32)
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid first name")
	}
}

func TestValidLastName(t *testing.T) {
	// Checks for valid last names
	user := domains.User{FirstName: "John", LastName: "Smith", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true}
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	user.LastName = "TestName"
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid last names
	user.LastName = ""
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid last name")
	}

	user.LastName = "Too long" + strings.Repeat("A", 32)
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid last name")
	}
}

func TestValidEmail(t *testing.T) {
	// Checks for valid emails
	user := domains.User{FirstName: "John", LastName: "Smith", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true}
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	user.Email = "test.test_test+123@example.com"
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid emails
	user.Email = "invalidEmail"
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	}

	user.Email = "email@email" + strings.Repeat("A", 64)
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	}
}

func TestValidPhone(t *testing.T) {
	// Checks for valid phone numbers
	user := domains.User{FirstName: "John", LastName: "Smith", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true}
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	user.Phone = "+1 (555) 555 0100"
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid phone numbers
	user.Phone = "#$^3rdg4@#4&%$^%8dfg^&*^%^45#$%"
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid phone")
	}

	user.Phone = "555" + strings.Repeat("5", 24)
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid phone")
	}
}

func TestValidGuardianId(t *testing.T) {
	// Checks for valid guardian ids
	user := domains.User{
		FirstName:  "John",
		LastName:   "Smith",
		Email:      "gmail@gmail.com",
		Phone:      "555-555-0100",
		IsGuardian: false,
		GuardianId: domains.NewNullUint(2),
	}
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid guardian ids
	user.GuardianId = domains.NewNullUint(0)
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid guardian id")
	}
}
