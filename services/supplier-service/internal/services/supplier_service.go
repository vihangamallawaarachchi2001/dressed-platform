/**
 * Business logic for supplier lifecycle.
 */
package services

import (
	"errors"

	"supplier-service/internal/models"
	"supplier-service/internal/repositories"

	"github.com/google/uuid"
)

type SupplierService struct {
	repo *repositories.SupplierRepository
}

func NewSupplierService(r *repositories.SupplierRepository) *SupplierService {
	return &SupplierService{repo: r}
}

func (s *SupplierService) Register(userID, company, desc, caps string) (*models.Supplier, error) {
	supplier := &models.Supplier{
		ID:           uuid.NewString(),
		UserID:       userID,
		CompanyName:  company,
		Description:  desc,
		Capabilities: caps,
		Status:       "PENDING",
		Availability: "OFFLINE",
	}
	return supplier, s.repo.Create(supplier)
}

func (s *SupplierService) Activate(userID string) error {
	supplier, err := s.repo.FindByUserID(userID)
	if err != nil {
		return err
	}
	supplier.Status = "ACTIVE"
	supplier.Availability = "AVAILABLE"
	return s.repo.Update(supplier)
}

func (s *SupplierService) UpdateAvailability(userID, availability string) error {
	if availability != "AVAILABLE" && availability != "BUSY" && availability != "OFFLINE" {
		return errors.New("invalid availability")
	}
	supplier, err := s.repo.FindByUserID(userID)
	if err != nil {
		return err
	}
	supplier.Availability = availability
	return s.repo.Update(supplier)
}

func (s *SupplierService) GetProfile(userID string) (*models.Supplier, error) {
	return s.repo.FindByUserID(userID)
}
