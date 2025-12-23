// Package models contains the refresh token entity.
package models

import "time"

// RefreshToken stores refresh tokens for session continuation.
type RefreshToken struct {
	ID        string    `gorm:"primaryKey"`
	UserID    string    `gorm:"index;not null"`
	TokenHash string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}
