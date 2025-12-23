package repositories

import (
	"order-service/internal/database"
	"order-service/internal/models"
)

type QuoteRepository struct{}

func (r *QuoteRepository) Create(q *models.Quote) error {
	return database.DB.Create(q).Error
}

func (r *QuoteRepository) FindByID(id string) (*models.Quote, error) {
	var q models.Quote
	err := database.DB.First(&q, "id = ?", id).Error
	return &q, err
}

func (r *QuoteRepository) Update(q *models.Quote) error {
	return database.DB.Save(q).Error
}
