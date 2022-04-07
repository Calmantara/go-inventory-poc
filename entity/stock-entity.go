package entity

import "github.com/google/uuid"

type StockEntity struct {
	ID           *uuid.UUID `json:"id,omitempty"`
	Name         string     `json:"name"`
	Price        float64    `json:"price" binding:"gte=1"`
	Availability int        `json:"availability" binding:"gte=0"`
	IsActive     bool       `json:"is_active"`
}
