package domains

import (
	"testing"
)

func TestHash(t *testing.T) {
	hash, err := NewHash("potato")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if match, _ := hash.Compare("pineapple"); match {
		t.Error("Invalid comparison, got: true, expected: false")
	}

	if match, _ := hash.Compare("sesame"); match {
		t.Error("Invalid comparison, got: true, expected: false")
	}

	if match, _ := hash.Compare("#@)!@(*@#!$#@!#@!#"); match {
		t.Error("Invalid comparison, got: true, expected: false")
	}

	if match, _ := hash.Compare("Potato"); match {
		t.Error("Invalid comparison, got: true, expected: false")
	}

	if match, err := hash.Compare("potato"); !match {
		t.Errorf("Invalid comparison, got: false, expected: true, error: %v", err)
	}
}
