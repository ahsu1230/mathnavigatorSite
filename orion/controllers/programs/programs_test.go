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

// TODO: add more tests for all programs and add programs that should not work
func TestValidProgramId(t *testing.T) {
  program := Program{ProgramId: "1", Name: "Test Program", Grade1: 1, Grade2: 12}
  err := CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s.", err.Error())
  }
}

func TestValidProgramNames(t *testing.T) {
  program := Program{ProgramId: "1", Name: "Test Program", Grade1: 1, Grade2: 12}
  err := CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s.", err.Error())
  }
}

func TestValidGrades(t *testing.T) {
  program := Program{ProgramId: "1", Name: "Test Program", Grade1: 1, Grade2: 12}
  err := CheckValidProgram(program)
  if err != nil {
    t.Errorf("Check was incorrect, got: %s.", err.Error())
  }
}