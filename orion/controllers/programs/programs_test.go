package programs

import (
	"testing"
)

func TestValidProgramId(t *testing.T) {
	program := Program{ProgramId: "test_program", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.ProgramId = "a_a_a_a_a_a"
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.ProgramId = "This_8932792_352_IS_3589_A_vaLiD_pRogRAm_iD_6A783B3S"
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.ProgramId = "a__a"
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program id")
	}

	program.ProgramId = "_"
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program id")
	}

	program.ProgramId = ""
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program id")
	}
}

func TestProgramName(t *testing.T) {
	program := Program{ProgramId: "ok", Name: "AP Calculus BC", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "AP Language & Composition (#2)"
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "SAT-Prep & Learning_Test (This works)"
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "Cooking 101"
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "100 Ways To Become A Millionaire"
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Name = "Test__"
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}

	program.Name = ""
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}

	program.Name = "Test )("
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}

	program.Name = "40 @40"
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid program name")
	}
}

func TestValidGrades(t *testing.T) {
	program := Program{ProgramId: "test_program", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Grade1 = 6
	program.Grade2 = 6
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Grade1 = 7
	program.Grade2 = 8
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	program.Grade1 = 0
	program.Grade2 = 12
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}

	program.Grade1 = 1
	program.Grade2 = 13
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}

	program.Grade1 = 12
	program.Grade2 = 0
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid grades.")
	}
}

func TestDescription(t *testing.T) {
	// Success example
	program := Program{ProgramId: "ok", Name: "Calculus BC", Grade1: 1, Grade2: 12, Description: "This is a description"}
	if err := CheckValidProgram(program); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Failure example
	program.Description = ""
	if err := CheckValidProgram(program); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid description.")
	}
}
