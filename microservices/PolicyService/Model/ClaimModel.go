package model

import "time"

type Policy struct {
	ID          int       `json:"id"`
	PolicyNumber string    `json:"Policy_number"`
	PolicyType   string    `json:"Policy_type"`
	PolicyAmount float64   `json:"Policy_amount"` 
	PolicyDate   time.Time `json:"Policy_date"`   
}
