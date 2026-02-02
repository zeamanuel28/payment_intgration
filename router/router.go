package router

import (
	_ "payment-integration/docs"
	"payment-integration/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configures the application routes and returns a Gin Engine
func SetupRoutes(paymentHandler *handlers.PaymentHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/pay", paymentHandler.HandlePay)
	r.GET("/verify", paymentHandler.HandleVerify)
	r.POST("/webhook", paymentHandler.HandleWebhook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
