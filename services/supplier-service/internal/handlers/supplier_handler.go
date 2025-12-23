package handlers

import (
	"net/http"
	"supplier-service/internal/database"
	"supplier-service/internal/models"
	"supplier-service/internal/repositories"
	"supplier-service/internal/services"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	supplierService *services.SupplierService
	quoteService    *services.QuoteService
	designRepo      *repositories.DesignRepository
}

// Update NewSupplierHandler
func NewSupplierHandler(
	s *services.SupplierService,
	qs *services.QuoteService,
	dr *repositories.DesignRepository,
) *SupplierHandler {
	return &SupplierHandler{
		supplierService: s,
		quoteService:    qs,
		designRepo:      dr,
	}
}

// ListPublicDesigns lists all submitted designs
func (h *SupplierHandler) ListPublicDesigns(c *gin.Context) {
	designs, err := h.designRepo.ListPublic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, designs)
}

// SubmitQuote handles quote submission
func (h *SupplierHandler) SubmitQuote(c *gin.Context) {
	supplierID := c.GetString("user_id")
	if supplierID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Supplier not authenticated"})
		return
	}

	var req struct {
		DesignID   string  `json:"design_id" binding:"required"`
		Price      float64 `json:"price" binding:"required,gt=0"`
		ETA        int     `json:"eta_days" binding:"required,gte=1"`
		Notes      string  `json:"notes"` 
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch design to get DesignerID
	var design models.Design
	if err := database.DB.Where("id = ? AND status = ?", req.DesignID, "SUBMITTED").First(&design).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Design not found or not available for quoting"})
		return
	}

	err := h.quoteService.SubmitQuote(req.DesignID, design.DesignerID, supplierID, req.Price, req.ETA)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Quote submitted successfully"})
}

func (h *SupplierHandler) Register(c *gin.Context) {
	var req struct {
		CompanyName  string `json:"company_name"`
		Description  string `json:"description"`
		Capabilities string `json:"capabilities"`
	}
	c.BindJSON(&req)

	s, err := h.supplierService.Register(
		c.GetString("user_id"),
		req.CompanyName,
		req.Description,
		req.Capabilities,
	)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, s)
}

func (h *SupplierHandler) Activate(c *gin.Context) {
	err := h.supplierService.Activate(c.GetString("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}

func (h *SupplierHandler) UpdateAvailability(c *gin.Context) {
	var req struct {
		Availability string `json:"availability"`
	}
	c.BindJSON(&req)

	err := h.supplierService.UpdateAvailability(c.GetString("user_id"), req.Availability)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}

func (h *SupplierHandler) Profile(c *gin.Context) {
	s, err := h.supplierService.GetProfile(c.GetString("user_id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "supplier not found"})
		return
	}
	c.JSON(200, s)
}
