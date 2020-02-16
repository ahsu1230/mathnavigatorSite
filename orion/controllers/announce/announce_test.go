package announce

import (
	"testing"
)

func TestAnnounceId(t *testing.T) {
	// Success examples
	announce := Announce{AnnounceId: "ok", Title: "Title", Message: "not blank"}
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.AnnounceId = "series_of_undersCores"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.AnnounceId = "j_asdf_erD_pq_898864"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	// Failure examples
	announce.AnnounceId = "_asdf"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid announcement id")
	}
	
	announce.AnnounceId = "asdf_"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid announcement id")
	}
	
	announce.AnnounceId = "asdf__asdf"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid announcement id")
	}
	
	announce.AnnounceId = "valid_id_until__here"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid announcement id")
	}
	
	announce.AnnounceId = "_"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid announcement id")
	}
}

func TestTitle(t *testing.T) {
	// Success examples
	announce := Announce{AnnounceId: "ok", Title: "Valid Title", Message: "not blank"}
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.Title = "New semester tuition"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.Title = "Borrowing & Lending"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.Title = "1000 Tips on Studying"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.Title = "X"
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	// Failure examples
	announce.Title = "&&"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}
	
	announce.Title = "lowercase Start"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}
	
	announce.Title = ""
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}
	
	announce.Title = "JH %"
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: invalid title")
	}
}

func TestMessage(t *testing.T) {
	// Success examples
	announce := Announce{AnnounceId: "ok", Title: "Valid Title", Message: "not blank"}
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	
	announce.Message = "                 a   "
	if err := CheckValidAnnouncement(announce); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Failure examples
	announce.Message = ""
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: empty message")
	}
	
	announce.Message = "                                                    "
	if err := CheckValidAnnouncement(announce); err == nil {
		t.Error("Check was incorrect, got: nil, expected: empty message")
	}
}