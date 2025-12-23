package main

import (
	"log"

	"order-service/internal/database"
	"order-service/internal/handlers"
	"order-service/internal/middleware"
	"order-service/internal/repositories"
	"order-service/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	qRepo := &repositories.QuoteRepository{}
	oRepo := &repositories.OrderRepository{}
	svc := service.NewQuoteService(qRepo, oRepo)
	h := handlers.NewQuoteHandler(svc)

	r := gin.Default()

	api := r.Group("/quotes")
	api.Use(middleware.RequireAuth())
	api.POST("", middleware.RequireRole("DESIGNER"), h.Request)
	api.POST("/:id/respond", middleware.RequireRole("SUPPLIER"), h.Submit)
	api.POST("/:id/accept", middleware.RequireRole("DESIGNER"), h.Accept)
	// Add under protected routes
	api.GET("/my-quotes", middleware.RequireRole("DESIGNER"), h.ListQuotesForDesigner)

	log.Println("📦 Quote-service running on :8003")
	r.Run(":8003")
}
