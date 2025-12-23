package models

import (
	"time"

	"gorm.io/gorm"
)

type Design struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	DesignerID  string    `gorm:"not null" json:"designerId"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Category    string    `gorm:"not null" json:"category"`
	FilePath    string    `gorm:"not null" json:"filePath"`
	Status      string    `gorm:"not null" json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (d *Design) BeforeCreate(_ *gorm.DB) error {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Design) BeforeUpdate(_ *gorm.DB) error {
	d.UpdatedAt = time.Now()
	return nil
}