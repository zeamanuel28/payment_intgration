package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"payment-integration/models"
)

func TestInitiatePayment(t *testing.T) {
	// Mock response from Chapa
	mockResponse := models.ChapaResponse{
		Message: "Success",
		Status:  "success",
	}
	mockResponse.Data.CheckoutURL = "https://checkout.chapa.co/checkout/payment/12345"

	jsonResp, _ := json.Marshal(mockResponse)

	// Create service with mock client
	// To properly mock *http.Client, we usually use a custom Transport.

	// Create service with mock client
	// Note: We need to modify ChapaService to accept an interface for http.Client to make it testable,
	// OR we can just replace the Client field since it's a *http.Client struct.
	// However, *http.Client is a struct, not an interface.
	// To properly mock *http.Client, we usually use a custom Transport.

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

	req := models.ChapaRequest{
		Amount:   "100",
		Currency: "ETB",
		Email:    "test@example.com",
		TxRef:    "test-ref-123",
	}

	resp, err := service.InitiatePayment(req)
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

type MockTransport struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}
