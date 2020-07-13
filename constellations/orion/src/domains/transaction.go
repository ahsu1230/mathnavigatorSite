package domains

import (
	"errors"
	"time"
)

var TABLE_TRANSACTIONS = "transactions"

type Transaction struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	Amount      int       `json:"amount"`
	PaymentType string    `json:"paymentType" db:"payment_type"`
}

func (transaction *Transaction) Validate() error {
	amount := transaction.Amount
	paymentType := transaction.PaymentType

	if paymentType != "pay_paypal" && paymentType != "pay_check" && paymentType != "pay_cash" && paymentType != "charge" && paymentType != "refund" {
		return errors.New("invalid payment type")
	}

	if paymentType != "charge" && amount < 0 {
		return errors.New("amount less than 0")
	}

	if paymentType == "charge" && amount > 0 {
		return errors.New("charge greater than 0")
	}

	return nil
}
