//ID,password and primary contact email

package domains_test

import (
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func Test_ValidPrimaryEmail(t *testing.T) {
	// Checks for valid emails
	family := domains.Family{
		PrimaryEmail: "gmail@gmail.com",
		Password:     "password",
	}
	if err := family.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	family.PrimaryEmail = "test.test_test+123@example.com"
	if err := family.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid emails
	family.PrimaryEmail = "invalidEmail"
	if err := family.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	}
}

func Test_ValidPassword(t *testing.T) {
	// Checks for passwords
	family := domains.Family{
		PrimaryEmail: "gmail@gmail.com",
		Password:     "password",
	}
	if err := family.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	family.Password = "TestName"
	if err := family.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	family.Password = ""
	if err := family.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid password")
	}

	family.Password = "1234567"
	if err := family.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid password")
	}
}
