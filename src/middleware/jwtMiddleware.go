package middleware

import (
	"context"
	"fmt"
	"gin-simple-api/src/config"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"code":    "UNAUTHORIZED",
				"message": "No authorization header provided",
			})
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"code":    "UNAUTHORIZED",
				"message": "Invalid token",
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(c.Request.Context(), config.UserIDKey, claims["userID"].(string))
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"code":    "UNAUTHORIZED",
				"message": "Invalid token",
			})
			c.Abort()
		}
	}
}
