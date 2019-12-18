package controllers

import (
  "testing"
)

func TestSum(t *testing.T) {
  total := 5 + 5
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
  // t.Error() // to indicate test failed
}
