package programs

import (
  "testing"
  "math/rand"
  "fmt"
)

func TestExample(t *testing.T) {
  total := 5 + 5
  if total != 10 {
     t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
  }
}

func Rand(a, b int) uint {
	return uint(rand.Int() % (b - a) + a)
}

func TestCheckProgram(t *testing.T) {
	// success examples
	var s string
	for i := 0; i < 1000; i++ {
		s += "Word "
		var grade1, grade2 uint = Rand(1, 12), Rand(1, 12)
		if grade1 > grade2 {
		    grade1, grade2 = grade2, grade1
		}
		program := Program{ProgramId: string(i), Name: s, Grade1: grade1, Grade2: grade2, Description: "a"}
		if err := CheckValidProgram(program); err != nil {
			fmt.Println(err)
			return
		}
	}
	
	// failure examples
	program := Program{ProgramId: "", Name: "Asdf", Grade1: 1, Grade2: 12, Description: "a"}
	if err := CheckValidProgram(program); err == nil {
		fmt.Println("Fail")
		return
	}
	program = Program{ProgramId: "0", Name: "asdf", Grade1: 1, Grade2: 12, Description: "a"}
	if err := CheckValidProgram(program); err == nil {
		fmt.Println("Fail")
		return
	}
	program = Program{ProgramId: "0", Name: "Asdf", Grade1: 1, Grade2: 12, Description: ""}
	if err := CheckValidProgram(program); err == nil {
		fmt.Println("Fail")
		return
	}
	program = Program{ProgramId: "0", Name: "Asdf", Grade1: 17, Grade2: 12, Description: "a"}
	if err := CheckValidProgram(program); err == nil {
		fmt.Println("Fail")
		return
	}
	program = Program{ProgramId: "0", Name: "Asdf", Grade1: 0, Grade2: 12, Description: "a"}
	if err := CheckValidProgram(program); err == nil {
		fmt.Println("Fail")
		return
	}
	program = Program{ProgramId: "0", Name: "Asdf", Grade1: 5, Grade2: 2, Description: "a"}
	if err := CheckValidProgram(program); err == nil {
		fmt.Println("Fail")
		return
	}
}