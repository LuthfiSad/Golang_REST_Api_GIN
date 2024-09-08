package databases

import (
	"fmt"
	"gin-simple-api/models"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Generate random data for seeding
func GenerateRandomProduct() models.Product {
	return models.Product{
		ID:          uuid.New(),
		Name:        fmt.Sprintf("Product %d", rand.Intn(1000)),
		Description: "Sample description",
		Price:       rand.Intn(1000) + 1,
		Image:       "http://localhost:8080/src/product/images/4043435a-8cc9-433d-98a8-3ce984ee0adaIMG-20240105-WA0010.jpg",
		Quantity:    rand.Intn(100) + 1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// SeedProducts seeds the database with 50 products
func SeedProducts(db *gorm.DB) error {
	for i := 0; i < 50; i++ {
		product := GenerateRandomProduct()
		if err := db.Create(&product).Error; err != nil {
			return err
		}
	}
	return nil
}
