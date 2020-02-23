package achieve

import (
	"testing"
)

func TestValidYear(t *testing.T) {
	achieve := Achieve{Year: 2020, Message: "This is a message"}
	if err := CheckValidAchievement(achieve); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	achieve.Year = 100
	if err := CheckValidAchievement(achieve); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid year")
	}
}

func TestValidMessage(t *testing.T) {
	achieve := Achieve{Year: 2050, Message: "Hello World!"}
	if err := CheckValidAchievement(achieve); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	achieve.Message = ""
	if err := CheckValidAchievement(achieve); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid message")
	}
}