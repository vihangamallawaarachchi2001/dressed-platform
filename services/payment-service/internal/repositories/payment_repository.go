package repositories

import (
	"payment-service/internal/database"
	"payment-service/internal/models"
)

type PaymentRepository struct{}

func (r *PaymentRepository) Create(p *models.Payment) error {
	return database.DB.Create(p).Error
}

func (r *PaymentRepository) FindByID(id string) (*models.Payment, error) {
	var p models.Payment
	err := database.DB.First(&p, "id = ?", id).Error
	return &p, err
}

func (r *PaymentRepository) Update(p *models.Payment) error {
	return database.DB.Save(p).Error
}
