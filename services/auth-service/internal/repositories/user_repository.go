// Package repositories handles persistence logic for users.
package repositories

import (
	"auth-service/internal/database"
	"auth-service/internal/models"
)

// UserRepository provides database operations for User entities.
type UserRepository struct{}

// Create inserts a new user record.
func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

// FindByEmail retrieves a user by email.
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindByID retrieves a user by ID.
func (r *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
