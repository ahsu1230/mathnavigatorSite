package domains

import (
	"errors"
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
	Id           uint       `json:"id"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt    NullTime   `json:"-" db:"deleted_at"`
	Amount       int        `json:"amount"`
	PaymentType  string     `json:"paymentType" db:"payment_type"`
	PaymentNotes NullString `json:"paymentNotes" db:"payment_notes"`
	AccountId    uint       `json:"accountId" db:"account_id"`
}

func (transaction *Transaction) Validate() error {
	amount := transaction.Amount
	paymentType := transaction.PaymentType

	if paymentType != PAY_PAYPAL && paymentType != PAY_CHECK && paymentType != PAY_CASH && paymentType != CHARGE && paymentType != REFUND {
		return errors.New("invalid payment type")
	}

	if paymentType != CHARGE && amount < 0 {
		return errors.New("amount less than 0")
	}

	if paymentType == CHARGE && amount > 0 {
		return errors.New("charge greater than 0")
	}

	return nil
}
