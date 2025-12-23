// Package service contains core authentication business logic.
package service

import (
	"errors"
	"time"

	"auth-service/internal/database"
	"auth-service/internal/models"
	"auth-service/internal/repositories"
	"auth-service/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuthService orchestrates authentication workflows.
type AuthService struct {
	userRepo *repositories.UserRepository
}

// NewAuthService creates a new AuthService.
func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

// Register creates a new user account.
func (s *AuthService) Register(email, password, role string) error {
	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("user already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: hashed,
		Role:     role,
	}

	return s.userRepo.Create(user)
}

// Login authenticates a user and returns auth tokens and user role.
func (s *AuthService) Login(email, password string) (token string, role string, err error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !utils.ComparePassword(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}

	token, err = utils.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	return token, user.Role, nil
}

// GetCurrentUser returns user profile for /me endpoint.
func (s *AuthService) GetCurrentUser(userID string) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}

// CreateRefreshToken generates and stores a refresh token.
func (s *AuthService) CreateRefreshToken(userID string) (string, error) {
	raw := uuid.NewString()
	hash, _ := utils.HashPassword(raw)

	rt := models.RefreshToken{
		ID:        uuid.NewString(),
		UserID:    userID,
		TokenHash: hash,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := database.DB.Create(&rt).Error; err != nil {
		return "", err
	}

	return raw, nil
}
