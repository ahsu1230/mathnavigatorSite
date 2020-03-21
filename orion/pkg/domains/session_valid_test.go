package domains_test

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"testing"
)

func TestValidClassId(t *testing.T) {
	// Checks for valid class IDs
	session := domains.Session{ClassId: "valid_id", Notes: "ok"}
	if err := session.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	session.ClassId = "l_o_t_s_o_f_u_n_d_e_r_s_c_o_r_e_s"
	if err := session.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	session.ClassId = "asdf123"
	if err := session.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid class IDs
	session.ClassId = "two__underscores"
	if err := session.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class id")
	}

	session.ClassId = "@"
	if err := session.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class id")
	}

	session.ClassId = ""
	if err := session.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid class id")
	}
}

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
