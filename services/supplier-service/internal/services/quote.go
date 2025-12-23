package services

import (
	"supplier-service/internal/models"
	"supplier-service/internal/repositories"
	"github.com/google/uuid"
)

type QuoteService struct {
	repo *repositories.QuoteRepository
}

func NewQuoteService(r *repositories.QuoteRepository) *QuoteService {
	return &QuoteService{repo: r}
}

func (s *QuoteService) SubmitQuote(designID, designerID, supplierID string, price float64, eta int) error {
	q := &models.Quote{
		ID:         uuid.NewString(),
		DesignID:   designID,
		DesignerID: designerID,
		SupplierID: supplierID,
		Price:      price,
		ETA:        eta,
		Status:     "QUOTED",
	}
	return s.repo.Create(q)
}