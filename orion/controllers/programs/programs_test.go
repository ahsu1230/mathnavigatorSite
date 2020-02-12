package programs

import (
  "testing"
)

func TestExample(t *testing.T) {
  total := 5 + 5
  if total != 10 {
     t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
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