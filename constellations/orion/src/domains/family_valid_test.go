//ID,password and primary contact email

package domains_test

import (
	"strings"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidPrimaryEmail(t *testing.T) {
	// Checks for valid emails
	family := domains.User{
		PrimaryEmail: "gmail@gmail.com",
		Password: "password",
	}
	if err := family.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	family.PrimaryEmail = "test.test_test+123@example.com"
	if err := user.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid emails
	family.PrimaryEmail = "invalidEmail"
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	}

	family.PrimaryEmail = "email@email" + strings.Repeat("A", 64)
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	} 
}

func TestValidPassword(t *testing.T) {
	// Checks for passwords
	family := domains.User{
		PrimaryEmail: "gmail@gmail.com",
		Password: "password",
	}
	if err := family.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	family.Password = "TestName"
	if err := family.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}


	family.Password = ""
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid password")
	}

	famiy.Password = "Too long" + strings.Repeat("A", 32)
	if err := user.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid password")
	}

}
