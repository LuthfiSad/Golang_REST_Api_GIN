package product

import (
	"gin-simple-api/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	perPage, err := strconv.Atoi(c.DefaultQuery("perPage", "10"))
	if err != nil {
		perPage = 10
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	search := c.Query("search")

	products, meta, err := GetAllProducts(perPage, page, search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Products retrieved successfully",
		"code":    "SUCCESS",
		"data":    products,
		"meta":    meta,
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Product retrieved successfully",
		"code":    "SUCCESS",
		"data":    product,
	})
}

func CreateProduct(c *gin.Context) {
	// Ambil input dari form
	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	quantityStr := c.PostForm("quantity")

	// Validasi produk
	product, err := ValidateProduct(name, description, priceStr, quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    "BAD_REQUEST",
		})
		return
	}

	// Upload gambar menggunakan service
	imagePath, err := services.ImageServices("product", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    "BAD_REQUEST",
		})
		return
	}

	// Set image path ke product
	product.Image = imagePath

	// Panggil service untuk menyimpan produk ke database
	_, err = CreateProductServices(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	// Kembalikan response sukses
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Product created successfully",
		"code":    "SUCCESS",
		// "data":    createdProduct,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	productById, err := GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	// Ambil input dari form
	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")
	quantityStr := c.PostForm("quantity")

	// Validasi produk
	product, err := ValidateProduct(name, description, priceStr, quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    "BAD_REQUEST",
		})
		return
	}

	// Upload gambar menggunakan service
	imagePath, err := services.ImageServices("product", c, productById.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    "BAD_REQUEST",
		})
		return
	}

	// Set image path ke product
	product.Image = imagePath

	// Panggil service untuk menyimpan produk ke database
	_, err = UpdateProductServices(productById.ID, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	// Kembalikan response sukses
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Product updated successfully",
		"code":    "SUCCESS",
		// "data":    updatedProduct,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	if err := services.RemoveImage(product.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	if err := DeleteProductServices(product.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Product deleted successfully",
		"code":    "SUCCESS",
	})
}
