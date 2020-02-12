package programs

import (
  "testing"
)

func TestProgramName(t *testing.T) {
	// Success examples
	program := Program{ProgramId: "ok", Name: "Calculus BC", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	program.Name = "AP Language And Composition"
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	program.Name = "EvErY First CHARACTER Of W0rds@are CapitalizeD45"
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	program.Name = "Cooking 101"
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	program.Name = "100 Ways To Become A Millionare"
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	// Failure examples
	program.Name = "calculus BC"
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
	
	program.Name = ""
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
	
	program.Name = "Calculus bC"
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
	
	program.Name = "40 @40"
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
}

func TestDescription(t *testing.T) {
	// Success example
	program := Program{ProgramId: "ok", Name: "Calculus BC", Grade1: 1, Grade2: 12, Description: "This is a description"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	// Failure example
	program.Description = ""
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
}

func TestValidProgramId(t *testing.T) {
  program := Program{ProgramId: "test_program", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
  err := CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
  }

  program.ProgramId = "a_a_a_a_a_a"
  err = CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
  }

  program.ProgramId = "This_8932792_352_IS_3589_A_vaLiD_pRogRAm_iD_6A783B3S"
  err = CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
  }

  program.ProgramId = "a__a"
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid program id")
  }

  program.ProgramId = "_"
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid program id")
  }

  program.ProgramId = ""
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid program id")
  }
}

func TestValidGrades(t *testing.T) {
  program := Program{ProgramId: "test_program", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
  err := CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s.", err.Error())
  }

  program.Grade1 = 6
  program.Grade2 = 6
  err = CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s.", err.Error())
  }

  program.Grade1 = 7
  program.Grade2 = 8
  err = CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s.", err.Error())
  }

  program.Grade1 = 0
  program.Grade2 = 12
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid grades.")
  }

  program.Grade1 = 1
  program.Grade2 = 13
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid grades.")
  }

  program.Grade1 = 12
  program.Grade2 = 0
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid grades.")
  }
}