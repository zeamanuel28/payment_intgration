package main

import (
	"log"
	"net/http"
	"os"

	"payment-integration/handlers"
	"payment-integration/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Chapa service
	chapaService := service.NewChapaService()
	paymentHandler := handlers.NewPaymentHandler(chapaService)

	// Set up routes
	http.HandleFunc("/pay", paymentHandler.HandlePay)
	http.HandleFunc("/verify", paymentHandler.HandleVerify)
	http.HandleFunc("/webhook", paymentHandler.HandleWebhook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
