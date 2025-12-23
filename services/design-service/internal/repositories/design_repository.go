package repositories

import (
	"design-service/internal/models"

	"gorm.io/gorm"
)

type DesignRepository struct {
	DB *gorm.DB
}

func NewDesignRepository(db *gorm.DB) *DesignRepository {
	return &DesignRepository{DB: db}
}

func (r *DesignRepository) Create(d *models.Design) error {
	return r.DB.Create(d).Error
}

func (r *DesignRepository) FindByID(id string) (*models.Design, error) {
	var d models.Design
	err := r.DB.First(&d, "id = ?", id).Error
	return &d, err
}

func (r *DesignRepository) FindByDesigner(designerID string) ([]models.Design, error) {
	var designs []models.Design
	err := r.DB.Where("designer_id = ?", designerID).Find(&designs).Error
	return designs, err
}

func (r *DesignRepository) Update(d *models.Design) error {
	return r.DB.Save(d).Error
}

func (r *DesignRepository) Delete(id string) error {
	return r.DB.Delete(&models.Design{}, "id = ?", id).Error
}