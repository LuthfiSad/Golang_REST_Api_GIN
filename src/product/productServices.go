package product

import (
	"errors"
	"gin-simple-api/databases"
	"gin-simple-api/models"
	"gin-simple-api/src/types"
	"math"
	"strconv"

	"github.com/google/uuid"
)

func GetAllProducts(perPage, page int, search string) ([]models.Product, types.PaginationMeta, error) {
	var products []models.Product
	var totalData int64
	var dbQuery = databases.DB.Model(&models.Product{})

	if search != "" {
		dbQuery = dbQuery.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := dbQuery.Count(&totalData).Error; err != nil {
		return nil, types.PaginationMeta{}, err
	}

	if err := dbQuery.Limit(perPage).Offset((page - 1) * perPage).Find(&products).Error; err != nil {
		return nil, types.PaginationMeta{}, err
	}

	// if err := databases.DB.Model(&models.Product{}).Count(&totalData).Error; err != nil {
	// 	return nil, types.PaginationMeta{}, err
	// }

	// if err := databases.DB.Limit(perPage).Offset((page - 1) * perPage).Find(&products).Error; err != nil {
	// 	return nil, types.PaginationMeta{}, err
	// }

	totalPages := int(math.Ceil(float64(totalData) / float64(perPage)))

	meta := types.PaginationMeta{
		PerPage:    perPage,
		Page:       page,
		TotalData:  int(totalData),
		TotalPages: totalPages,
	}

	return products, meta, nil
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product
	if err := databases.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return models.Product{}, errors.New("product not found")
	}
	return product, nil
}

func ValidateProduct(name, description, priceStr, quantityStr string) (models.Product, error) {
	if name == "" {
		return models.Product{}, errors.New("name are required")
	}

	if description == "" {
		return models.Product{}, errors.New("description are required")
	}

	if priceStr == "" {
		return models.Product{}, errors.New("price are required")
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return models.Product{}, errors.New("invalid price")
	}

	if quantityStr == "" {
		return models.Product{}, errors.New("quantity are required")
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return models.Product{}, errors.New("invalid quantity")
	}

	product := models.Product{Name: name, Description: description, Price: price, Quantity: quantity}
	return product, nil
}

func CreateProductServices(product models.Product) (models.Product, error) {
	if err := databases.DB.Create(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func UpdateProductServices(id uuid.UUID, product models.Product) (models.Product, error) {
	if err := databases.DB.Model(&product).Where("id = ?", id).Updates(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func DeleteProductServices(id uuid.UUID) error {
	if err := databases.DB.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}
