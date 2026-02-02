package handlers

import (
	"fmt"
	"net/http"
	"payment-integration/models"
	"payment-integration/service"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	Service service.PaymentGateway
}

func NewPaymentHandler(service service.PaymentGateway) *PaymentHandler {
	return &PaymentHandler{Service: service}
}

// HandlePay godoc
// @Summary      Initiate a payment
// @Description  Initiates a payment transaction with Chapa
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        transaction  body      models.Transaction  true  "Transaction Request"
// @Success      200          {object}  models.ChapaResponse
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /pay [post]
func (h *PaymentHandler) HandlePay(c *gin.Context) {
	var req models.Transaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Generate a TxRef if missing
	if req.TxRef == "" {
		req.TxRef = fmt.Sprintf("tx-%d", time.Now().UnixNano())
	}

	resp, err := h.Service.InitiatePayment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate payment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// HandleVerify godoc
// @Summary      Verify a payment
// @Description  Verifies a payment transaction using Chapa
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        tx_ref  query     string  true  "Transaction Reference"
// @Success      200     {object}  models.ChapaVerifyResponse
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /verify [get]
func (h *PaymentHandler) HandleVerify(c *gin.Context) {
	txRef := c.Query("tx_ref")
	if txRef == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tx_ref is required"})
		return
	}

	resp, err := h.Service.VerifyPayment(txRef)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify payment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
