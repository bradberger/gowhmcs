package whmcs

import (
	"errors"
)

// Various error responses
var (
	ErrTransactionAmountsEmpty       = errors.New("Transaction amounts empty")
	ErrTransactionPaymentMethodEmpty = errors.New("Transaction payment method empty")
	ErrTransactionDateEmpty          = errors.New("Transaction date empty")
)

// Transaction is a struct containing parameters which can be sent to create a new transaction.
type Transaction struct {
	// Required fields
	AmountIn      float64 `json:"amountin,string" xml:"amountin"`
	AmountOut     float64 `json:"amountout,string" xml:"amountout"`
	PaymentMethod string  `json:"paymentmethod" xml:"paymentmethod"`
	Date          Date    `json:"date" xml:"date"`

	// Optional fields
	UserID      int64   `json:"userid,string,omitempty" xml:"userid,omitempty"`
	InvoiceID   int64   `json:"invoiceid,string,omitempty" xml:"invoiceid,omitempty"`
	Description string  `json:"description,omitempty" xml:"description"`
	Fees        float64 `json:"fees,string",omitempty xml:"fees"`
	TransID     string  `json:"transid,omitempty" xml:"transid"`
	Credit      bool    `json:"credit,string" xml:"credit"`
}

func (t *Transaction) Error() error {
	if t.AmountIn == 0 && t.AmountOut == 0 {
		return ErrTransactionAmountsEmpty
	}
	if t.PaymentMethod == "" {
		return ErrTransactionPaymentMethodEmpty
	}
	if t.Date.Time().IsZero() {
		return ErrTransactionDateEmpty
	}
	return nil
}
