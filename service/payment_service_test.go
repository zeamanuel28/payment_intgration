package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"payment-integration/models"
)

// MockTransport allows mocking http.Client behaviors
type MockTransport struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

func TestInitiatePayment(t *testing.T) {
	// Mock response from Chapa
	mockResponse := models.ChapaResponse{
		Message: "Success",
		Status:  "success",
	}
	mockResponse.Data.CheckoutURL = "https://checkout.chapa.co/checkout/payment/12345"

	jsonResp, _ := json.Marshal(mockResponse)

	service := NewChapaService()
	service.Client.Transport = &MockTransport{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(jsonResp)),
				Header:     make(http.Header),
			}, nil
		},
	}

	// Use Transaction model instead of ChapaRequest
	tx := models.Transaction{
		TxRef:     "test-ref-123",
		Amount:    100,
		Currency:  "ETB",
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	resp, err := service.InitiatePayment(tx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.Status != "success" {
		t.Errorf("Expected status success, got %s", resp.Status)
	}

	if resp.Data.CheckoutURL != "https://checkout.chapa.co/checkout/payment/12345" {
		t.Errorf("Expected checkout URL, got %s", resp.Data.CheckoutURL)
	}
}
