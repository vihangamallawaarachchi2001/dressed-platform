package service

import (
	"errors"

	"design-service/internal/models"
	"design-service/internal/repositories"

	"github.com/google/uuid"
)

type DesignService struct {
	repo *repositories.DesignRepository
}

func NewDesignService(r *repositories.DesignRepository) *DesignService {
	return &DesignService{repo: r}
}

func (s *DesignService) Create(designerID, title, desc, category, filePath string) (*models.Design, error) {
	allowedCategories := map[string]bool{
		"Men": true, "Women": true, "Boy": true,
		"Girl": true, "Unisex": true,
	}
	if !allowedCategories[category] {
		return nil, errors.New("invalid category")
	}

	d := &models.Design{
		ID:          uuid.NewString(),
		DesignerID:  designerID,
		Title:       title,
		Description: desc,
		Category:    category,
		FilePath:    filePath,
		Status:      "DRAFT",
	}
	err := s.repo.Create(d)
	return d, err
}

func (s *DesignService) Submit(id, designerID string) error {
	d, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if d.DesignerID != designerID {
		return errors.New("forbidden")
	}
	d.Status = "SUBMITTED"
	return s.repo.Update(d)
}

func (s *DesignService) ListByDesigner(designerID string) ([]models.Design, error) {
	return s.repo.FindByDesigner(designerID)
}