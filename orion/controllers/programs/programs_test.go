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

  program = Program{ProgramId: "a_a_a_a_a_a", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
  err = CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
  }

  program = Program{ProgramId: "This_8932792_352_IS_3589_A_vaLiD_pRogRAm_iD_6A783B3S", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
  err = CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
  }

  program = Program{ProgramId: "a__a", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid program id")
  }

  program = Program{ProgramId: "_", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid program id")
  }

  program = Program{ProgramId: "", Name: "Test Program", Grade1: 1, Grade2: 12, Description: "Description"}
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

  program = Program{ProgramId: "test_program", Name: "Test Program", Grade1: 6, Grade2: 6, Description: "Description"}
  err = CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s.", err.Error())
  }

  program = Program{ProgramId: "test_program", Name: "Test Program", Grade1: 7, Grade2: 8, Description: "Description"}
  err = CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s.", err.Error())
  }

  program = Program{ProgramId: "test_program", Name: "Test Program", Grade1: 0, Grade2: 12, Description: "Description"}
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid grades.")
  }

  program = Program{ProgramId: "test_program", Name: "Test Program", Grade1: 1, Grade2: 13, Description: "Description"}
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid grades.")
  }

  program = Program{ProgramId: "test_program", Name: "Test Program", Grade1: 12, Grade2: 0, Description: "Description"}
  err = CheckValidProgram(program)
  if err == nil {
    t.Errorf("Check was incorrect, got: nil, expected: invalid grades.")
  }
}