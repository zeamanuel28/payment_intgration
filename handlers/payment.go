package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payment-integration/models"
	"payment-integration/service"
	"time"
)

type PaymentHandler struct {
	Service service.PaymentGateway
}

func NewPaymentHandler(service service.PaymentGateway) *PaymentHandler {
	return &PaymentHandler{Service: service}
}

func (h *PaymentHandler) HandlePay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a TxRef if missing
	if req.TxRef == "" {
		req.TxRef = fmt.Sprintf("tx-%d", time.Now().UnixNano())
	}

	resp, err := h.Service.InitiatePayment(req)
	if err != nil {
		http.Error(w, "Failed to initiate payment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *PaymentHandler) HandleVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	txRef := r.URL.Query().Get("tx_ref")
	if txRef == "" {
		http.Error(w, "tx_ref is required", http.StatusBadRequest)
		return
	}

	resp, err := h.Service.VerifyPayment(txRef)
	if err != nil {
		http.Error(w, "Failed to verify payment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
