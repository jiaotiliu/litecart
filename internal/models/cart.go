package models

import "github.com/shurco/litecart/pkg/litepay"

type Cart struct {
	Core
	Email         string                `json:"email"`
	Cart          []CartProduct         `json:"cart,omitempty"`
	AmountTotal   int                   `json:"amount_total"`
	Currency      string                `json:"currency"`
	PaymentID     string                `json:"payment_id"`
	PaymentStatus litepay.Status        `json:"payment_status"`
	PaymentSystem litepay.PaymentSystem `json:"payment_system"`
}

type CartProduct struct {
	ProductID string `json:"id"`
	Quantity  int    `json:"quantity"`
}

type Payment struct {
	Email    string                `json:"email"`
	Provider litepay.PaymentSystem `json:"provider"`
	Products []CartProduct         `json:"products"`
}
