/**
 * Quote represents a supplier quotation.
 */
package models

import "time"

type Quote struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	DesignID   string    `gorm:"index;not null" json:"design_id"`
	DesignerID string    `gorm:"index;not null" json:"designer_id"`
	SupplierID string    `gorm:"index" json:"supplier_id"`
	Price      float64   `json:"price"`
	ETA        int       `json:"eta_days"`
	Status     string    `gorm:"not null" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
