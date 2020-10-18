package domains_test

import (
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidClassKey(t *testing.T) {
	// Checks for valid class keys
	class := domains.Class{
		ClassKey:     domains.NewNullString("final_review"),
		TimesStr:     "3 pm - 5 pm",
		PriceLumpSum: domains.NewNullUint(1000),
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

func TestValidTimesStr(t *testing.T) {
	// Checks for valid times
	class := domains.Class{
		ClassKey:     domains.NewNullString("final_review"),
		TimesStr:     "6 pm - 8 pm",
		PriceLumpSum: domains.NewNullUint(1000),
	}
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	class.TimesStr = "8"
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid times
	class.TimesStr = ""
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class name")
	}

	class.TimesStr = "Hi"
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class name")
	}
}

func TestValidPrice(t *testing.T) {
	// Check valid prices
	class := domains.Class{
		ClassKey:     domains.NewNullString("final_review"),
		TimesStr:     "6 pm - 8 pm",
		PriceLumpSum: domains.NewNullUint(1000),
	}

	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Valid PricePerSession and no PriceLumpSum
	class.PricePerSession = domains.NewNullUint(100)
	class.PriceLumpSum = domains.NewNullUint(0)
	if err := class.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Both PricePerSession and PriceLumpSum
	class.PricePerSession = domains.NewNullUint(100)
	class.PriceLumpSum = domains.NewNullUint(1000)
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid price: choose one or the other", err.Error())
	}

	// Both empty
	class.PriceLumpSum = domains.NewNullUint(0)
	class.PricePerSession = domains.NewNullUint(0)
	if err := class.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid price: choose one or the other")
	}
}
