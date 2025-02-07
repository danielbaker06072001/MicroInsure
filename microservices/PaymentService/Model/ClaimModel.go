package model

import "time"

type Payment struct {
	ID          int       `json:"id"`
	PaymentNumber string    `json:"Payment_number"`
	PaymentType   string    `json:"Payment_type"`
	PaymentAmount float64   `json:"Payment_amount"` // Changed from int to float64
	PaymentDate   time.Time `json:"Payment_date"`   // Changed from string to time.Time
}
