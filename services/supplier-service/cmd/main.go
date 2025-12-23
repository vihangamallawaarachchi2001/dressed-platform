package main

import (
	"log"
	"supplier-service/internal/database"
	"supplier-service/internal/handlers"
	"supplier-service/internal/middleware"
	"supplier-service/internal/repositories"
	"supplier-service/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	// Supplier profile
	supplierRepo := &repositories.SupplierRepository{}
	supplierService := services.NewSupplierService(supplierRepo)

	// Quote submission
	quoteRepo := &repositories.QuoteRepository{}
	quoteService := services.NewQuoteService(quoteRepo)

	// Public designs
	designRepo := &repositories.DesignRepository{}

	handler := handlers.NewSupplierHandler(supplierService, quoteService, designRepo)

	r := gin.Default()

	// Public: anyone can view designs (later restrict to registered suppliers)
	r.GET("/designs", handler.ListPublicDesigns)

	// Protected: supplier actions
	supplierGroup := r.Group("/supplier")
	supplierGroup.Use(middleware.RequireAuth(), middleware.RequireRole("SUPPLIER"))

	supplierGroup.POST("/register", handler.Register)
	supplierGroup.POST("/activate", handler.Activate)
	supplierGroup.PATCH("/availability", handler.UpdateAvailability)
	supplierGroup.GET("/profile", handler.Profile)

	// Quote submission
	supplierGroup.POST("/quotes", handler.SubmitQuote)

	log.Println("🏭 Supplier-service running on :8004")
	r.Run(":8004")
}