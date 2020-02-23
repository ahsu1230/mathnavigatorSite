package semesters

import (
	"testing"
)

func TestValidSemesterId(t *testing.T) {
	semester := Semester{SemesterId: "2019_fall", Title: "Fall 2019"}
	if err := CheckValidSemester(semester); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	semester.SemesterId = "2050_winter"
	if err := CheckValidSemester(semester); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	semester.SemesterId = "0999_spring"
	if err := CheckValidSemester(semester); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid semester id")
	}

	semester.SemesterId = "2019_wrong"
	if err := CheckValidSemester(semester); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid semester id")
	}
}

func TestSemesterName(t *testing.T) {
	semester := Semester{SemesterId: "2020_spring", Title: "Test title"}
	if err := CheckValidSemester(semester); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	semester.Title = "Winter 2222 (#1 class)"
	if err := CheckValidSemester(semester); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	semester.Title = "This title is too long. 123456789012345678901234567890123456789012345678901234567890"
	if err := CheckValidSemester(semester); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid semester title")
	}

	semester.Title = "Fall_ _2020"
	if err := CheckValidSemester(semester); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid semester title")
	}
}
