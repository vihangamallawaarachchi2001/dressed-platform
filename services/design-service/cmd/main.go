package main

import (
	"design-service/internal/database"
	"design-service/internal/handlers"
	"design-service/internal/middleware"
	"design-service/internal/repositories"
	service "design-service/internal/services"
	"log"

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
	r.Use(middleware.RequireAuth(), middleware.RequireRole("designer"))
	r.GET("/designs", h.ListDesigns)
	r.POST("/designs", h.CreateDesign)
	r.PATCH("/designs/:id/submit", h.Submit)

	log.Println("🎨 Design-service running on :8002")
	r.Run(":8002")
}