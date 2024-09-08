package routes

import (
	"gin-simple-api/src/middleware"
	"gin-simple-api/src/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.Static("/src/product/images", "./src/product/images")
	router_product := router.Group("/product")
	router_product.Use(middleware.JwtMiddleware())
	{
		router_product.GET("/", product.GetProducts)
		router_product.GET("/:id", product.GetProduct)
		router_product.POST("/", product.CreateProduct)
		router_product.PUT("/:id", product.UpdateProduct)
		router_product.DELETE("/:id", product.DeleteProduct)
	}
}
