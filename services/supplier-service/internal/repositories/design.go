package repositories

import (
	"supplier-service/internal/database"
	"supplier-service/internal/models"
)

type DesignRepository struct{}

func (r *DesignRepository) ListPublic() ([]models.Design, error) {
	var designs []models.Design
	err := database.DB.Where("status = ?", "SUBMITTED").Find(&designs).Error
	return designs, err
}