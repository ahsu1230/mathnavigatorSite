package announce

import (
	"testing"
)

func TestValidAuthor(t *testing.T) {
	// Check for valid authors
	announce := Announce{Author: "Valid Author", Message: "not blank"}
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	announce.Author = "legitimate author name"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Check for invalid authors
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

func TestValidMessage(t *testing.T) {
	// Check for valid messages
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

	// Check for invalid messages
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
