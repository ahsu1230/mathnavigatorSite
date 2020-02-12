package programs

import (
  "testing"
)

func TestProgramName(t *testing.T) {
	// success examples
	program := Program{ProgramId: "ok", Name: "Calculus BC", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	program = Program{ProgramId: "ok", Name: "AP Language And Composition", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	program = Program{ProgramId: "ok", Name: "EvErY First CHARACTER Of W0rds@are CapitalizeD45", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	program = Program{ProgramId: "ok", Name: "Cooking 101", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	program = Program{ProgramId: "ok", Name: "100 Ways To Become A Millionare", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	// failure examples
	program = Program{ProgramId: "ok", Name: "calculus BC", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
	
	program = Program{ProgramId: "ok", Name: "", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
	
	program = Program{ProgramId: "ok", Name: "Calculus bC", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
	
	program = Program{ProgramId: "ok", Name: "40 @40", Grade1: 1, Grade2: 12, Description: "ok"}
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
}

func TestDescription(t *testing.T) {
	// success example
	program := Program{ProgramId: "ok", Name: "Calculus BC", Grade1: 1, Grade2: 12, Description: "This is a description"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	// failure example
	program = Program{ProgramId: "ok", Name: "Calculus BC", Grade1: 1, Grade2: 12, Description: ""}
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
}