package databases

import (
	"gin-simple-api/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	if err := SeedProducts(DB); err != nil {
		log.Fatalf("failed to seed database: %v", err)
	}
	DB.AutoMigrate(&models.User{}, &models.Product{})
}
