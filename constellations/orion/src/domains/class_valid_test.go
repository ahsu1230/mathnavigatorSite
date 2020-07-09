package domains_test

import (
	"testing"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidClassKey(t *testing.T) {
	// Checks for valid class keys
	now := time.Now().UTC()
	later := now.Add(time.Hour * 24 * 100)
	class := domains.Class{
		ClassKey:  domains.NewNullString("final_review"),
		Times:     "3 pm - 5 pm",
		StartDate: now,
		EndDate:   later,
	}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.ClassKey = domains.NewNullString("")
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.ClassKey = domains.NewNullString("Valid_Class_100")
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid class keys
	class.ClassKey = domains.NewNullString("a__a")
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class key")
	}

	class.ClassKey = domains.NewNullString("_")
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class key")
	}
}

func TestValidTimes(t *testing.T) {
	// Checks for valid times
	now := time.Now().UTC()
	later := now.Add(time.Hour * 24 * 100)
	class := domains.Class{
		ClassKey:  domains.NewNullString("final_review"),
		Times:     "6 pm - 8 pm",
		StartDate: now,
		EndDate:   later,
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
	now := time.Now().UTC()
	later := now.Add(time.Hour * 24 * 30)
	class := domains.Class{
		ClassKey:  domains.NewNullString(""),
		Times:     "3 pm - 5 pm",
		StartDate: now,
		EndDate:   later,
	}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.StartDate = now
	class.EndDate = now.Add(time.Hour * 24 * 365 * 100)
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid dates
	class.StartDate = time.Unix(0, 0)
	class.EndDate = now
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}

	class.StartDate = now.Add(time.Hour * 24)
	class.EndDate = now
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}
}

func TestValidPrice(t *testing.T) {
	//Check valid prices
	now := time.Now().UTC()
	later := now.Add(time.Hour * 24 * 30)

	validClass := domains.Class{
		ClassKey:  domains.NewNullString("final_review"),
		Times:     "6 pm - 8 pm",
		StartDate: now,
		EndDate:   later,
	}
	if err := validClass.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	//Valid PriceLump
	validClass.PriceLump = domains.NewNullUint(100)
	if err := validClass.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	validClass.PriceLump = domains.NewNullUint(0)
	if err := validClass.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	//Valid PricePerSession
	validClass.PricePerSession = domains.NewNullUint(100)
	if err := validClass.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	validClass.PricePerSession = domains.NewNullUint(0)
	if err := validClass.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	//Check invalid prices
	invalidClass := domains.Class{
		ClassKey:  domains.NewNullString("final_review"),
		Times:     "6 pm - 8 pm",
		StartDate: now,
		EndDate:   later,
	}
	if err := invalidClass.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	//Both valid
	invalidClass.PriceLump = domains.NewNullUint(10)
	invalidClass.PricePerSession = domains.NewNullUint(50)
	if err := invalidClass.Validate(); err == nil {
		t.Errorf("Check was incorrect, got: nil, expected: invalid price: both valid")
	}
}
