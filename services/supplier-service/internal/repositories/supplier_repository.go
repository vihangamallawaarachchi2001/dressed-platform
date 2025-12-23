package repositories

import (
	"supplier-service/internal/database"
	"supplier-service/internal/models"
)

type SupplierRepository struct{}

func (r *SupplierRepository) Create(s *models.Supplier) error {
	return database.DB.Create(s).Error
}

func (r *SupplierRepository) FindByUserID(userID string) (*models.Supplier, error) {
	var s models.Supplier
	err := database.DB.First(&s, "user_id = ?", userID).Error
	return &s, err
}

func (r *SupplierRepository) Update(s *models.Supplier) error {
	return database.DB.Save(s).Error
}
