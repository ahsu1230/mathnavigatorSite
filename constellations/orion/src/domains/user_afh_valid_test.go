package domains_test

import (
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidUserId(t *testing.T) {
	// Checks for valid userIds
	userAfh := domains.UserAfh{UserId: 3, AfhId: 7}
	if err := userAfh.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}

	// Checks for invalid userIds
}
