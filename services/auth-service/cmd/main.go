// @title Auth Service API
// @version 1.0
// @description Authentication & Authorization service for Dressed™ platform
// @termsOfService https://dressed.com/terms

// @contact.name API Support
// @contact.email support@dressed.com

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// Auth-service entry point.
package main

import (
	"log"

	"auth-service/internal/database"
	"auth-service/internal/handlers"
	"auth-service/internal/middleware"
	"auth-service/internal/repositories"
	"auth-service/internal/services"

	"github.com/gin-gonic/gin"

	_ "auth-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.Connect()

	repo := &repositories.UserRepository{}
	authService := service.NewAuthService(repo)
	handler := handlers.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	protected := r.Group("/")
	protected.Use(middleware.RequireAuth())
	protected.GET("/me", handler.Me)

	log.Println("🚀 Auth service running on :8001")
	r.Run(":8001")
}
