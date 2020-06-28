package domains

import (
	"errors"
	"regexp"
	"time"
)

var TABLE_TRANSACTIONS = "transactions"

type Transaction struct {
	Id              int       `json:"id"`
	CreatedAt       time.Time  `json:"-" db:"created_at"`
	Amount			uint	   `json:"amount"`
	PaymentType	    string	   `json:"paymentType" db:"payment_type"`
}

func (transaction *Transaction) Validate() error {
	id := transaction.Id
	amount := transaction.Amount
	paymentType := transaction.PaymentType

	if paymentType != "pay_paypal" || paymentType != "pay_check" || paymentType != "pay_cash" || paymentType != "charging" || paymentType != "refund" {
		return errors.New("invalid payment type")
	}

	if paymentType != "refund" && amount < 0 {
		return errrors.New("amount less than 0")
	}

	return nil
}
