//ID,password and primary contact email

package domains_test

import (
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func Test_ValidPrimaryEmail(t *testing.T) {
	// Checks for valid emails
	account := domains.Account{
		PrimaryEmail: "gmail@gmail.com",
		Password:     "password",
	}
	if err := account.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	account.PrimaryEmail = "test.test_test+123@example.com"
	if err := account.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid emails
	account.PrimaryEmail = "invalidEmail"
	if err := account.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid email")
	}
}

func Test_ValidPassword(t *testing.T) {
	// Checks for passwords
	account := domains.Account{
		PrimaryEmail: "gmail@gmail.com",
		Password:     "password",
	}
	if err := account.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	account.Password = "TestName"
	if err := account.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	account.Password = ""
	if err := account.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid password")
	}

	account.Password = "1234567"
	if err := account.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid password")
	}
}
