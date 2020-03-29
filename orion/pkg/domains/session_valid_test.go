package domains_test

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"testing"
)

func TestValidNotes(t *testing.T) {
	// Checks for valid notes
	session := domains.Session{ClassId: "valid_id", Notes: "ok"}
	if err := session.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid notes
	session.Notes = ""
	if err := session.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid notes")
	}
}
