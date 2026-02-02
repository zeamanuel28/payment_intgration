package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChapaWebhookRequest represents the payload sent by Chapa webhook
// Note: Adjust fields based on actual Chapa webhook payload
type ChapaWebhookRequest struct {
	TxRef  string `json:"tx_ref"`
	Status string `json:"status"`
	// Add other fields as needed
}

func (h *PaymentHandler) HandleWebhook(c *gin.Context) {
	// Verify signature if Chapa provides one (TODO)
	// signature := c.GetHeader("Chapa-Signature")

	var payload ChapaWebhookRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("Received webhook for tx: %s, status: %s\n", payload.TxRef, payload.Status)

	// Update transaction status in database (mocked)
	// err := h.Service.UpdateTransactionStatus(payload.TxRef, payload.Status)

	c.Status(http.StatusOK)
}
