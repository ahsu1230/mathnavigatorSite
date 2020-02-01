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
  isValid := CheckValidProgram()
  if (isValid == false) {
    t.Errorf("Check was incorrect, got: %t, expected: %t.", isValid, false)
  }
}
