package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ChapaWebhookRequest represents the payload sent by Chapa webhook
// Note: Adjust fields based on actual Chapa webhook payload
type ChapaWebhookRequest struct {
	TxRef  string `json:"tx_ref"`
	Status string `json:"status"`
	// Add other fields as needed
}

func (h *PaymentHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Verify signature if Chapa provides one (TODO)
	// signature := r.Header.Get("Chapa-Signature")

	var payload ChapaWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received webhook for tx: %s, status: %s\n", payload.TxRef, payload.Status)

	// Update transaction status in database (mocked)
	// err := h.Service.UpdateTransactionStatus(payload.TxRef, payload.Status)

	w.WriteHeader(http.StatusOK)
}
