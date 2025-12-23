package models

import (
	"time"
	"gorm.io/gorm"
)

type Design struct {
	ID          string `gorm:"primaryKey"`
	DesignerID  string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	Category    string `gorm:"not null"`
	FilePath    string `gorm:"not null"`
	Status      string `gorm:"not null"` // DRAFT, SUBMITTED
	CreatedAt   int64
	UpdatedAt   int64
}

func (d *Design) BeforeCreate(_ *gorm.DB) error {
	now := time.Now().Unix()
	d.CreatedAt = now
	d.UpdatedAt = now
	return nil
}

func (d *Design) BeforeUpdate(_ *gorm.DB) error {
	d.UpdatedAt = time.Now().Unix()
	return nil
}