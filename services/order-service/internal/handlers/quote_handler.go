package handlers

import (
	"net/http"
	"order-service/internal/database"
	"order-service/internal/models"
	"order-service/internal/services"

	"github.com/gin-gonic/gin"
)

type QuoteHandler struct {
	service *service.QuoteService
}

func NewQuoteHandler(s *service.QuoteService) *QuoteHandler {
	return &QuoteHandler{service: s}
}

func (h *QuoteHandler) Request(c *gin.Context) {
	var req struct {
		DesignID string `json:"design_id"`
	}
	c.BindJSON(&req)

	q, err := h.service.RequestQuote(req.DesignID, c.GetString("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, q)
}

func (h *QuoteHandler) Submit(c *gin.Context) {
	var req struct {
		Price float64 `json:"price"`
		ETA   int     `json:"eta_days"`
	}
	c.BindJSON(&req)

	err := h.service.SubmitQuote(c.Param("id"), c.GetString("user_id"), req.Price, req.ETA)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}

func (h *QuoteHandler) Accept(c *gin.Context) {
	o, err := h.service.AcceptQuote(c.Param("id"), c.GetString("user_id"))
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, o)
}

func (h *QuoteHandler) ListQuotesForDesigner(c *gin.Context) {
	designerID := c.GetString("user_id")
	var quotes []models.Quote
	err := database.DB.Where("designer_id = ?", designerID).Find(&quotes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}
