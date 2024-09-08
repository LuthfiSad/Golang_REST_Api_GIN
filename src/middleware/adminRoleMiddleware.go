package middleware

import (
	"gin-simple-api/src/auth"
	"gin-simple-api/src/config"

	"github.com/gin-gonic/gin"
)

func AdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Request.Context().Value(config.UserIDKey).(string)
		if !exists {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "Wrong role provided",
				"code":    "UNAUTHORIZED",
			})
			c.Abort()
			return
		}
		user, err := auth.GetUserProfile(userID)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Failed to retrieve user profile",
				"code":    "INTERNAL_SERVER_ERROR",
			})
			c.Abort()
			return
		}
		if user.Role != "admin" {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "Wrong role provided",
				"code":    "UNAUTHORIZED",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
