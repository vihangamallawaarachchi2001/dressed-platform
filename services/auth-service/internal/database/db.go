/**
*   Package database is responsible for initializing and exposing
*   the database connection used by the auth-service.
*/
package database

import (
	"log"
	"os"

	"auth-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection instance.
var DB *gorm.DB

// Connect initializes the PostgreSQL database connection using
// individual environment variables and performs schema migration.
func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || dbname == "" {
		log.Fatal("database environment variables are not properly set")
	}

	dsn := "host=" + host +
		" port=" + port +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.RefreshToken{}); err != nil {
		log.Fatal("database migration failed:", err)
	}

	DB = db
	log.Println("✅ Database connected and migrated")
}

