package handlers

import (

	"payment-service/internal/services"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(s *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: s}
}

func (h *PaymentHandler) Initiate(c *gin.Context) {
	var req struct {
		OrderID string  `json:"order_id"`
		Amount  float64 `json:"amount"`
	}
	c.BindJSON(&req)

	p, err := h.service.Initiate(req.OrderID, req.Amount)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, p)
}

func (h *PaymentHandler) Confirm(c *gin.Context) {
	id := c.Param("id")

	p, err := h.service.Confirm(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "payment not found"})
		return
	}
	c.JSON(200, p)
}

func (h *PaymentHandler) Status(c *gin.Context) {
	id := c.Param("id")

	p, err := h.service.GetStatus(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "payment not found"})
		return
	}
	c.JSON(200, p)
}
