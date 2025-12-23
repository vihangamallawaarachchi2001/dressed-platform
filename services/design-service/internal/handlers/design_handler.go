package handlers

import (
	"net/http"
	"os"
	"strings"

	service "design-service/internal/services"
	"design-service/internal/utils"

	"github.com/gin-gonic/gin"
)

type DesignHandler struct {
	designService *service.DesignService
}

func NewDesignHandler(s *service.DesignService) *DesignHandler {
	return &DesignHandler{designService: s}
}

func (h *DesignHandler) CreateDesign(c *gin.Context) {
	designerID := c.GetString("user_id")
	if designerID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
		return
	}

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	category := c.PostForm("category")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	if title == "" || category == "" || description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title, description, and category are required"})
		return
	}

	filePath, err := utils.SaveFile(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	design, err := h.designService.Create(designerID, title, description, category, filePath)
	if err != nil {
		// Clean up file on failure
		os.Remove(strings.TrimPrefix(filePath, "/"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          design.ID,
		"title":       design.Title,
		"category":    design.Category,
		"description": design.Description,
		"file_url":    design.FilePath,
		"status":      design.Status,
	})
}

func (h *DesignHandler) ListDesigns(c *gin.Context) {
	designerID := c.GetString("user_id")
	designs, err := h.designService.ListByDesigner(designerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, designs)
}

func (h *DesignHandler) Submit(c *gin.Context) {
	designerID := c.GetString("user_id")
	id := c.Param("id")
	err := h.designService.Submit(id, designerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Design submitted successfully"})
}