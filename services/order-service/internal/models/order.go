/**
 * Order represents an accepted quote turned into a job.
 */
package models

import "time"

type Order struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	QuoteID    string    `gorm:"index;not null" json:"quote_id"`
	DesignerID string    `gorm:"index;not null" json:"designer_id"`
	SupplierID string    `gorm:"index;not null" json:"supplier_id"`
	Status     string    `gorm:"not null" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
