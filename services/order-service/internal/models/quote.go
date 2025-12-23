/**
 * Quote represents a supplier quotation.
 */
package models

import "time"

type Quote struct {
	ID         string    `json:"id"`
	DesignID   string    `json:"designId"`
	DesignerID string    `json:"designerId"`
	SupplierID string    `json:"supplierId"`
	Price      float64   `json:"price"`
	ETA        int       `json:"eta"`
	Notes      string    `json:"notes"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}
