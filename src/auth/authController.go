package auth

import (
	"gin-simple-api/src/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	perPage, err := strconv.Atoi(c.DefaultQuery("perPage", "10"))
	if err != nil {
		perPage = 10
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	search := c.Query("search")

	users, meta, err := GetAllUsers(perPage, page, search)
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
		"message": "Users retrieved successfully",
		"code":    "SUCCESS",
		"data":    users,
		"meta":    meta,
	})
}

func Register(c *gin.Context) {
	var payload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    "BAD_REQUEST",
		})
		return
	}

	if err := ValidateRegister(payload.Name, payload.Email, payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    "BAD_REQUEST",
		})
		return
	}

	user, err := RegisterUser(payload.Name, payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created successfully",
		"code":    "CREATED",
		"data":    user,
	})
}

func Login(c *gin.Context) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    "BAD_REQUEST",
		})
		return
	}

	if err := ValidateLogin(payload.Email, payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    "BAD_REQUEST",
		})
		return
	}

	user, err := LoginUser(payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
			"code":    "UNAUTHORIZED",
		})
		return
	}

	tokenString, err := GenerateJWT(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	// Kirim token sebagai response
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Login successful",
		"code":    "SUCCESS",
		"data":    gin.H{"token": tokenString},
	})
}

func Profile(c *gin.Context) {
	userID, exists := c.Request.Context().Value(config.UserIDKey).(string)
	if !exists {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "You are not authorized",
			"code":    "UNAUTHORIZED",
		})
		return
	}

	user, err := GetUserProfile(userID)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Failed to retrieve user profile",
			"code":    "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Successfully retrieved user profile",
		"code":    "SUCCESS",
		"data":    user,
	})
}
