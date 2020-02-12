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
	
	// failure examples
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
	// success example
	program := Program{ProgramId: "ok", Name: "Calculus BC", Grade1: 1, Grade2: 12, Description: "This is a description"}
	if err := CheckValidProgram(program); err != nil {
		t.Error(err)
	}
	
	// failure example
	program.Description = ""
	if err := CheckValidProgram(program); err == nil {
		t.Error("Should have detected error but didn't.")
	}
}