package models

import "time"

type Quote struct {
	ID         string    `gorm:"primaryKey"`
	DesignID   string    `gorm:"index;not null"`
	DesignerID string    `gorm:"index;not null"`
	SupplierID string    `gorm:"index"`
	Price      float64   `json:"price"`
	ETA        int       `json:"eta_days"`
	Status     string    `gorm:"not null"`
	Notes      string    `json:"notes"` 
	CreatedAt  time.Time
}