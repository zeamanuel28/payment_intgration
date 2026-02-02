package models

import (
	"fmt"
	"time"
)

// Transaction represents an internal payment transaction record
type Transaction struct {
	ID            string    `json:"id"`
	TxRef         string    `json:"tx_ref"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ToChapaRequest converts a Transaction to a ChapaRequest
func (t *Transaction) ToChapaRequest(callbackURL, returnURL string) ChapaRequest {
	return ChapaRequest{
		Amount:      fmt.Sprintf("%f", t.Amount),
		Currency:    t.Currency,
		Email:       t.Email,
		FirstName:   t.FirstName,
		LastName:    t.LastName,
		TxRef:       t.TxRef,
		CallbackURL: callbackURL,
		ReturnURL:   returnURL,
	}
}
