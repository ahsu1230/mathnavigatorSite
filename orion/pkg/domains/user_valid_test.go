package domains_test

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"strings"
	"testing"
)

func TestValidFirstName(t *testing.T) {
	// Checks for valid first names
	program := domains.User{FirstName: "John", LastName: "Smith", MiddleName: "Middle", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true, GuardianId: 1}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.FirstName = "TestName"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid first names
	program.FirstName = ""
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid first name")
	}

	program.FirstName = "Too long" + strings.Repeat("A", 32)
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid first name")
	}
}

func TestValidLastName(t *testing.T) {
	// Checks for valid last names
	program := domains.User{FirstName: "John", LastName: "Smith", MiddleName: "Middle", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true, GuardianId: 1}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.LastName = "TestName"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid last names
	program.LastName = ""
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid last name")
	}

	program.LastName = "Too long" + strings.Repeat("A", 32)
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid last name")
	}
}

func TestValidMiddleName(t *testing.T) {
	// Checks for valid middle names
	program := domains.User{FirstName: "John", LastName: "Smith", MiddleName: "Middle", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true, GuardianId: 1}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.MiddleName = ""
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid middle names
	program.MiddleName = "Too long" + strings.Repeat("A", 32)
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid middle name")
	}
}

func TestValidEmail(t *testing.T) {
	// Checks for valid emails
	program := domains.User{FirstName: "John", LastName: "Smith", MiddleName: "Middle", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true, GuardianId: 1}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Email = "test.test_test+123@example.com"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid emails
	program.Email = "invalidEmail"
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	}

	program.Email = "email@email" + strings.Repeat("A", 64)
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	}
}

func TestValidPhone(t *testing.T) {
	// Checks for valid phone numbers
	program := domains.User{FirstName: "John", LastName: "Smith", MiddleName: "Middle", Email: "gmail@gmail.com", Phone: "555-555-0100", IsGuardian: true, GuardianId: 1}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Phone = "+1 (555) 555 0100"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid phone numbers
	program.Phone = "#$^3rdg4@#4&%$^%8dfg^&*^%^45#$%"
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid phone")
	}

	program.Phone = "555" + strings.Repeat("5", 24)
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid phone")
	}
}
