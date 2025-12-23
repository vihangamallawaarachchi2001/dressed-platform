/**
 * Payment represents a mocked payment transaction.
 */
package models

import "time"

type Payment struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	OrderID   string    `json:"order_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"` // PENDING, SUCCESS, FAILED
	CreatedAt time.Time `json:"created_at"`
}
