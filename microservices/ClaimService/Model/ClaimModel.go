package model

import "time"

type Claim struct {
	ID          int       `json:"id"`
	ClaimNumber string    `json:"claim_number"`
	ClaimType   string    `json:"claim_type"`
	ClaimAmount float64   `json:"claim_amount"` // Changed from int to float64
	ClaimDate   time.Time `json:"claim_date"`   // Changed from string to time.Time
}
