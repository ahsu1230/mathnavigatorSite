package announce

import (
	"testing"
)

func TestAuthor(t *testing.T) {
	// Success examples
	announce := Announce{Author: "Valid Author", Message: "not blank"}
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.Author = "legitimate author name"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	// Failure examples
	announce.Author = "&&"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid author")
	}
	
	announce.Author = "  @1 + 4"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid author")
	}
	
	announce.Author = ""
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid author")
	}
}

func TestMessage(t *testing.T) {
	// Success examples
	announce := Announce{Author: "Valid Author", Message: "not blank"}
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.Message = "                 a   "
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.Message = "I'm a filled message with speci@l ch@r&cterS"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Failure examples
	announce.Message = ""
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid message")
	}
	
	announce.Message = "                                                    "
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid message")
	}
	
	announce.Message = "     123%                                           "
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid message")
	}
}