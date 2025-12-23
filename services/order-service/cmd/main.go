package main

import (
	"log"

	"order-service/internal/database"
	"order-service/internal/handlers"
	"order-service/internal/middleware"
	"order-service/internal/repositories"
	service "order-service/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	qRepo := &repositories.QuoteRepository{}
	oRepo := &repositories.OrderRepository{}
	svc := service.NewQuoteService(qRepo, oRepo)
	h := handlers.NewQuoteHandler(svc)

	r := gin.Default()

	r.Use(middleware.RequireAuth())
	r.POST("", middleware.RequireRole("designer"), h.Request)
	r.POST("/:id/respond", middleware.RequireRole("supplier"), h.Submit)
	r.POST("/:id/accept", middleware.RequireRole("designer"), h.Accept)
	// Add under protected routes
	r.GET("/my-quotes", middleware.RequireRole("designer"), h.ListQuotesForDesigner)

	log.Println("📦 Quote-service running on :8003")
	r.Run(":8003")
}
