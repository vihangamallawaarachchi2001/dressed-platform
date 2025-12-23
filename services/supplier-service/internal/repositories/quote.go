package repositories

import (
	"supplier-service/internal/database"
	"supplier-service/internal/models"
)

type QuoteRepository struct{}

func (r *QuoteRepository) Create(q *models.Quote) error {
	return database.DB.Create(q).Error
}