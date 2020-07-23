package domains_test

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"testing"
	"time"
)

func TestValidPaymentTypes(t *testing.T) {
	now := time.Now().UTC()
	transaction := domains.Transaction{
		CreatedAt: now,
	}
	//Test valid payment types
	transaction.Amount = 10
	transaction.PaymentType = "pay_paypal"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	transaction.PaymentType = "pay_check"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	transaction.PaymentType = "pay_cash"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	transaction.PaymentType = "refund"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	transaction.Amount = -10
	transaction.PaymentType = "charge"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	//Test invalid payment types
	transaction.PaymentType = "asdf"
	if err := transaction.Validate(); err == nil {
		t.Errorf("Check was incorrect, got: nil, expected: invalid payment type")
	}
	transaction.PaymentType = "bitcoin"
	if err := transaction.Validate(); err == nil {
		t.Errorf("Check was incorrect, got: nil, expected: invalid payment type")
	}
}

func TestValidAmounts(t *testing.T) {
	now := time.Now().UTC()
	transaction := domains.Transaction{
		CreatedAt: now,
	}
	//Test valid charge amount
	transaction.Amount = -10
	transaction.PaymentType = "charge"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	//Test invalid charge amount
	transaction.Amount = 10
	if err := transaction.Validate(); err == nil {
		t.Errorf("Check was incorrect, got: nil, expected: charge greater than 0")
	}
	//Test valid other amount
	transaction.Amount = 10
	transaction.PaymentType = "pay_paypal"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	transaction.PaymentType = "pay_check"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	transaction.PaymentType = "pay_cash"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	transaction.PaymentType = "refund"
	if err := transaction.Validate(); err != nil {
		t.Errorf("Check was incorrect, got: %s, expected: nil", err.Error())
	}
	//Test invalid other amount
	transaction.Amount = -10
	transaction.PaymentType = "pay_paypal"
	if err := transaction.Validate(); err == nil {
		t.Errorf("Check was incorrect, got: nil, expected: amount less than 0")
	}
}
