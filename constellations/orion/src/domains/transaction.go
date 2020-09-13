package domains

import (
	"fmt"
	"time"
)

var TABLE_TRANSACTIONS = "transactions"

const (
	PAY_CASH   = "pay_cash"
	PAY_CHECK  = "pay_check"
	PAY_PAYPAL = "pay_paypal"
	REFUND     = "refund"
	CHARGE     = "charge"
)

var ALL_TRANSACTION_TYPES = []string{PAY_PAYPAL, PAY_CASH, PAY_CHECK, CHARGE, REFUND}

type Transaction struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt NullTime   `json:"-" db:"deleted_at"`
	AccountId uint       `json:"accountId" db:"account_id"`
	Amount    int        `json:"amount"`
	Type      string     `json:"type" db:"type"`
	Notes     NullString `json:"notes" db:"notes"`
}

func (transaction *Transaction) Validate() error {
	messageFmt := "Invalid Transaction: %s"

	amount := transaction.Amount
	paymentType := transaction.Type

	if paymentType != PAY_PAYPAL &&
		paymentType != PAY_CHECK &&
		paymentType != PAY_CASH &&
		paymentType != CHARGE &&
		paymentType != REFUND {
		return fmt.Errorf(messageFmt, "Unrecognized payment type")
	}

	if paymentType != CHARGE && amount < 0 {
		return fmt.Errorf(messageFmt, "Cannot have payment amount < 0. Must be positive.")
	}

	if paymentType == CHARGE && amount > 0 {
		return fmt.Errorf(messageFmt, "Cannot have charge amount > 0. Must be negative.")
	}

	return nil
}
