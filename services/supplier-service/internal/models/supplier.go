/**
 * Supplier represents a manufacturing partner.
 */
package models

import "time"

type Supplier struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	UserID       string    `gorm:"uniqueIndex;not null" json:"user_id"`
	CompanyName  string    `json:"company_name"`
	Description  string    `json:"description"`
	Capabilities string    `json:"capabilities"` // e.g. "3D Printing, CNC"
	Status       string    `json:"status"`
	Availability string    `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
}
