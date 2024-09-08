package routes

import (
	"gin-simple-api/src/auth"
	"gin-simple-api/src/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth_router := router.Group("/")
	auth_router.POST("/login", auth.Login)
	auth_router.POST("/register", auth.Register)
	auth_router.Use(middleware.JwtMiddleware())
	{
		auth_router.GET("/profile", auth.Profile)
		auth_router.GET("/users", middleware.AdminRoleMiddleware(), auth.GetUsers)
	}
}
