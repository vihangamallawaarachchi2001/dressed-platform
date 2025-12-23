package main

import (
	"log"

	"payment-service/internal/database"
	"payment-service/internal/handlers"
	"payment-service/internal/middleware"
	"payment-service/internal/repositories"
	"payment-service/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	repo := &repositories.PaymentRepository{}
	svc := service.NewPaymentService(repo)
	h := handlers.NewPaymentHandler(svc)

	r := gin.Default()

	api := r.Group("/payments")
	api.Use(middleware.RequireAuth())

	api.POST("/initiate", h.Initiate)
	api.POST("/confirm/:id", h.Confirm)
	api.GET("/:id", h.Status)

	log.Println("💳 Payment-service running on :8005")
	r.Run(":8005")
}
