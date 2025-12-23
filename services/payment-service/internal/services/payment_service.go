/**
 * Business logic for mocked payment processing.
 */
package service

import (
	"math/rand"
	"time"

	"payment-service/internal/models"
	"payment-service/internal/repositories"

	"github.com/google/uuid"
)

type PaymentService struct {
	repo *repositories.PaymentRepository
}

func NewPaymentService(r *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{repo: r}
}

func (s *PaymentService) Initiate(orderID string, amount float64) (*models.Payment, error) {
	p := &models.Payment{
		ID:      uuid.NewString(),
		OrderID: orderID,
		Amount:  amount,
		Status:  "PENDING",
	}
	return p, s.repo.Create(p)
}

func (s *PaymentService) Confirm(paymentID string) (*models.Payment, error) {
	p, err := s.repo.FindByID(paymentID)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		p.Status = "SUCCESS"
	} else {
		p.Status = "FAILED"
	}

	return p, s.repo.Update(p)
}

func (s *PaymentService) GetStatus(paymentID string) (*models.Payment, error) {
	return s.repo.FindByID(paymentID)
}
