package models

// ChapaRequest represents the payload sent to Chapa to initiate a transaction
type ChapaRequest struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	TxRef         string `json:"tx_ref"`
	CallbackURL   string `json:"callback_url"`
	ReturnURL     string `json:"return_url"`
	Customization struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"customization"`
}

// ChapaResponse represents the success response from Chapa initialization
type ChapaResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

// ChapaVerifyResponse represents the response from verifying a transaction
type ChapaVerifyResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		TxRef         string  `json:"tx_ref"`
		Amount        float64 `json:"amount"`
		Currency      string  `json:"currency"`
		Status        string  `json:"status"`
		PaymentMethod string  `json:"payment_method"`
	} `json:"data"`
}
