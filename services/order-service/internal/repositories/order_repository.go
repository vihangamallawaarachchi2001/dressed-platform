package repositories

import (
	"order-service/internal/database"
	"order-service/internal/models"
)

type OrderRepository struct{}

func (r *OrderRepository) Create(o *models.Order) error {
	return database.DB.Create(o).Error
}
