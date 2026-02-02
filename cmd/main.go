package main

import (
	"log"
	"os"

	"payment-integration/handlers"
	"payment-integration/router"
	"payment-integration/service"

	"github.com/joho/godotenv"
)

// @title           Payment Integration API
// @version         1.0
// @description     This is a payment integration service using Chapa.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Chapa service
	chapaService := service.NewChapaService()
	paymentHandler := handlers.NewPaymentHandler(chapaService)

	// Set up routes
	mux := router.SetupRoutes(paymentHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := mux.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
