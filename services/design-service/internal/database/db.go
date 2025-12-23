/**
 * Package database initializes PostgreSQL connection for design-service.
 */
package database

import (
	"log"
	"os"

	"design-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || user == "" || dbname == "" {
		log.Fatal("database env vars not set")
	}

	dsn := "host=" + host +
		" port=" + port +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&models.Design{}); err != nil {
		log.Fatal(err)
	}

	DB = db
	log.Println("✅ design-service DB connected")
}
