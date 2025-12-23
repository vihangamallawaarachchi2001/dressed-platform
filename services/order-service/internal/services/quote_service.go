/**
 * Business logic for quote & order lifecycle.
 */
package service

import (
	"errors"

	"order-service/internal/models"
	"order-service/internal/repositories"

	"github.com/google/uuid"
)

type QuoteService struct {
	quotes *repositories.QuoteRepository
	orders *repositories.OrderRepository
}

func NewQuoteService(q *repositories.QuoteRepository, o *repositories.OrderRepository) *QuoteService {
	return &QuoteService{quotes: q, orders: o}
}

func (s *QuoteService) RequestQuote(designID, designerID string) (*models.Quote, error) {
	q := &models.Quote{
		ID:         uuid.NewString(),
		DesignID:   designID,
		DesignerID: designerID,
		Status:     "REQUESTED",
	}
	return q, s.quotes.Create(q)
}

func (s *QuoteService) SubmitQuote(id, supplierID string, price float64, eta int) error {
	q, err := s.quotes.FindByID(id)
	if err != nil {
		return err
	}
	q.SupplierID = supplierID
	q.Price = price
	q.ETA = eta
	q.Status = "QUOTED"
	return s.quotes.Update(q)
}

func (s *QuoteService) AcceptQuote(id, designerID string) (*models.Order, error) {
	q, err := s.quotes.FindByID(id)
	if err != nil || q.DesignerID != designerID {
		return nil, errors.New("forbidden")
	}
	q.Status = "ACCEPTED"
	s.quotes.Update(q)

	o := &models.Order{
		ID:         uuid.NewString(),
		QuoteID:    q.ID,
		DesignerID: q.DesignerID,
		SupplierID: q.SupplierID,
		Status:     "CREATED",
	}
	return o, s.orders.Create(o)
}
