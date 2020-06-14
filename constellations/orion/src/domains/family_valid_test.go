//ID,password and primary contact email

package domains_test

import (
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidPrimaryEmail(t *testing.T) {
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

	family.PrimaryEmail = "email@email" + strings.Repeat("A", 64)
	if err := family.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	}
}

func TestValidPassword(t *testing.T) {
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

	// family.Password = "Too long" + strings.Repeat("A", 32)
	// if err := family.Validate(); err == nil {
	// 	t.Error("Check was incorrect, got: nil, expected: invalid password")
	// }

}
