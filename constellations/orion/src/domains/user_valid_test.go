package domains_test

import (
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidFirstName(t *testing.T) {
	// Checks for valid first names
	user := domains.User{
		FirstName:  "John",
		LastName:   "Smith",
		Email:      "gmail@gmail.com",
		Phone:      "555-555-0100",
		IsGuardian: true,
	}
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

}

func TestValidLastName(t *testing.T) {
	// Checks for valid last names
	user := domains.User{
		FirstName:  "John",
		LastName:   "Smith",
		Email:      "gmail@gmail.com",
		Phone:      "555-555-0100",
		IsGuardian: true,
	}
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

}

func TestValidEmail(t *testing.T) {
	// Checks for valid emails
	user := domains.User{
		FirstName:  "John",
		LastName:   "Smith",
		Email:      "gmail@gmail.com",
		Phone:      "555-555-0100",
		IsGuardian: true,
	}
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

}

func TestValidPhone(t *testing.T) {
	// Checks for valid phone numbers
	user := domains.User{
		FirstName:  "John",
		LastName:   "Smith",
		Email:      "gmail@gmail.com",
		Phone:      "555-555-0100",
		IsGuardian: true,
	}
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

}

func TestValidSchool(t *testing.T) {
	user := domains.User{
		FirstName:  "John",
		LastName:   "Smith",
		Email:      "gmail@gmail.com",
		Phone:      "555-555-0100",
		IsGuardian: false,
	}
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	//Test valid school
	user.School = domains.NewNullString("Churchill H.S.")
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	user.School = domains.NewNullString("Winston churchill High school")
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	user.School = domains.NewNullString("Montgomery Blair High school")
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	user.School = domains.NewNullString("Thomas-Jefferson H.S.")
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	//Test invalid school
	user.School = domains.NewNullString("school12992")
	if err := user.Validate(); err == nil {
		t.Errorf("Check was incorrect, got: nil, expected: school contains non alphabetical characters")
	}
}

func TestValidGradYear(t *testing.T) {
	user := domains.User{
		FirstName:  "John",
		LastName:   "Smith",
		Email:      "gmail@gmail.com",
		Phone:      "555-555-0100",
		IsGuardian: false,
	}
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	//Test valid year
	user.GraduationYear = domains.NewNullUint(2022)
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	user.GraduationYear = domains.NewNullUint(2018)
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	user.GraduationYear = domains.NewNullUint(2031)
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	//Test invalid school
	user.GraduationYear = domains.NewNullUint(1992)
	if err := user.Validate(); err == nil {
		t.Errorf("Check was incorrect, got: nil, expected: invalid graduation year")
	}

}
