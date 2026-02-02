package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"payment-integration/models"
)

type ChapaService struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

func NewChapaService() *ChapaService {
	return &ChapaService{
		APIKey:  os.Getenv("CHAPA_SECRET_KEY"),
		BaseURL: "https://api.chapa.co/v1",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// InitiatePayment initializes a payment transaction with Chapa
func (s *ChapaService) InitiatePayment(tx models.Transaction) (*models.ChapaResponse, error) {
	url := fmt.Sprintf("%s/transaction/initialize", s.BaseURL)

	// Convert Transaction to ChapaRequest
	// TODO: Caller should provide callback URLs, or we configure them here/env
	req := tx.ToChapaRequest("http://localhost:8080/webhook", "http://localhost:8080/success")

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+s.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("chapa api returned status: %d", resp.StatusCode)
	}

	var chapaResp models.ChapaResponse
	if err := json.NewDecoder(resp.Body).Decode(&chapaResp); err != nil {
		return nil, err
	}

	return &chapaResp, nil
}

func (s *ChapaService) VerifyPayment(txRef string) (*models.ChapaVerifyResponse, error) {
	url := fmt.Sprintf("%s/transaction/verify/%s", s.BaseURL, txRef)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+s.APIKey)

	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("chapa api returned status: %d", resp.StatusCode)
	}

	var verifyResp models.ChapaVerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&verifyResp); err != nil {
		return nil, err
	}

	return &verifyResp, nil
}
