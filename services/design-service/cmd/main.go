package main

import (
	"log"
	"design-service/internal/database"
	"design-service/internal/handlers"
	"design-service/internal/middleware"
	"design-service/internal/repositories"
	"design-service/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	// Initialize repository with DB
	repo := repositories.NewDesignRepository(database.DB)
	svc := service.NewDesignService(repo)
	h := handlers.NewDesignHandler(svc)

	r := gin.Default()

	// Serve uploaded files
	r.Static("/uploads", "./uploads")

	// Protected routes for designers
	designerGroup := r.Group("/designs")
	designerGroup.Use(middleware.RequireAuth(), middleware.RequireRole("DESIGNER"))
	designerGroup.GET("", h.ListDesigns)
	designerGroup.POST("", h.CreateDesign)
	designerGroup.PATCH("/:id/submit", h.Submit)

	log.Println("🎨 Design-service running on :8002")
	r.Run(":8002")
}