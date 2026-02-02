package service

import "payment-integration/models"

// PaymentGateway defines the interface that all payment providers must implement
type PaymentGateway interface {
	InitiatePayment(req models.Transaction) (*models.ChapaResponse, error)
	VerifyPayment(txRef string) (*models.ChapaVerifyResponse, error)
}
