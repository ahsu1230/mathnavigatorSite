package domains_test

import (
	"testing"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
)

func TestValidProgramId(t *testing.T) {
	// Checks for valid program IDs
	program := domains.Program{ProgramId: "test_program", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.ProgramId = "a_a_a_a_a_a"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.ProgramId = "This_8932792_352_IS_3589_A_vaLiD_pRogRAm_iD_6A783B3S"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid program IDs
	program.ProgramId = "a__a"
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program id")
	}

	program.ProgramId = "_"
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program id")
	}

	program.ProgramId = ""
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program id")
	}
}

func TestProgramName(t *testing.T) {
	// Checks for valid names
	program := domains.Program{ProgramId: "ok", Name: "AP Calculus BC", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "AP Language & Composition (#2)"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "SAT-Prep & Learning_Test (This works)"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "Cooking 101"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "100 Ways To Become A Millionaire"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "Mommy & me"
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid names
	program.Name = "Test__"
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}

	program.Name = ""
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}

	program.Name = "Test )("
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}

	program.Name = "40 @40"
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}

	program.Name = "A0 ^40"
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}
}

func TestValidGrades(t *testing.T) {
	// Checks for valid grades
	program := domains.Program{ProgramId: "test_program", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Grade1 = 6
	program.Grade2 = 6
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Grade1 = 7
	program.Grade2 = 8
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid grades
	program.Grade1 = 0
	program.Grade2 = 12
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}

	program.Grade1 = 1
	program.Grade2 = 13
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}

	program.Grade1 = 12
	program.Grade2 = 0
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}
}

func TestDescription(t *testing.T) {
	// Checks for valid descriptions
	program := domains.Program{ProgramId: "ok", Name: "Calculus BC", Grade1: 1, Grade2: 12, Description: "This is a description"}
	if err := program.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid descriptions
	program.Description = ""
	if err := program.Validate(); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid description.")
	}
}