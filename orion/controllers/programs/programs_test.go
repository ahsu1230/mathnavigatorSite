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

func TestCheckProgram(t *testing.T) {
  program := Program{Name: "Test Program", Grade1: 1, Grade2: 12}
  isValid := CheckValidProgram(program)
  if !isValid == false {
    t.Errorf("Check was incorrect, got: %t, expected: %t.", isValid, true)
  }
}
