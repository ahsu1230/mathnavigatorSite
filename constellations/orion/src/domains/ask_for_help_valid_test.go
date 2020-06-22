package domains_test

import (
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func TestValidAFHTitle(t *testing.T) {
	askForHelp := domains.AskForHelp{}
	if err := askForHelp.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
}
